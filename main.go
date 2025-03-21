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

//GET by ID
router.GET("/users/:id", func(c *gin.Context) {
    var user User
    id := c.Param("id")

    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
})

//DELETE by ID
router.DELETE("/users/:id", func(c *gin.Context) {
    var user User
    id := c.Param("id")

    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }

    if err := db.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
})

    router.Run(":3000")
}
