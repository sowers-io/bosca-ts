import {TableCell, TableRow} from "@/components/ui/table";
import {Badge} from "@/components/ui/badge";
import moment from "moment/moment";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {MoreHorizontal} from "lucide-react";
import {DeleteMenuItem} from "@/components/browse/deleteMenuItem";
import {ReactNode} from "react";
import {ICollectionItem} from "@/api/collections";
import Link from "next/link";
import {FolderIcon} from "@heroicons/react/24/outline";
import {ReprocessMenuItem} from "@/components/browse/reprocessMenuItem";

interface FolderItemProps {
  item: ICollectionItem
}

function getDate(item: ICollectionItem): Date {
  return new Date(item.created * 1000)
}

function getLanguage(item: ICollectionItem): ReactNode {
  if (item.languageTag) {
    return <Badge variant="outline">{item.languageTag}</Badge>
  }
  return <span></span>
}

export function FolderItem({item}: FolderItemProps) {
  return (
    <TableRow key={item.id}>
      <TableCell className="font-medium">
        <Link href={item.collectionType ?
          '/dashboard/browse?collection=' + item.id :
          '/dashboard/browse/metadata?metadata=' + item.id
        } className="inline-flex">
          {item.collectionType ? <FolderIcon className="w-6 text-gray-400 mr-2"/> : <FolderIcon className="w-6 text-gray-400 mr-2 invisible"/>}
          <span className="m-4">{item.name}</span>
        </Link>
      </TableCell>
      <TableCell>
        {item.status ? <Badge variant="outline">{item.status}</Badge> : <></>}
      </TableCell>
      <TableCell>
        {item.workflowStateId ? <Badge variant="outline">{item.workflowStateId}</Badge> : <></>}
      </TableCell>
      <TableCell>
        {item.collectionType ?? item.contentType}
      </TableCell>
      <TableCell>
        {getLanguage(item)}
      </TableCell>
      <TableCell className="hidden md:table-cell">{moment(getDate(item)).startOf('hour').fromNow()}</TableCell>
      <TableCell>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button aria-haspopup="true" size="icon" variant="ghost">
              <MoreHorizontal className="h-4 w-4"/>
              <span className="sr-only">Toggle menu</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <ReprocessMenuItem id={item.id}/>
            <DeleteMenuItem id={item.id} collection={item.collectionType !== undefined}/>
          </DropdownMenuContent>
        </DropdownMenu>
      </TableCell>
    </TableRow>
  )
}