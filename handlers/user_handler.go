package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	service "www/services"
)

type UsuarioHandler struct {
	service *service.UsuarioService
}

func NewUsuarioHandler(
	service *service.UsuarioService,
) *UsuarioHandler {
	return &UsuarioHandler{
		service: service,
	}
}

type UsuarioInput struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func responderJSON(
	w http.ResponseWriter,
	status int,
	dados any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Printf("Erro ao escrever resposta JSON: %v", err)
	}
}

func responderErro(
	w http.ResponseWriter,
	status int,
	mensagem string,
) {
	responderJSON(w, status, map[string]string{
		"erro": mensagem,
	})
}

func (h *UsuarioHandler) HandleUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodPost:
		h.Create(w, r)

	case http.MethodGet:
		h.GetAll(w, r)

	default:
		w.Header().Set("Allow", "GET, POST")
		responderErro(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

func (h *UsuarioHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input UsuarioInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responderErro(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	usuario, err := h.service.Create(input.Nome)
	if err != nil {
		h.handleError(w, err)
		return
	}

	fmt.Println(usuario)

	responderJSON(w, http.StatusCreated, usuario)
}

func (h *UsuarioHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	usuarios, err := h.service.GetAll()
	if err != nil {
		h.handleError(w, err)
		return
	}

	responderJSON(w, http.StatusOK, usuarios)
}

func (h *UsuarioHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	switch {
	case errors.Is(err, service.ErrDadosInvalidos):
		responderErro(w, http.StatusBadRequest, err.Error())

	case errors.Is(err, service.ErrNaoEncontrado):
		responderErro(w, http.StatusNotFound, err.Error())

	case errors.Is(err, sql.ErrNoRows):
		responderErro(w, http.StatusNotFound, "Registro não encontrado")

	default:
		log.Printf("Erro interno: %v", err)
		responderErro(w, http.StatusInternalServerError, "Erro interno do servidor")
	}
}
