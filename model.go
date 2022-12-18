package main

import (
	"database/sql"
)

type cake struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func (c *cake) getCake(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM cakes WHERE id = ?",
		c.ID).Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.Image, &c.CreatedAt, &c.UpdatedAt)
}

func (c *cake) updateCake(db *sql.DB) error {
	_, err := db.Exec("UPDATE cakes SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?",
		c.Title, c.Description, c.Rating, c.Image, c.ID)

	return err
}

func (c *cake) deleteCake(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM cakes WHERE id = ?", c.ID)
	return err
}

func (c *cake) createCake(db *sql.DB) error {
	res, err := db.Exec("INSERT INTO cakes(title, description, rating, image, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())",
		c.Title, c.Description, c.Rating, c.Image, c.CreatedAt, c.UpdatedAt)

	if err != nil {
		return err
	}

	// Get the ID of the inserted row
	lastID, err := res.LastInsertId()

	if err != nil {
		return err
	}

	// Select the inserted row
	row := db.QueryRow("SELECT id, title, description, rating, image FROM cakes WHERE id = ?",
		lastID).Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.Image, &c.CreatedAt, &c.UpdatedAt)

	return row
}

func getCakes(db *sql.DB) ([]cake, error) {
	rows, err := db.Query(
		"SELECT * FROM cakes ORDER BY rating DESC, title ASC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cakes := []cake{}

	for rows.Next() {
		var c cake
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Rating, &c.Image, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}

		cakes = append(cakes, c)
	}
	return cakes, nil
}
