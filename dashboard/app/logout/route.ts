import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { validateRequest } from "../auth/validate";
import { lucia } from "../auth/lucia.server";
import { NextResponse } from "next/server";

export async function GET(): Promise<Response> {
	const { session } = await validateRequest();
	if (!session) {
		return NextResponse.json({
			error: "Unauthorized"
		});
	}

	await lucia.invalidateSession(session.id);

	const sessionCookie = lucia.createBlankSessionCookie();
	cookies().set(sessionCookie.name, sessionCookie.value, sessionCookie.attributes);

  // Clear auth cookies for application
  cookies().delete("apikeyper-auth");

	return redirect("/login");
}