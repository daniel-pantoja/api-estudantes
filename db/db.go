package db

import (
	"log"
	"fmt"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Estudante struct {
	gorm.Model
	Nome  string
	CPF   int
	Email string
	Idade int
	Ativo bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("estudante.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Estudante{})

	return db
}

func AddEstudante() {
	db := Init()

	estudante := Estudante {
		Nome:  "Daniel",
		CPF:   12345,
		Email: "daniel@gmail.com",
		Idade: 19,
		Ativo: true,
	}

	if result := db.Create(&estudante); result.Error != nil {
		fmt.Println("Error ao criar estudante")
	}

	fmt.Println("Estudante criado!")
}