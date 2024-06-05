package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"projct/internal/handler"
)

func SignUp(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.User{}

		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := u.SignUp(c, conn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": u})
	}
}

func SignIn(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.User{}
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := u.SignIn(c, conn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "signed", "user": u})
	}
}

func Users(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.User{}
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		users, err := handler.Users(c, conn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func UpdateUser(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.User{}
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := u.Update(c, conn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": u})
	}
}

func DeleteUser(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.User{}
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := u.Delete(c, conn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": u})
	}
}

func ResetPassword(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := handler.User{}
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := u.ResetPassword(c, conn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": u})
	}
}
