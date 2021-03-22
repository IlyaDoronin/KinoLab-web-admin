package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

func createFilmObjects(c *gin.Context) {
	c.HTML(200, "create_film.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createAuthorObjects(c *gin.Context) {
	c.HTML(200, "create_author.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createActorObjects(c *gin.Context) {
	c.HTML(200, "create_actor.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createGenreObjects(c *gin.Context) {
	c.HTML(200, "create_genre.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createFilmCommentObjects(c *gin.Context) {
	c.HTML(200, "create_film_comment.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createFilmActorObjects(c *gin.Context) {
	c.HTML(200, "create_film_actor.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createFilmAuthorObjects(c *gin.Context) {
	c.HTML(200, "create_film_author.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}

func createFilmGenreObjects(c *gin.Context) {
	c.HTML(200, "create_film_genre.html", gin.H{
		"Host": conf.Api.Host,
		"Port": conf.Api.Port,
	})
}
