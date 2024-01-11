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
			email: data.email
		};
	}
});

// UTIL FUNCTIONS

const getUserByEmail = async (email: string) => {
	const user_query_result = await db.select().from(user).where(eq(user.email, email));
	if (user_query_result.length === 0) return null;
	return user_query_result[0];
};

const createOAuthUser = async (userId: string, providerId: string, providerUserId: string) => {};

const userHasLoggedInWithProviderBefore = async (
	luciaUser: typeof user.$inferSelect,
	providerKey: string
) => {
	const key_query_result = await db
		.select()
		.from(key)
		.where(and(eq(key.id, providerKey), eq(key.userId, luciaUser.id)));
	return key_query_result.length > 0;
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
		// createOAuthUser()
		return;
	}
	// Check if the user has an oauth account with the provider.
	if (await userHasLoggedInWithProviderBefore(existingUser, providerKey)) {
		// Return the user.
		return existingUser;
	}
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
	redirectUri: 'http://localhost:5173/login/google/callback'
});

export const facebookAuth = facebook(auth, {
	clientId: FACEBOOK_CLIENT_ID,
	clientSecret: FACEBOOK_CLIENT_SECRET,
	redirectUri: 'http://localhost:5173/login/facebook/callback',
	scope: ['email']
});

export type Auth = typeof auth;
