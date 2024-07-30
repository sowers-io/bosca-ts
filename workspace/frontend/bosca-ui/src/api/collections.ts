"use server";

import "server-only";

import { useClient } from "@bosca/common";
import {
  AddCollectionRequest,
  Collection,
  CollectionItem,
  CollectionItems,
  CollectionType,
  ContentService,
  Empty,
  IdRequest,
} from "@bosca/protobufs";

export async function onDelete(id: string, collection: boolean) {
  const service = useClient(ContentService);
  const request = new IdRequest({ id: id });
  if (collection) {
    await service.deleteCollection(request);
  } else {
    await service.deleteMetadata(request);
  }
}

export async function onReprocess(id: string) {
  const service = useClient(ContentService);
  const request = new IdRequest({ id: id });
  await service.setMetadataReady(request);
}

export async function onCreateCollection(
  collection: string | undefined,
  name: string
) {
  const service = useClient(ContentService);
  await service.addCollection(
    new AddCollectionRequest({
      parent: collection,
      collection: new Collection({
        name: name,
        type: CollectionType.folder,
      }),
    })
  );
}

export interface ICollectionItems {
  items: ICollectionItem[];
}

export interface ICollectionItem {
  id: string;
  name: string;
  created: number;
  collectionType: string | undefined;
  contentType: string;
  languageTag: string | undefined;
  status: string | undefined;
  workflowStateId: string | undefined;
}

function getType(item: CollectionItem) {
  if (item.Item.case == "collection") {
    switch (item.Item.value.type) {
      case CollectionType.standard:
        return "Standard";
      case CollectionType.folder:
        return "Folder";
      case CollectionType.root:
        return "Root";
    }
  }
  return "Metadata";
}

function getLanguageTag(item: CollectionItem) {
  if (item.Item.case == "metadata") {
    return item.Item.value.languageTag || "en";
  }
  return undefined;
}

function getStatus(item: CollectionItem) {
  if (item.Item.case == "metadata") {
    return item.Item.value?.workflowStateId;
  }
  return undefined;
}

function mapTo(item: CollectionItem): ICollectionItem {
  switch (item.Item.case) {
    case "collection":
      return {
        id: item.Item.value?.id ?? "",
        name: item.Item.value?.name ?? "",
        created: Number(item.Item.value?.created ?? 0),
        collectionType: getType(item),
        contentType: "",
        workflowStateId: "",
        languageTag: getLanguageTag(item),
        status: getStatus(item),
      };
    case "metadata":
      return {
        id: item.Item.value.id ?? "",
        name: item.Item.value.name ?? "",
        created: Number(item.Item.value?.created ?? 0),
        collectionType: getType(item),
        contentType: item.Item.value?.contentType,
        languageTag: getLanguageTag(item),
        status: getStatus(item),
        workflowStateId: item.Item.value.workflowStateId,
      };
    default:
      throw new Error("unsupported case");
  }
}

export async function getCollection(
  id: string
): Promise<ICollectionItem | null> {
  const service = useClient(ContentService);
  const collection = await service.getCollection(new IdRequest({ id: id }));
  if (!collection) return null;
  return mapTo(
    new CollectionItem({ Item: { case: "collection", value: collection } })
  );
}

export async function getItems(
  id: string | null | undefined
): Promise<ICollectionItems> {
  const service = useClient(ContentService);
  const items: CollectionItems = !id
    ? await service.getRootCollectionItems(new Empty())
    : await service.getCollectionItems(new IdRequest({ id: id }));
  return {
    items: items?.items.map(mapTo) ?? [],
  };
}
