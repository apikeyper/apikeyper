import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Header from "@/components/ui/header";
import { Sidebar } from "@/components/sidebar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "ApiKeyper",
  description: "Api Keys as a Service",
};

export default async function DashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className={`${inter.className} bg-zinc-200 dark:bg-zinc-800`}>
      <Sidebar />
      <Header />
      <div className="flex items-start justify-between">
        <main className="grid h-full w-full pl-[300px]">
          <div className="p-8">
            {children}
          </div>
        </main>
      </div>
    </main>
  );
}
