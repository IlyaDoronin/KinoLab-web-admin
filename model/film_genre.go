package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

type FilmGenre struct {
	ID        int
	FilmID    int
	GenreID   int
	FilmName  string
	GenreName string
	Table     string
	Title     string
}

func (fg FilmGenre) GetEdit(id int) map[string]FilmGenre {

	var film_genre map[string]FilmGenre

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/film_genre?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&film_genre)

	return film_genre
}

func (fg FilmGenre) GetPageElements(page int) map[string][]FilmGenre {

	var films_genres map[string][]FilmGenre

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/films_genres?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&films_genres)

	for i := 0; i < len(films_genres["films_genres"]); i++ {
		films_genres["films_genres"][i].Table = "films_genres"
		films_genres["films_genres"][i].Title = fmt.Sprintf("%s film(%s) id(%d)",
			films_genres["films_genres"][i].GenreName,
			films_genres["films_genres"][i].FilmName,
			films_genres["films_genres"][i].ID,
		)
	}

	return films_genres
}
