import { dbClient } from "./client";
import { userTable, UserType } from "./schema";
import { eq } from "drizzle-orm/expressions";

export async function getUserByGithubId(githubId: string): Promise<typeof UserType | null> {
  const result = await dbClient.select().from(userTable).where(eq(userTable.githubId, githubId)).limit(1);
  return result[0] ?? null;
}


export async function createGithubUser(userId: string, githubId: string, username: string) {
  await dbClient.insert(userTable).values({
    id: userId,
    githubId,
    username
  });
}