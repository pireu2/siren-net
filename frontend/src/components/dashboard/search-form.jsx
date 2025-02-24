import { Search } from "lucide-react"

import { Label } from "@/components/ui/label"
import {
  SidebarGroup,
  SidebarGroupContent,
  SidebarInput,
} from "@/components/ui/sidebar"

export function SearchForm({
  ...props
}) {
  return (
    // aici am comentat pentru ca nu vreau searchbar, da nici sa distrug asta, idk il las asa
    // (<form {...props}>
    //   <SidebarGroup className="py-0">
    //     <SidebarGroupContent className="relative">
    //       <Label htmlFor="search" className="sr-only">
    //         Search
    //       </Label>
    //       <SidebarInput id="search" placeholder="Search the docs..." className="pl-8" />
    //       <Search
    //         className="pointer-events-none absolute top-1/2 left-2 size-4 -translate-y-1/2 opacity-50 select-none" />
    //     </SidebarGroupContent>
    //   </SidebarGroup>
    // </form>)
    <></>
  );
}
