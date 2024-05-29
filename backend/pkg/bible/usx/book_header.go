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

import "encoding/xml"

type BookHeaderType int

const (
	BookHeaderTypeUnknown BookHeaderType = iota
	// BookHeaderTypeIDE File encoding information
	BookHeaderTypeIDE
	// BookHeaderTypeH Running header text for a book
	BookHeaderTypeH
	// BookHeaderTypeH1 Running header text (DEPRECATED)
	BookHeaderTypeH1
	// BookHeaderTypeH2 Running header text, left side of page (DEPRECATED)
	BookHeaderTypeH2
	// BookHeaderTypeH3 Running header text, right side of page (DEPRECATED)
	BookHeaderTypeH3
	// BookHeaderTypeTOC1 Long table of contents text
	BookHeaderTypeTOC1
	// BookHeaderTypeTOC2 Short table of contents text
	BookHeaderTypeTOC2
	// BookHeaderTypeTOC3 Book Abbreviation
	BookHeaderTypeTOC3
	// BookHeaderTypeTOCA1 Alternative language long table of contents text
	BookHeaderTypeTOCA1
	// BookHeaderTypeTOCA2 Alternative language short table of contents text
	BookHeaderTypeTOCA2
	// BookHeaderTypeTOCA3 Alternative language book Abbreviation
	BookHeaderTypeTOCA3
	// BookHeaderTypeREM Remark
	BookHeaderTypeREM
	// BookHeaderTypeUSFM USFM
	BookHeaderTypeUSFM
)

func parseBookHeader(attr xml.Attr) BookHeaderType {
	switch attr.Value {
	case "ide":
		return BookHeaderTypeIDE
	case "h":
		return BookHeaderTypeH
	case "h1":
		return BookHeaderTypeH1
	case "h2":
		return BookHeaderTypeH2
	case "h3":
		return BookHeaderTypeH3
	case "toc1":
		return BookHeaderTypeTOC1
	case "toc2":
		return BookHeaderTypeTOC2
	case "toc3":
		return BookHeaderTypeTOC3
	case "toca1":
		return BookHeaderTypeTOCA1
	case "toca2":
		return BookHeaderTypeTOCA2
	case "toca3":
		return BookHeaderTypeTOCA3
	case "rem":
		return BookHeaderTypeREM
	case "usfm":
		return BookHeaderTypeUSFM
	default:
		return BookHeaderTypeUnknown
	}
}
