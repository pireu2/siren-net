import { AppSidebar } from "@/components/dashboard-container/app-sidebar"
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
import ImageGeneratorDashboard from "../ai-prompting/sd-dashboard"
import { useState, useEffect, useContext } from 'react';
import TeamPage from "../team/teamPage"
import { AuthContext } from "../auth/auth-handler";
import { getCookie } from 'react-use-cookie';


export default function Dashboard() 
{

  const [pageState, setState] = useState(()=>{
    return localStorage.getItem('pageState') || "idle"
  });

  const possibleStates = {0:"idle", 1: "SD", 2:"AI", 3:"TP", 4:"AN"}
  const {logout} = useContext(AuthContext);

  
  useEffect(() => {
    localStorage.setItem('pageState', pageState);
  }, [pageState]);


  useEffect(() => { 
      const makeProtectedRequest = async () => { 
      const token = getCookie('token');
      console.log(`The cookie is ${token}`)
      if (!token) {
        console.error("No token found, user is not logged in");
        return;
      }
    
      try {
        const response = await fetch("protected/", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
        });
        
        const data = await response.json(); 
        if(data.error)
        {
          logout();
        }
      } catch (error) {
        console.error("Error making protected request:", error);
      }
    };
    makeProtectedRequest();
    }, []);

  function changeState(key)
  {
    setState(possibleStates[key])
  }

  const onLogout = ()=>{
    logout();
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


  function PlaceHolder()
  {
    return(
      <>
      <div>
        <h1 className="text-3xl font-bold mb-4">
          Welcome to the <i>siren-net</i> Dashboard
        </h1>
        <div>
          <button onClick = {()=> setState(possibleStates[0])}>
            Back
          </button>
        </div>
      </div>
    </>
    );
  }

  return (
    (<SidebarProvider>
      <AppSidebar webState = {pageState} onStateChange = {arg => changeState(arg)}/>
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
          <SidebarTrigger className="-ml-1" />
          <Separator orientation="vertical" className="mr-2 data-[orientation=vertical]:h-4" />
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem className="hidden md:block">
                <BreadcrumbLink onClick = {onLogout}>
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
            (pageState === possibleStates[1]) // pentru stable diffusion
            ? <ImageGeneratorDashboard onBack = {()=> setState(possibleStates[0])}/>
            :
            (pageState === possibleStates[2]) // pentru ai
            ? <AiPromptingDashboard onBack = {()=> setState(possibleStates[0])}/>
            :
            (pageState === possibleStates[3]) // pentru team page
            ? <TeamPage onBack = {()=> setState(possibleStates[0])}/>
            :
            <PlaceHolder />
        }
        </div>
      </SidebarInset>
    </SidebarProvider>)
  );
}
