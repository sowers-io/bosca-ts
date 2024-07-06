export class ManifestName {
  private name: any

  constructor(name: any) {
    this.name = name
  }

  get id(): string {
    return this.name.$.id
  }

  get abbreviation(): string {
    return this.name.abbr
  }

  get short(): string {
    return this.name.short
  }

  get long(): string {
    return this.name.long
  }
}

export class MetadataSystemId {
  private systemId: any

  constructor(systemId: any) {
    this.systemId = systemId
  }

  get id(): string {
    for (const id of this.systemId) {
      if (id.$.type === 'paratext') {
        return id.id[0]
      }
    }
    throw new Error('unknown id')
  }
}

export class MetadataIdentification {
  private identification: any
  readonly systemId: MetadataSystemId

  constructor(identification: any) {
    this.identification = identification
    this.systemId = new MetadataSystemId(identification.systemId)
  }

  get name(): string {
    return this.identification.name[0]
  }

  get nameLocal(): string {
    return this.identification.nameLocal[0]
  }

  get description(): string {
    return this.identification.description[0]
  }

  get abbreviation(): string {
    return this.identification.abbreviation[0]
  }

  get abbreviationLocal(): string {
    return this.identification.abbreviationLocal[0]
  }
}

export class MetadataLanguage {
  private language: any

  constructor(language: any) {
    this.language = language
  }

  get iso(): string {
    return this.language.iso[0]
  }

  get name(): string {
    return this.language.name[0]
  }

  get nameLocal(): string {
    return this.language.nameLocal[0]
  }

  get script(): string {
    return this.language.script[0]
  }

  get scriptCode(): string {
    return this.language.scriptCode[0]
  }

  get scriptDirection(): string {
    return this.language.scriptDirection[0]
  }
}

export class Publication {
  readonly names: ManifestName[]
  readonly contents: { [id: string]: PublicationContent } = {}

  constructor(names: ManifestName[], publication: any) {
    this.names = names
    for (const content of publication.structure[0].content) {
      const p = new PublicationContent(content)
      this.contents[p.id] = p
    }
  }
}

export class PublicationContent {
  private readonly content: any

  constructor(content: any) {
    this.content = content
  }

  get id(): string {
    return this.content.$.name
  }

  get file(): string {
    return this.content.$.src
  }

  get usfm(): string {
    return this.content.$.role
  }
}

export class BibleMetadata {
  readonly identification: MetadataIdentification
  readonly language: MetadataLanguage
  readonly publication!: Publication

  constructor(metadata: any) {
    this.identification = new MetadataIdentification(metadata.identification[0])
    this.language = new MetadataLanguage(metadata.language[0])
    const names: ManifestName[] = []
    for (const name of metadata.names[0].name) {
      names.push(new ManifestName(name))
    }
    this.publication = new Publication(names, metadata.publications[0].publication[0])
  }
}