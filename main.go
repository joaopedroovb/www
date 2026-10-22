package main

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"www/database"
	"www/handlers"
	"www/repositories"
	"www/services"
)

func main() {
	db := database.Connect()
	defer db.Close()

	database.Migrate(db)

	usuarioRepository := repositories.NewUsuarioRepository(db)

	usuarioService := services.NewUsuarioService(usuarioRepository)

	usuariohandler := handlers.NewUsuarioHandler(usuarioService)

	mux := http.NewServeMux()

	mux.HandleFunc("/usuarios", usuariohandler.HandleUser)

	log.Println("Servidor iniciado em http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
