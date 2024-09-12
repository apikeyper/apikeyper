'use server'

export async function createDefaultWorkspaceForUser(userId: string, sessionId: string): Promise<string> {
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;
  const response = await fetch(`${baseUrl}/workspace`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${sessionId}`,
    },
    body: JSON.stringify({
      // "userId": userId,
      "name": `"Default-${userId}"`,
    }),
  })

  if (!response.ok) {
    throw new Error("Failed to create a new workspace");
  }

  const { workspaceId }: { workspaceId: string } = await response.json();
  return workspaceId;
}

