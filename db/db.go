package db

import (
	"log"
	"fmt"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Estudante struct {
	gorm.Model
	Nome  string `json:"name"`
	CPF   int `json:"cpf"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Ativo bool `json:"ativo"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("estudante.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Estudante{})

	return db
}

func AddEstudante(estudante Estudante) {
	db := Init()


	if result := db.Create(&estudante); result.Error != nil {
		fmt.Println("Error ao criar estudante")
	}

	fmt.Println("Estudante criado!")
}