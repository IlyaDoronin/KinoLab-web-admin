package server

import (
	"github.com/gin-gonic/gin"
)

func createFilmObjects(c *gin.Context) {
	c.HTML(200, "create_film.html", nil)
}

func createAuthorObjects(c *gin.Context) {
	c.HTML(200, "create_author.html", nil)
}

func createActorObjects(c *gin.Context) {
	c.HTML(200, "create_actor.html", nil)
}

func createGenreObjects(c *gin.Context) {
	c.HTML(200, "create_genre.html", nil)
}

func createFilmCommentObjects(c *gin.Context) {
	c.HTML(200, "create_film_comment.html", nil)
}

func createFilmActorObjects(c *gin.Context) {
	c.HTML(200, "create_film_actor.html", nil)
}

func createFilmAuthorObjects(c *gin.Context) {
	c.HTML(200, "create_film_author.html", nil)
}

func createFilmGenreObjects(c *gin.Context) {
	c.HTML(200, "create_film_genre.html", nil)
}
