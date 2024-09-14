-- Create "organization_admins" table
CREATE TABLE "organization_admins" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "organization_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organization_admins_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_organization_admins_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_organization_admins_id" to table: "organization_admins"
CREATE UNIQUE INDEX "idx_organization_admins_id" ON "organization_admins" ("id");
-- Create index "idx_unique_organization_admins" to table: "organization_admins"
CREATE UNIQUE INDEX "idx_unique_organization_admins" ON "organization_admins" ("organization_id", "user_id");
-- Create "projects" table
CREATE TABLE "projects" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  "slug" character varying(16) NOT NULL,
  "organization_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_projects_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_projects_id" to table: "projects"
CREATE UNIQUE INDEX "idx_projects_id" ON "projects" ("id");
-- Create index "idx_unique_projects" to table: "projects"
CREATE UNIQUE INDEX "idx_unique_projects" ON "projects" ("organization_id", "slug");
-- Create "project_managers" table
CREATE TABLE "project_managers" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "project_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_project_managers_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_project_managers_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_project_managers_id" to table: "project_managers"
CREATE UNIQUE INDEX "idx_project_managers_id" ON "project_managers" ("id");
-- Create "teams" table
CREATE TABLE "teams" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  "slug" character varying(16) NOT NULL,
  "organization_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_teams_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_teams_id" to table: "teams"
CREATE UNIQUE INDEX "idx_teams_id" ON "teams" ("id");
-- Create index "idx_unique_teams" to table: "teams"
CREATE UNIQUE INDEX "idx_unique_teams" ON "teams" ("organization_id", "slug");
-- Create "project_teams" table
CREATE TABLE "project_teams" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "project_id" uuid NOT NULL,
  "team_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_project_teams_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_project_teams_team" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_project_teams_id" to table: "project_teams"
CREATE UNIQUE INDEX "idx_project_teams_id" ON "project_teams" ("id");
-- Create index "idx_unique_project_teams" to table: "project_teams"
CREATE UNIQUE INDEX "idx_unique_project_teams" ON "project_teams" ("project_id", "team_id");
-- Create "team_members" table
CREATE TABLE "team_members" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "team_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_team_members_team" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_team_members_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_team_members_id" to table: "team_members"
CREATE UNIQUE INDEX "idx_team_members_id" ON "team_members" ("id");
-- Create index "idx_unique_team_members" to table: "team_members"
CREATE UNIQUE INDEX "idx_unique_team_members" ON "team_members" ("team_id", "user_id");
-- Create "team_leaders" table
CREATE TABLE "team_leaders" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "team_id" uuid NOT NULL,
  "team_member_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_team_leaders_team" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_team_leaders_team_member" FOREIGN KEY ("team_member_id") REFERENCES "team_members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_team_leaders_id" to table: "team_leaders"
CREATE UNIQUE INDEX "idx_team_leaders_id" ON "team_leaders" ("id");
-- Create index "idx_unique_team_leaders" to table: "team_leaders"
CREATE UNIQUE INDEX "idx_unique_team_leaders" ON "team_leaders" ("team_id", "team_member_id");
