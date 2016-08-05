// Package main is the example package.
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janek-bieser/gint"
)

type user struct {
	FirstName string
	LastName  string
	Email     string
}

var users = []user{
	{"Jack", "Sparrow", "jack.sparrow@pirates.com"},
	{"Spider", "Man", "spider.man@animals.com"},
}

func main() {
	r := gin.Default()

	// Set our custom HTMLRender
	r.HTMLRender = gint.NewHTMLRender()

	r.GET("/", homePage)
	r.GET("/users", usersPage)
	r.GET("/users/:id", userDetailPage)

	r.Run()
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "home",
	})
}

func usersPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index", gin.H{
		"title": "users",
		"users": users,
	})
}

func userDetailPage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if id < 0 || err != nil {
		c.Redirect(http.StatusMovedPermanently, "/users/0")
	} else if id >= len(users) {
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/users/%d", len(users)-1))
	} else {
		c.HTML(http.StatusOK, "users/detail", gin.H{
			"title": "detail",
			"user":  users[id],
		})
	}
}
