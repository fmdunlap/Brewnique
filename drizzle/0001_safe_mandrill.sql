ALTER TABLE "auth_user" ALTER COLUMN "username" DROP NOT NULL;--> statement-breakpoint
ALTER TABLE "auth_user" ADD COLUMN "email" text NOT NULL;--> statement-breakpoint
ALTER TABLE "auth_user" ADD CONSTRAINT "auth_user_email_unique" UNIQUE("email");