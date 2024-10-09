package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmadramadhannn/bluesky-clone/config"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

var conn *pgx.Conn

func connectDB(url string) error {
	var err error
	ctx := context.Background()
	conn, err = pgx.Connect(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to connect to database : %w", err)
	}
	return nil
}

func getUsers(c echo.Context) error {
	var users []User

	rows, err := conn.Query(c.Request().Context(), "SELECT username, email FROM users")
	if err != nil {
		return err
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Username, &user.Email); err != nil {
			return err
		}
		users = append(users, user)
	}

	defer rows.Close()

	return c.JSON(http.StatusOK, map[string]interface{}{"data": users, "status": "OK"})
}
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error", err.Error())
	}

	if err := connectDB(config.DATABASEURL); err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close(context.Background())

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", getUsers)

	e.GET("/greeting", func(c echo.Context) error {
		var greeting string
		if err := conn.QueryRow(c.Request().Context(), "select 'Hello, World!'").Scan(&greeting); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"data": greeting})
	})

	e.Logger.Fatal(e.Start(":2000"))
}
