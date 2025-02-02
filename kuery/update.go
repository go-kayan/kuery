package kuery

import (
	"fmt"
	"strings"
)

type KUpdate struct {
	query    string
	setParts []string
	where    string
}

func Update(table string) *KUpdate {
	return &KUpdate{
		query: fmt.Sprintf("UPDATE %s", table),
	}
}

func (k *KUpdate) Set(values map[string]interface{}) *KUpdate {
	setClauses := []string{}
	for column, value := range values {
		setClauses = append(setClauses, fmt.Sprintf("%s = '%v'", column, value))
	}
	k.setParts = append(k.setParts, strings.Join(setClauses, ", "))
	return k
}

func (k *KUpdate) Where(condition string) *KUpdate {
	k.where = fmt.Sprintf(" WHERE %s", condition)
	return k
}

func (k *KUpdate) Build() string {
	finalQuery := k.query
	if len(k.setParts) > 0 {
		finalQuery += " SET " + strings.Join(k.setParts, ", ")
	}
	if k.where != "" {
		finalQuery += k.where
	}
	return finalQuery
}
