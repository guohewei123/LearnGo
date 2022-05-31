# GORM 连接、创建表、插入记录、更新记录

- 官方文档：https://gorm.io/zh_CN/docs/

## 注意事项

### 零值注意事项
1. db.Update() 是会更新零值的
2. db.Updates() 一般不会更新零值，需要定义表结构体类型使用指针或者sql.NullString等时，才会更新零值

例子：
```go
// Product 定义商品表结构
type Product struct {
    gorm.Model
    Price uint           `gorm:"column:product_price;not null"`
    Code  string         `gorm:"column:product_code;not null"`
    Name  sql.NullString `gorm:"default:'apple';index:idx_name;unique;type:varchar(60)"`
    Age   *uint          `gorm:"default:18"`
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
```