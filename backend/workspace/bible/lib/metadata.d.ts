export declare class ManifestName {
    private name;
    constructor(name: any);
    get id(): string;
    get abbreviation(): string;
    get short(): string;
    get long(): string;
}
export declare class MetadataSystemId {
    private systemId;
    constructor(systemId: any);
    get id(): string;
}
export declare class MetadataIdentification {
    private identification;
    readonly systemId: MetadataSystemId;
    constructor(identification: any);
    get name(): string;
    get nameLocal(): string;
    get description(): string;
    get abbreviation(): string;
    get abbreviationLocal(): string;
}
export declare class MetadataLanguage {
    private language;
    constructor(language: any);
    get iso(): string;
    get name(): string;
    get nameLocal(): string;
    get script(): string;
    get scriptCode(): string;
    get scriptDirection(): string;
}
export declare class Publication {
    readonly names: ManifestName[];
    readonly contents: {
        [id: string]: PublicationContent;
    };
    constructor(names: ManifestName[], publication: any);
}
export declare class PublicationContent {
    private readonly content;
    constructor(content: any);
    get id(): string;
    get file(): string;
    get usfm(): string;
}
export declare class BibleMetadata {
    readonly identification: MetadataIdentification;
    readonly language: MetadataLanguage;
    readonly publication: Publication;
    constructor(metadata: any);
}
//# sourceMappingURL=metadata.d.ts.map