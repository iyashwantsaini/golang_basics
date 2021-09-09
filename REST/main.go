package main

import (
	"context"
	"fmt"
	"restapp/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {

	//database connection
	conn, err := connectDB()
	if err != nil {
		return
	}

	router := gin.Default()

	router.Use(dbMiddleware(*conn))

	// /users/
	usersGroup := router.Group("users")
	{
		// /users/register
		usersGroup.POST("register", routes.UserRegister)
	}

	router.Run(":8080")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:postgres@localhost:5432/restapp")
	if err != nil {
		fmt.Println("Error connection : ", err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
