package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	handler "www/handlers"
	repository "www/repositories"
	service "www/services"
)

func main() {
	db, err := sql.Open("sqlite", "./banco.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS usuarios (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("sql"); err != nil {
		log.Fatal("Pasta sql não encontrada: ", err)
	}

	usuarioRepository := repository.NewUsuarioRepository(db)

	usuarioService := service.NewUsuarioService(usuarioRepository)

	usuarioHandler := handler.NewUsuarioHandler(usuarioService)

	mux := http.NewServeMux()

	mux.HandleFunc("/usuarios", usuarioHandler.HandleCollection)

	log.Println("Servidor iniciado em http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
