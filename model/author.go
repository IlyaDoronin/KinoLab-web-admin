package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

//Author структура описывающая таблицу author в БД
type Author struct {
	ID    int
	FName string
	LName string
	Table string
	Title string
}

func (au *Author) GetEdit(id int) map[string]Author {

	var author map[string]Author

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/edit/author?id=%d", conf.Api.Host, conf.Api.Port, id))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&author)

	return author
}

func (au *Author) GetPageElements(page int) map[string][]Author {

	var authors map[string][]Author

	r, err := http.Get(fmt.Sprintf("http://%s:%s/get/table/authors?page=%d", conf.Api.Host, conf.Api.Port, page))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&authors)

	for i := 0; i < len(authors["authors"]); i++ {
		authors["authors"][i].Table = "author"
		authors["authors"][i].Title = fmt.Sprintf("%s %s id(%d)", authors["authors"][i].LName, authors["authors"][i].FName, authors["authors"][i].ID)
	}

	return authors
}
