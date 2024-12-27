"use client";

import React from "react";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { useToast } from "@/components/ui/use-toast";

export async function copyToClipboardWithMeta(value: string) {
  navigator.clipboard.writeText(value);
}

export function MinimalButtonWithCopy({
  displayText,
  text,
  className,
}: {
  displayText?: string;
  text: string;
  className?: string;
}) {
  const { toast } = useToast();
  return (
    <div>
      <Button
        size="icon"
        variant="ghost"
        className={cn("w-max items-center space-x-2 rounded p-2 ", className)}
        onClick={() => {
          copyToClipboardWithMeta(text);
          toast({
            title: "Copied to clipboard",
            duration: 1000,
          });
        }}
      >
        <span className="sr-only">Copy</span>
        {displayText || text}
      </Button>
    </div>
  );
}
