package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Teacher 多对多 老师 学生
type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers;"`
}

// Class 一对多 学生 教师
type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

// Student 多对多 学生 老师
// 一对多 学生 教师
// 一对一 学生 ID卡
type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers;"`
}

// IDCard 一对一 学生 ID卡
type IDCard struct {
	gorm.Model
	StudentID uint
	Num       string
}

func crateTables(db *gorm.DB) {
	// 创建表，顺序不对会报错
	err := db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Create table OK.")
}

func AddRecord(db *gorm.DB) {
	idCard := IDCard{
		Num: "YH200000",
	}

	t := Teacher{
		TeacherName: "C老师",
		//Students:   []Student{s},
	}

	s := Student{
		StudentName: "HeHe2",
		IDCard:      idCard,
		Teachers:    []Teacher{t},
	}

	c := Class{
		ClassName: "小教师",
		Students:  []Student{s},
	}

	_ = db.Create(&c).Error
	fmt.Println("Add record success")
}

func initDBConn() *gorm.DB {
	dsn := "root:123456@tcp(192.168.10.53:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := initDBConn()
	crateTables(db)

	c := Class{
		ClassName: "小教师",
	}
	_ = db.Create(&c).Error

	//AddRecord(db)
	router := gin.Default()
	router.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.ShouldBindJSON(&student)
		_ = db.Create(&student).Error
		c.JSON(200, gin.H{
			"msg": "Create student success",
		})
	})

	router.GET("/student/ :ID", func(c *gin.Context) {
		studentID := c.Param("ID")
		var student Student
		db.Preload("Teachers").Preload("IDCard").Where("id = ?", studentID).First(&student)
		c.JSON(200, gin.H{
			"student": student,
		})
	})

	router.GET("/class/:ID", func(c *gin.Context) {
		classID := c.Param("ID")
		var class Class
		db.Preload("Students").Preload("Students.IDCard").Preload("Students.Teachers").Where("id = ?", classID).First(&class)
		c.JSON(200, gin.H{
			"class": class,
		})
	})
	_ = router.Run(":8080")
}
