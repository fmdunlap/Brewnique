import type { Session, SupabaseClient } from '@supabase/supabase-js';
import type { Database } from '$lib/types/supabaseDB';

export async function getUserProfile(session: Session, supabase: SupabaseClient<Database>) {
	return (
		await supabase
			.from('profile')
			.select('*')
			.eq('id', session?.user.id)
			.single()
	).data;
}