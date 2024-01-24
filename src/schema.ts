import {
	pgTable,
	bigint,
	varchar,
	text,
	pgEnum,
	timestamp,
	boolean,
	real,
	smallint,
	uuid
} from 'drizzle-orm/pg-core';

// AUTH

export const onboardingStatus = pgEnum('onboarding_status', [
	'PENDING_EMAIL_VERIFICATION',
	'PENDING_USERNAME',
	'PENDING_AVATAR',
	'PENDING_BIO',
	'COMPLETE'
]);

export const user = pgTable('auth_user', {
	id: varchar('id', {
		length: 15
	}).primaryKey(),
	email: text('email').notNull().unique(),
	username: text('username').unique(),
	avatarUrl: text('avatar_url').notNull(),
	bio: text('bio'),
	onboardingStatus: onboardingStatus('onboarding_status').notNull().default('PENDING_USERNAME')
});

export const session = pgTable('user_session', {
	id: varchar('id', {
		length: 128
	}).primaryKey(),
	userId: varchar('user_id', {
		length: 15
	})
		.notNull()
		.references(() => user.id),
	activeExpires: bigint('active_expires', {
		mode: 'number'
	}).notNull(),
	idleExpires: bigint('idle_expires', {
		mode: 'number'
	}).notNull()
});

export const key = pgTable('user_key', {
	id: varchar('id', {
		length: 255
	}).primaryKey(),
	userId: varchar('user_id', {
		length: 15
	})
		.notNull()
		.references(() => user.id, { onDelete: 'cascade', onUpdate: 'cascade' }),
	hashedPassword: varchar('hashed_password', {
		length: 255
	})
});

export const emailVerification = pgTable('email_verification_token', {
	id: varchar('id', {
		length: 255
	}).primaryKey(),
	expires: bigint('expires', {
		mode: 'number'
	}).notNull(),
	userId: varchar('user_id', {
		length: 15
	})
		.notNull()
		.references(() => user.id, { onDelete: 'cascade', onUpdate: 'cascade' })
});

// MODELS

export const difficulty = pgEnum('difficulty', ['EASY', 'MEDIUM', 'HARD']);
export const brewType = pgEnum('brew_type', [
	'Ale',
	'Lager',
	'Stout',
	'IPA',
	'Mead',
	'Melomel',
	'Cyser',
	'Hydromel',
	'Metheglin',
	'Cider',
	'Fruit Wine',
	'Other'
]);
export const ingredientType = pgEnum('ingredient_type', [
	'Grain',
	'Hops',
	'Yeast',
	'Fruit',
	'Spice',
	'Honey',
	'Sugar',
	'Nutrient',
	'Additives',
	'Other'
]);
export const unitOfMeasurement = pgEnum('unit_of_measurement', [
	'g',
	'kg',
	'oz',
	'lb',
	'ml',
	'liter',
	'tsp',
	'tbsp',
	'cup',
	'pint',
	'quart',
	'gal',
	'barrel'
]);

export const recipe = pgTable('recipe', {
	id: uuid('id').primaryKey().notNull(),
	ownerId: varchar('owner_id', { length: 15 })
		.notNull()
		.references(() => user.id, { onDelete: 'cascade', onUpdate: 'cascade' }),
	createdAt: timestamp('created_at').notNull().defaultNow(),
	updatedAt: timestamp('updated_at').notNull().defaultNow(),
	name: text('name'),
	description: text('description'),
	published: boolean('published').notNull().default(false),
	difficulty: difficulty('difficulty').default('EASY'),
	brewType: brewType('brew_type'),
	originalGravity: real('original_gravity'),
	finalGravity: real('final_gravity'),
	sweetenedGravity: real('sweetened_gravity'),
	process: text('process').array(),
	rating: real('rating').default(0),
	batchSize: smallint('batch_size').default(0),
	batchUnit: unitOfMeasurement('batch_unit').default('gal'),
	images: text('images').array(),
	notes: text('notes')
});

export const recipeIngredient = pgTable('recipe_ingredient', {
	id: uuid('id').primaryKey().notNull(),
	recipeId: uuid('recipe_id')
		.notNull()
		.references(() => recipe.id, { onDelete: 'cascade', onUpdate: 'cascade' }),
	name: text('name'),
	quantity: real('quantity'),
	unit: unitOfMeasurement('unit')
});
