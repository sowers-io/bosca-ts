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
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.USXProcessor = void 0;
const decompress_1 = __importDefault(require("decompress"));
const xml2js_1 = require("xml2js");
const metadata_1 = require("./metadata");
const book_processor_1 = require("./book_processor");
const node_util_1 = require("node:util");
const fs = __importStar(require("node:fs"));
const readFile = (0, node_util_1.promisify)(fs.readFile);
class USXProcessor {
    metadata;
    books = [];
    async processBook(name, content, file) {
        const data = await readFile(file, 'utf-8');
        const processor = new book_processor_1.BookProcessor();
        return await processor.process(name, content, data.toString());
    }
    async processChapter(name, content, file) {
        const data = await readFile(file, 'utf-8');
        const processor = new book_processor_1.BookProcessor();
        const book = await processor.process(name, content, data.toString());
        return book.chapters[0];
    }
    async process(file) {
        const files = await (0, decompress_1.default)(file);
        const filesMap = {};
        for (const file of files) {
            if (file.path == 'metadata.xml') {
                const metadata = await (0, xml2js_1.parseStringPromise)(file.data);
                this.metadata = new metadata_1.BibleMetadata(metadata.DBLMetadata);
            }
            else {
                filesMap[file.path] = file;
            }
        }
        const processor = new book_processor_1.BookProcessor();
        for (const name of this.metadata.publication.names) {
            const content = this.metadata.publication.contents[name.id];
            const file = filesMap[content.file];
            const book = await processor.process(name, content, file.data.toString());
            this.books.push(book);
        }
    }
}
exports.USXProcessor = USXProcessor;
//# sourceMappingURL=processor.js.map