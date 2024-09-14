table "project_teams" {
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

  column "team_id" {
    null = false
    type = uuid
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_project_teams_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_project_teams_project" {
    columns = [column.project_id]
    ref_columns = [table.projects.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_project_teams_team" {
    columns = [column.team_id]
    ref_columns = [table.teams.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  index "idx_unique_project_teams" {
    unique = true
    columns = [column.project_id, column.team_id]
  }
}