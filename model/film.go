package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

type Film struct {
	ID          int
	FilmName    string
	Description string
	FilmYear    string
	Budget      int
	FileURL     string
	PosterURL   string
	BannerURL   string
	Table       string
	Title       string
}

func (f *Film) GetEdit(id int) map[string]Film {
	var film map[string]Film

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/film?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json_err := json.NewDecoder(r.Body).Decode(&film)
	if json_err != nil {
		fmt.Println(json_err)
	}

	return film
}

func (f *Film) GetPageElements(page int) map[string][]Film {
	var films map[string][]Film

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/films?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json_err := json.NewDecoder(r.Body).Decode(&films)
	if json_err != nil {
		fmt.Println(json_err)
	}

	for i := 0; i < len(films["films"]); i++ {
		films["films"][i].Table = "film"
		films["films"][i].Title = fmt.Sprintf("%s id(%d)", films["films"][i].FilmName, films["films"][i].ID)
	}

	return films
}

func (f *Film) GetAll() map[string][]Film {
	var films map[string][]Film

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/films/all", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json_err := json.NewDecoder(r.Body).Decode(&films)
	if json_err != nil {
		fmt.Println(json_err)
	}

	for i := 0; i < len(films["films"]); i++ {
		films["films"][i].Table = "film"
		films["films"][i].Title = fmt.Sprintf("%s id(%d)", films["films"][i].FilmName, films["films"][i].ID)
	}

	return films
}
