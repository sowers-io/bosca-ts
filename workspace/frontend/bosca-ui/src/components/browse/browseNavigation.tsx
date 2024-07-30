"use client"

import {Button} from "@/components/ui/button"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import {ArrowUpIcon, Plus, RefreshCw} from "lucide-react";
import AddContent from "@/components/browse/addContent";
import UploadContent from "@/components/browse/uploadContent";
import {useRouter} from "next/navigation";

interface AddContentPopoverProps {
  collection: string
}

export function BrowseNavigation({collection}: AddContentPopoverProps) {
  const router = useRouter()
  return (
    <>
      {collection != '00000000-0000-0000-0000-000000000000' ?
        <Button onClick={() => router.back()} style={{"marginRight": "10px"}}>
          <ArrowUpIcon className="h-4 w-4"/>
        </Button> :
        <></>
      }
      <Button onClick={() => window.dispatchEvent(new Event('files-changed'))} style={{"marginRight": "10px"}}>
        <RefreshCw className="h-4 w-4"/>
      </Button>
      <Popover>
        <PopoverTrigger asChild>
          <Button>
            <Plus className="h-4 w-4"/>
          </Button>
        </PopoverTrigger>
        <PopoverContent style={{"width": "450px"}} className="mr-10">
          <AddContent collection={collection}/>
          <UploadContent collection={collection}/>
        </PopoverContent>
      </Popover>
    </>
  )
}
