package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"projct/internal/handler"
	"strconv"
)

func Products(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := handler.Product(db)
		if err != nil {
			panic(err)

			return
		}

		c.JSON(200, products)
	}
}

func Create(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.Products{}
		if err := c.ShouldBindJSON(&u); err != nil {
			return
		}

		if createErr := u.Create(db); createErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product created", "product": u})
	}
}

func Update(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		u := handler.Products{ID: int64(idNum)}
		if err := c.ShouldBindJSON(&u); err != nil {
			return
		}

		if updateErr := u.Update(db); updateErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product updated", "product": u})
	}
}

func Delete(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		p := handler.Products{ID: int64(idNum)}
		if deleteErr := p.Delete(db); deleteErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product deleted", "product": p})
	}
}

func Get(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		idNum, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		p := handler.Products{ID: int64(idNum)}
		if getErr := p.Get(db); getErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product found", "product": p})
	}
}
