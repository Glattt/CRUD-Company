package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"os"
	"projct/internal/router"
)

func init() {
	viper.SetConfigName("config")                // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./cmd")                 // path to look for the config file in
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

func main() {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	productsGroup := r.Group("/products")
	{
		productsGroup.GET("/", router.Products(conn))
		productsGroup.GET("/:id", router.Get(conn))
		productsGroup.PUT("/:id", router.Update(conn))
		productsGroup.DELETE("/:id", router.Delete(conn))
		productsGroup.POST("/new", router.Create(conn))
	}

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", router.Users(conn))
		userGroup.POST("/sign_up", router.SignUp(conn))
		userGroup.POST("/sign_in", router.SignIn(conn))
		userGroup.POST("/Reset", router.ResetPassword(conn))
		userGroup.POST("/Update/:id", router.Update(conn))
		userGroup.DELETE("/Delete/:id", router.Delete(conn))
	}

	r.Run(cfg.Server.Addr) // listen and serve on 0.0.0.0:8080
}
