import { drizzle } from 'drizzle-orm/postgres-js';
import { migrate } from 'drizzle-orm/postgres-js/migrator';
import postgres from 'postgres';

const pg_db_url = process.env.POSTGRES_DB_URL;

if (!pg_db_url) {
	throw new Error('POSTGRES_DB_URL env var not set');
}

console.log('Connecting to DB');

const migrationClient = postgres(pg_db_url, { max: 1 });

console.log('Running migrations');
await migrate(drizzle(migrationClient), { migrationsFolder: './drizzle' });

console.log('Done!');
process.exit(0);
