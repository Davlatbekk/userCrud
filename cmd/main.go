package main

import (
	"fmt"
	"log"

	"app/config"
	"app/controller"
	"app/models"
	"app/storage"
)

func main() {

	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)

	id, err := c.CreateUser(
		&models.CreateUser{
			Name:    "Abduqodir",
			Surname: "Musayev",
		},
	)

	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}

	fmt.Println(id)
}
