"use client";

import { ThemeToggle } from "@/components/theme-toggle";

export default function Header() {
  return (
    <div className="dark:border-b-1 flex w-full gap-4 border-b border-b-black/25 p-4 dark:border-b-white/25">
      <div className="ml-auto">
        <ThemeToggle />
      </div>
    </div>
  );
}
