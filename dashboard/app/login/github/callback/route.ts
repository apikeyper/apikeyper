import { cookies } from "next/headers";
import { OAuth2RequestError } from "arctic";
import { generateIdFromEntropySize } from "lucia";
import { githubAuth } from "@/app/auth/oauth";
import { lucia } from "@/app/auth/lucia.server";
import { createGithubUser, getUserByGithubId } from "@/lib/db/query";

export async function GET(request: Request): Promise<Response> {

	console.log("Received callback request");

	const url = new URL(request.url);
	const code = url.searchParams.get("code");
	const state = url.searchParams.get("state");


	const storedState = cookies().get("github_oauth_state")?.value ?? null;

	if (!code || !state || !storedState || state !== storedState) {
		return new Response(null, {
			status: 400
		});
	}

	try {
		const tokens = await githubAuth.validateAuthorizationCode(code);
		const githubUserResponse = await fetch("https://api.github.com/user", {
			headers: {
				Authorization: `Bearer ${tokens.accessToken}`
			}
		});
		const githubUser: GitHubUser = await githubUserResponse.json();

    // Fetch existing user
		const existingUser = await getUserByGithubId(githubUser.id);
		console.log("Existing user: ", existingUser);

		if (existingUser !== null) {
			const session = await lucia.createSession(existingUser.id, {});
			const sessionCookie = lucia.createSessionCookie(session.id);
			cookies().set(sessionCookie.name, sessionCookie.value, sessionCookie.attributes);
			return new Response(null, {
				status: 302,
				headers: {
					Location: "/onboarding"
				}
			});
		}

    // Create a new user
		console.log("Creating new user");
		const userId = generateIdFromEntropySize(10); // 16 characters long
		await createGithubUser(userId, githubUser.id, githubUser.login);

		const session = await lucia.createSession(userId, {});
		const sessionCookie = lucia.createSessionCookie(session.id);
		cookies().set(sessionCookie.name, sessionCookie.value, sessionCookie.attributes);
		return new Response(null, {
			status: 302,
			headers: {
				Location: "/onboarding"
			}
		});
	} catch (e) {
		// the specific error message depends on the provider
		if (e instanceof OAuth2RequestError) {
			// invalid code
			return new Response(null, {
				status: 400
			});
		}
		return new Response(null, {
			status: 500
		});
	}
}

interface GitHubUser {
	id: string;
	login: string;
}