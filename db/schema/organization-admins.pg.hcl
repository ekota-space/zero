table "organization_admins" {
  schema = schema.public

  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")   
  }

  column "created_at" {
    null = false
    type = timestamptz
    default = sql("now()")
    
  }

  column "organization_id" {
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
  
  index "idx_organization_admins_id" {
    unique = true
    columns = [column.id]
  }

  foreign_key "fk_organization_admins_organization" {
    columns = [column.organization_id]
    ref_columns = [table.organizations.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_organization_admins_user" {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  index "idx_unique_organization_admins" {
    unique = true
    columns = [column.organization_id, column.user_id]
  }
}