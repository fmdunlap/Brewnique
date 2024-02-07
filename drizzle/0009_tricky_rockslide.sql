ALTER TABLE "recipe_comment" ADD COLUMN "parent_id" uuid;--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe_comment" ADD CONSTRAINT "recipe_comment_parent_id_recipe_comment_id_fk" FOREIGN KEY ("parent_id") REFERENCES "recipe_comment"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
