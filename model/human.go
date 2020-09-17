package main

import "github.com/jinzhu/gorm"

func main() {
	type Human struct {
		gorm.Model
		Name   string
		Sex    byte
		Status string
	}
}
