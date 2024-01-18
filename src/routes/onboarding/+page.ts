export async function load({ parent, data }) {
	const { session } = await parent();

	return {
		...data,
		session
	};
}
