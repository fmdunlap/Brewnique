import { defineConfig } from 'drizzle-kit';

const connectionString = process.env.POSTGRES_DB_URL;

if (!connectionString) {
	throw new Error('POSTGRES_DB_URL env var not set');
}

export default defineConfig({
	driver: 'pg',
	out: './drizzle',
	dbCredentials: {
		connectionString:
			'postgresql://postgres:?%p~dW+4LU9mPq~@db.jefjlzwehchcvsqsdgit.supabase.co:5432/postgres'
	},
	verbose: true,
	strict: true
});
