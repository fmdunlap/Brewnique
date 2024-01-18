ALTER TABLE "recipe_ingredient" ALTER COLUMN "name" DROP NOT NULL;--> statement-breakpoint
ALTER TABLE "recipe_ingredient" ALTER COLUMN "quantity" DROP NOT NULL;--> statement-breakpoint
ALTER TABLE "recipe_ingredient" ALTER COLUMN "unit" DROP NOT NULL;--> statement-breakpoint
ALTER TABLE "recipe_ingredient" DROP COLUMN IF EXISTS "type";