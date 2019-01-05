package main

import (
	"fmt"

	"github.com/iamgoangle/go-capcha/captcha"
)

func main() {
	fmt.Println(captcha.Do(1, 1, 1, 1))
	fmt.Println(captcha.Do(1, 1, 1, 2))
}
