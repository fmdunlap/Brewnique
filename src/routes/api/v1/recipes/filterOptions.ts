export type FilterOptions = {
	maxAbv: number;
	minAbv: number;
	minBatchSize: number;
	maxBatchSize: number;
	rating: number[];
};

export const DEFAULT_FILTER_OPTIONS: FilterOptions = {
	maxAbv: 100,
	minAbv: 0,
	minBatchSize: 0,
	maxBatchSize: 100,
	rating: []
};

export type SortByValue = 'NameAsc' | 'NameDesc' | 'Newest' | 'Oldest' | 'RatingAsc' | 'RatingDesc';

export function isSortByValue(value: string): value is SortByValue {
	return ['NameAsc', 'NameDesc', 'Newest', 'Oldest', 'RatingAsc', 'RatingDesc'].includes(value);
}

export interface SortByOption {
	value: SortByValue;
	label: string;
}

export const SortByOptions: Record<SortByValue, SortByOption> = {
	NameAsc: {
		value: 'NameAsc',
		label: 'Name (A-Z)'
	},
	NameDesc: {
		value: 'NameDesc',
		label: 'Name (Z-A)'
	},
	Newest: {
		value: 'Newest',
		label: 'Newest'
	},
	Oldest: {
		value: 'Oldest',
		label: 'Oldest'
	},
	RatingAsc: {
		value: 'RatingAsc',
		label: 'Rating (Lowest)'
	},
	RatingDesc: {
		value: 'RatingDesc',
		label: 'Rating (Highest)'
	}
};

export interface SearchOptions {
	sortBy: SortByValue;
	filter: FilterOptions;
}

export const DEFAULT_SEARCH_OPTIONS: SearchOptions = {
	sortBy: 'Newest',
	filter: DEFAULT_FILTER_OPTIONS
};
