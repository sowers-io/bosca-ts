"use client"

import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card"
import {Table, TableBody, TableHead, TableHeader, TableRow} from "@/components/ui/table"

import {FolderItem} from "@/components/browse/folderItem";
import {useEffect, useState} from "react";
import {getCollection, getItems, ICollectionItems} from "@/api/collections";
import {useRouter} from "next/navigation";
import {BrowseNavigation} from "@/components/browse/browseNavigation";

interface FoldersProps {
  collection: string
}

export default function Folders({collection}: FoldersProps) {
  const router = useRouter()
  const [current, setCurrent] = useState<string | null>(null)
  const [items, setItems] = useState<ICollectionItems | null>(null)

  useEffect(() => {
    async function fetchItems() {
      try {
        let collectionId = collection
        if (!collectionId) {
          collectionId = "00000000-0000-0000-0000-000000000000"
        }
        const item = await getCollection(collectionId)
        setCurrent(item?.name ?? "Root Collection")
        const data = await getItems(collectionId)
        setItems(data)
      } catch (e: any) {
        if (e && e.message && e.message.startsWith("16 UNAUTHENTICATED")) {
          router.replace("/account/login")
          return
        }
        throw e
      }
    }
    function onFilesChanged() {
      fetchItems()
    }
    if (!items) {
      fetchItems();
    }
    window.addEventListener('files-changed', onFilesChanged)
    return () => {
      window.removeEventListener('files-changed', onFilesChanged)
    }
  }, [router, collection, items, setItems]);

  return (
    <Card>
      <div className="relative float-right top-4 right-4">
        <BrowseNavigation key={'add-content-popover-' + collection} collection={collection}/>
      </div>
      <CardHeader>
        <CardTitle>
          {current ?? "..."}
        </CardTitle>
        <CardDescription>
          View and manage your content.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>State</TableHead>
              <TableHead>Type</TableHead>
              <TableHead>Language</TableHead>
              <TableHead className="hidden md:table-cell">Created</TableHead>
              <TableHead>
                <span className="sr-only">Actions</span>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items?.items.map((item) => <FolderItem key={item.id} item={item}/>)}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}
