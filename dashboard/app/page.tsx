import Link from "next/link";
import { Button } from "@/components/ui/button";
import UserItem from "@/components/userItem";

export default async function Home() {
  return (
    <main>
      <div className="flex flex-col items-center space-y-3 p-10 text-center">
        <UserItem />
        <Link href="/apis">
          <Button className="bg-zinc-300 text-black dark:bg-zinc-900 dark:text-white">
            Go to APIs
          </Button>
        </Link>
      </div>
    </main>
  );
}
