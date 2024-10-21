package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type TODO struct {
	Name string
	Done bool
}

func main() {
	//mysql://root:123456@localhost:3306/todo_list
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/homework")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	api := r.Group("/api")
	api.GET("/post", func(c *gin.Context) {
		n := c.PostForm("name")
		_, err := db.Exec("INSERT INTO todo_list (name, done) VALUES (?, ?)", n, "false")

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, "Added")
	})
	api.GET("/list", func(c *gin.Context) {
		var result string
		sel := c.PostForm("select")
		if sel == "true" {
			result = ""
			rows, err := db.Query("SELECT * FROM todo_list WHERE done = ?", "true")
			if err != nil {
				panic(err)
			}
			for rows.Next() {
				var name string
				var done string
				err = rows.Scan(&name, &done)
				if err != nil {
					panic(err)
				}
				result += name + " " + done + "\n"
			}

			c.JSON(http.StatusOK, result)
		} else if sel == "false" {
			result = ""
			rows, err := db.Query("SELECT * FROM todo_list WHERE done = ?", "false")
			if err != nil {
				panic(err)
			}
			for rows.Next() {
				var name string
				var done string
				err = rows.Scan(&name, &done)
				if err != nil {
					panic(err)
				}
				result += name + " " + done + "\n"
			}
			c.JSON(http.StatusOK, result)
		} else {
			result = ""
			rows, err := db.Query("SELECT * FROM todo_list")
			if err != nil {
				panic(err)
			}
			for rows.Next() {
				var name string
				var done string
				err = rows.Scan(&name, &done)
				if err != nil {
					panic(err)
				}
				result += name + " " + done
			}
			c.JSON(http.StatusOK, result)
		}
	})

	api.GET("/put", func(c *gin.Context) {
		n := c.PostForm("name")
		mark := c.PostForm("mark")
		_, err := db.Exec("UPDATE todo_list SET done = ? WHERE name = ?", mark, n)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, "Updated")
	})

	api.GET("/delete", func(c *gin.Context) {
		n := c.PostForm("name")
		_, err := db.Exec("DELETE FROM todo_list WHERE name = ?", n)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, "Deleted")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
