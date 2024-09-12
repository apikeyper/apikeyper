export async function getNewRootKey(workspaceId: string, keyName?: string): Promise<string> {
  // Check if the root key is already stored in a cookie
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;

  const response = await fetch(`${baseUrl}/rootKey`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      "name": keyName ?? "web-client-default",
      "workspaceId": workspaceId,
    }),
  })

  if (!response.ok) {
    throw new Error("Failed to fetch root key");
  }

  const respJson = await response.json();
  const newRootKey = respJson.rootKey as string;
  return newRootKey;
}