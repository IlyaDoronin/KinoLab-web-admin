package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
)

//Start Запускает роутинг
func Start() {
	r := gin.New()

	r.Static("/css", "ui/assets/css")
	r.Static("/js", "ui/assets/js")
	r.Static("/img", "ui/assets/img")
	r.LoadHTMLGlob("ui/*/*.html")

	r.GET("/", index)

	// editGroup := r.Group("/edit")
	// {
	// 	editGroup.GET("/author")
	// 	editGroup.GET("/actor")
	// 	editGroup.GET("/film")
	// 	editGroup.GET("/genre")
	// 	editGroup.GET("/film_genre")
	// 	editGroup.GET("/film_author")
	// 	editGroup.GET("/film_actor")
	// 	editGroup.GET("/film_comment")
	// }

	objectsGroup := r.Group("/table")
	{
		objectsGroup.GET("/author", getAuthorObjects)
		objectsGroup.GET("/actor", getActorObjects)
		objectsGroup.GET("/film", getFilmObjects)
		objectsGroup.GET("/genre", getGenreObjects)
		objectsGroup.GET("/film_genre", getFilmGenreObjects)
		objectsGroup.GET("/film_author", getFilmAuthorObjects)
		objectsGroup.GET("/film_actor", getFilmActorObjects)
		objectsGroup.GET("/film_comments", getFilmCommentObjects)
	}

	createGroup := r.Group("/create")
	{
		createGroup.GET("/author", createAuthorObjects)
		createGroup.GET("/actor", createActorObjects)
		createGroup.GET("/film", createFilmObjects)
		createGroup.GET("/genre", createGenreObjects)
		createGroup.GET("/film_genre", createFilmGenreObjects)
		createGroup.GET("/film_author", createFilmAuthorObjects)
		createGroup.GET("/film_actor", createFilmActorObjects)
		createGroup.GET("/film_comments", createFilmCommentObjects)
	}

	//Server run
	r.Run(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)) // listen and serve

}
