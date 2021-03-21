package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

type FilmAuthor struct {
	ID         int
	FilmID     int
	AuthorID   int
	FilmName   string
	AuthorName string
	Table      string
	Title      string
}

func (f_au FilmAuthor) GetEdit(id int) map[string]FilmAuthor {

	var film_author map[string]FilmAuthor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/film_author?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&film_author)

	return film_author
}

func (f_au FilmAuthor) GetPageElements(page int) map[string][]FilmAuthor {

	var films_authors map[string][]FilmAuthor

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/films_authors?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&films_authors)

	for i := 0; i < len(films_authors["films_authors"]); i++ {
		films_authors["films_authors"][i].Table = "films_authors"
		films_authors["films_authors"][i].Title = fmt.Sprintf("%s film(%s) id(%d)",
			films_authors["films_authors"][i].AuthorName,
			films_authors["films_authors"][i].FilmName,
			films_authors["films_authors"][i].ID,
		)
	}

	return films_authors
}
