package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
  )
  
  var DB *gorm.DB

  var err error

  //DATABASE CONNECTION STRING
  var DSN string = "host=localhost user=test5 password=test5 dbname=test5 port=5432 sslmode=disable"

  //OPEN DB CONNECTION FUNCTION
  func Connect(){

		//OPEN DATABASE CONNECTION
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

		//Print Error if DB Doesn't Connect
		if err != nil {
			panic("failed to connect database")
		  } else {
			fmt.Println("Database is Connected")
			fmt.Println(DB)
		  }

		//MIGRATE PERSON
		DB.AutoMigrate(&Person{})



  }