package kuery

import (
	"testing"
)

func TestUpdateBasic(t *testing.T) {
	query := Update("usuarios").Build()
	expected := "UPDATE usuarios"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestUpdateSetSingle(t *testing.T) {
	query := Update("usuarios").
		Set(map[string]interface{}{"nome": "João"}).
		Build()

	expected := "UPDATE usuarios SET nome = 'João'"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestUpdateSetMultiple(t *testing.T) {
	query := Update("usuarios").
		Set(map[string]interface{}{
			"nome":  "João",
			"email": "joao@email.com",
		}).
		Build()

	expected := "UPDATE usuarios SET nome = 'João', email = 'joao@email.com'"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}

func TestUpdateWithWhere(t *testing.T) {
	query := Update("usuarios").
		Set(map[string]interface{}{"nome": "João"}).
		Where("id = 1").
		Build()

	expected := "UPDATE usuarios SET nome = 'João' WHERE id = 1"

	if query != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, query)
	}
}
