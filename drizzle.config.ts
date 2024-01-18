import { defineConfig } from 'drizzle-kit';

const connectionString = process.env.POSTGRES_DB_URL;

if (!connectionString) {
	throw new Error('POSTGRES_DB_URL env var not set');
}

export default defineConfig({
	schema: './src/schema.ts',
	driver: 'pg',
	out: './drizzle',
	dbCredentials: {
		connectionString: connectionString
	},
	verbose: true,
	strict: true
});
