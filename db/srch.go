package db

import (
	"database/sql"
	"kops/query"
)

type SearchSQL struct {
	sql          string
	lastPosition int
}

func (s *SearchSQL) AddPosition(idx int, word string) {
	s.lastPosition++
	s.sql += "JOIN token as t3 ON t3.id = t2.id - 1 "
}

func NewSearchSQL() *SearchSQL {
	query := "SELECT array_agg(t1.id ORDER BY t1.id), array_agg(t1.pa_word order by t1.id) " +
		"FROM token AS t1 " +
		"JOIN token AS t2 ON t2.id - 3  <= t1.id " +
		"AND t1.id <= t2.id + 3 "
	return &SearchSQL{
		sql: query,
	}
}

type Search struct {
	srchStmt *sql.Stmt
}

func (s *Search) run() {
	// FOO
}

func NewSearch(db *sql.DB, query *query.Query) (*Search, error) {
	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}
	sql := "SELECT array_agg(t1.id ORDER BY t1.id), array_agg(t1.pa_word order by t1.id) " +
		"FROM token AS t1 " +
		"JOIN token AS t2 ON t2.id - 3  <= t1.id " +
		"AND t1.id <= t2.id + 3 " +
		"JOIN token as t3 ON t3.id = t2.id - 1 " +
		"WHERE t2.pa_word = 'je' AND t3.pa_word = 'to' " +
		"GROUP BY t2.id " +
		"LIMIT 50 OFFSET 0"
	stmt, err := txn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return &Search{
		srchStmt: stmt,
	}, nil
}
