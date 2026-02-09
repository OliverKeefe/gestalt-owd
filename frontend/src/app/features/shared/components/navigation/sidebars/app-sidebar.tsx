import * as React from "react"
import {
    Blocks,
    Calendar,
    File,
    MessageCircleQuestion,
    Search,
    Settings2,
    Image,
    Trash2,
    CircleUser,
    School,
    BriefcaseBusiness,
    FolderClosed,
} from "lucide-react"

import { NavFavorites } from "@/app/features/shared/components/navigation/nav-favorites.tsx"
import { NavMain } from "@/app/features/shared/components/navigation/nav-main.tsx"
import { NavSecondary } from "@/app/features/shared/components/navigation/nav-secondary.tsx"
import { NavWorkspaces } from "@/app/features/shared/components/navigation/nav-workspaces.tsx"
import { TeamSwitcher } from "@/app/features/shared/components/navigation/team-switcher.tsx"
import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar.tsx"
import { UploadDialog } from "@/app/features/shared/components/dialog/upload-dialog.tsx";
import {RestHandler} from "@/app/features/shared/api/rest/rest-handler.ts";
import { ScrollArea } from "@/components/ui/scroll-area"

interface FolderData {
  id: string;
  name: string;
  url: string;
  subFolder: string;
}

interface FavoritesData {
  id: string;
  name: string;
  url: string;
}

const data = {
  teams: [
    {
      name: "Personal",
      logo: CircleUser,
      plan: "Enterprise",
    },
    {
      name: "University Of Essex",
      logo: School,
      plan: "Startup",
    },
    {
      name: "Company ltd",
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
      title: "Files",
      url: "/",
      icon: FolderClosed,
      isActive: true,
    },
    {
      title: "Photos",
      url: "/photos",
      icon: Image,
    },
    {
      title: "Documents",
      url: "/documents",
      icon: File,
      badge: "10",
    },
  ],
  navSecondary: [
    {
      title: "Calendar",
      url: "#",
      icon: Calendar,
    },
    {
      title: "Settings",
      url: "/settings",
      icon: Settings2,
    },
    {
      title: "Apps",
      url: "#",
      icon: Blocks,
    },
    {
      title: "Rubbish",
      url: "#",
      icon: Trash2,
    },
    {
      title: "Help",
      url: "#",
      icon: MessageCircleQuestion,
    },
  ],
  favorites: [
    {
      name: "budget_tracker.xlsx",
      url: "#",
      emoji: "📄",
    },
    {
      name: "CV.docx",
      url: "#",
      emoji: "📄",
    },
    {
      name: "Assignment 1 Essay.docx",
      url: "#",
      emoji: "📄",
    },
  ],
  workspaces: [
    {
      name: "CE303 Study Group",
      emoji: "🧑‍🎓",
      pages: [
        {
          name: "Daily Journal & Reflection.note",
          url: "#",
          emoji: "📁",
        },
        {
          name: "List_of_localPubs.docx",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Contents Insurance FilesPage",
          url: "#",
          emoji: "📁",
        },
      ],
    },
    {
      name: "House Move",
      emoji: "🏡",
      pages: [
        {
          name: "Career Objectives & Milestones",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Skill Acquisition & Training Log",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Networking Contacts & Events",
          url: "#",
          emoji: "📁",
        },
      ],
    },
    {
      name: "Creative Projects",
      emoji: "🖌️",
      pages: [
        {
          name: "Writing Ideas & Story Outlines",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Art & Design Portfolio",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Music Composition & Practice Log",
          url: "#",
          emoji: "📁",
        },
      ],
    },
    {
      name: "Home Management",
      emoji: "🪴",
      pages: [
        {
          name: "Household Budget & Expense Tracking.docx",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Home Maintenance Schedule & Tasks.docx",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Family Calendar & Event Planning",
          url: "#",
          emoji: "📁",
        },
      ],
    },
    {
      name: "Travel & Holidays",
      emoji: "✈️",
      pages: [
        {
          name: "Trip Planning & Itineraries",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Travel Bucket List & Inspiration",
          url: "#",
          emoji: "📁",
        },
        {
          name: "Travel Journal & Photo Gallery",
          url: "#",
          emoji: "📁",
        },
      ],
    },
  ],
}

const client: RestHandler = new RestHandler(`http://localhost:8081`);

function GetFavorites(): Promise<FavoritesData[]> {
  return client.handleGet<FavoritesData[]>(`files/favorites`)
}

function GetFolders(): Promise<FolderData[]> {
  return client.handleGet<FolderData[]>(`files/folders`);
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
    return (
        <Sidebar className="border-r-0 left-16" {...props}>
            <SidebarHeader>
                <TeamSwitcher teams={data.teams} />
                <NavMain items={data.navMain} />
            </SidebarHeader>

            <SidebarContent className="min-h-0">
                <ScrollArea className="h-full [&_[data-radix-scroll-area-scrollbar]]:w-1" >
                    <NavFavorites favorites={data.favorites} />
                    <NavWorkspaces workspaces={data.workspaces} />
                    <NavSecondary items={data.navSecondary} className="mt-auto" />
                </ScrollArea>
            </SidebarContent>

            <SidebarRail />
        </Sidebar>
    )
}

