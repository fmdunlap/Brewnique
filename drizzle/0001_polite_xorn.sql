ALTER TYPE "unit_of_measurement" ADD VALUE 'barrel';--> statement-breakpoint
ALTER TABLE "recipe" ALTER COLUMN "name" DROP NOT NULL;--> statement-breakpoint
ALTER TABLE "recipe" ALTER COLUMN "difficulty" DROP NOT NULL;--> statement-breakpoint
ALTER TABLE "recipe" ALTER COLUMN "rating" DROP NOT NULL;