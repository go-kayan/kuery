package kuery

import (
	"fmt"
	"strings"
)

type KDelete struct {
	query string
	where string
}

func Delete(table string) *KDelete {
	return &KDelete{
		query: fmt.Sprintf("DELETE FROM %s", table),
	}
}

func (k *KDelete) Where(condition string) *KDelete {
	k.where = fmt.Sprintf(" WHERE %s", condition)
	return k
}

func (k *KDelete) Build() string {
	finalQuery := k.query
	if k.where != "" {
		finalQuery += k.where
	}
	return strings.TrimSpace(finalQuery)
}
