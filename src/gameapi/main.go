package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GlobalDB is global database pointer
var GlobalDB *gorm.DB

// SetupDB creates a db
func SetupDB() (err error) {
	GlobalDB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"root",
		"10.50.0.210",
		"3306",
		"gamedb")), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}

// User defines the user in db
type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement",json:"id"`
	Username string `gorm:"unique;not null",json:"username"`
	Password string `gorm:"not null",json:"password"`
	Data     string `json:"data"`
	gorm.Model
}

// UserService is struct
type UserService struct{}

// LoginPayload login body
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdatePayload body
type UpdatePayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Data     string `json:"data"`
}

// CreateObject to Create Table Object
func (user *User) CreateObject() error {
	err := GlobalDB.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}

// Create creates a user record in the db
func (user *User) Create() error {
	result := GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (user *User) Verify() error {
	user.Username = strings.Trim(user.Username, "")
	user.Password = strings.Trim(user.Password, "")

	if user.Username == "" || user.Password == "" {
		return errors.New("Username and Password cannot be empty.")
	}

	return nil
}

func (us UserService) InitUser(c *gin.Context) {
	var user User
	err := user.CreateObject()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Create Obejct Failed!"})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"message": "Init DB Successfully!",
	})
	return
}

func (us UserService) Retrieve(c *gin.Context) {
	var payload User
	var user User

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := GlobalDB.Where("username=?", payload.Username).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials, user does not exists",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)

	if err != nil {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials, password does not match user",
		})
		c.Abort()
		return
	}
	c.JSON(200, &user)
	return
}

func (us UserService) CreateUser(c *gin.Context) {
	var user User
	// bind json data from request
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		c.Abort()

		return
	}

	err = user.Verify()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		c.Abort()

		return
	}
	// crypt user password
	err = user.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	err = user.Create()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		c.Abort()
	}

	c.JSON(200, gin.H{
		"message": "Create User Successfully!",
	})
}

func (us UserService) UpdateUser(c *gin.Context) {
	var payload UpdatePayload
	var user User

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}
	result := GlobalDB.Where("username=?", payload.Username).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials, user does not exists",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)

	if err != nil {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials, password does not match user",
		})
		c.Abort()
		return
	}

	user.Data = payload.Data
	GlobalDB.Save(&user)

	c.JSON(200, gin.H{
		"message": "Update User Successfully!",
	})
}

// CORSMiddleware is middleware to allow cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许Origin字段的域发送请求
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		// 设置预验请求有效期，单位秒
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// 设置允许请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,PUT,GET,PATCH,OPTIONS,DELETE,UPDATE")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// OPTIONS请求返回200
		if c.Request.Method == "OPTIONS" {
			fmt.Println(c.Request.Header)
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}

	}
}

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	var us UserService

	SetupDB()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/init", us.InitUser)
	r.POST("/create", us.CreateUser)
	r.POST("/retrieve", us.Retrieve)
	r.PUT("/update", us.UpdateUser)

	r.Run()
}
