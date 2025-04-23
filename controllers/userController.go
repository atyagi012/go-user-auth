package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/atyagi012/go-user-auth/config"
	"github.com/atyagi012/go-user-auth/db"
	"github.com/atyagi012/go-user-auth/models"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.User
		//Get name email password phone from req body
		err := c.Bind(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//Hash the pwd
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user := models.User{First_name: body.First_name, Last_name: body.Last_name, Email: body.Email, Phone: body.Phone, Password: string(hashPwd)}

		//Create user
		result := db.Database.Create(&user) // pass pointer of data to Create

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Unable to create user",
			})
			return
		}

		//Respond
		//c.JSON(http.StatusOK, user)
		c.JSON(http.StatusOK, gin.H{
			"success": "user created successfully",
		})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := pgx.Connect(context.Background(), config.Config.Postgres_URL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(context.Background())

		//getAllUsers(conn *pgx.Conn)

		var greeting string
		err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(greeting)

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.User

		//Get email password from body
		err := c.Bind(&body)
		if err != nil {
			panic("Unable to parse body")
		}

		var user = models.User{Email: body.Email}
		//fmt.Println("user = ", user)

		db.Database.First(&user, "email = ?", body.Email)

		if user.ID == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "User not found",
			})
			return
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"message":  "User found",
		// 	"user":     user.First_name,
		// 	"password": user.Password,
		// })

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User name or password is incorrect",
			})
			return
		}

		//Grant access - generate jwt token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(config.Config.SecretKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create token - " + err.Error(),
			})
			return
		}

		//create cookie
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 60*60*24*30, "", "", false, true)

		//Respond
		c.JSON(http.StatusOK, gin.H{
			//"token": tokenString,
		})
	}
}

// func Validate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "I am logged in.",
// 		})
// 	}
// }

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	//user.(models.User).Email   //Can access user props 

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
