"use client"

import Link from "next/link"

import {
  Breadcrumb,
  BreadcrumbEllipsis,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import {Button} from "@/components/ui/button"
import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "@/components/ui/drawer"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {useSelectedLayoutSegments} from "next/navigation";
import {useState} from "react";

const ITEMS_TO_DISPLAY = 3

export default function DashboardBreadcrumb() {
  const [open, setOpen] = useState(false)
  const isDesktop = true

  let segments = useSelectedLayoutSegments()

  if (segments.length === 0) {
    segments = ['dashboard']
  }
  if (segments[0] !== 'dashboard') {
    segments = ['dashboard'].concat(segments)
  }

  function getHref(segment: string) {
    if (segment === segments[segments.length - 1]) {
      return '#'
    }
    let href = '/'
    for (let i = 0; i < segments.length; i++) {
      href += segments[i] + '/'
      if (segments[i] === segment) {
        break
      }
    }
    return href
  }

  function getTitle(segment: string) {
    return segment.slice(0, 1).toLocaleUpperCase() + segment.slice(1)
  }

  return (
    <Breadcrumb>
      <BreadcrumbList>

      </BreadcrumbList>
    </Breadcrumb>
  )
}
