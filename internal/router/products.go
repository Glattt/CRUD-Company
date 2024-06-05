package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"projct/internal/handler"
	"strconv"
)

func Create(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.Product{}
		if err := c.ShouldBindJSON(&u); err != nil {
			return
		}

		if createErr := u.Create(c, db); createErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product created", "product": u})
	}
}

func Get(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		idNum, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		p := handler.Product{ID: int64(idNum)}
		if getErr := p.Get(c, db); getErr != nil {
			c.JSON(200, gin.H{"message": "product not found"})

			return
		}

		c.JSON(200, gin.H{"message": "product found", "product": p})
	}
}

func Delete(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		p := handler.Product{ID: int64(idNum)}
		if deleteErr := p.Delete(c, db); deleteErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product deleted", "product": p})
	}
}

func Update(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			return
		}

		u := handler.Product{ID: int64(idNum)}
		if err := c.ShouldBindJSON(&u); err != nil {
			return
		}

		if updateErr := u.Update(c, db); updateErr != nil {
			return
		}

		c.JSON(200, gin.H{"message": "product updated", "product": u})
	}
}

func Products(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := handler.Products(c, db)
		if err != nil {
			panic(err)

			return
		}

		c.JSON(200, products)
	}
}
