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
	"encoding/xml"
	"errors"
)

type CrossReferenceType int

const (
	FootnoteTypeUnknown FootnoteType = iota
	FootnoteTypeF
	FootnoteTypeFE
	FootnoteTypeEF
)

const (
	CrossReferenceTypeUnknown CrossReferenceType = iota
	CrossReferenceTypeX
	CrossReferenceTypeXE
)

func parseFootnoteType(attr xml.Attr) FootnoteType {
	switch attr.Value {
	case "f":
		return FootnoteTypeF
	case "fe":
		return FootnoteTypeFE
	case "ef":
		return FootnoteTypeEF
	default:
		return FootnoteTypeUnknown
	}
}

func parseCrossReferenceType(attr xml.Attr) CrossReferenceType {
	switch attr.Value {
	case "x":
		return CrossReferenceTypeX
	case "xe":
		return CrossReferenceTypeXE
	default:
		return CrossReferenceTypeUnknown
	}
}

func (p *Parser) onNote(noteTag xml.StartElement, current NodeContainer) (bool, error) {
	for _, attr := range noteTag.Attr {
		if attr.Name.Local == "style" {
			ft := parseFootnoteType(attr)
			if ft != FootnoteTypeUnknown {
				return p.onFootnote(noteTag, ft, current)
			}
			crt := parseCrossReferenceType(attr)
			if crt != CrossReferenceTypeUnknown {
				return p.onCrossReference(noteTag, crt, current)
			}
		}
	}
	return false, errors.New("onNote: unexpected note type")
}
