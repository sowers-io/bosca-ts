import type {Metadata} from "next"
import "@/app/globals.css"

import {TooltipProvider} from "@/components/ui/tooltip"
import SideNavigation from "@/components/dashboard/sideNavigation";
import Header from "@/components/dashboard/header";
import {SearchBox} from "@/components/dashboard/searchBox";

export const metadata: Metadata = {
  title: "Bosca",
  description: "Bosca: The AI Content Management System",
  manifest: '/site.webmanifest',
}

export default function RootLayout({children}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
    <body>
    <TooltipProvider>
      <div className="flex min-h-screen w-full flex-col bg-muted/40">
        <SideNavigation/>
        <div className="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
          <Header/>
          <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
            {children}
          </main>
        </div>
      </div>
      <SearchBox/>
    </TooltipProvider>
    </body>
    </html>
  )
}
