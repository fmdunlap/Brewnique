import { user } from '$src/schema';
import { eq } from 'drizzle-orm';
import { db } from './db';

export async function getUserProfileByUsername(username: string) {
	const userQueryResult = await db.select().from(user).where(eq(user.username, username));

	if (userQueryResult.length < 1) {
		return null;
	}

	return userQueryResult[0];
}

export async function getUserProfileById(userId: string) {
	const userQueryResult = await db.select().from(user).where(eq(user.id, userId));

	if (userQueryResult.length < 1) {
		return null;
	}

	return userQueryResult[0];
}
