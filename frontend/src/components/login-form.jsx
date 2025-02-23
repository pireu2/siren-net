import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Link } from "react-router-dom";

export function LoginForm({
  className,
  onClick,
  ...props
}) {
  return (
    <div className={cn("flex flex-col gap-6 items-center justify-center w-full max-w-md h-[500px]", className)} {...props}>
      <Card className="w-full h-full flex flex-col justify-between">
        <CardHeader className="text-center" style={{ marginTop: '14%' }}>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription style={{ marginTop: '14%' }}>
            Enter your email below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent className="flex-1 flex flex-col justify-center">
          <form className="flex flex-col gap-6">
            <div className="grid gap-3">
              <Label htmlFor="email">Username</Label>
              <Input id="email" type="email" placeholder="ExampleUser" required />
            </div>
            <div className="grid gap-3">
              <Label htmlFor="password">Password</Label>
              <Input id="password" type="password" placeholder="example123" required />
            </div>
            <Button type="submit" className="w-full">Login</Button>
          </form>
        </CardContent>
        <div className="mt-4 text-center text-sm">
          Don't have an account?{" "}
          <Link className="underline underline-offset-4" onClick={onClick}>
            Sign up
          </Link>
        </div>
      </Card>
    </div>
  );
}
