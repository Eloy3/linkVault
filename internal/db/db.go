package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Eloy3/LinkVault/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func Init() error {

	err = godotenv.Load()
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("user=%s dbname=%s host=%s sslmode=disable", user, dbname, host)
	DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return fmt.Errorf("failed to open DB: %w", err)
	}

	err = DB.Ping()

	if err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	fmt.Println("Connected to the database")

	return nil
}

func GetBookmarks() ([]models.Bookmark, error) {
	rows, err := DB.Query("SELECT * FROM bookmark")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []models.Bookmark
	for rows.Next() {
		var b models.Bookmark
		err = rows.Scan(&b.ID, &b.Title, &b.URL, &b.Description, &b.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		bookmarks = append(bookmarks, b)
	}

	return bookmarks, nil
}
