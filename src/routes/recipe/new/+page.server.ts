export const load = async ({ parent }) => {
	return {
		...(await parent()),
		page: 'new'
	};
};
