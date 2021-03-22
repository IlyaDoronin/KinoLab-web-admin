package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

type FilmActor struct {
	ID        int
	FilmID    int
	ActorID   int
	FilmName  string
	ActorName string
	Table     string
	Title     string
}

func (f_ac FilmActor) GetEdit(id int) map[string]FilmActor {

	var film_actor map[string]FilmActor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/film_actor?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&film_actor)

	return film_actor
}

func (f_ac FilmActor) GetPageElements(page int) map[string][]FilmActor {

	var films_actors map[string][]FilmActor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/films_actors?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&films_actors)

	for i := 0; i < len(films_actors["films_actors"]); i++ {
		films_actors["films_actors"][i].Table = "film_actor"
		films_actors["films_actors"][i].Title = fmt.Sprintf("%s film(%s) id(%d)",
			films_actors["films_actors"][i].ActorName,
			films_actors["films_actors"][i].FilmName,
			films_actors["films_actors"][i].ID,
		)
	}

	return films_actors
}
