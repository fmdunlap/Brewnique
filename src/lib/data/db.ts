import { drizzle } from 'drizzle-orm/postgres-js';
import { POSTGRES_DB_URL } from '$env/static/private';
import postgres from 'postgres';

const sql = postgres(POSTGRES_DB_URL, { ssl: 'require' });
export const db = drizzle(sql);
