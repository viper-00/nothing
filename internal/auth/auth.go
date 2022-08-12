package auth

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "test-client"
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	tokenString, err := token.SignedString([]byte(GetKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetKey returns keys from file if exist or generate key
func GetKey() string {
	if _, err := os.Stat("key"); err == nil {
		b, err := ioutil.ReadFile("key")
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}
		return strings.TrimSpace(string(b))
	}
	// generate key
	key := keyGen()
	file, err := os.Create("key")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	w := bufio.NewWriter(file)
	_, err = fmt.Fprintf(w, "%v", key)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	w.Flush()

	return key
}

func keyGen() string {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
