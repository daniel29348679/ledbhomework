package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TODO struct {
	Name string
	Done bool
}

func main() {

	todolist := []TODO{}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	api := r.Group("/api")
	api.POST("/post", func(c *gin.Context) {
		n := c.PostForm("name")
		fmt.Println(n)
		todolist = append(todolist, TODO{Name: n, Done: false})
		c.JSON(http.StatusOK, "Added")
		fmt.Println(todolist)
	})
	api.POST("/list", func(c *gin.Context) {
		sel := c.PostForm("select")
		if sel == "all" {
			c.JSON(http.StatusOK, todolist)
		} else if sel == "true" {
			done := []TODO{}
			for _, v := range todolist {
				if v.Done {
					done = append(done, v)
				}
			}
			c.JSON(http.StatusOK, done)
		} else if sel == "false" {
			notdone := []TODO{}
			for _, v := range todolist {
				if !v.Done {
					notdone = append(notdone, v)
				}
			}
			c.JSON(http.StatusOK, notdone)
		}
	})

	api.POST("/put", func(c *gin.Context) {
		n := c.PostForm("name")
		mark := c.PostForm("mark")
		for i, v := range todolist {
			if v.Name == n {
				if mark == "true" {
					todolist[i].Done = true
				} else {
					todolist[i].Done = false
				}
				break
			}
		}
		c.JSON(http.StatusOK, "Updated")
	})

	api.POST("/delete", func(c *gin.Context) {
		n := c.PostForm("name")
		for i, v := range todolist {
			if v.Name == n {
				todolist = append(todolist[:i], todolist[i+1:]...)
				break
			}
		}
		c.JSON(http.StatusOK, "Deleted")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
