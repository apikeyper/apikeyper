import { pgSchema, text, timestamp, uniqueIndex } from "drizzle-orm/pg-core";

export const authSchema = pgSchema("auth");

export const userTable = authSchema.table("user", {
	id: text("id").primaryKey(),
	username: text("username")
		.notNull(),
	githubId: text("github_id")
		.notNull(),
},
	(table) => ({
		githubIdUnqIdx: uniqueIndex('github_id_unq_idx').on(table.githubId)
	})
);
export const UserType = userTable.$inferSelect

export const sessionTable = authSchema.table("session", {
	id: text("id").primaryKey(),
	userId: text("user_id")
		.notNull()
		.references(() => userTable.id),
	expiresAt: timestamp("expires_at", {
		withTimezone: true,
		mode: "date"
	}).notNull()
});
