package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-admin/model"
)

var (
	ActorObj       model.Actor
	AuthorObj      model.Author
	FilmObj        model.Film
	GenreObj       model.Genre
	FilmCommentObj model.FilmComments
	FilmAuthorObj  model.FilmAuthor
	FilmActorObj   model.FilmActor
	FilmGenreObj   model.FilmGenre
)

func index(c *gin.Context) {
	tables := model.GetTables()

	c.HTML(200, "admin_nofields.html", gin.H{
		"Tables": tables["tables"],
	})
}

func film(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	film := FilmObj.GetEdit(ID)

	c.HTML(200, "film_page.html", gin.H{
		"ID":          film["film"].ID,
		"FilmName":    film["film"].FilmName,
		"Description": film["film"].Description,
		"FilmYear":    film["film"].FilmYear,
		"Budget":      film["film"].Budget,
		"FileURL":     film["film"].FileURL,
		"PosterURL":   film["film"].PosterURL,
		"BannerURL":   film["film"].BannerURL,
	})
}
