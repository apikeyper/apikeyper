import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers";
import { getNewRootKey } from "@/lib/apiKeyService/rootKey";
import { sealData } from 'iron-session';


export async function middleware(request: NextRequest, next: NextResponse) {
  const sessionId = cookies().get("auth_session")
	if (!sessionId) {
    return NextResponse.redirect(new URL("/login", request.nextUrl))
	}

  const response = NextResponse.next();
  // If onboarding, redirect to onboarding page
  if (request.nextUrl.pathname === "/onboarding") {
    return response
  }


  // Check if root key in cookies
  if (!cookies().get("apikeyper-auth")) {
    return NextResponse.redirect(new URL("/login", request.nextUrl))
  }

  return response
}

// Match against the pages
export const config = { matcher: ['/((?!api|callback|login|home|_next/static|_next/image|favicon.ico).*)', "/apis/:path*"] };