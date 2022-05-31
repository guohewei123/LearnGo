package main

import (
    "database/sql"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
)

// Product 定义商品表结构
type Product struct {
    gorm.Model
    Price uint           `gorm:"column:product_price;not null"`
    Code  string         `gorm:"column:product_code;not null"`
    Name  sql.NullString `gorm:"default:'apple';index:idx_name;unique;type:varchar(60)"`
    Age   *uint          `gorm:"default:18"`
}

// TableName 定义表名
func (p Product) TableName() string {
    return "my_products"
}

// crateTables 创建表
// SELECT DATABASE()
// SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'learn_gorm%' ORDER BY SCHEMA_NAME='learn_gorm' DESC limit 1
// CREATE TABLE `my_products` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`product_price` bigint unsigned NOT NULL,`product_code` longtext NOT NULL,`name` longtext,PRIMARY KEY (`id`),INDEX idx_my_products_deleted_at (`deleted_at`))
func crateTables(db *gorm.DB) {
    err := db.AutoMigrate(&Product{})
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
        Logger: newLogger,
    })
    if err != nil {
        panic(err)
    }
    return db
}

// insertRecord 插入表数据
func insertRecord(db *gorm.DB) {
    // INSERT INTO `my_products` (`created_at`,`updated_at`,`deleted_at`,`product_price`,`product_code`,`name`,`age`) VALUES ('2022-04-30 12:23:49.174','2022-04-30 12:23:49.174',NULL,100,'D42','苹果',18)
    db.Create(&Product{Code: "D42", Price: 100, Name: sql.NullString{
        String: "苹果",
        Valid:  true,
    }})

    // INSERT INTO `my_products` (`created_at`,`updated_at`,`deleted_at`,`product_price`,`product_code`,`name`,`age`) VALUES ('2022-04-30 12:23:49.279','2022-04-30 12:23:49.279',NULL,100,'D43',NULL,18)
    db.Create(&Product{Code: "D43", Price: 100, Name: sql.NullString{
        String: "苹果",
        Valid:  false,
    }})
}

// updateRecord 更新记录
func updateRecord(db *gorm.DB) {
    // UPDATE `my_products` SET `product_price`=30,`updated_at`='2022-04-30 12:25:38.203' WHERE id = 1 AND `my_products`.`deleted_at` IS NULL
    db.Model(&Product{}).Where("id = ?", 1).Update("product_price", 30)

    // UPDATE `my_products` SET `updated_at`='2022-04-30 12:27:34.818',`product_price`=200,`product_code`='p10001',`name`='香蕉',`age`=23 WHERE id = 2 AND `my_products`.`deleted_at` IS NULL
    age := uint(23)
    db.Model(&Product{}).Where("id = ?", 2).Updates(Product{
        Price: 200,      // 仅更新非零值字段
        Code:  "p10001", // 仅更新非零值字段
        Age:   &age,
        Name: sql.NullString{
            String: "香蕉",
            Valid:  true,
        },
    })

    // UPDATE `my_products` SET `updated_at`='2022-04-30 12:27:34.909',`name`='',`age`=0 WHERE id = 1 AND `my_products`.`deleted_at` IS NULL
    age = 0
    db.Model(&Product{}).Where("id = ?", 1).Updates(Product{
        Price: 0,  // 仅更新非零值字段
        Code:  "", // 仅更新非零值字段
        Name: sql.NullString{ // 使用 NullString 可以更新零值(Valid=true)
            String: "",
            Valid:  true,
        },
        Age: &age, // 使用指针变量更新0值
    })
}

func main() {
    db := dbConn()
    //crateTables(db)
    //insertRecord(db)
    updateRecord(db)

    /* 根据整形主键查找
       var product Product
       db.First(&product, 1)
       fmt.Println(product)
    */

    /* 通过where条件查询
       var product Product
       db.First(&product, "code = ?", "C100123")  // 根据整形主键查找
       fmt.Println(product)
       fmt.Println("ok")*/

    /* 通过过滤查询所有
       var products []Product
       db.Find(&products, "price <= ?", 20)
       fmt.Println(products)*/

    /* 通过Where过滤查询所有
       var products []Product
       //db.Where("price < ? and code like ?", 20, "%C1%").Find(&products)
       db.Where("price < ?", 20).Or("code like ?", "%C2%").Find(&products)
       fmt.Println(products)*/

    // 修改数据
    // db.Model(&Product{}).Where("id <= ?", 4).Update("price", 2)
    //db.Model(&Product{}).Where("id <= ?", 4).Update("price", 2).Updates(Product{
    //	Code:  "C1111",
    //	Price: 5,
    //})

    /* 删除数据
       db.Where("id = ?", 3).Delete(&Product{})            // 软删除
       db.Where("id = ?", 3).Unscoped().Delete(&Product{}) // 硬删除*/

    fmt.Println("OK")
}
