import { DataSource } from '@bosca/common'
import { Collection, Metadata, Source, Trait } from '@bosca/protobufs'
import { WorkflowDataSource } from './workflow'
import { Pool } from 'pg'

export const RootCollectionId = '00000000-0000-0000-0000-000000000000'

export class ContentDataSource extends DataSource {
  private readonly workflows: WorkflowDataSource

  constructor(pool: Pool, workflows: WorkflowDataSource) {
    super(pool)
    this.workflows = workflows
  }

  async getSources(): Promise<Source[]> {
    return await this.queryAndMap(() => new Source(), 'select id, name, description, configuration from sources')
  }

  async getSource(id: string): Promise<Source | null> {
    return await this.queryAndMapFirst(() => new Source(), 'select * from sources where id = $1::uuid', [id])
  }

  async getTraits(): Promise<Trait[]> {
    return await this.queryAndMap(
      () => new Trait(),
      'select id, name, description, array(select workflow_id from trait_workflows where trait_id = id) as workflow_ids from traits'
    )
  }

  async getTrait(id: string): Promise<Trait | null> {
    return await this.queryAndMapFirst(
      () => new Trait(),
      'select id, name, description, array(select workflow_id from trait_workflows where trait_id = id) as workflow_ids from traits where id = $1',
      [id]
    )
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
    collection.traitIds = (
      await this.query('select trait_id from collection_traits where collection_id = $1::uuid', [id])
    ).rows.map((r) => r.trait_id)
    return collection
  }

  async addRootCollection(): Promise<boolean> {
    const root = await this.getCollection(RootCollectionId)
    if (root) return false
    await this.workflows.initialize()
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
    return record.rows[0].id
  }

  async getMetadata(id: string): Promise<Metadata | null> {
    return await this.queryAndMapFirst(
      () => new Metadata(),
      'select * from metadata where id = $1::uuid',
      [id],
      (r) => {
        r.modified = r.modified.toISOString()
        r.created = r.created.toISOString()
      }
    )
  }

  async deleteMetadata(id: string): Promise<void> {
    await this.query('delete from metadata where id = $1::uuid', [id])
  }
}
