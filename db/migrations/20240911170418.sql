-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "first_name" character varying(255) NOT NULL,
  "last_name" character varying(255) NOT NULL,
  "username" character varying(16) NOT NULL,
  "email" text NOT NULL,
  "password" text NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "verified_at" timestamptz NULL,
  "password_reset_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
-- Create index "idx_users_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_id" ON "users" ("id");
-- Create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "users" ("username");
-- Create "organizations" table
CREATE TABLE "organizations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" character varying(255) NULL,
  "owner_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organizations_owner" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_organizations_id" to table: "organizations"
CREATE UNIQUE INDEX "idx_organizations_id" ON "organizations" ("id");
