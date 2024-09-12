-- Modify "organizations" table
ALTER TABLE "organizations" ALTER COLUMN "description" TYPE text, ADD COLUMN "slug" character varying(16) NOT NULL;
-- Create index "idx_organizations_slug" to table: "organizations"
CREATE UNIQUE INDEX "idx_organizations_slug" ON "organizations" ("slug");
