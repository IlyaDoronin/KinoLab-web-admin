package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

func GetTables() map[string][]string {

	var tables map[string][]string

	r, err := http.Get(fmt.Sprintf("http://%s:%s/tables", conf.Api.Host, conf.Api.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&tables)

	return tables
}

func GetPageCount(tableName string) map[string]int {

	var pageCount map[string]int

	r, err := http.Get(fmt.Sprintf("http://%s:%s/pageCount?table=%s", conf.Api.Host, conf.Api.Port, tableName))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	decoder_err := json.NewDecoder(r.Body).Decode(&pageCount)
	if decoder_err != nil {
		fmt.Println(decoder_err)
	}

	return pageCount

}
