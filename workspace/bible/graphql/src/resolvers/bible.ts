/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Resolvers, Verse } from '../generated/resolvers'
import { executeGraphQL, executeHttpRequest, getGraphQLHeaders, GraphQLRequestContext, useClient } from '@bosca/common'
import {
  ContentService,
  FindCollectionRequest,
  FindMetadataRequest,
  IdRequest,
  SupplementaryIdRequest,
} from '@bosca/protobufs'

interface BibleRequestContext extends GraphQLRequestContext {
  systemId: string
  version: string
  language: string
}

export const resolvers: Resolvers<BibleRequestContext> = {
  Query: {
    bibles: async (_, __, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const bibles = await service.findCollection(
          new FindCollectionRequest({
            attributes: {
              'bible.type': 'bible',
            },
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        return bibles.collections.map((collection) => {
          return {
            systemId: collection.attributes['bible.system.id'],
            version: collection.attributes['bible.version'],
            language: collection.attributes['bible.language'],
            id: collection.id,
            abbreviation: collection.attributes['bible.abbreviation'],
            name: collection.name,
            books: [],
          }
        })
      })
    },
    chapter: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const chapters = await service.findMetadata(
          new FindMetadataRequest({
            attributes: {
              'bible.type': 'chapter',
              'bible.system.id': args.systemId,
              'bible.version': args.version,
              'bible.chapter.usfm': args.usfm,
            },
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        if (chapters.metadata.length === 0) {
          return null
        }
        return chapters.metadata.map((metadata) => {
          return {
            id: metadata.id,
            usfm: metadata.attributes['bible.chapter.usfm'],
            name: metadata.name,
            number: metadata.attributes['bible.chapter.usfm'].split('.')[1],
          }
        })[0]
      })
    },
    verses: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const allVerses: Verse[] = []
        for (const usfm of args.usfm) {
          const verses = await service.findMetadata(
            new FindMetadataRequest({
              attributes: {
                'bible.type': 'verse',
                'bible.system.id': args.systemId,
                'bible.version': args.version,
                'bible.verse.usfm': usfm,
              },
            }),
            {
              headers: await getGraphQLHeaders(context),
            },
          )
          if (verses.metadata.length === 0) {
            return null
          }
          allVerses.push(
            verses.metadata.map((metadata) => {
              return {
                id: metadata.id,
                usfm: metadata.attributes['bible.verse.usfm'],
                name: metadata.name,
                number: metadata.attributes['bible.verse.usfm'].split('.')[2],
              }
            })[0],
          )
        }
        return allVerses
      })
    },
    verse: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const verses = await service.findMetadata(
          new FindMetadataRequest({
            attributes: {
              'bible.type': 'verse',
              'bible.system.id': args.systemId,
              'bible.version': args.version,
              'bible.verse.usfm': args.usfm,
            },
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        if (verses.metadata.length === 0) {
          return null
        }
        return verses.metadata.map((metadata) => {
          return {
            id: metadata.id,
            usfm: metadata.attributes['bible.verse.usfm'],
            name: metadata.name,
            number: metadata.attributes['bible.verse.usfm'].split('.')[2],
          }
        })[0]
      })
    },
  },
  Chapter: {
    html: async (chapter, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataSupplementaryDownloadUrl(
          new SupplementaryIdRequest({
            id: chapter.id,
            key: 'html',
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        const response = await executeHttpRequest(url)
        return response.toString()
      })
    },
    usx: async (chapter, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataDownloadUrl(
          new IdRequest({
            id: chapter.id,
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        const response = await executeHttpRequest(url)
        return response.toString()
      })
    },
  },
  Verse: {
    html: async (verse, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataSupplementaryDownloadUrl(
          new SupplementaryIdRequest({
            id: verse.id,
            key: 'html',
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        const response = await executeHttpRequest(url)
        return response.toString()
      })
    },
    text: async (verse, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataSupplementaryDownloadUrl(
          new SupplementaryIdRequest({
            id: verse.id,
            key: 'text',
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        const response = await executeHttpRequest(url)
        return response.toString()
      })
    },
    usx: async (verse, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataDownloadUrl(
          new IdRequest({
            id: verse.id,
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        const response = await executeHttpRequest(url)
        return response.toString()
      })
    },
  },
  BibleMetadata: {
    books: async (bible, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const books = await service.findMetadata(
          new FindMetadataRequest({
            attributes: {
              'bible.type': 'book',
              'bible.system.id': bible.systemId,
              'bible.version': bible.version,
            },
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        books.metadata.sort((a, b) => {
          const a1 = parseInt(a.attributes['bible.book.order'])
          const b1 = parseInt(b.attributes['bible.book.order'])
          return a1 - b1
        })
        context.systemId = bible.systemId
        context.version = bible.version
        return books.metadata.map((collection) => {
          return {
            id: collection.id,
            usfm: collection.attributes['bible.book.usfm'],
            name: collection.name,
            chapters: [],
          }
        })
      })
    },
  },
  BookMetadata: {
    chapters: async (book, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const books = await service.findMetadata(
          new FindMetadataRequest({
            attributes: {
              'bible.type': 'chapter',
              'bible.system.id': context.systemId,
              'bible.version': context.version,
              'bible.book.usfm': book.usfm,
            },
          }),
          {
            headers: await getGraphQLHeaders(context),
          },
        )
        books.metadata.sort((a, b) => {
          const a1 = parseInt(a.attributes['bible.chapter.order'])
          const b1 = parseInt(b.attributes['bible.chapter.order'])
          return a1 - b1
        })
        return books.metadata.map((metadata) => {
          return {
            id: metadata.id,
            usfm: metadata.attributes['bible.chapter.usfm'],
            name: metadata.name,
            verses: [],
          }
        })
      })
    },
  },
}
