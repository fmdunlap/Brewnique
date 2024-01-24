import { redirect } from '@sveltejs/kit';

export const load = async ({ url, parent }) => {
	const { session } = await parent();

	if (!session) {
		throw redirect(302, '/login');
	}

	if (session.user.onboardingStatus == 'COMPLETE') {
		throw redirect(302, '/');
	}

	if (
		session.user.onboardingStatus == 'PENDING_USERNAME' &&
		url.pathname !== '/onboarding/username'
	) {
		throw redirect(302, '/onboarding/username');
	}

	if (session.user.onboardingStatus == 'PENDING_BIO' && url.pathname !== '/onboarding/bio') {
		throw redirect(302, '/onboarding/bio');
	}

	if (
		session.user.onboardingStatus === 'PENDING_EMAIL_VERIFICATION' &&
		url.pathname !== '/onboarding/email'
	) {
		throw redirect(302, '/onboarding/email');
	}

	return {
		session
	};
};
