import { ManifestName, PublicationContent } from '../metadata'
import { Chapter } from './chapter'

export class Book {

  readonly name: ManifestName
  readonly content: PublicationContent
  readonly chapters: Chapter[] = []
  readonly raw: string

  constructor(name: ManifestName, content: PublicationContent, raw: string) {
    this.name = name
    this.content = content
    this.raw = raw
  }

  get usfm(): string {
    return this.content.usfm
  }
}
