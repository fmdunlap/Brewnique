import type { Database } from '$lib/types/supabaseDB.js';
import type { SupabaseClient } from '@supabase/supabase-js';

export async function userHasAvatar(user_id: string, supabase: SupabaseClient<Database>) {
	const avatar_url = (
		await supabase.from('profile').select('avatar_url').eq('id', user_id).single()
	).data?.avatar_url;

	return avatar_url != null;
}

async function downloadDicebearAvatar(user_id: string) {
	const url = 'https://api.dicebear.com/7.x/thumbs/svg?seed=' + user_id + '.svg';
	const response = await fetch(url);
	const svg = await response.blob();
	return svg;
}

async function uploadAvatarToStorage(
	user_id: string,
	avatar_blob: Blob,
	supabase: SupabaseClient<Database>
) {
	await supabase.storage.from('avatars').upload(user_id + '.svg', avatar_blob);
}

export async function addDefaultAvatarToStorage(
	user_id: string,
	supabase: SupabaseClient<Database>
) {
	const avatar_blob = await downloadDicebearAvatar(user_id);
	uploadAvatarToStorage(user_id, avatar_blob, supabase);
	await supabase
		.from('profile')
		.update({
			avatar_url: await supabase.storage.from('avatars').getPublicUrl(user_id + '.svg').data
				.publicUrl
		})
		.eq('id', user_id);
}
