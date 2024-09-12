package main

import (
	"fmt"
	"strconv"

	"github.com/ekota-space/zero/pkgs/common"
	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

func main() {
	common.SetupEnvironmentVars()
	port, err := strconv.Atoi(common.Env.PostgresPort)

	if err != nil {
		panic(err)
	}

	err = postgres.Generate(
		"./pkgs/root/db",
		postgres.DBConnection{
			Host:       common.Env.PostgresHost,
			Port:       port,
			User:       common.Env.PostgresUser,
			Password:   common.Env.PostgresPassword,
			DBName:     common.Env.PostgresDB,
			SslMode:    "disable",
			SchemaName: "public",
		},
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)

									omitEmpty := ""
									if columnMetaData.IsNullable {
										omitEmpty = ",omitempty"
									}

									return defaultTableModelField.UseTags(
										fmt.Sprintf(`json:"%s%s"`, columnMetaData.Name, omitEmpty),
									)
								})
						}),
					// .UseView(func(table metadata.Table) template.ViewModel {
					// 	return template.DefaultViewModel(table).
					// 		UseField(func(columnMetaData metadata.Column) template.TableModelField {
					// 			defaultTableModelField := template.DefaultTableModelField(columnMetaData)
					// 			if table.Name == "actor_info" && columnMetaData.Name == "actor_id" {
					// 				return defaultTableModelField.UseTags(`sql:"primary_key"`)
					// 			}
					// 			return defaultTableModelField
					// 		})
					// }),
					)
			}),
	)

	if err != nil {
		panic(err)
	}
}
