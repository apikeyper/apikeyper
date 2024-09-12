import { GitHub } from "arctic";

export const githubAuth = new GitHub(process.env.GITHUB_CLIENT_ID!, process.env.GITHUB_CLIENT_SECRET!);