package kuery

import (
	"fmt"
	"strings"
)

type KInsert struct {
	table            string
	columns          []string
	values           [][]interface{}
	returningColumns []string
}

func Insert(table string) *KInsert {
	return &KInsert{table: table}
}

func (k *KInsert) Columns(columns ...string) *KInsert {
	k.columns = columns
	return k
}

func (k *KInsert) Values(values ...interface{}) *KInsert {
	k.values = append(k.values, values)
	return k
}

// Novo método para adicionar RETURNING
func (k *KInsert) Returning(columns ...string) *KInsert {
	k.returningColumns = columns
	return k
}

func (k *KInsert) Build() string {
	if len(k.columns) == 0 || len(k.values) == 0 {
		return ""
	}

	columnStr := fmt.Sprintf("(%s)", strings.Join(k.columns, ", "))
	valueStrings := []string{}

	for _, row := range k.values {
		valueParts := []string{}
		for _, v := range row {
			valueParts = append(valueParts, fmt.Sprintf("'%v'", v))
		}
		valueStrings = append(valueStrings, fmt.Sprintf("(%s)", strings.Join(valueParts, ", ")))
	}

	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		k.table,
		columnStr,
		strings.Join(valueStrings, ", "),
	)

	// Se houver colunas em RETURNING, adiciona ao final da query
	if len(k.returningColumns) > 0 {
		query += fmt.Sprintf(" RETURNING %s", strings.Join(k.returningColumns, ", "))
	}

	return query
}
