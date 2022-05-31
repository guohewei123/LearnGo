package main

import (
    "bytes"
    "database/sql"
    "encoding/json"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
    "log"
    "os"
    "time"
)

// User 定义用户表结构
type User struct {
    gorm.Model
    Name         string
    Email        *string
    Age          uint8
    Birthday     *time.Time
    MemberNumber sql.NullString
    ActivatedAt  sql.NullTime
}

// crateTables 创建表
func crateTables(db *gorm.DB) {
    err := db.AutoMigrate(&User{})
    if err != nil {
        panic(err)
    }
}

// dbConn 建立表连接
func dbConn() *gorm.DB {
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
        logger.Config{
            SlowThreshold:             time.Second, // 慢 SQL 阈值
            LogLevel:                  logger.Info, // 日志级别
            IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
            Colorful:                  true,        // 禁用彩色打印
        },
    )

    // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
    // [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
    dsn := "root:123456@tcp(192.168.10.53:3306)/learn_gorm?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
        Logger: newLogger,
    })
    if err != nil {
        panic(err)
    }
    return db
}

// insertRecord 插入表数据
func insertRecord(db *gorm.DB) {
    //INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`) VALUES ('2022-04-30 23:27:42.704','2022-04-30 23:27:42.704',NULL,'A','123@123.com',18,'2022-04-30 15:27:42.654','3','2022-04-30 15:27:42.654')
    email := "123@123.com"
    birthday := time.Now().UTC()
    db.Create(&User{
        Name:     "A",
        Email:    &email,
        Age:      18,
        Birthday: &birthday,
        MemberNumber: sql.NullString{
            String: "3",
            Valid:  true,
        },
        ActivatedAt: sql.NullTime{
            Time:  time.Now().UTC(),
            Valid: true,
        },
    })

    db.Create(&User{Name: "B", Age: 15})
    db.Create(&User{Name: "C", Age: 16})
    db.Create(&User{Name: "D", Age: 17})
}

// query sql injection 查询条件
// 官方文档 https://gorm.io/zh_CN/docs/security.html
func queryInjection1(db *gorm.DB) {
    var user User
    userInput := "jinzhu;drop table users;"
    // 使用问号占位符会被转义，不会sql注入
    // SELECT * FROM `users` WHERE name = 'jinzhu;drop table users;' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
    db.Model(&user).Where("name = ?", userInput).First(&user)
    printStruct(user)

    // 不使用问号占位符，不会被转义，sql注入成功
    // SELECT * FROM `users` WHERE user = jinzhu;drop table users; AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
    db.Model(&user).Where(fmt.Sprintf("user = %v", userInput)).First(&user)
    printStruct(user)
}

// query sql injection 查询条件
// 官方文档 https://gorm.io/zh_CN/docs/security.html
func queryInjection2(db *gorm.DB) {
    var user User
    userInput := "jinzhu;drop table users;"
    // 会被转义
    // SELECT * FROM `users` WHERE name = 'jinzhu;drop table users;' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
    db.First(&user, "name = ?", userInput)
    printStruct(user)
    // SQL 注入
    // SELECT * FROM `users` WHERE name = jinzhu;drop table users; AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
    db.First(&user, fmt.Sprintf("name = %v", userInput))
    printStruct(user)

}

// query sql injection 查询条件
// 官方文档 https://gorm.io/zh_CN/docs/security.html
func queryInjection3(db *gorm.DB) {
    var user User
    //userInputID := "1"
    userInputID := "1=1;drop table users;"

    // 安全的，返回 err
    /*id, err := strconv.Atoi(userInputID)
      if err != nil {
          panic(err)
      }
      db.First(&user, id)
      printStruct(user)*/

    // SQL 注入
    //  SELECT * FROM `users` WHERE 1=1;drop table users; AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
    db.First(&user, userInputID)
    printStruct(user)
}

// query sql injection 查询条件
// 官方文档 https://gorm.io/zh_CN/docs/security.html
func queryInjection4(db *gorm.DB) {
    //userInputID := "1"
    // 会被转义
    userInputID := "1 OR 1=1"
    var users []User

    // SELECT * FROM `users` WHERE id = '1 OR 1=1' AND `users`.`deleted_at` IS NULL
    db.Find(&users, "id = ?", userInputID)
    printStruct(users)

    // SQL 注入
    // SELECT * FROM `users` WHERE (id = 1 OR 1=1) AND `users`.`deleted_at` IS NULL
    db.Find(&users, fmt.Sprintf("id = %s", userInputID))
    printStruct(users)

}

func main() {
    db := dbConn()
    //crateTables(db)
    //insertRecord(db)
    queryInjection4(db)

    fmt.Println("OK")
}

func printStruct(in interface{}) {
    bs, _ := json.Marshal(in)
    var out bytes.Buffer
    json.Indent(&out, bs, "", "\t")
    fmt.Printf("res = %v\n", out.String())
}
