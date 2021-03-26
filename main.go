package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func dbFunc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT name from abilities")
		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error reading abilities: %q", err))
			return
		}

		defer rows.Close()
		for rows.Next() {
			var ability string
			if err := rows.Scan(&ability); err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error scanning abilities: %q", err))
				return
			}
			c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", ability))
		}
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/db", dbFunc(db))

	router.Run(":" + port)
}
