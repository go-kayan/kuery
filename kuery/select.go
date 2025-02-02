package kuery

import (
	"fmt"
	"strings"
)

type KSelect struct {
	query string
}

func Select(columns ...string) *KSelect {
	selection := "*"
	if len(columns) > 0 {
		selection = strings.Join(columns, ", ")
	}
	query := fmt.Sprintf("SELECT %s", selection)
	return &KSelect{query: query}
}

func (s *KSelect) From(table string) *KSelect {
	s.query = fmt.Sprintf("%s FROM %s", s.query, table)
	return s
}

func (k *KSelect) LeftJoin(table, condition string) *KSelect {
	k.query = fmt.Sprintf("%s LEFT JOIN %s ON %s", k.query, table, condition)
	return k
}

func (k *KSelect) RightJoin(table, condition string) *KSelect {
	k.query = fmt.Sprintf("%s RIGHT JOIN %s ON %s", k.query, table, condition)
	return k
}

func (k *KSelect) InnerJoin(table, condition string) *KSelect {
	k.query = fmt.Sprintf("%s INNER JOIN %s ON %s", k.query, table, condition)
	return k
}

func (k *KSelect) Where(condition string) *KSelect {
	k.query = fmt.Sprintf("%s WHERE %s", k.query, condition)
	return k
}

func (k *KSelect) And(condition string) *KSelect {
	k.query = fmt.Sprintf("%s AND %s", k.query, condition)
	return k
}

func (k *KSelect) Or(condition string) *KSelect {
	k.query = fmt.Sprintf("%s OR %s", k.query, condition)
	return k
}

func (k *KSelect) OrderBy(columns ...string) *KSelect {
	if len(columns) == 0 {
		return k
	}
	orderClause := "ORDER BY " + strings.Join(columns, ", ")
	k.query = fmt.Sprintf("%s %s", k.query, orderClause)
	return k
}

func (k *KSelect) Limit(limit int) *KSelect {
	k.query = fmt.Sprintf("%s LIMIT %d", k.query, limit)
	return k
}

func (k *KSelect) GroupBy(columns ...string) *KSelect {
	if len(columns) == 0 {
		return k
	}
	groupByClause := "GROUP BY " + strings.Join(columns, ", ")
	k.query = fmt.Sprintf("%s %s", k.query, groupByClause)
	return k
}

func (k *KSelect) Build() string {
	return k.query
}
