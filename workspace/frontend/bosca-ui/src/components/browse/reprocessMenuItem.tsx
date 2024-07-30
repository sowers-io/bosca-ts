"use client"

import {DropdownMenuItem} from "@/components/ui/dropdown-menu";
import {onDelete, onReprocess} from "@/api/collections";

interface ReprocessMenuItemProps {
  id: string
}

export function ReprocessMenuItem({id}: ReprocessMenuItemProps) {
  return <DropdownMenuItem onSelect={async (ev) => {
    ev.preventDefault()
    ev.stopImmediatePropagation()
    await onReprocess(id)
    window.dispatchEvent(new Event('files-changed'))
  }}>Reprocess</DropdownMenuItem>
}