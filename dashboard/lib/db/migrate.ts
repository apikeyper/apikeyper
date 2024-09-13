import { migrate } from 'drizzle-orm/node-postgres/migrator';
import { drizzle } from "drizzle-orm/node-postgres";
import { Pool } from "pg";

export const pool = new Pool({
  connectionString: process.env.DATABASE_URL as string,
});

export const dbClient = drizzle(pool);

async function main() {
  await migrate(dbClient, { migrationsFolder: "./migrations" });
  await pool.end();
}


main().then(
  () => {
    console.log('Migration successful')
    pool.end()
    process.exit(0)
  }
).catch((e) => {
  console.error('Migration failed')
  console.error(e)
  process.exit(1)
})