import { db } from '$lib/data/db';
import { user } from '$src/schema';
import { json, type RequestHandler } from '@sveltejs/kit';
import { eq } from 'drizzle-orm';

export const GET: RequestHandler = async ({ url }) => {
	const user_id = url.searchParams.get('id');
	const username = url.searchParams.get('name');

	if (!user_id && !username) {
		return json({ message: 'Missing user id and user name. Must supply one.' }, { status: 400 });
	}

	if (user_id) {
		const resultUser = await db.select().from(user).where(eq(user.id, user_id));

		if (!resultUser) {
			return json({ message: `User with id ${user_id} not found` }, { status: 404 });
		}

		return json(resultUser);
	}

	if (!username) {
		return json({ message: 'Missing user name' }, { status: 400 });
	}

	const resultUser = await db.select().from(user).where(eq(user.username, username));

	if (!resultUser) {
		return json({ message: `User with name ${username} not found` }, { status: 404 });
	}

	return json(resultUser);
};
