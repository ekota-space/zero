-- Create "organization_members" table
CREATE TABLE "organization_members" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "organization_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organization_members_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_organization_members_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_organization_members_id" to table: "organization_members"
CREATE UNIQUE INDEX "idx_organization_members_id" ON "organization_members" ("id");
-- Create index "idx_unique_organization_members" to table: "organization_members"
CREATE UNIQUE INDEX "idx_unique_organization_members" ON "organization_members" ("organization_id", "user_id");
