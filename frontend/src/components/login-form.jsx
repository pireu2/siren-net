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
import { Eye , EyeOff} from 'lucide-react';


export function LoginForm({
  className,
  form,
  onShuffle,
  ...props
}) 
{

  const {message, setOpen, login, isOpen} = useContext(AuthContext);
  const [showPassword, setShowPassword] = useState(false);

  return (
    <div className={cn("flex flex-col gap-6 items-center justify-center w-full max-w-md h-[37.5rem]", className)} {...props}>
      <Card className="w-full h-full flex flex-col justify-between">
        <CardHeader className="text-center" style={{ marginTop: '20%' }}>
          <CardTitle>Login to your account</CardTitle>
          <Popup isOpen={isOpen} setOpenState={setOpen} message={message}/> 
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
                  <FormLabel>Password</FormLabel>
                  <FormControl className="relative">
                    <div className="relative">
                      <Input
                        type={showPassword ? "text" : "password"}
                        placeholder="Your password"
                        {...field}
                        required
                      />
                      <Button
                        type="button"
                        variant="ghost"
                        size="icon"
                        className="absolute right-0 top-0 h-full px-3 py-2 text-muted-foreground hover:text-foreground"
                        onClick={()=>setShowPassword(!showPassword) }
                      >
                        {showPassword ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
                        <span className="sr-only">{showPassword ? "Hide password" : "Show password"}</span>
                      </Button>
                    </div>
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
