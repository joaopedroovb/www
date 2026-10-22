package repository

import (
	"database/sql"
	"os"
	models "www/models"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{
		db: db,
	}
}

func lerSQL(arquivo string) (string, error) {
	query, err := os.ReadFile("sql/" + arquivo)

	if err != nil {
		return "", err
	}

	return string(query), nil
}

func (r *UsuarioRepository) Create(
	nome string,
) (*models.Usuario, error) {
	query, err := lerSQL("create.sql")
	if err != nil {
		return nil, err
	}

	var usuario models.Usuario

	err = r.db.QueryRow(query, nome).Scan(
		&usuario.ID,
		&usuario.Nome,
	)
	if err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (r *UsuarioRepository) GetAll() ([]models.Usuario, error) {
	query, err := lerSQL("list.sql")
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []models.Usuario{}

	for rows.Next() {
		var usuario models.Usuario

		if err := rows.Scan(
			&usuario.ID,
			&usuario.Nome,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}
