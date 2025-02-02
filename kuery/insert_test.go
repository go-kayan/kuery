package kuery

import (
	"testing"
)

func TestKInsert(t *testing.T) {
	query := Insert("usuarios").
		Columns("nome", "email").
		Values("João", "joao@email.com").
		Values("Maria", "maria@email.com").
		Returning("id", "nome").
		Build()

	expected := "INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com'), ('Maria', 'maria@email.com') RETURNING id, nome"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}
