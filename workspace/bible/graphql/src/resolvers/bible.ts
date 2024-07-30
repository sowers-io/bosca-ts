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

import { Resolvers } from '../generated/resolvers'
import { RequestContext } from '../context'
import { useClient } from '@bosca/common'
import { ContentService, FindCollectionRequest, FindMetadataRequest } from '@bosca/protobufs'
import { execute, getHeaders } from '../util/requests'

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    bibles: async (_, __, context) => {
      return await execute(async () => {
        const service = useClient(ContentService)
        const bibles = await service.findCollection(
          new FindCollectionRequest({
            attributes: {
              'bible.type': 'bible',
            },
          }),
          {
            headers: getHeaders(context),
          }
        )
        return bibles.collections.map((collection) => {
          return {
            metadata: {
              systemId: collection.attributes['bible.system.id'],
              version: collection.attributes['bible.version'],
              language: collection.attributes['bible.language'],
            },
            id: collection.id,
            abbreviation: collection.attributes['bible.abbreviation'],
            name: collection.name,
            books: [],
          }
        })
      })
    },
    verse: async (_, args, context) => {
      return await execute(async () => {
        const service = useClient(ContentService)
        const verses = await service.findMetadata(
          new FindMetadataRequest({
            attributes: {
              'bible.type': 'verse',
              'bible.system.id': args.systemId,
              // 'bible.version': args.version,
              'bible.verse.usfm': args.usfm,
            },
          }),
          {
            headers: getHeaders(context),
          }
        )
        if (verses.metadata.length === 0) {
          return null
        }
        return verses.metadata.map((metadata) => {
          return {
            metadata: {
              systemId: metadata.attributes['bible.system.id'],
              version: metadata.attributes['bible.version'],
              language: metadata.attributes['bible.language'],
            },
            id: metadata.id,
            usfm: metadata.attributes['bible.verse.usfm'],
            name: metadata.name,
          }
        })[0]
      })
    }
  },
  Bible: {
    books: async (bible, _, context) => {
      return await execute(async () => {
        const service = useClient(ContentService)
        const books = await service.findCollection(
          new FindCollectionRequest({
            attributes: {
              'bible.type': 'book',
              'bible.system.id': bible.metadata.systemId,
              'bible.version': bible.metadata.version,
            },
          }),
          {
            headers: getHeaders(context),
          }
        )
        books.collections.sort((a, b) => {
          const a1 = parseInt(a.attributes['bible.book.order'])
          const b1 = parseInt(b.attributes['bible.book.order'])
          return a1 - b1
        })
        return books.collections.map((collection) => {
          return {
            metadata: {
              systemId: collection.attributes['bible.system.id'],
              version: collection.attributes['bible.version'],
              language: collection.attributes['bible.language'],
            },
            id: collection.id,
            usfm: collection.attributes['bible.book.usfm'],
            name: collection.name,
            chapters: [],
          }
        })
      })
    },
  },
  Book: {
    chapters: async (book, _, context) => {
      return await execute(async () => {
        const service = useClient(ContentService)
        const books = await service.findMetadata(
          new FindMetadataRequest({
            attributes: {
              'bible.type': 'chapter',
              'bible.system.id': book.metadata.systemId,
              'bible.version': book.metadata.version,
              'bible.book.usfm': book.usfm,
            },
          }),
          {
            headers: getHeaders(context),
          }
        )
        books.metadata.sort((a, b) => {
          const a1 = parseInt(a.attributes['bible.chapter.order'])
          const b1 = parseInt(b.attributes['bible.chapter.order'])
          return a1 - b1
        })
        return books.metadata.map((metadata) => {
          return {
            metadata: {
              systemId: metadata.attributes['bible.system.id'],
              version: metadata.attributes['bible.version'],
              language: metadata.attributes['bible.language'],
            },
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
