package main

import "fmt"
import motor "github.com/rubencosta/ev3dev-lang-go/motor"

func main() {
	if motor, err := motor.Connect("outA", ""); err != nil {
		fmt.Println("connect failed\nerr:", err)
	} else {
		fmt.Println("connect success", motor)
	}
}
