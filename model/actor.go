package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

type Actor struct {
	ID    int    `json:"ID,omitempty"`
	FName string `json:"FName,omitempty"`
	LName string `json:"LName,omitempty"`
	Table string
	Title string
}

func (ac *Actor) GetEdit(id int) map[string]Actor {

	var actor map[string]Actor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/actor?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&actor)

	return actor
}

func (ac *Actor) GetPageElements(page int) map[string][]Actor {

	var actors map[string][]Actor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/actors?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&actors)

	for i := 0; i < len(actors["actors"]); i++ {
		actors["actors"][i].Table = "actor"
		actors["actors"][i].Title = fmt.Sprintf("%s %s id(%d)", actors["actors"][i].LName, actors["actors"][i].FName, actors["actors"][i].ID)
	}

	return actors
}

func (ac *Actor) GetAll() map[string][]Actor {

	var actors map[string][]Actor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/actors/all", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&actors)

	for i := 0; i < len(actors["actors"]); i++ {
		actors["actors"][i].Table = "actor"
		actors["actors"][i].Title = fmt.Sprintf("%s %s id(%d)", actors["actors"][i].LName, actors["actors"][i].FName, actors["actors"][i].ID)
	}

	return actors
}
