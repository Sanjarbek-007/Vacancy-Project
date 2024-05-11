package main

import (
	"fmt"
	"hh/login"
	"hh/menyu"
)

func main() {
	fmt.Println("Welcome to our company ")
	AorU := login.LoginFunc()
	menyu.Menyu(AorU)
}
