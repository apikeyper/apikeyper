import { getRootKey } from "@/app/auth/auth.client";
import { ActivityTable } from "@/components/activityTable";
import { Card, CardHeader, CardContent } from "@/components/ui/card";
import { type Metadata } from "next";

export const metadata: Metadata = {
  title: "Activity logs",
  description: "Api Keys as a Service",
};

interface ActivityRecord {
  keyId: string;
  apiId: string;
  usage: string;
  timestamp: string;
}

async function getActivityData(): Promise<ActivityRecord[]> {
  const rootKey = await getRootKey();
  const baseUrl = process.env.NEXT_PUBLIC_API_URL;

  const response = await fetch(`${baseUrl}/apiKey/activity`, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${rootKey}`,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch activity data");
  }

  return response.json();
}

export default async function ActivityPage() {
  const activities = await getActivityData();

  return (
    <div className="container mx-auto py-6">
      <Card>
        <CardHeader>
          <h1 className="text-2xl font-bold">Activity</h1>
        </CardHeader>
        <CardContent>
          <ActivityTable activities={activities} />
        </CardContent>
      </Card>
    </div>
  );
}
