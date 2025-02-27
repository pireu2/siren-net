import { AppSidebar } from "@/components/dashboard/app-sidebar"
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import { Separator } from "@/components/ui/separator"
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar"
import TitleCard from "./TitleCards"
import { dashboardSections } from "./TitleCards"
import AiPromptingDashboard from "../ai-prompting/ai-dashboard"
import { useState, useEffect } from 'react';

export default function Dashboard() {
  const [pageState, setState] = useState(()=>{
    return localStorage.getItem('pageState') || "idle"
  });
  const possibleStates = {0:"idle", 1: "SD", 2:"AI"}


  useEffect(() => {
    localStorage.setItem('pageState', pageState);
  }, [pageState]);


  function changeState(key)
  {
    setState(possibleStates[key])
  }

  function IdleComponent()
  {
    return(
      <>
      <h1 className="text-3xl font-bold mb-4">Welcome to the <i>siren-net</i> Dashboard</h1>
      <div className="grid gap-6 md:grid-cols-1">
        {dashboardSections.map((section) => (
          <TitleCard key={section.id} {...section} onClick = {()=> changeState(section.id)}/>
        ))}
      </div>
      </>
    );
  }

  return (
    (<SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
          <SidebarTrigger className="-ml-1" />
          <Separator orientation="vertical" className="mr-2 data-[orientation=vertical]:h-4" />
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem className="hidden md:block">
                <BreadcrumbLink href="#">
                  Log-off
                </BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem>
                <BreadcrumbPage>Dashboard</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </header>
        
        <div className="flex flex-1 flex-col gap-6 p-6 bg-white text-gray-900">
        {
          (pageState === possibleStates[0])
            ? <IdleComponent />
            : 
          (pageState === possibleStates[2]) // pentru ai
            ? <AiPromptingDashboard onBack = {()=> setState(possibleStates[0])}/>
            :
          <h1 className="text-3xl font-bold mb-4">
                Welcome to the <i>siren-net</i> Dashboard
              </h1>
        }
        </div>
      </SidebarInset>
    </SidebarProvider>)
  );
}
