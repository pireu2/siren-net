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
import { Form, FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form";
import Popup from "./popUp/pop-up";
import { jwtDecode } from "jwt-decode";
import React, { useContext,useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "./auth/auth-handler";

export function LoginForm({
  className,
  onShuffle,
  ...props
}) {
  const form = useForm({
    defaultValues: {
    username: '',
    password: ''
  }
});

  const {message, setOpen, login, isOpen} = useContext(AuthContext);


  return (
    <div className={cn("flex flex-col gap-6 items-center justify-center w-full max-w-md h-[600px]", className)} {...props}>
      <Card className="w-full h-full flex flex-col justify-between">
        <CardHeader className="text-center" style={{ marginTop: '20%' }}>
          <CardTitle>Login to your account</CardTitle>
          <Popup isOpen = {isOpen} setOpenState={setOpen} message={message}/> 
          <CardDescription style={{ marginTop: '14%' }}>
            Enter your username below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent className="flex-1 flex flex-col justify-center">
        <Form {...form}>
            <form onSubmit={form.handleSubmit(login)} className="flex flex-col gap-12">
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
