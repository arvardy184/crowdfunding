package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("apa isi ",db,err)
	
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("connected")
	//create table
	// 
	//select data pada table users
	
	
}

//create table

