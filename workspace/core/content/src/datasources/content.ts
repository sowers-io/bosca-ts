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

import { DataSource, Subject } from '@bosca/common'
import {
  Collection,
  CollectionType,
  Metadata,
  MetadataRelationship,
  MetadataSupplementary,
  Source,
  Trait
} from '@bosca/protobufs'
import { proto3 } from '@bufbuild/protobuf'

export const RootCollectionId = '00000000-0000-0000-0000-000000000000'

export interface IdName {
  id: string
  name: string
}

export class ContentDataSource extends DataSource {
  async getSources(): Promise<Source[]> {
    return await this.queryAndMap(() => new Source(), 'select id, name, description, configuration from sources')
  }

  async getSource(id: string): Promise<Source | null> {
    try {
      const source = await this.queryAndMapFirst(() => new Source(), 'select * from sources where id = $1::uuid', [id])
      if (source) return source
    } catch (ignore) {}
    return await this.queryAndMapFirst(() => new Source(), 'select * from sources where name = $1', [id])
  }

  async getTraits(): Promise<Trait[]> {
    return await this.queryAndMap(
      () => new Trait(),
      'select id, name, description, array(select workflow_id from trait_workflows where trait_id = id) as workflow_ids from traits'
    )
  }

  async getTraitWorkflowIds(traitId: string): Promise<string[]> {
    const records = await this.query('select workflow_id from trait_workflows where trait_id = $1', [traitId])
    return records.rows.map((r) => r.workflow_id)
  }

  async getTrait(id: string): Promise<Trait | null> {
    return await this.queryAndMapFirst(
      () => new Trait(),
      'select id, name, description, array(select workflow_id from trait_workflows where trait_id = id) as workflow_ids from traits where id = $1',
      [id]
    )
  }

  async getCollectionIdNames(collectionId: string): Promise<IdName[]> {
    const records = await this.query(
      'select id, name from collections where id in (select child_id from collection_collection_items where collection_id = $1)',
      [collectionId]
    )
    return records.rows.map((r) => ({ id: r.id, name: r.name }))
  }

  async getCollectionIdName(collectionId: string, name: string): Promise<IdName[]> {
    const records = await this.query(
      'select id, name from collections where id in (select child_id from collection_collection_items where collection_id = $1) and lower(name) = lower($2)',
      [collectionId, name]
    )
    return records.rows.map((r) => ({ id: r.id, name: r.name }))
  }

  async getMetadataIdNames(collectionId: string) {
    const records = await this.query(
      'select id, name from metadata where id in (select metadata_id from collection_metadata_items where collection_id = $1)',
      [collectionId]
    )
    return records.rows.map((r) => ({ id: r.id, name: r.name }))
  }

  async getMetadataIdName(collectionId: string, name: string): Promise<IdName[]> {
    const records = await this.query(
      'select id, name from metadata where id in (select metadata_id from collection_metadata_items where collection_id = $1) and lower(name) = lower($2)',
      [collectionId, name]
    )
    return records.rows.map((r) => ({ id: r.id, name: r.name }))
  }

  async getCollectionCollectionItemIds(id: string): Promise<string[]> {
    const records = await this.query('select child_id from collection_collection_items where collection_id = $1', [id])
    return records.rows.map((r) => r.child_id)
  }

  async getCollectionMetadataItemIds(id: string): Promise<string[]> {
    const records = await this.query('select metadata_id from collection_metadata_items where collection_id = $1', [id])
    return records.rows.map((r) => r.metadata_id)
  }

  private buildFindWhere(attributes: { [key: string]: string }): string {
    let where = ''
    let i = 1
    for (const _ in attributes) {
      if (where.length > 0) {
        where += ' and '
      }
      where += `attributes->>(\$${i}::varchar) = \$${i + 1}::varchar`
      i += 2
    }
    return where
  }

  private async mapCollection(collection: Collection) {
    collection.traitIds = (
      await this.query('select trait_id from collection_traits where collection_id = $1::uuid', [collection.id])
    ).rows.map((r) => r.trait_id)
  }

  async findCollection(attributes: { [key: string]: string }): Promise<Collection[]> {
    let args = []
    for (const key in attributes) {
      args.push(key)
      args.push(attributes[key])
    }
    const query = 'select * from collections where ' + this.buildFindWhere(attributes)
    const collections = await this.queryAndMap(
      () => new Collection(),
      query,
      args,
      (r) => {
        r.modified = r.modified.toISOString()
        r.created = r.created.toISOString()
      }
    )
    for (const collection of collections) {
      await this.mapCollection(collection)
    }
    return collections
  }

  async addCollection(collection: Collection): Promise<string> {
    const CollectionTypeEnum = proto3.getEnumType(CollectionType)
    const records = await this.query(
      'insert into collections (name, type, labels, attributes) values ($1, $2, $3, ($4)::jsonb) returning id',
      [collection.name, CollectionTypeEnum.findNumber(collection.type)?.name, collection.labels, JSON.stringify(collection.attributes)]
    )
    const collectionId = records.rows[0].id
    for (const traitId in collection.traitIds) {
      await this.query('insert into collection_traits (collection_id, trait_id) values ($1, $2)', [
        collectionId,
        traitId,
      ])
    }
    return collectionId
  }

  async addCollectionCollectionItem(collectionId: string, itemId: string): Promise<void> {
    await this.query('insert into collection_collection_items (collection_id, child_id) values ($1::uuid, $2::uuid)', [
      collectionId,
      itemId,
    ])
  }

  async addCollectionMetadataItem(collectionId: string, itemId: string): Promise<void> {
    await this.query('insert into collection_metadata_items (collection_id, metadata_id) values ($1::uuid, $2::uuid)', [
      collectionId,
      itemId,
    ])
  }

  async deleteCollection(id: string): Promise<void> {
    await this.query('delete from collections where id = $1::uuid', [id])
  }

  async getCollection(id: string): Promise<Collection | null> {
    const collection = await this.queryAndMapFirst(
      () => new Collection(),
      'select * from collections where id = $1::uuid',
      [id],
      (r) => {
        r.modified = r.modified.toISOString()
        r.created = r.created.toISOString()
      }
    )
    if (!collection) return null
    await this.mapCollection(collection)
    return collection
  }

  async addRootCollection(): Promise<boolean> {
    const root = await this.getCollection(RootCollectionId)
    if (root) return false
    await this.query(
      "insert into collections (id, name, type, workflow_state_id) values ($1, 'Root', 'root', 'published')",
      [RootCollectionId]
    )
    return true
  }

  async addMetadata(metadata: Metadata): Promise<string> {
    const record = await this.query(
      'insert into metadata (name, content_type, content_length, labels, attributes, source_id, source_identifier, language_tag) values ($1, $2, $3, $4, ($5)::jsonb, $6, $7, $8) returning id',
      [
        metadata.name,
        metadata.contentType,
        metadata.contentLength ? Number(metadata.contentLength) : null,
        metadata.labels,
        JSON.stringify(metadata.attributes),
        metadata.sourceId,
        metadata.sourceIdentifier,
        metadata.languageTag,
      ]
    )
    const metadataId = record.rows[0].id
    for (const traitId of metadata.traitIds) {
      await this.addMetadataTrait(metadataId, traitId)
    }
    for (const categoryId of metadata.categoryIds) {
      await this.addMetadataCategory(metadataId, categoryId)
    }
    return metadataId
  }

  async addMetadataTrait(metadataId: string, traitId: string): Promise<void> {
    await this.query('insert into metadata_traits (metadata_id, trait_id) values ($1, $2)', [metadataId, traitId])
  }

  async addMetadataCategory(metadataId: string, categoryId: string): Promise<void> {
    await this.query('insert into metadata_categories (metadata_id, category_id) values ($1, $2)', [
      metadataId,
      categoryId,
    ])
  }

  async findMetadata(attributes: { [key: string]: string }): Promise<Metadata[]> {
    let args = []
    for (const key in attributes) {
      args.push(key)
      args.push(attributes[key])
    }
    return await this.queryAndMap(
      () => new Metadata(),
      'select * from metadata where ' + this.buildFindWhere(attributes),
      args,
      (r) => {
        r.modified = r.modified.toISOString()
        r.created = r.created.toISOString()
      }
    )
  }

  async getMetadata(id: string): Promise<Metadata | null> {
    const metadata = await this.queryAndMapFirst(
      () => new Metadata(),
      'select * from metadata where id = $1::uuid',
      [id],
      (r) => {
        r.modified = r.modified.toISOString()
        r.created = r.created.toISOString()
      }
    )
    if (!metadata) return null
    metadata.traitIds = (
      await this.query('select trait_id from metadata_traits where metadata_id = $1::uuid', [id])
    ).rows.map((r) => r.trait_id)
    metadata.categoryIds = (
      await this.query('select category_id from metadata_categories where metadata_id = $1::uuid', [id])
    ).rows.map((r) => r.category_id)
    return metadata
  }

  async addMetadataRelationship(metadataId1: string, metadataId2: string, relationship: string): Promise<void> {
    await this.query(
      'insert into metadata_relationship (metadata1_id, metadata2_id, relationship) values ($1, $2, $3)',
      [metadataId1, metadataId2, relationship]
    )
  }

  async getMetadataRelationships(metadataId: string, relationship: string): Promise<MetadataRelationship[]> {
    return await this.queryAndMap(
      () => new MetadataRelationship(),
      'select * from metadata_relationship where relationship = $1 and (metadata1_id = $2 or metadata2_id = $2)',
      [relationship, metadataId]
    )
  }

  async addMetadataSupplementary(
    metadataId: string,
    key: string,
    name: string,
    contentType: string,
    contentLength: number,
    traitIds: string[],
    sourceId: string | null,
    sourceIdentifier: string | null
  ) {
    await this.query(
      'insert into metadata_supplementary (metadata_id, "key", name, content_type, content_length, source_id, source_identifier) values ($1::uuid, $2, $3, $4, $5, $6, $7)',
      [metadataId, key, name, contentType, contentLength, sourceId, sourceIdentifier]
    )
    for (const traitId of traitIds) {
      await this.query('insert into metadata_supplementary_traits (metadata_id, key, trait_id) values ($1, $2, $3)', [
        metadataId,
        key,
        traitId,
      ])
    }
  }

  async getMetadataSupplementary(metadataId: string, key: string): Promise<MetadataSupplementary | null> {
    const supplementary = await this.queryAndMapFirst(
      () => new MetadataSupplementary(),
      'select * from metadata_supplementary where metadata_id = $1::uuid and "key" = $2',
      [metadataId, key]
    )
    if (!supplementary) return null
    supplementary.traitIds = (
      await this.query('select trait_id from metadata_supplementary_traits where metadata_id = $1::uuid and "key" = $2', [metadataId, key])
    ).rows.map((r) => r.trait_id)
    return supplementary
  }

  async getMetadataSupplementaries(metadataId: string): Promise<MetadataSupplementary[]> {
    return this.queryAndMap(
      () => new MetadataSupplementary(),
      'select "key", name, content_type, content_length, source_id, source_identifier from metadata_supplementary where metadata_id = $1::uuid',
      [metadataId]
    )
  }

  async setMetadataSupplementaryReady(metadataId: string, key: string): Promise<void> {
    await this.query('update metadata_supplementary set uploaded = now() where metadata_id = $1::uuid and "key" = $2', [
      metadataId,
      key,
    ])
  }

  async deleteMetadataSupplementary(metadataId: string, key: string): Promise<void> {
    await this.query('delete from metadata_supplementary where metadata_id = $1::uuid and "key" = $2', [
      metadataId,
      key,
    ])
  }

  async deleteMetadata(id: string): Promise<void> {
    await this.query('delete from metadata where id = $1::uuid', [id])
  }

  async setWorkflowState(
    subject: Subject,
    metadataId: string,
    fromStateId: string,
    toStateId: string,
    status: string,
    success: boolean,
    complete: boolean
  ): Promise<void> {
    await this.transaction(async (client) => {
      await client.query(
        'insert into metadata_workflow_transition_history (metadata_id, from_state_id, to_state_id, subject, status, success, complete) values ($1::uuid, $2, $3, $4, $5, $6, $7)',
        [metadataId, fromStateId, toStateId, subject.id, status, success, complete]
      )
      if (!success) {
        await client.query('update metadata set workflow_state_pending_id = null where id = $1::uuid', [metadataId])
      } else {
        if (complete) {
          await client.query(
            'update metadata set workflow_state_id = $1, workflow_state_pending_id = null where id = $2::uuid',
            [toStateId, metadataId]
          )
        } else {
          await client.query('update metadata set workflow_state_pending_id = $1 where id = $2::uuid', [
            toStateId,
            metadataId,
          ])
        }
      }
    })
  }
}
