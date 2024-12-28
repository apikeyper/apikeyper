import Link from "next/link";
import { getRootKey } from "@/app/auth/auth.client";
import { CreateNewApiSheet } from "@/components/newApiSheet";
import { type Metadata } from "next";

export const metadata: Metadata = {
  title: "Apis",
  description: "Api Keys as a Service",
};

interface Api {
  apiId: string;
  workspaceId: string;
  apiName: string;
  createdAt: string;
  updatedAt: string;
}

async function getApis(): Promise<Api[]> {
  const rootKey = await getRootKey();
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;

  const response = await fetch(`${baseUrl}/api/list`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${rootKey}`,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch APIs");
  }

  const respJson = await response.json();

  return respJson;
}

export default async function ApisPage() {
  const apis = await getApis();

  if (!apis) {
    return (
      <div className="flex flex-col">
        <div className="p-5">
          <CreateNewApiSheet />
        </div>
        <p className="text-center text-lg">Create a new API to get started</p>
      </div>
    );
  }

  return (
    <main className="flex flex-col space-y-5">
      <div>
        <CreateNewApiSheet />
      </div>
      <div className="flex flex-row space-x-2">
        {apis.map((api) => (
          <Link key={api.apiId} href={`/apis/${api.apiId}`}>
            <div
              key={api.apiId}
              className="cursor-pointer rounded-lg bg-zinc-300 p-10 shadow-sm shadow-black/50 hover:ring-1 hover:ring-black hover:ring-opacity-50 dark:bg-zinc-900 dark:text-white dark:hover:ring-1 dark:hover:ring-white dark:hover:ring-opacity-50"
            >
              {api.apiName}
            </div>
          </Link>
        ))}
      </div>
    </main>
  );
}
