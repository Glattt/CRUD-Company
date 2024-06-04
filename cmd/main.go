package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"projct/internal/router"
)

func main() {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/", router.Products(db))
	r.POST("/new", router.Create(db))
	r.PUT("/:id", router.Update(db))
	r.DELETE("/:id", router.Delete(db))
	r.GET("/:id", router.Get(db))
	//router.Create(r, db)

	r.Run() // listen and serve on 0.0.0.0:8080

}
