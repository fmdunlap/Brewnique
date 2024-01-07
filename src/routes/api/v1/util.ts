export function parseParamAsNumber(param: string | null, def: number): number {
	if (!param) {
		return def;
	}

	const parsed = Number.parseInt(param);

	if (Number.isNaN(parsed)) {
		return def;
	}

	return parsed;
}
