package utils_token

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user_id int) (string, error) {
	token_life, err := strconv.Atoi(os.Getenv("TOKE_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_life)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		log.Println("[utils][token][TokeValid] error parsing token : ", err)
		return err
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	token := c.Query("token")
	log.Println("token : ", token)
	if token != "" {
		return token
	}
	bearerToken := c.Get("Authorization")
	log.Println("token : ", bearerToken)
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}