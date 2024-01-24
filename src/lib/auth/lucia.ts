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
import { emailVerification, key, user } from '$src/schema';
import { and, eq } from 'drizzle-orm';
import { generateRandomString, isWithinExpiration } from 'lucia/utils';
import { addDefaultAvatarToStorage, userAvatarUrl } from '$lib/data/avatar';

const EMAIL_EXPIRE_TIME = 1000 * 60 * 60 * 2; // 2 hours

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

// EMAIL VERIFICATION

export const generateEmailVerificationToken = async (userId: string) => {
	const storedUserTokens = await db
		.select()
		.from(emailVerification)
		.where(eq(emailVerification.userId, userId));
	if (storedUserTokens.length > 0) {
		const reuseableToken = storedUserTokens.find((token) => {
			return isWithinExpiration(Number(token.expires) - EMAIL_EXPIRE_TIME / 2);
		});
		if (reuseableToken) return reuseableToken.id;
	}
	const newToken = generateRandomString(63);
	await db.insert(emailVerification).values({
		id: newToken,
		userId: userId,
		expires: Date.now() + EMAIL_EXPIRE_TIME
	});

	return newToken;
};

export const validateEmailVerificationToken = async (token: string) => {
	const storedToken = await db.transaction(async (tx) => {
		const storedToken = await tx
			.select()
			.from(emailVerification)
			.where(eq(emailVerification.id, token));
		if (!storedToken || storedToken.length === 0) throw new Error('Invalid Token');
		await tx.delete(emailVerification).where(eq(emailVerification.id, token));
		return storedToken[0];
	});
	const tokenExpires = Number(storedToken.expires);
	if (!isWithinExpiration(tokenExpires)) throw new Error('Token Expired');
	return storedToken.userId;
};

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

const createEmailUser = async (email: string) => {
	const newUserId = generateRandomString(15);
	const newUser: typeof user.$inferInsert = {
		id: generateRandomString(15),
		email: email,
		username: null,
		avatarUrl: userAvatarUrl(newUserId),
		bio: null,
		onboardingStatus: 'PENDING_EMAIL_VERIFICATION'
	};
	await db.insert(user).values(newUser);
	await addDefaultAvatarToStorage(newUserId);
	return newUser;
};

const createEmailKey = async (userId: string, email: string, password: string) => {
	await auth.createKey({
		userId,
		providerId: 'email',
		providerUserId: email.toLowerCase(),
		password: password
	});
};

const userHasProviderKey = async (luciaUser: typeof user.$inferSelect, providerKey: string) => {
	const keyQueryResult = await db
		.select()
		.from(key)
		.where(and(eq(key.id, providerKey), eq(key.userId, luciaUser.id)));
	return keyQueryResult.length > 0;
};

export const getOrCreateEmailUser = async (email: string, password: string) => {
	const existingUser = await getUserByEmail(email);
	if (existingUser) {
		return existingUser;
	}
	// Create & return a new email user.
	const newUser = await createEmailUser(email);
	await createEmailKey(newUser.id, email, password);
	return newUser;
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

export const githubAuth = (urlOrigin: string) => {
	return github(auth, {
		clientId: GITHUB_CLIENT_ID,
		clientSecret: GITHUB_CLIENT_SECRET,
		redirectUri: urlOrigin + '/login/github/callback'
	});
};

export const googleAuth = (urlOrigin: string) => {
	return google(auth, {
		clientId: GOOGLE_CLIENT_ID,
		clientSecret: GOOGLE_CLIENT_SECRET,
		redirectUri: urlOrigin + '/login/google/callback',
		scope: ['email']
	});
};

export const facebookAuth = (urlOrigin: string) => {
	return facebook(auth, {
		clientId: FACEBOOK_CLIENT_ID,
		clientSecret: FACEBOOK_CLIENT_SECRET,
		redirectUri: urlOrigin + '/login/facebook/callback',
		scope: ['email']
	});
};

export type Auth = typeof auth;
