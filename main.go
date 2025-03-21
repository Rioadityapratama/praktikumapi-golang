package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
    ID        uint      `gorm:"column:id;primaryKey"`
    Name      string    `gorm:"column:name"`
    Email     string    `gorm:"column:email"`
    Age       string    `gorm:"column:age"`
    CreatedAt time.Time `gorm:"column:createdAt"`
    UpdatedAt time.Time `gorm:"column:updatedAt"`
}



func main() {
    dsn := "root:@tcp(localhost:3306)/openapi?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("failed to connect database")
    }

    router := gin.Default()
//Get
    router.GET("/users", func(c *gin.Context) {
        var users []User
        db.Find(&users)
        c.JSON(http.StatusOK, gin.H{"data": users})
    })
//POST
	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		db.Create(&user)
		c.JSON(http.StatusCreated, gin.H{"data": user})
	})

	//
    router.Run(":3000")
}
