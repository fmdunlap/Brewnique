import { createClient } from '@supabase/supabase-js';
import { env } from '$env/dynamic/private';
import { type Database } from './types/supabaseDB';

if (!env.SUPABASE_URL || !env.SUPABASE_ANON_KEY) {
	throw new Error('Missing env variables for Supabase');
}

export const supabase = createClient<Database>(env.SUPABASE_URL, env.SUPABASE_ANON_KEY);
