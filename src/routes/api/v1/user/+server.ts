import { json, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ url, locals: { supabase } }) => {
	const user_id = url.searchParams.get('id');
	const user_name = url.searchParams.get('name');

	if (!user_id && !user_name) {
		return json({ message: 'Missing user id and user name. Must supply one.' }, { status: 400 });
	}

	if (user_id) {
		const { data: profile, error } = await supabase
			.from('profile')
			.select('*')
			.eq('id', user_id)
			.single();

		if (error) {
			return json({ message: error.message }, { status: 500 });
		}

		if (!profile) {
			return json({ message: `User with id ${user_id} not found` }, { status: 404 });
		}

		return json(profile);
	}

	const { data: profile, error } = await supabase
		.from('profile')
		.select('*')
		.eq('display_name', user_name)
		.single();

	if (error) {
		return json({ message: error.message }, { status: 500 });
	}

	if (!profile) {
		return json({ message: `User with name ${user_name} not found` }, { status: 404 });
	}

	return json(profile);
};
