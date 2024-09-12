import { drizzle } from "drizzle-orm/node-postgres";
import { Pool } from "pg";
const pool = new Pool({
  connectionString: "postgres://postgres:pass@localhost:5438/postgres",
});

export const dbClient = drizzle(pool);