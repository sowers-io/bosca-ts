"use strict";
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
Object.defineProperty(exports, "__esModule", { value: true });
exports.BibleMetadata = exports.PublicationContent = exports.Publication = exports.MetadataLanguage = exports.MetadataIdentification = exports.MetadataSystemId = exports.ManifestName = void 0;
class ManifestName {
    name;
    constructor(name) {
        this.name = name;
    }
    get id() {
        return this.name.$.id;
    }
    get abbreviation() {
        return this.name.abbr;
    }
    get short() {
        return this.name.short;
    }
    get long() {
        return this.name.long;
    }
}
exports.ManifestName = ManifestName;
class MetadataSystemId {
    systemId;
    constructor(systemId) {
        this.systemId = systemId;
    }
    get id() {
        for (const id of this.systemId) {
            if (id.$.type === 'paratext') {
                return id.id[0];
            }
        }
        throw new Error('unknown id');
    }
}
exports.MetadataSystemId = MetadataSystemId;
class MetadataIdentification {
    identification;
    systemId;
    constructor(identification) {
        this.identification = identification;
        this.systemId = new MetadataSystemId(identification.systemId);
    }
    get name() {
        return this.identification.name[0];
    }
    get nameLocal() {
        return this.identification.nameLocal[0];
    }
    get description() {
        return this.identification.description[0];
    }
    get abbreviation() {
        return this.identification.abbreviation[0];
    }
    get abbreviationLocal() {
        return this.identification.abbreviationLocal[0];
    }
}
exports.MetadataIdentification = MetadataIdentification;
class MetadataLanguage {
    language;
    constructor(language) {
        this.language = language;
    }
    get iso() {
        return this.language.iso[0];
    }
    get name() {
        return this.language.name[0];
    }
    get nameLocal() {
        return this.language.nameLocal[0];
    }
    get script() {
        return this.language.script[0];
    }
    get scriptCode() {
        return this.language.scriptCode[0];
    }
    get scriptDirection() {
        return this.language.scriptDirection[0];
    }
}
exports.MetadataLanguage = MetadataLanguage;
class Publication {
    names;
    contents = {};
    constructor(names, publication) {
        this.names = names;
        for (const content of publication.structure[0].content) {
            const p = new PublicationContent(content);
            this.contents[p.id] = p;
        }
    }
}
exports.Publication = Publication;
class PublicationContent {
    content;
    constructor(content) {
        this.content = content;
    }
    get id() {
        return this.content.$.name;
    }
    get file() {
        return this.content.$.src;
    }
    get usfm() {
        return this.content.$.role;
    }
}
exports.PublicationContent = PublicationContent;
class BibleMetadata {
    identification;
    language;
    publication;
    constructor(metadata) {
        this.identification = new MetadataIdentification(metadata.identification[0]);
        this.language = new MetadataLanguage(metadata.language[0]);
        const names = [];
        for (const name of metadata.names[0].name) {
            names.push(new ManifestName(name));
        }
        this.publication = new Publication(names, metadata.publications[0].publication[0]);
    }
}
exports.BibleMetadata = BibleMetadata;
//# sourceMappingURL=metadata.js.map