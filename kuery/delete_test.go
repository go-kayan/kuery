package kuery

import (
	"testing"
)

func TestDeleteBasic(t *testing.T) {
	query := Delete("usuarios").Build()
	expected := "DELETE FROM usuarios"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestDeleteWithWhere(t *testing.T) {
	query := Delete("usuarios").
		Where("id = 1").
		Build()

	expected := "DELETE FROM usuarios WHERE id = 1"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestDeleteWithLimit(t *testing.T) {
	query := Delete("usuarios").
		Limit(1).
		Build()

	expected := "DELETE FROM usuarios LIMIT (1)"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestDeleteFull(t *testing.T) {
	query := Delete("usuarios").
		Where("id = 1").
		Limit(1).
		Build()

	expected := "DELETE FROM usuarios WHERE id = 1 LIMIT (1)"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}
