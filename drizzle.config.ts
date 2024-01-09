import { defineConfig } from 'drizzle-kit';

const connectionString = process.env.NEONDB_URL;

if (!connectionString) {
	throw new Error('NEONDB_URL env var not set');
}

export default defineConfig({
	schema: './src/schema.ts',
	driver: 'pg',
	dbCredentials: {
		connectionString: connectionString
	},
	verbose: true,
	strict: true
});
