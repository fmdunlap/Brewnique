import { db } from './db';
import { user } from '$src/schema';
import { eq } from 'drizzle-orm';
import { S3 } from './s3';
import { PutObjectCommand } from '@aws-sdk/client-s3';

export function userAvatarUrl(userId: string) {
	return 'https://cdn.brewnique.io/avatars/' + userId + '.svg';
}

export async function userHasAvatar(userId: string) {
	const avatarSelect = await db
		.select({ avatarUrl: user.avatarUrl })
		.from(user)
		.where(eq(user.id, userId));
	return avatarSelect.length > 0;
}

export async function downloadDicebearAvatar(userId: string) {
	const url = 'https://api.dicebear.com/7.x/thumbs/svg?seed=' + userId + '.svg';
	const response = await fetch(url);
	const svg = await response.blob();
	return svg;
}

export async function uploadAvatarToStorage(userId: string, avatarBlob: Blob) {
	S3.send(
		new PutObjectCommand({
			Bucket: 'brewnique',
			Key: 'avatars/' + userId + '.svg',
			Body: Buffer.from(await avatarBlob.arrayBuffer()),
			ContentType: 'image/svg+xml',
			ContentLength: avatarBlob.size
		})
	);
}

export async function addDefaultAvatarToStorage(user_id: string) {
	const avatar_blob = await downloadDicebearAvatar(user_id);
	await uploadAvatarToStorage(user_id, avatar_blob);
}
