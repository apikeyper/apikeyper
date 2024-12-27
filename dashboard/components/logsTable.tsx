"use client";
import { useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useRouter } from "next/navigation";

interface UsageRecord {
  intervalStart: string;
  success: number;
  failed: number;
  revoked: number;
  rateLimited: number;
}

interface UsageResponse {
  apiKeyId: string;
  records: UsageRecord[];
  count: number;
}

interface LogsTableProps {
  initialData: UsageResponse;
}

export function LogsTable({ initialData }: LogsTableProps) {
  const [interval, setInterval] = useState("5");
  const router = useRouter();

  const handleIntervalChange = (newInterval: string) => {
    setInterval(newInterval);
    router.refresh(); // This will trigger a new server-side fetch
  };

  return (
    <div>
      <div className="mb-4 flex justify-end">
        <Select value={interval} onValueChange={handleIntervalChange}>
          <SelectTrigger className="w-32">
            <SelectValue placeholder="Interval" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="1">1 min</SelectItem>
            <SelectItem value="5">5 mins</SelectItem>
            <SelectItem value="15">15 mins</SelectItem>
            <SelectItem value="30">30 mins</SelectItem>
            <SelectItem value="60">1 hour</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Time Interval</TableHead>
            <TableHead>Successful</TableHead>
            <TableHead>Failed</TableHead>
            <TableHead>Revoked</TableHead>
            <TableHead>Rate Limited</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {initialData.records.map((record) => (
            <TableRow key={record.intervalStart}>
              <TableCell>
                {new Date(record.intervalStart).toLocaleString()}
              </TableCell>
              <TableCell>
                <span className="rounded-full bg-green-100 px-2 py-1 text-sm text-green-800">
                  {record.success}
                </span>
              </TableCell>
              <TableCell>
                <span className="rounded-full bg-red-100 px-2 py-1 text-sm text-red-800">
                  {record.failed}
                </span>
              </TableCell>
              <TableCell>
                <span className="rounded-full bg-orange-100 px-2 py-1 text-sm text-orange-800">
                  {record.revoked}
                </span>
              </TableCell>
              <TableCell>
                <span className="rounded-full bg-yellow-100 px-2 py-1 text-sm text-yellow-800">
                  {record.rateLimited}
                </span>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
