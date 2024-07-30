"use server";

import "server-only";
import { useClient } from "@bosca/common";
import {
  AddMetadataRequest,
  ContentService,
  IdRequest,
  Metadata,
  SignedUrl,
} from "@bosca/protobufs";
import { protoInt64 } from "@bufbuild/protobuf";

export async function getUploadUrl(
  fileName: string,
  contentType: string,
  contentLength: number,
  collection: string
): Promise<SignedUrl> {
  const client = useClient(ContentService);
  const request = new AddMetadataRequest({
    metadata: new Metadata({
      contentType: contentType || "application/octet-stream",
      contentLength: protoInt64.parse(contentLength),
      name: fileName,
    }),
    collection: collection,
  });
  const response = await client.addMetadata(request);
  return await client.getMetadataUploadUrl(
    new IdRequest({ id: response.id })
  );
}
