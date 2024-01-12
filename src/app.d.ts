/* eslint-disable @typescript-eslint/ban-types */
import { AuthRequest } from 'lucia';
import { onboardingStatus } from './schema';

declare global {
	namespace Lucia {
		type Auth = import('$lib/auth/lucia').Auth;
		// AIUI, these must match the column names in the db. Not the schema property names.
		type DatabaseUserAttributes = {
			username: string | null;
			email: string;
			avatar_url: string;
			bio: string | null;
			onboarding_status: onboardingStatus;
		};
		type DatabaseSessionAttributes = {};
	}

	namespace App {
		interface Locals {
			auth: AuthRequest;
		}
		interface PageState {
			loginOpen: boolean;
		}
	}
}

export {};
