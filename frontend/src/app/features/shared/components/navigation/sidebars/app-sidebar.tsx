import * as React from "react"
import {
    Search,
    Settings2,
    Trash2,
    CircleUser,
    School,
    BriefcaseBusiness,
    FolderClosed,
    Handshake, Laptop, Cloud,
} from "lucide-react"

import { NavMain } from "@/app/features/shared/components/navigation/nav-main.tsx"
import { NavSecondary } from "@/app/features/shared/components/navigation/nav-secondary.tsx"
import { TeamSwitcher } from "@/app/features/shared/components/navigation/team-switcher.tsx"
import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar.tsx"
import { ScrollArea } from "@/components/ui/scroll-area"
import type {MeterGaugeSegment} from "@/app/features/shared/components/gauges/meter-gauge.tsx";
import {MeterGauge} from "@/app/features/shared/components/gauges/meter-gauge.tsx";
import {Button} from "@/components/ui/button.tsx";

const data = {
  teams: [
    {
      name: "Personal",
      logo: CircleUser,
      plan: "Enterprise",
    },
    {
      name: "Org 1",
      logo: School,
      plan: "Startup",
    },
    {
      name: "Org 2",
      logo: BriefcaseBusiness,
      plan: "Free",
    },
  ],
  navMain: [
    {
      title: "Search",
      url: "#",
      icon: Search,
    },
    {
      title: "My Files",
      url: "/",
      icon: FolderClosed,
      isActive: true,
    },
    {
      title: "Devices",
      url: "/photos",
      icon: Laptop,
    },
    {
      title: "Shared With Me",
      url: "/documents",
      icon: Handshake,
      badge: "10",
    },
  ],
  navSecondary: [
      {
          title: "Rubbish",
          url: "#",
          icon: Trash2,
      },
      {
        title: "Settings",
        url: "/settings",
        icon: Settings2,
      },
      //{
      //  title: "Storage",
      //  url: "#",
      //  icon: Cloud,
      //},
   ],
}

const segdat: MeterGaugeSegment[] = [
    {
        label: "Photos",
        value: 200,
        color: "bg-yellow-500",
        percentage: 0.25,
    },
    {
        label: "Documents",
        value: 15.6,
        color: "bg-blue-500",
        percentage: 0.156,
    },
    {
        label: "Code Repos",
        value: 3.9,
        color: "bg-green-500",
        percentage: 0.039,
    }
];

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
    return (
        <Sidebar className="border-r-0 left-16" {...props}>
            <SidebarHeader>
                <TeamSwitcher teams={data.teams} />
                <NavMain items={data.navMain} />
            </SidebarHeader>
            <SidebarContent className="min-h-0">
                <ScrollArea className="h-full [&_[data-radix-scroll-area-scrollbar]]:w-1" >
                    <NavSecondary items={data.navSecondary} className="mt-auto" />
                </ScrollArea>
                <div className="p-1 ml-3 mr-3 pb-2">
                    <div className="flex flex-row items-center">
                        <Cloud className="mr-2 mb-2"/>
                        <p className="font-sans text-sm text-foreground mb-2">45 GB / 250 GB Remaining.</p>
                    </div>
                    <MeterGauge
                        segmentData={segdat}
                        total={250}
                    >
                    </MeterGauge>
                    <div className="p-1 mt-2">
                        <Button
                            className="mx-auto block rounded-full"
                            variant="outline">Get more storage
                        </Button>
                    </div>
                </div>
            </SidebarContent>
            <SidebarRail />
        </Sidebar>
    )
}
