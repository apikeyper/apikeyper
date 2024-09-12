import { defineConfig } from 'drizzle-kit'

export default defineConfig({
  schema: "./lib/db/schema.ts",
  dialect: 'postgresql',
  dbCredentials: {
    host: '0.0.0.0',
    port: 5438,
    user: 'postgres',
    password: 'pass',
    database: 'postgres',
  },
  schemaFilter: "auth",
  verbose: true,
  strict: true,
})

