import { Skeleton } from "@/components/ui/skeleton";

export function ApiCardsSkeleton() {
  return (
    <div className="flex flex-row space-x-2">
      {[1, 2].map((_, index) => (
        <Skeleton
          key={index}
          className="w-1/6 rounded-lg bg-zinc-200 p-10 shadow-sm shadow-black/50 dark:bg-zinc-700 dark:text-white"
        />
      ))}
    </div>
  );
}
