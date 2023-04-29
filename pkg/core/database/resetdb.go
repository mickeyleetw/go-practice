package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
)

var (
	//go:embed sqlscripts/*.sql
	sqlScriptFS embed.FS
)

type ResetDBHelper struct {
	dropSchemaSqlScript   string
	createSchemaSqlScript string
	createTablesSqlScript string
}

func NewResetDBHelper(schema string) (*ResetDBHelper, error) {
	resetDBHelper := ResetDBHelper{}
	if err := resetDBHelper.initPureSQL(schema); err != nil {
		return nil, err
	}
	return &resetDBHelper, nil
}

// init ResetDBHelper Struct with pure sql
func (o *ResetDBHelper) initPureSQL(schema string) error {
	o.dropSchemaSqlScript = fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE;", schema)
	o.createSchemaSqlScript = fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schema)

	createTablesSqlByte, err := sqlScriptFS.ReadFile("sqlscripts/create-tables.sql")
	if err != nil {
		return err
	}
	o.createTablesSqlScript = string(createTablesSqlByte)

	return nil
}

func (o *ResetDBHelper) Reset(ctx context.Context, db *sql.DB) error {
	sqlScripts := []string{
		o.dropSchemaSqlScript,
		o.createSchemaSqlScript,
		o.createTablesSqlScript,
	}
	if err := ExecSqlScripts(ctx, db, sqlScripts); err != nil {
		return err
	}
	return nil
}

func ExecSqlScripts(ctx context.Context, db *sql.DB, sqlScripts []string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, sqlScript := range sqlScripts {
		if _, err := tx.Exec(sqlScript); err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}
