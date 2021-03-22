package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

func editFilmObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	film := FilmObj.GetEdit(id)
	c.HTML(200, "edit_film.html", gin.H{
		"Film": film["film"],
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func editAuthorObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	author := AuthorObj.GetEdit(id)
	c.HTML(200, "edit_author.html", gin.H{
		"Author": author["author"],
		"Host":   conf.Api.Host,
		"Port":   conf.Api.Port,
	})
}

func editActorObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	actor := ActorObj.GetEdit(id)

	c.HTML(200, "edit_actor.html", gin.H{
		"Actor": actor["actor"],
		"Host":  conf.Api.Host,
		"Port":  conf.Api.Port,
	})
}

func editGenreObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	genre := GenreObj.GetEdit(id)
	c.HTML(200, "edit_genre.html", gin.H{
		"Genre": genre["genre"],
		"Host":  conf.Api.Host,
		"Port":  conf.Api.Port,
	})
}

func editFilmCommentObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	film_comments := FilmCommentObj.GetEdit(id)
	c.HTML(200, "edit_film_comment.html", gin.H{
		"FilmComment": film_comments["film_comment"],
		"Host":        conf.Api.Host,
		"Port":        conf.Api.Port,
	})
}

func editFilmActorObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	film_actor := FilmActorObj.GetEdit(id)
	c.HTML(200, "edit_film_actor.html", gin.H{
		"FilmActor": film_actor["film_actor"],
		"Host":      conf.Api.Host,
		"Port":      conf.Api.Port,
	})
}

func editFilmAuthorObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	film_author := FilmAuthorObj.GetEdit(id)
	c.HTML(200, "edit_film_author.html", gin.H{
		"FilmAuthor": film_author["film_author"],
		"Host":       conf.Api.Host,
		"Port":       conf.Api.Port,
	})
}

func editFilmGenreObjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
	}
	film_genre := FilmGenreObj.GetEdit(id)
	c.HTML(200, "edit_film_genre.html", gin.H{
		"FilmGenre": film_genre["film_genre"],
		"Host":      conf.Api.Host,
		"Port":      conf.Api.Port,
	})
}
