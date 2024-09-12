'use server'

import { apiKeyperUrl } from "./config";

export async function createDefaultWorkspaceForUser(githubId: string, sessionId: string): Promise<string> {
  const response = await fetch(`${apiKeyperUrl}/workspace`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${sessionId}`,
    },
    body: JSON.stringify({
      // "userId": userId,
      "name": `"Default-${githubId}"`,
    }),
  })

  if (!response.ok) {
    throw new Error("Failed to create a new workspace");
  }

  const { workspaceId }: { workspaceId: string } = await response.json();
  return workspaceId;
}

