import React, { useContext } from "react";
import { useForm } from "react-hook-form";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Form, FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form";
import Popup from "./popUp/pop-up";
import { AuthContext } from "./auth/auth-handler";

export function Register({ onShuffle, loginForm }) {
  const form = useForm({
    defaultValues: {
    username: '',
    email: '',
    password: '',
    confirm_password: ''
  }
});

  const {message, setOpen, register, isOpen} = useContext(AuthContext);

  async function submitHandler(props)
  {
    const res = await register(props);
    if(!res)
    {
      loginForm.setValue('username', props.username);
      loginForm.setValue('password', props.password);
      onShuffle();
    }
  }

  return (
    <div className={cn("flex flex-col gap-6 items-center justify-center w-full max-w-md h-[37.5rem]")}>
      <Card className="w-full h-full flex flex-col justify-between">
        <CardHeader className="text-center">
          <CardTitle>Create a new account</CardTitle>
          <Popup isOpen = {isOpen} setOpenState={setOpen} message={message}/> 
          <CardDescription style={{ marginTop: '14%' }}> 
            Enter your details below to create a new account
          </CardDescription>
        </CardHeader>
        <CardContent className="flex-1 flex flex-col justify-center">
          <Form {...form}>
            <form onSubmit={form.handleSubmit((p)=>submitHandler(p))} className="flex flex-col gap-6">
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
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="email">Email</FormLabel>
                    <FormControl>
                      <Input id="email" type="email" placeholder="example@example.com" {...field} required />
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
              <FormField
                control={form.control}
                name="confirm_password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel htmlFor="confirm_password">Confirm Password</FormLabel>
                    <FormControl>
                      <Input id="confirm_password" type="password" placeholder="Your password" {...field} required />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <div className="flex flex-col gap-3">
                <Button type="submit" className="w-full">
                  Register
                </Button>
              </div>
              <div className="mt-4 text-center text-sm">
                Already have an account?{" "}
                <a className="underline underline-offset-4" onClick={onShuffle}>
                  Login
                </a>
              </div>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
}

export default Register;
