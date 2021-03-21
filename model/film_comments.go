package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

type FilmComments struct {
	ID        int
	FilmID    int
	Name      string
	Text      string
	CreatedAt string
	Table     string
	Title     string
}

func (f_c FilmComments) GetEdit(id int) map[string]FilmComments {

	var film_comment map[string]FilmComments

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/film_comment?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&film_comment)

	return film_comment
}

func (f_au FilmComments) GetPageElements(page int) map[string][]FilmComments {

	var film_comments map[string][]FilmComments

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/film_comments?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&film_comments)

	for i := 0; i < len(film_comments["film_comments"]); i++ {
		film_comments["film_comments"][i].Table = "film_comments"
		film_comments["film_comments"][i].Title = fmt.Sprintf("%s film_id(%d) id(%d)",
			film_comments["film_comments"][i].Name,
			film_comments["film_comments"][i].FilmID,
			film_comments["film_comments"][i].ID,
		)
	}

	return film_comments
}
