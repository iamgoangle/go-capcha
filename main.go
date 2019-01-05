package main

import (
	"fmt"

	"github.com/iamgoangle/go-capcha/captcha"
)

func main() {
	fmt.Println(captcha.Encode(1, 1, 1, 1))
	fmt.Println(captcha.Encode(1, 1, 1, 2))
	fmt.Println(captcha.Encode(1, 1, 1, 3))
	fmt.Println(captcha.Encode(1, 1, 1, 4))
	fmt.Println(captcha.Encode(1, 1, 1, 5))
}
