import { createKeyId, lucia } from 'lucia';
import { sveltekit } from 'lucia/middleware';
import { dev } from '$app/environment';
import { postgres as postgresAdapter } from '@lucia-auth/adapter-postgresql';

import { facebook, github, google } from '@lucia-auth/oauth/providers';
import { db, queryClient } from '../data/db';
import {
	FACEBOOK_CLIENT_ID,
	FACEBOOK_CLIENT_SECRET,
	GITHUB_CLIENT_ID,
	GITHUB_CLIENT_SECRET,
	GOOGLE_CLIENT_ID,
	GOOGLE_CLIENT_SECRET
} from '$env/static/private';
import { key, user } from '$src/schema';
import { and, eq } from 'drizzle-orm';
import { generateRandomString } from 'lucia/utils';
import { addDefaultAvatarToStorage, userAvatarUrl } from '$lib/data/avatar';

// AUTH

export const auth = lucia({
	env: dev ? 'DEV' : 'PROD',
	adapter: postgresAdapter(queryClient, {
		user: 'auth_user',
		session: 'user_session',
		key: 'user_key'
	}),
	middleware: sveltekit(),
	getUserAttributes: (data) => {
		return {
			username: data.username,
			email: data.email,
			avatarUrl: data.avatar_url,
			bio: data.bio,
			onboardingStatus: data.onboarding_status
		};
	}
});

// UTIL FUNCTIONS

const getUserByEmail = async (email: string) => {
	const user_query_result = await db.select().from(user).where(eq(user.email, email));
	if (user_query_result.length === 0) return null;
	return user_query_result[0];
};

const createOAuthKey = async (userId: string, providerKey: string) => {
	await db.insert(key).values({
		id: providerKey,
		userId: userId
	});
};

const createOAuthUser = async (email: string) => {
	const newUserId = generateRandomString(15);
	await db.insert(user).values({
		id: newUserId,
		email: email,
		username: null,
		avatarUrl: userAvatarUrl(newUserId),
		bio: null,
		onboardingStatus: 'PENDING_USERNAME'
	});
	await addDefaultAvatarToStorage(newUserId);
	return newUserId;
};

const userHasProviderKey = async (luciaUser: typeof user.$inferSelect, providerKey: string) => {
	const keyQueryResult = await db
		.select()
		.from(key)
		.where(and(eq(key.id, providerKey), eq(key.userId, luciaUser.id)));
	return keyQueryResult.length > 0;
};

export const getOrCreateOAuthUser = async (
	email: string,
	providerId: string,
	providerUserId: string
) => {
	const existingUser = await getUserByEmail(email);
	const providerKey = createKeyId(providerId, providerUserId);
	if (!existingUser) {
		// Create & return a new oauth user.
		const newUserId = await createOAuthUser(email);
		await createOAuthKey(newUserId, providerKey);
		return {
			id: newUserId,
			email: email,
			username: null
		};
	}
	// Check if the user has an oauth account with the provider.
	if (!(await userHasProviderKey(existingUser, providerKey))) {
		// Create a new oauth key for the user for the given provider.
		createOAuthKey(existingUser.id, providerKey);
	}
	// Return the user.
	return existingUser;
};

// PROVIDERS

export const githubAuth = github(auth, {
	clientId: GITHUB_CLIENT_ID,
	clientSecret: GITHUB_CLIENT_SECRET,
	redirectUri: 'http://localhost:5173/login/github/callback'
});

export const googleAuth = google(auth, {
	clientId: GOOGLE_CLIENT_ID,
	clientSecret: GOOGLE_CLIENT_SECRET,
	redirectUri: 'http://localhost:5173/login/google/callback',
	scope: ['email']
});

export const facebookAuth = facebook(auth, {
	clientId: FACEBOOK_CLIENT_ID,
	clientSecret: FACEBOOK_CLIENT_SECRET,
	redirectUri: 'http://localhost:5173/login/facebook/callback',
	scope: ['email']
});

export type Auth = typeof auth;
