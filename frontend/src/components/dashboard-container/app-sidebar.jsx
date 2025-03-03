"use client"
import * as React from "react"
import {BarChart3, Home, Settings, Users  } from "lucide-react"
import { useState , useEffect} from "react"

import { SearchForm } from "@/components/dashboard-container/search-form"
import {
} from "@/components/ui/collapsible"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"

const data = {
  navMain: [
    {
      title: "Homepage",
      url: "/",
      icon: Home,
      id: 0 //ca sa corespunda cu id-ul de stateuri din dashboard
    },
    {
      title: "Team",
      url: "/",
      icon: Users,
      id: 3 //ca sa corespunda cu id-ul de stateuri din dashboard
    },
    {
      title: "Analytics",
      url: "/",
      icon: BarChart3,
      id: 4 //ca sa corespunda cu id-ul de stateuri din dashboard
    },
  ],
}

export function AppSidebar({webState, onStateChange, ...props}) {
  const [activeSection, setActiveSection] = useState(data.navMain[0].title)

  function handleSectionChange(title)
  {
    setActiveSection(title)
  }


  useEffect(() => {
    console.log("S a schimbat team stateu")
    if(webState === "TP")
      handleSectionChange("Team") 
    else handleSectionChange("Homepage"); // Aici defapt ne luam dupa title cardurile din dashboard
    return () => {};
  }, [webState]);

  return (
    (<Sidebar {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton 
              size="lg" 
              asChild>
              <a href="#">
                <div
                  className="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg">
                  <img src="src\images\sirennet.webp" alt="Siren Net" className="size-6" />
                </div>
                <div className="flex flex-col gap-0.5 leading-none">
                  <span className="font-medium">siren-net</span>
                  <span className="">v1.0</span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
        <SearchForm />
      </SidebarHeader>
      <SidebarContent>
        <SidebarMenu className="space-y-2">
          {data.navMain.map((item) => (
            <SidebarMenuItem key={item.title}>
              <SidebarMenuButton
                asChild
                isActive={activeSection === item.title}
                onClick={() => {
                  handleSectionChange(item.title); 
                  onStateChange(item.id);
                  }}>
                <div> {/* Aici se va pune Link penttru routing! */}
                  <item.icon className="mr-2 h-4 w-4" />
                  <span>{item.title}</span>
                </div>
              </SidebarMenuButton>
            </SidebarMenuItem>
          ))}
        </SidebarMenu>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton asChild>
              <a href="#">
                <Settings className="mr-2 h-4 w-4" />
                <span>Settings</span>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>)
  );
}
