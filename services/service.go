package service

import (
	"errors"
	"strings"
	models "www/models"
	repository "www/repositories"
)

var (
	ErrDadosInvalidos = errors.New("nome é obrigatórios")
	ErrNaoEncontrado  = errors.New("usuário não encontrado")
)

type UsuarioService struct {
	repository *repository.UsuarioRepository
}

func NewUsuarioService(
	repository *repository.UsuarioRepository,
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
