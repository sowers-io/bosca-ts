import { DataSource } from '@bosca/common'
import { Metadata, Source } from '@bosca/protobufs'

export class ContentDataSource extends DataSource {
  async getSources(): Promise<Source[]> {
    const result = await this.query('select id, name, description, configuration from sources')
    return result.rows.map((row) => Source.fromJson(row, { ignoreUnknownFields: true }))
  }

  async getSource(id: string): Promise<Source | null> {
    const record = await this.query('select * from sources where id = $1::uuid', [id])
    if (!record || !record.rowCount) return null
    return Source.fromJson(record.rows[0], { ignoreUnknownFields: true })
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
    const record = await this.query('select * from metadata where id = $1::uuid', [id])
    if (!record || !record.rowCount) return null
    const row = record.rows[0]
    row.modified = row.modified.toISOString()
    row.created = row.created.toISOString()
    return Metadata.fromJson(row, { ignoreUnknownFields: true })
  }

  async deleteMetadata(id: string): Promise<void> {
    await this.query('delete from metadata where id = $1::uuid', [id])
  }
}
