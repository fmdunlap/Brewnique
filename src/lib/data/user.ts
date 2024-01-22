import { user } from '$src/schema';
import { eq } from 'drizzle-orm';
import { db } from './db';

export async function getUserProfile(username: string) {
	const userQueryResult = await db.select().from(user).where(eq(user.username, username));

	if (userQueryResult.length < 1) {
		return null;
	}

	return userQueryResult[0];
}
