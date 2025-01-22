package database

import (
	"database/sql"
	"log"

	"github.com/alladiobb/ticket/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{DB: db}
}

func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.DB.Prepare("SELECT ID, Name, Email, CreatedAt FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientDB) Save(client *entity.Client) error {
	stmt, err := c.DB.Prepare("INSERT INTO clients (ID, Name, Email, CreatedAt) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		log.Fatalf("Erro ao salvar cliente: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("%d linhas inseridas", rowsAffected)
	// defer stmt.Close()

	return nil
}
