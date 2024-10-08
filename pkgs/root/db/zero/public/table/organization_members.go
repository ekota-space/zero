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

var OrganizationMembers = newOrganizationMembersTable("public", "organization_members", "")

type organizationMembersTable struct {
	postgres.Table

	// Columns
	ID             postgres.ColumnString
	CreatedAt      postgres.ColumnTimestampz
	OrganizationID postgres.ColumnString
	UserID         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrganizationMembersTable struct {
	organizationMembersTable

	EXCLUDED organizationMembersTable
}

// AS creates new OrganizationMembersTable with assigned alias
func (a OrganizationMembersTable) AS(alias string) *OrganizationMembersTable {
	return newOrganizationMembersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrganizationMembersTable with assigned schema name
func (a OrganizationMembersTable) FromSchema(schemaName string) *OrganizationMembersTable {
	return newOrganizationMembersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrganizationMembersTable with assigned table prefix
func (a OrganizationMembersTable) WithPrefix(prefix string) *OrganizationMembersTable {
	return newOrganizationMembersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrganizationMembersTable with assigned table suffix
func (a OrganizationMembersTable) WithSuffix(suffix string) *OrganizationMembersTable {
	return newOrganizationMembersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrganizationMembersTable(schemaName, tableName, alias string) *OrganizationMembersTable {
	return &OrganizationMembersTable{
		organizationMembersTable: newOrganizationMembersTableImpl(schemaName, tableName, alias),
		EXCLUDED:                 newOrganizationMembersTableImpl("", "excluded", ""),
	}
}

func newOrganizationMembersTableImpl(schemaName, tableName, alias string) organizationMembersTable {
	var (
		IDColumn             = postgres.StringColumn("id")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		OrganizationIDColumn = postgres.StringColumn("organization_id")
		UserIDColumn         = postgres.StringColumn("user_id")
		allColumns           = postgres.ColumnList{IDColumn, CreatedAtColumn, OrganizationIDColumn, UserIDColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, OrganizationIDColumn, UserIDColumn}
	)

	return organizationMembersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		OrganizationID: OrganizationIDColumn,
		UserID:         UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
