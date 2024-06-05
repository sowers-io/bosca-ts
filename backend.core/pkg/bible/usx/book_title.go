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

type BookTitleType int

const (
	BookTitleTypeUnknown BookTitleType = iota
	// BookTitleTypeMT is the main title of the book (if single level)
	BookTitleTypeMT
	// BookTitleTypeMT1 is the main title of the book (if single level)
	BookTitleTypeMT1
	// BookTitleTypeMT2 is the main title of the book (if multiple levels)
	BookTitleTypeMT2
	// BookTitleTypeMT3 is a secondary title usually occurring before the main title
	BookTitleTypeMT3
	// BookTitleTypeMT4 is a tertiary title occurring after the main title
	BookTitleTypeMT4
	// BookTitleTypeIMT is introduction major title, level 1 (if single level)
	BookTitleTypeIMT
	// BookTitleTypeIMT1 is introduction major title, level 1 (if multiple levels)
	BookTitleTypeIMT1
	// BookTitleTypeIMT2 is introduction major title, level 2
	BookTitleTypeIMT2
	// BookTitleTypeREM is a remark
	BookTitleTypeREM
)

func parseBookTitle(attr xml.Attr) BookTitleType {
	switch attr.Value {
	case "mt":
		return BookTitleTypeMT
	case "mt1":
		return BookTitleTypeMT1
	case "mt2":
		return BookTitleTypeMT2
	case "mt3":
		return BookTitleTypeMT3
	case "mt4":
		return BookTitleTypeMT4
	case "imt":
		return BookTitleTypeIMT
	case "imt1":
		return BookTitleTypeIMT1
	case "imt2":
		return BookTitleTypeIMT2
	case "rem":
		return BookTitleTypeREM
	default:
		return BookTitleTypeUnknown
	}
}
