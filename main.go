package main

import (
	"net/http"
	"strconv"

	"github.com/jumadimuhammad/http-request/model"
	"github.com/labstack/echo"
)

func main() {
	store := model.NewArticleStoreInMemory()

	e := echo.New()

	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap

		return c.JSON(http.StatusOK, articles)
	})
	e.GET("/articles/:id", func(c echo.Context) error {
		articles := store.ArticleMap
		id, _ := strconv.Atoi(c.Param("id"))

		return c.JSON(http.StatusOK, articles[id-1])
	})

	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")

		article, _ := model.CreateArticle(title, body)

		store.Save(article)

		return c.JSON(http.StatusOK, article)
	})

	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")
		id, _ := strconv.Atoi(c.Param("id"))

		store.Put(id, title, body)

		return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		store.Delete(id)

		return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
