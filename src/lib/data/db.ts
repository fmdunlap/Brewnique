import { drizzle } from 'drizzle-orm/postgres-js';
import { NEONDB_URL } from '$env/static/private';
import postgres from 'postgres';

const sql = postgres(NEONDB_URL, { ssl: 'require' });
export const db = drizzle(sql);
