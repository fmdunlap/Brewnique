import { db } from '$lib/data/db';
import { user } from '$src/schema';
import { RegExpMatcher, englishDataset, englishRecommendedTransformers } from 'obscenity';
import { eq } from 'drizzle-orm';

export const obscenityMatcher = new RegExpMatcher({
	...englishDataset.build(),
	...englishRecommendedTransformers
});

export async function advanceOnboardingState(userId: string) {
	const currentState = (await db.select().from(user).where(eq(user.id, userId)))[0]
		.onboardingStatus;

	// The state machine is just a linear progression.
	//
	// email_verification_pending -> display_name_pending -> bio_pending -> avatar_pending -> completed
	let nextState: typeof currentState | null = null;
	switch (currentState) {
		case 'PENDING_USERNAME':
			nextState = 'PENDING_BIO';
			break;
		// For now we skip the avatar upload
		case 'PENDING_BIO':
		case 'PENDING_AVATAR':
			nextState = 'COMPLETE';
			break;
		case 'COMPLETE':
		default:
			nextState = null;
	}
	if (!nextState) {
		return;
	}

	return await db.update(user).set({ onboardingStatus: nextState }).where(eq(user.id, userId));
}
