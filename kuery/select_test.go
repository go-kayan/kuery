package kuery

import (
	"testing"
)

func TestSelectWithoutColumns(t *testing.T) {
	q := Select().Build()
	expected := "SELECT *"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestSelectWithColumns(t *testing.T) {
	q := Select("id", "nome").Build()
	expected := "SELECT id, nome"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestFrom(t *testing.T) {
	q := Select("id").From("usuarios").Build()
	expected := "SELECT id FROM usuarios"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestLeftJoin(t *testing.T) {
	q := Select("u.id", "p.nome").
		From("usuarios u").
		LeftJoin("pedidos p", "u.id = p.usuario_id").
		Build()
	expected := "SELECT u.id, p.nome FROM usuarios u LEFT JOIN pedidos p ON u.id = p.usuario_id"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestRightJoin(t *testing.T) {
	q := Select("u.id", "p.nome").
		From("usuarios u").
		RightJoin("pedidos p", "u.id = p.usuario_id").
		Build()
	expected := "SELECT u.id, p.nome FROM usuarios u RIGHT JOIN pedidos p ON u.id = p.usuario_id"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestInnerJoin(t *testing.T) {
	q := Select("u.id", "p.nome").
		From("usuarios u").
		InnerJoin("pedidos p", "u.id = p.usuario_id").
		Build()
	expected := "SELECT u.id, p.nome FROM usuarios u INNER JOIN pedidos p ON u.id = p.usuario_id"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestWhereAndOr(t *testing.T) {
	q := Select("id", "nome").
		From("usuarios").
		Where("id = 1").
		And("status = 'ativo'").
		Or("admin = true").
		Build()
	expected := "SELECT id, nome FROM usuarios WHERE id = 1 AND status = 'ativo' OR admin = true"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestOrderBy(t *testing.T) {
	q := Select("id", "nome").
		From("usuarios").
		OrderBy("nome ASC", "id DESC").
		Build()
	expected := "SELECT id, nome FROM usuarios ORDER BY nome ASC, id DESC"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestGroupBy(t *testing.T) {
	q := Select("status", "COUNT(*)").
		From("usuarios").
		GroupBy("status").
		Build()
	expected := "SELECT status, COUNT(*) FROM usuarios GROUP BY status"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestLimit(t *testing.T) {
	q := Select("id", "nome").
		From("usuarios").
		Limit(10).
		Build()
	expected := "SELECT id, nome FROM usuarios LIMIT 10"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}

func TestFullQuery(t *testing.T) {
	q := Select("id", "nome").
		From("usuarios").
		LeftJoin("pedidos", "usuarios.id = pedidos.usuario_id").
		Where("usuarios.status = 'ativo'").
		And("pedidos.total > 100").
		OrderBy("usuarios.nome ASC").
		Limit(5).
		Build()

	expected := "SELECT id, nome FROM usuarios LEFT JOIN pedidos ON usuarios.id = pedidos.usuario_id WHERE usuarios.status = 'ativo' AND pedidos.total > 100 ORDER BY usuarios.nome ASC LIMIT 5"

	if q != expected {
		t.Errorf("Esperado: %s, Obtido: %s", expected, q)
	}
}
