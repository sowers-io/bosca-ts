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

package usx

import (
	"archive/zip"
	"encoding/xml"
	"os"
)

type Bundle struct {
	metadata *Metadata
	books    []*USX
}

func (b *Bundle) Metadata() *Metadata {
	return b.metadata
}

func (b *Bundle) Books() []*USX {
	return b.books
}

func OpenBundle(file *os.File) (*Bundle, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return nil, err
	}

	var metadataFile *zip.File

	contents := make(map[string]*zip.File)
	for _, file := range zipReader.File {
		if file.Name == "metadata.xml" {
			metadataFile = file
		} else {
			contents[file.Name] = file
		}
	}

	metadata, err := parseMetadata(metadataFile)
	if err != nil {
		return nil, err
	}

	structure := make(map[string]MetadataStructureContent)
	for _, content := range metadata.Publications.Publication.Structure.Content {
		structure[content.Role] = content
	}

	books := make([]*USX, len(metadata.Source.CanonicalContent.Book))
	for b, book := range metadata.Source.CanonicalContent.Book {
		contentStructure := structure[book.Code]
		content := contents[contentStructure.Src]
		books[b], err = parseUSX(content)
		if err != nil {
			return nil, err
		}
	}

	return &Bundle{
		metadata: metadata,
		books:    books,
	}, err
}

func parseMetadata(file *zip.File) (*Metadata, error) {
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	metadata := &Metadata{}
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func parseUSX(file *zip.File) (*USX, error) {
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	decoder := xml.NewDecoder(reader)
	return NewParser(decoder).Parse()
}
