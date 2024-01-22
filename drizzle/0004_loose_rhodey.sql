ALTER TYPE "unit_of_measurement" ADD VALUE 'liter';--> statement-breakpoint
ALTER TABLE "user_key" DROP CONSTRAINT "user_key_user_id_auth_user_id_fk";
--> statement-breakpoint
ALTER TABLE "recipe" DROP CONSTRAINT "recipe_owner_id_auth_user_id_fk";
--> statement-breakpoint
ALTER TABLE "recipe_ingredient" DROP CONSTRAINT "recipe_ingredient_recipe_id_recipe_id_fk";
--> statement-breakpoint
ALTER TABLE "user_session" DROP CONSTRAINT "user_session_user_id_auth_user_id_fk";
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "user_key" ADD CONSTRAINT "user_key_user_id_auth_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "auth_user"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe" ADD CONSTRAINT "recipe_owner_id_auth_user_id_fk" FOREIGN KEY ("owner_id") REFERENCES "auth_user"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe_ingredient" ADD CONSTRAINT "recipe_ingredient_recipe_id_recipe_id_fk" FOREIGN KEY ("recipe_id") REFERENCES "recipe"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "user_session" ADD CONSTRAINT "user_session_user_id_auth_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "auth_user"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
