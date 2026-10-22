package services

import (
	"errors"
	"strings"
	"www/models"
	"www/repositories"
)

var (
	ErrDadosInvalidos = errors.New("nome é obrigatórios")
	ErrNaoEncontrado  = errors.New("usuário não encontrado")
)

type UsuarioService struct {
	repository *repositories.UsuarioRepository
}

func NewUsuarioService(
	repository *repositories.UsuarioRepository,
) *UsuarioService {
	return &UsuarioService{
		repository: repository,
	}
}

func (s *UsuarioService) Create(
	nome string,
) (*models.Usuario, error) {
	nome = strings.TrimSpace(nome)

	if nome == "" {
		return nil, ErrDadosInvalidos
	}

	return s.repository.Create(nome)
}

func (s *UsuarioService) GetAll() ([]models.Usuario, error) {
	return s.repository.GetAll()
}
