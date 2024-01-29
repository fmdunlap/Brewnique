export type SortByValue = 'NameAsc' | 'NameDesc' | 'Newest' | 'Oldest' | 'RatingAsc' | 'RatingDesc';

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

export type BatchSizeValue = 'gallonOrLess' | 'gallonToBarrel' | 'barrelOrMore';

export const BatchSizeOptions: { value: BatchSizeValue; label: string }[] = [
	{
		value: 'gallonOrLess',
		label: '1 gallon or less'
	},
	{
		value: 'gallonToBarrel',
		label: '1 gallon - 1 barrel'
	},
	{
		value: 'barrelOrMore',
		label: '1 barrel or more'
	}
];

export type AbvFilterValue = 'underSeven' | 'sevenToFifteen' | 'overFifteen';

export const AbvFilterOptions: { value: AbvFilterValue; label: string }[] = [
	{
		value: 'underSeven',
		label: 'Under 7%'
	},
	{
		value: 'sevenToFifteen',
		label: '7% - 15%'
	},
	{
		value: 'overFifteen',
		label: 'Over 15%'
	}
];

export interface SearchSidebarOptions {
	sortBy: SortByValue;
	filter: {
		size: BatchSizeValue[];
		rating: number[];
		abv: AbvFilterValue[];
	};
}

export const DefaultSearchSidebarOptions: SearchSidebarOptions = {
	sortBy: 'NameAsc',
	filter: {
		size: [],
		rating: [],
		abv: []
	}
};
