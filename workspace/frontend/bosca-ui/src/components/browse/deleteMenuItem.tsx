"use client"

import {DropdownMenuItem} from "@/components/ui/dropdown-menu";
import {onDelete} from "@/api/collections";

interface DeleteMenuItemProps {
  id: string
  collection: boolean
}

export function DeleteMenuItem({id, collection}: DeleteMenuItemProps) {
  return <DropdownMenuItem onSelect={async (ev) => {
    ev.preventDefault()
    ev.stopImmediatePropagation()
    await onDelete(id, collection)
    window.dispatchEvent(new Event('files-changed'))
  }}>Delete</DropdownMenuItem>
}