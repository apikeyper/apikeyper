import { dbClient } from "@/lib/db/client";
import { sessionTable, userTable } from "@/lib/db/schema";
import { DrizzlePostgreSQLAdapter } from "@lucia-auth/adapter-drizzle";
import { Lucia } from "lucia";


export const dbAdapter = new DrizzlePostgreSQLAdapter(dbClient, sessionTable, userTable);

export const lucia = new Lucia(dbAdapter, {
	sessionCookie: {
		attributes: {
			// set to `true` when using HTTPS
			secure: process.env.NODE_ENV === "production"
		}
	},
	getUserAttributes: (attributes) => {
		return {
			// attributes has the type of DatabaseUserAttributes
			githubId: attributes.githubId,
			username: attributes.username
		};
	}
});


// IMPORTANT!
declare module "lucia" {
	interface Register {
		Lucia: typeof lucia;
		DatabaseUserAttributes: DatabaseUserAttributes;
	}
}

interface DatabaseUserAttributes {
	githubId: string;
	username: string;
}