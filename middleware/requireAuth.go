package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/atyagi012/go-user-auth/config"
	"github.com/atyagi012/go-user-auth/db"
	"github.com/atyagi012/go-user-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	fmt.Println("In middleware")

	//get cookie of req.
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.SecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//Find user with token sub
		var user models.User
		db.Database.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//Attach to req.
		c.Set("user", user)

		fmt.Println(`claims["exp"], claims["sub"] = `, claims["exp"], claims["sub"])

		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
