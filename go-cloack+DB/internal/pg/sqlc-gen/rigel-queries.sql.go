// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: rigel-queries.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/sqlc-dev/pqtype"
)

const checkSchemaExists = `-- name: CheckSchemaExists :one
SELECT EXISTS(SELECT 1 FROM schema WHERE id=$1)
`

func (q *Queries) CheckSchemaExists(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkSchemaExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createConfig = `-- name: CreateConfig :one
INSERT INTO config (name, description, active, tags, schema_id, values, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, description, active, schema_id, values, tags, created_by, updated_by, created_at, updated_at
`

type CreateConfigParams struct {
	Name        string                `json:"name"`
	Description sql.NullString        `json:"description"`
	Active      sql.NullBool          `json:"active"`
	Tags        pqtype.NullRawMessage `json:"tags"`
	SchemaID    int32                 `json:"schema_id"`
	Values      json.RawMessage       `json:"values"`
	CreatedBy   string                `json:"created_by"`
	UpdatedBy   string                `json:"updated_by"`
}

func (q *Queries) CreateConfig(ctx context.Context, arg CreateConfigParams) (Config, error) {
	row := q.db.QueryRowContext(ctx, createConfig,
		arg.Name,
		arg.Description,
		arg.Active,
		arg.Tags,
		arg.SchemaID,
		arg.Values,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i Config
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.SchemaID,
		&i.Values,
		&i.Tags,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createSchema = `-- name: CreateSchema :one
INSERT INTO schema (name, description, tags, fields, created_by, updated_by, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
RETURNING id, name, description, fields, tags, created_by, updated_by, created_at, updated_at
`

type CreateSchemaParams struct {
	Name        string                `json:"name"`
	Description sql.NullString        `json:"description"`
	Tags        pqtype.NullRawMessage `json:"tags"`
	Fields      json.RawMessage       `json:"fields"`
	CreatedBy   string                `json:"created_by"`
	UpdatedBy   string                `json:"updated_by"`
}

func (q *Queries) CreateSchema(ctx context.Context, arg CreateSchemaParams) (Schema, error) {
	row := q.db.QueryRowContext(ctx, createSchema,
		arg.Name,
		arg.Description,
		arg.Tags,
		arg.Fields,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i Schema
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Fields,
		&i.Tags,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getConfig = `-- name: GetConfig :one
SELECT id, name, description, active, schema_id, values, tags, created_by, updated_by, created_at, updated_at FROM config WHERE id=$1
`

func (q *Queries) GetConfig(ctx context.Context, id int32) (Config, error) {
	row := q.db.QueryRowContext(ctx, getConfig, id)
	var i Config
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.SchemaID,
		&i.Values,
		&i.Tags,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getConfigByNameAndSchema = `-- name: GetConfigByNameAndSchema :one
SELECT id, name, description, active, schema_id, values, tags, created_by, updated_by, created_at, updated_at FROM config
WHERE config.name = $1 AND schema_id = (SELECT id FROM schema WHERE schema.name = $2 LIMIT 1)
`

type GetConfigByNameAndSchemaParams struct {
	ConfigName string `json:"config_name"`
	SchemaName string `json:"schema_name"`
}

func (q *Queries) GetConfigByNameAndSchema(ctx context.Context, arg GetConfigByNameAndSchemaParams) (Config, error) {
	row := q.db.QueryRowContext(ctx, getConfigByNameAndSchema, arg.ConfigName, arg.SchemaName)
	var i Config
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Active,
		&i.SchemaID,
		&i.Values,
		&i.Tags,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
