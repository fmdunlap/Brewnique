ALTER TYPE "onboarding_status" ADD VALUE 'PENDING_EMAIL_VERIFICATION';--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "email_verification_token" (
	"id" varchar(255) PRIMARY KEY NOT NULL,
	"expires" bigint NOT NULL,
	"user_id" varchar(15) NOT NULL
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "email_verification_token" ADD CONSTRAINT "email_verification_token_user_id_auth_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "auth_user"("id") ON DELETE cascade ON UPDATE cascade;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
