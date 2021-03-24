package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-admin/model"
	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

func getFilmObjects(c *gin.Context) {
	tables := model.GetTables()
	films := FilmObj.GetPageElements(1)
	pageCount := model.GetPageCount("film")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": films["films"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getAuthorObjects(c *gin.Context) {
	tables := model.GetTables()
	authors := AuthorObj.GetPageElements(1)
	pageCount := model.GetPageCount("author")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": authors["authors"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getActorObjects(c *gin.Context) {
	tables := model.GetTables()
	actors := ActorObj.GetPageElements(1)
	pageCount := model.GetPageCount("actor")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": actors["actors"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getGenreObjects(c *gin.Context) {
	tables := model.GetTables()
	genres := GenreObj.GetPageElements(1)
	pageCount := model.GetPageCount("genre")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": genres["genres"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getFilmCommentObjects(c *gin.Context) {
	tables := model.GetTables()
	film_comments := FilmCommentObj.GetPageElements(1)
	pageCount := model.GetPageCount("film_comments")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": film_comments["film_comments"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getFilmActorObjects(c *gin.Context) {
	tables := model.GetTables()
	films_actors := FilmActorObj.GetPageElements(1)
	pageCount := model.GetPageCount("film_actor")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": films_actors["films_actors"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getFilmAuthorObjects(c *gin.Context) {
	tables := model.GetTables()
	films_authors := FilmAuthorObj.GetPageElements(1)
	pageCount := model.GetPageCount("film_author")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": films_authors["films_authors"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func getFilmGenreObjects(c *gin.Context) {
	tables := model.GetTables()
	films_genres := FilmGenreObj.GetPageElements(1)
	pageCount := model.GetPageCount("film_genre")

	c.HTML(200, "admin_fields.html", gin.H{
		"Tables":  tables["tables"],
		"Objects": films_genres["films_genres"],
		"Pages":   pageCount["page_count"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}
