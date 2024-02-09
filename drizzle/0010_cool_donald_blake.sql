CREATE TABLE IF NOT EXISTS "recipe_save" (
	"id" uuid PRIMARY KEY NOT NULL,
	"user_id" varchar(15) NOT NULL,
	"recipe_id" uuid NOT NULL
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe_save" ADD CONSTRAINT "recipe_save_user_id_auth_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "auth_user"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe_save" ADD CONSTRAINT "recipe_save_recipe_id_recipe_id_fk" FOREIGN KEY ("recipe_id") REFERENCES "recipe"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
