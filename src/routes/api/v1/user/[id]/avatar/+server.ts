import { uploadAvatarToStorage, userAvatarUrl } from '$lib/data/avatar';
import { db } from '$lib/data/db';
import { convertBase64ToFile } from '$lib/data/util';
import { user } from '$src/schema';
import { type RequestHandler } from '@sveltejs/kit';
import { eq } from 'drizzle-orm';

export const POST: RequestHandler = async ({ request, params }) => {
	const userId = params.id;

	if (!userId) {
		return new Response('User ID not provided', { status: 400 });
	}

	const base64Image = (await request.json()).b64img;
	const filetype = base64Image.split(';')[0].split('/')[1];

	if (!filetype || !['png', 'jpeg', 'jpg', 'svg'].includes(filetype)) {
		return new Response('Invalid image', { status: 400 });
	}

	const imageFile = await convertBase64ToFile(base64Image);
	await uploadAvatarToStorage(userId, filetype, imageFile);
	await db
		.update(user)
		.set({ avatarUrl: userAvatarUrl(userId) })
		.where(eq(user.id, userId));

	return new Response('Avatar uploaded', { status: 200 });
};
