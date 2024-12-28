"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { MinimalButtonWithCopy } from "@/components/ui/minimal-copy-button";

interface ActivityRecord {
  keyId: string;
  keyName?: string;
  apiId: string;
  usage: string;
  timestamp: string;
}

interface ActivityTableProps {
  activities: ActivityRecord[];
}

export function ActivityTable({ activities }: ActivityTableProps) {
  const getUsageStyle = (usage: string) => {
    switch (usage) {
      case "success":
        return "bg-green-300 text-green-800";
      case "failed":
        return "bg-red-500 text-red-800";
      case "revoked":
        2;
        return "bg-red-300 text-red-800";
      case "rate_limited":
        return "bg-amber-300 text-amber-800";
      default:
        return "bg-gray-300 text-gray-800";
    }
  };

  const getRelativeTime = (timestamp: string) => {
    const date = new Date(timestamp);
    const now = new Date();
    const diffInSeconds = Math.floor((now.getTime() - date.getTime()) / 1000);

    if (diffInSeconds < 60) {
      return "just now";
    } else if (diffInSeconds < 3600) {
      const minutes = Math.floor(diffInSeconds / 60);
      return `${minutes} ${minutes === 1 ? "minute" : "minutes"} ago`;
    } else if (diffInSeconds < 86400) {
      const hours = Math.floor(diffInSeconds / 3600);
      return `${hours} ${hours === 1 ? "hour" : "hours"} ago`;
    } else {
      const days = Math.floor(diffInSeconds / 86400);
      return `${days} ${days === 1 ? "day" : "days"} ago`;
    }
  };

  const getFullTimestamp = (timestamp: string) => {
    return new Date(timestamp).toLocaleString("en-US", {
      year: "numeric",
      month: "numeric",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      fractionalSecondDigits: 3,
      hour12: false,
    });
  };

  return (
    <div className="relative overflow-x-auto">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Key</TableHead>
            <TableHead>Api Id</TableHead>
            <TableHead>Usage</TableHead>
            <TableHead>Time</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {activities.map((activity) => (
            <TableRow key={activity.timestamp}>
              <TableCell>
                <MinimalButtonWithCopy
                  displayText={activity.keyName}
                  text={activity.keyId}
                />
              </TableCell>
              <TableCell className="font-mono text-sm">
                <MinimalButtonWithCopy text={activity.keyId} />
              </TableCell>
              <TableCell>
                <span
                  className={`rounded-full px-2 py-1 text-sm ${getUsageStyle(activity.usage)}`}
                >
                  {activity.usage}
                </span>
              </TableCell>
              <TableCell>
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger className="text-sm text-gray-600">
                      {getRelativeTime(activity.timestamp)}
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>{getFullTimestamp(activity.timestamp)}</p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
