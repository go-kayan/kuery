package kuery

import (
	"testing"
)

func TestInsertSingleRow(t *testing.T) {
	query := Insert("usuarios").
		Columns("nome", "email").
		Values("João", "joao@email.com").
		Build()

	expected := "INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com')"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestInsertMultipleRows(t *testing.T) {
	query := Insert("usuarios").
		Columns("nome", "email").
		Values("João", "joao@email.com").
		Values("Maria", "maria@email.com").
		Build()

	expected := "INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com'), ('Maria', 'maria@email.com')"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestInsertNoColumns(t *testing.T) {
	query := Insert("usuarios").Build()

	if query != "" {
		t.Errorf("Esperado string vazia, Obtido: %s", query)
	}
}

func TestInsertNoValues(t *testing.T) {
	query := Insert("usuarios").
		Columns("nome", "email").
		Build()

	if query != "" {
		t.Errorf("Esperado string vazia, Obtido: %s", query)
	}
}
