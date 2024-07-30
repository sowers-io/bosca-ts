"use client"

import Folders from "@/components/browse/folders";
import {BrowseNavigation} from "@/components/browse/browseNavigation";
import {useEffect, useState} from "react";
import {useSearchParams} from "next/navigation";

interface SelectedCollection {
  collection: string
}

export default function Content() {
  const searchParams = useSearchParams()
  const [collection, setCollection] = useState<SelectedCollection>({
    collection: '00000000-0000-0000-0000-000000000000',
  })

  useEffect(() => {
    function onRouteChanged() {
      let collectionId = searchParams.get('collection')
      if (!collectionId) {
        collectionId = '00000000-0000-0000-0000-000000000000'
      }
      if (collectionId !== collection.collection) {
        setCollection({
          collection: collectionId
        })
      }
    }

    onRouteChanged();
  }, [searchParams, collection, setCollection]);

  return (
    <>
      <Folders key={'folders-' + collection.collection} collection={collection.collection}/>
    </>
  )
}
