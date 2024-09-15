import { drizzle } from "drizzle-orm/node-postgres";
import { Pool } from "pg";

export const pool = new Pool({
  connectionString: process.env.DATABASE_URL as string,
});

export const dbClient = drizzle(pool, {logger: process.env.NODE_ENV === "development"});