/* eslint-disable @typescript-eslint/ban-types */

declare global {
	namespace Lucia {
		type Auth = import('$lib/server/lucia').Auth;
		type DatabaseUserAttributes = { username: string | null; email: string };
		type DatabaseSessionAttributes = {};
	}

	namespace App {
		interface Locals {
			auth: import('lucia').AuthRequest;
		}
		interface PageState {
			loginOpen: boolean;
		}
	}
}

// THIS IS IMPORTANT!!!
export {};
