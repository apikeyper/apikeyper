import { apiKeyperUrl } from "./config";

export async function getNewRootKey(workspaceId: string, keyName?: string): Promise<string> {
  // Check if the root key is already stored in a cookie
  const response = await fetch(`${apiKeyperUrl}/rootKey`, {
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