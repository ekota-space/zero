table "team_members" {
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

  column "team_id" {
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

  index "idx_team_members_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_team_members_team" {
    columns = [column.team_id]
    ref_columns = [table.teams.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_team_members_user" {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  index "idx_unique_team_members" {
    unique = true
    columns = [column.team_id, column.user_id]
  }
}