import { getRootKey } from "@/app/auth/auth.client";
import { apiKeyperUrl } from "./config";


interface CreateApiProps {
  apiName: string;
}

export async function createApi(createApiProps: CreateApiProps) {
  const rootKey = await getRootKey();
  const response = await fetch(`${apiKeyperUrl}/api`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${rootKey}`,
    },
    body: JSON.stringify({
      "apiName": createApiProps.apiName,
    }),
  })

  const data = await response.json();

  if (!response.ok) {
    throw new Error("Failed to create a new api");
  }

  return data;
}