import { downloadDicebearAvatar, uploadAvatarToStorage } from '$lib/data/avatar';
import { json } from '@sveltejs/kit';

export const GET = async () => {
	const imgBlob = await downloadDicebearAvatar('test');
	uploadAvatarToStorage('test', imgBlob);
	return json({ message: 'Hello world!' });
};
