-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now();
-- Modify "organizations" table
ALTER TABLE "organizations" DROP CONSTRAINT "fk_organizations_owner", ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now(), ALTER COLUMN "owner_id" SET NOT NULL, ADD
 CONSTRAINT "fk_organizations_owner" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
