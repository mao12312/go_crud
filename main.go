package main

import (

        "strconv"
        "github.com/gin-gonic/gin"
        "github.com/jinzhu/gorm"
        _ "github.com/mattn/go-sqlite3"
        _ "net/http"
)

type Person struct {
        gorm.Model
        Name string
        Age  int
}

func db_init() {
        db, err := gorm.Open("sqlite3", "test.sqlite3")
        if err != nil {
                panic("failed to connect database\n")
        }

        db.AutoMigrate(&Person{})
}
func create(name string, age int) {
        db, err := gorm.Open("sqlite3", "test.sqlite3")
        if err != nil {
                panic("failed to connect database\n")
        }
        db.Create(&Person{Name: name, Age: age})
}
func get_all() []Person {
        db, err := gorm.Open("sqlite3", "test.sqlite3")
        if err != nil {
                panic("failed to connect database\n")
        }
        var people []Person
        db.Find(&people)
        return people

}
func main() {
        r := gin.Default()
        r.LoadHTMLGlob("html/*")
        db_init()
        r.GET("/", func(c *gin.Context) {
                people := get_all()
                c.HTML(200, "index.html", gin.H{
                        "people": people,
                })
        })
        r.POST("/new", func(c *gin.Context) {
                name := c.PostForm("name")
                age, _ := strconv.Atoi(c.PostForm("age"))
                create(name, age)
                c.Redirect(302, "/")
        })
        r.Run()
}