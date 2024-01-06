// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces

import type { SupabaseClient } from '@supabase/supabase-js';
import type { Session } from '@supabase/supabase-js';
import type { Database } from './types/supabaseDB';

type TypedSupabase = SupabaseClient<Database>;

declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			supabase: TypedSupabase;
			getSession: () => Promise<Session | null>;
		}
		interface PageData {
			session: Session | null;
		}
		interface PageState {
			loginOpen: boolean;
		}
		// interface Platform {}
	}
}

export {};
