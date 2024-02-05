CREATE TABLE IF NOT EXISTS "recipe_comment" (
	"id" uuid PRIMARY KEY NOT NULL,
	"recipe_id" uuid NOT NULL,
	"user_id" varchar(15) NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	"content" text
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe_comment" ADD CONSTRAINT "recipe_comment_recipe_id_recipe_id_fk" FOREIGN KEY ("recipe_id") REFERENCES "recipe"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "recipe_comment" ADD CONSTRAINT "recipe_comment_user_id_auth_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "auth_user"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
