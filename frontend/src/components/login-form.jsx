import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { useForm } from "react-hook-form";
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Link } from "react-router-dom";
import { Form, FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form";

export function LoginForm({
  className,
  onShuffle,
  ...props
}) {
  const form = useForm();

  const onSubmit = async (formData) => {
    console.log("Form Data:", formData);

    try {
      const response = await fetch("/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      const data = await response.json();
      if (response.ok) {
        alert("Registration successful!");
      } else {
        alert("Error: " + data.error);
      }
    } 
    
    catch (error) {
      console.error("Error:", error);
    }

  };


  return (
    <div className={cn("flex flex-col gap-6 items-center justify-center w-full max-w-md h-[500px]", className)} {...props}>
      <Card className="w-full h-full flex flex-col justify-between">
        <CardHeader className="text-center" style={{ marginTop: '14%' }}>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription style={{ marginTop: '14%' }}>
            Enter your username below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent className="flex-1 flex flex-col justify-center">
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-6">
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="username">Username</FormLabel>
                    <FormControl>
                      <Input id="username" type="text" placeholder="Your username" {...field} required />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="password">Password</FormLabel>
                    <FormControl>
                      <Input id="password" type="password" placeholder="Your password" {...field} required />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <div className="flex flex-col gap-3">
                <Button type="submit" className="w-full">
                  Login
                </Button>
              </div>
            </form>
        </Form>
        </CardContent>
        <div className="mt-4 text-center text-sm">
          Don't have an account?{" "}
          <a className="underline underline-offset-4" onClick={onShuffle}>
            Sign up
          </a>
        </div>
      </Card>
    </div>
  );
}
