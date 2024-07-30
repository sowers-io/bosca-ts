import {Sheet, SheetContent, SheetTrigger} from "@/components/ui/sheet";
import {Button} from "@/components/ui/button";
import {CircleUser, FolderTree, Home, LineChart, PanelLeft, Search} from "lucide-react";
import Link from "next/link";
import Image from "next/image";
import Bosca from "@/app/bosca.svg";
import DashboardBreadcrumb from "@/components/dashboard/breadcrumb";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import SearchButton from "@/components/dashboard/searchButton";

export default function Header() {
  return (
    <header className="sticky top-0 z-30 flex h-14 items-center gap-4 border-b bg-background px-4 sm:static sm:h-auto sm:border-0 sm:bg-transparent sm:px-6">
      <Sheet>
        <SheetTrigger asChild>
          <Button size="icon" variant="outline" className="sm:hidden">
            <PanelLeft className="h-5 w-5"/>
            <span className="sr-only">Toggle Menu</span>
          </Button>
        </SheetTrigger>
        <SheetContent side="left" className="sm:max-w-xs">
          <nav className="grid gap-6 text-lg font-medium">
            <Link href="#" className="group flex h-10 w-10 shrink-0 items-center justify-center gap-2 rounded-full bg-white text-lg font-semibold text-primary-foreground md:text-base">
              <Image src={Bosca} alt="Bosca" className="h-8 w-8 transition-all group-hover:scale-110"/>
              <span className="sr-only">Bosca</span>
            </Link>
            <Link href="/" className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground">
              <Home className="h-5 w-5"/>
              Dashboard
            </Link>
            <Link href="/dashboard/browse" className="flex items-center gap-4 px-2.5 text-foreground">
              <FolderTree className="h-5 w-5"/>
              Browse
            </Link>
            <Link href="/settings" className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground">
              <LineChart className="h-5 w-5"/>
              Settings
            </Link>
          </nav>
        </SheetContent>
      </Sheet>
      <DashboardBreadcrumb/>
      <SearchButton />
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            size="icon"
            className="overflow-hidden rounded-full">
            <CircleUser width={24} height={24} className="overflow-hidden stroke-lime-500 rounded-full"/>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>My Account</DropdownMenuLabel>
          <DropdownMenuSeparator/>
          <DropdownMenuItem>Settings</DropdownMenuItem>
          <DropdownMenuItem>Support</DropdownMenuItem>
          <DropdownMenuSeparator/>
          <DropdownMenuItem>Logout</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </header>
  )
}