table "projects" {
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

  column "updated_at" {
    null = false
    type = timestamptz
    default = sql("now()")
  }

  column "name" {
    null = false
    type = varchar(50)
  }

  column "description" {
    null = true
    type = text
  }

  column "slug" {
    null = false
    type = varchar(16)
  }

  column "organization_id" {
    null = false
    type = uuid
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_projects_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_projects_organization" {
    columns = [column.organization_id]
    ref_columns = [table.organizations.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  index "idx_unique_projects" {
    unique = true
    columns = [column.organization_id, column.slug]
  }
}