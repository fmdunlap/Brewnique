export function getUserProfile() {
	return {
		avatar_url: 'https://authjs.dev/img/logo/logo-xs.webp',
		bio: null,
		created_at: null,
		display_name: 'Guest',
		id: 'guest',
		onboarding_state: 'completed'
	};
	// return (
	// 	await supabase
	// 		.from('profile')
	// 		.select('*')
	// 		.eq('id', session?.user.id)
	// 		.single()
	// ).data;
}
