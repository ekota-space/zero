schema "public" {}

table "users" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "first_name" {
    null = false
    type = character_varying(255)
  }
  column "last_name" {
    null = false
    type = character_varying(255)
  }
  column "username" {
    null = false
    type = character_varying(16)
  }
  column "email" {
    null = false
    type = text
  }
  column "password" {
    null = true
    type = text
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
  column "verified_at" {
    null = true
    type = timestamptz
  }
  column "password_reset_at" {
    null = true
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_users_email" {
    unique  = true
    columns = [column.email]
  }
  index "idx_users_id" {
    unique  = true
    columns = [column.id]
  }
  index "idx_users_username" {
    unique  = true
    columns = [column.username]
  }
}