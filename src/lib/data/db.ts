import { dev } from '$app/environment';
import { DEV_POSTGRES_DB_URL, POSTGRES_DB_URL } from '$env/static/private';
import { drizzle, type PostgresJsDatabase } from 'drizzle-orm/postgres-js';
import postgres from 'postgres';

export const queryClient = postgres(dev ? DEV_POSTGRES_DB_URL : POSTGRES_DB_URL);

export const db: PostgresJsDatabase = drizzle(queryClient);
