table "teams" {
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
    type = character_varying(50)
  }

  column "description" {
    null = true
    type = text
  }

  column "slug" {
    null = false
    type = character_varying(16)
  }

  column "organization_id" {
    null = false
    type = uuid
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_teams_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_teams_organization" {
    columns = [column.organization_id]
    ref_columns = [table.organizations.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  index "idx_unique_teams" {
    unique = true
    columns = [column.organization_id, column.slug]
  }
}