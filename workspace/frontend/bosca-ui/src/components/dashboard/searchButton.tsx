"use client"

import {Search} from "lucide-react";
import {Button} from "@/components/ui/button";

export default function SearchButton() {
  return (
    <div className="relative ml-auto flex-1 md:grow-0">
      <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground"/>
      <Button className="w-full rounded-lg hover:bg-background bg-background text-gray-300 pl-8 md:w-[200px] lg:w-[320px]"
              onClick={() => window.dispatchEvent(new Event('search-box'))}>
        Search...
      </Button>
    </div>
  )
}