table "team_leaders" {
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

  column "team_member_id" {
    null = false
    type = uuid
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_team_leaders_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_team_leaders_team" {
    columns = [column.team_id]
    ref_columns = [table.teams.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_team_leaders_team_member" {
    columns = [column.team_member_id]
    ref_columns = [table.team_members.column.id]
    on_update = NO_ACTION
    on_delete = NO_ACTION
  }

  index "idx_unique_team_leaders" {
    unique = true
    columns = [column.team_id, column.team_member_id]
  }
}