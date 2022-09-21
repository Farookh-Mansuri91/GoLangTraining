package main

import (
	"fmt"
	"encoding/json"
)

type BankUser struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func main() {
	fmt.Println("Hello")
}
