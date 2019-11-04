package main

import (
	_ "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Person struct
type Person struct {
	gorm.Model
	Name string `validate:"required"`
	Age  int `validate:"required"`
}

func dbInit() {
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

func delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	var person Person
	db.First(&person, id)
	db.Delete(&person)
	db.Close()
}

func update(id int, name string, age int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	var person Person

	db.First(&person, id)
	person.Name = name
	person.Age = age
	db.Save(&person)
	db.Close()
}

func getAll() []Person {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	var people []Person
	db.Find(&people)
	db.Close()
	return people
}

// DbGetOne get a data
func DbGetOne(id int) Person {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database(DbGetOne)\n")
	}
	var person Person
	db.First(&person, id)
	db.Close()
	return person
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	dbInit()

	r.GET("/", func(c *gin.Context) {
		people := getAll()
		c.HTML(200, "index.html", gin.H{
			"people": people,
		})
	})

	// validate := validator.New()
	// errors := validate.Struct(user)
	r.GET("/delete_check/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		person := DbGetOne(id)
		c.HTML(200, "delete.html", gin.H{"person": person})
	})

	r.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		create(name, age)
		c.Redirect(302, "/")
	})

	r.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		delete(id)
		c.Redirect(302, "/")
	})

	r.GET("/edit/:id", func(c *gin.Context){
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil{
			panic("ERROR")
		}
		person := DbGetOne(id)
		c.HTML(200, "edit.html", gin.H{"person":person})

	})

	r.POST("update/:id", func(c *gin.Context){
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil{
			panic("ERROR")
		}
		name := c.PostForm("name")
		a := c.PostForm("age")
		age, errAge := strconv.Atoi(a)
		if errAge != nil {
			panic(errAge)
		}
		update(id, name, age)
		c.Redirect(302, "/")
	})
	r.Run()
}