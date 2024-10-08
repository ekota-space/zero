//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Organizations = newOrganizationsTable("public", "organizations", "")

type organizationsTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz
	UpdatedAt   postgres.ColumnTimestampz
	Name        postgres.ColumnString
	Description postgres.ColumnString
	OwnerID     postgres.ColumnString
	Slug        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrganizationsTable struct {
	organizationsTable

	EXCLUDED organizationsTable
}

// AS creates new OrganizationsTable with assigned alias
func (a OrganizationsTable) AS(alias string) *OrganizationsTable {
	return newOrganizationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrganizationsTable with assigned schema name
func (a OrganizationsTable) FromSchema(schemaName string) *OrganizationsTable {
	return newOrganizationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrganizationsTable with assigned table prefix
func (a OrganizationsTable) WithPrefix(prefix string) *OrganizationsTable {
	return newOrganizationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrganizationsTable with assigned table suffix
func (a OrganizationsTable) WithSuffix(suffix string) *OrganizationsTable {
	return newOrganizationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrganizationsTable(schemaName, tableName, alias string) *OrganizationsTable {
	return &OrganizationsTable{
		organizationsTable: newOrganizationsTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newOrganizationsTableImpl("", "excluded", ""),
	}
}

func newOrganizationsTableImpl(schemaName, tableName, alias string) organizationsTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		NameColumn        = postgres.StringColumn("name")
		DescriptionColumn = postgres.StringColumn("description")
		OwnerIDColumn     = postgres.StringColumn("owner_id")
		SlugColumn        = postgres.StringColumn("slug")
		allColumns        = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, NameColumn, DescriptionColumn, OwnerIDColumn, SlugColumn}
		mutableColumns    = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, NameColumn, DescriptionColumn, OwnerIDColumn, SlugColumn}
	)

	return organizationsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,
		OwnerID:     OwnerIDColumn,
		Slug:        SlugColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
