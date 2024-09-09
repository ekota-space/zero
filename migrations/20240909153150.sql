-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "first_name" character varying(255) NULL,
  "last_name" character varying(255) NULL,
  "username" character varying(16) NULL,
  "email" text NULL,
  "password" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "verified_at" timestamptz NULL,
  "password_reset_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_id" ON "public"."users" ("id");
-- Create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "public"."users" ("username");
