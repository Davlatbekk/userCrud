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
	//////////////////////////GITListUser////////////////////////////////
	users, e := c.GitListUser(&models.GetListRequest{
		Limit:  1,
		Offset: 1,
	})

	if e != nil {
		log.Println(e)
	}
	fmt.Println(users)
	/////////////////////UpdateUser///////////////////////////
	err = c.UpdateUser(&models.UpdateUser{
		Name:    "Davlat",
		Surname: "Jalolov",
	})
	if err != nil {
		log.Println(err)
	}

}
