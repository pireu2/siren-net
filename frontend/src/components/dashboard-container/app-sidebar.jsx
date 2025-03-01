"use client"
import * as React from "react"
import {BarChart3, Home, Settings, Users  } from "lucide-react"
import { useState } from "react"

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
      url: "#",
      icon: Home,
    },
    {
      title: "Team",
      url: "#",
      icon: Users,
    },
    {
      title: "Analytics",
      url: "#",
      icon: BarChart3,
    },
  ],
}

export function AppSidebar({...props}) {
  const [activeSection, setActiveSection] = useState(data.navMain[0].title)


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
                onClick={() => setActiveSection(item.title)}>
                <a href={item.url}> {/* Aici se va pune Link penttru routing! */}
                  <item.icon className="mr-2 h-4 w-4" />
                  <span>{item.title}</span>
                </a>
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
