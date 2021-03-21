package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

//Genre структура описывающая таблицу genre в БД
type Genre struct {
	ID        int
	GenreName string
	Table     string
	Title     string
}

func (g *Genre) GetEdit(id int) map[string]Genre {

	var genre map[string]Genre

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/genre?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&genre)

	return genre
}

func (g *Genre) GetPageElements(page int) map[string][]Genre {

	var genres map[string][]Genre

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/genres?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&genres)

	for i := 0; i < len(genres["genres"]); i++ {
		genres["genres"][i].Table = "genre"
		genres["genres"][i].Title = fmt.Sprintf("%s id(%d)", genres["genres"][i].GenreName, genres["genres"][i].ID)
	}

	return genres
}

func (g *Genre) GetAll() map[string][]Genre {

	var genres map[string][]Genre

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/genres/all", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&genres)

	for i := 0; i < len(genres["genres"]); i++ {
		genres["genres"][i].Table = "genre"
		genres["genres"][i].Title = fmt.Sprintf("%s id(%d)", genres["genres"][i].GenreName, genres["genres"][i].ID)
	}

	return genres
}
