table "project_managers" {
  schema = schema.public

  column "id" {
    null = false
    type = uuid
    default = sql("gen_random_uuid()")
  }

  column "created_at" {
    null = false
    type = timestamptz
    default = sql("now()")
  }

  column "project_id" {
    null = false
    type = uuid
  }

  column "user_id" {
    null = false
    type = uuid
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_project_managers_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_project_managers_project" {
    columns = [column.project_id]
    ref_columns = [table.projects.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_project_managers_user" {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}