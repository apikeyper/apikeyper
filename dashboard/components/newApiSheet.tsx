'use client'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm, useFormState } from "react-hook-form"
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet"

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { useToast } from "@/components/ui/use-toast"

const formSchema = z.object({
  apiName: z.string().min(2, {
    message: "Api name must be at least 2 characters.",
  }),
})

export function CreateNewApiSheet() {
  const { toast } = useToast();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      apiName: "",
    },
  })
  async function onSubmit(values: z.infer<typeof formSchema>) {
    toast({
      title: "Creating API...",
      duration: 1000,
    });

    const resp = await fetch(`/api/createApi`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        apiName: values.apiName,
      }),
    })

    if (resp.ok) {
      window.location.reload()
      return
    } else {
      toast({
        title: "Failed to create api. Please try again.",
        duration: 2000,
        variant: "destructive",
      });
    }
  }

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button className="bg-zinc-400 dark:bg-zinc-900 text-black dark:text-white hover:bg-zinc-700">Create API</Button>
      </SheetTrigger>
      <SheetContent className="w-3/4 bg-zinc-400 dark:bg-zinc-800 text-black dark:text-white">
        <SheetHeader>
          <SheetTitle className="text-black dark:text-white">Create a new API</SheetTitle>
          <SheetDescription>
            Define the properties of your new API.
          </SheetDescription>
        </SheetHeader>
        <Form {...form}>
          <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8 text-white">
            <FormField
              control={form.control}
              name="apiName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>API name</FormLabel>
                  <FormControl>
                    <Input {...field} className="text-black dark:bg-zinc-900 dark:text-white" />
                  </FormControl>
                  <FormDescription>
                    This is your api display name.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button className="text-black dark:text-white bg-zinc-400 dark:bg-zinc-900" type="submit">Submit</Button>
          </form>
        </Form>
      </SheetContent>
    </Sheet>
  )
}
