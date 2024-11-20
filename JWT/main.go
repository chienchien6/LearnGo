package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	mySigningKey := []byte("AllYourBase")

	// Create the Claims 把結構體實例化
	claims := MyCustomClaims{
		"Chien",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 5,
			Issuer:    "test",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(ss) //token
	//time.Sleep(6 * time.Second) //讓當前執行执行当前 Goroutine（协程）睡6秒,再繼續往下執行

	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Println(token.Claims.(*MyCustomClaims).Username) //token is expired by 1s
}
