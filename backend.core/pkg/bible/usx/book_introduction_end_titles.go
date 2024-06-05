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

type BookIntroductionEndTitleType int

const (
	BookIntroductionEndTitleUnknown BookIntroductionEndTitleType = iota
	// BookIntroductionEndTitleMT The main title of the book (if single level)
	BookIntroductionEndTitleMT
	// BookIntroductionEndTitleMT1 The main title of the book (if multiple levels)
	BookIntroductionEndTitleMT1
	// BookIntroductionEndTitleMT2 A secondary title usually occurring before the main title
	BookIntroductionEndTitleMT2
	// BookIntroductionEndTitleMT3 A tertiary title occurring after the main title
	BookIntroductionEndTitleMT3
	// BookIntroductionEndTitleMT4 A small secondary title sometimes occuring within parentheses
	BookIntroductionEndTitleMT4
	// BookIntroductionEndTitleIMT Introduction major title, level 1 (if single level)
	BookIntroductionEndTitleIMT
	// BookIntroductionEndTitleIMT1 Introduction major title, level 1 (if multiple levels)
	BookIntroductionEndTitleIMT1
	// BookIntroductionEndTitleIMT2 Introduction major title, level 2
	BookIntroductionEndTitleIMT2
)

func parseBookIntroductionEndTitle(attr xml.Attr) BookIntroductionEndTitleType {
	switch attr.Value {
	case "mt":
		return BookIntroductionEndTitleMT
	case "mt1":
		return BookIntroductionEndTitleMT1
	case "mt2":
		return BookIntroductionEndTitleMT2
	case "mt3":
		return BookIntroductionEndTitleMT3
	case "mt4":
		return BookIntroductionEndTitleMT4
	case "imt":
		return BookIntroductionEndTitleIMT
	case "imt1":
		return BookIntroductionEndTitleIMT1
	case "imt2":
		return BookIntroductionEndTitleIMT2
	default:
		return BookIntroductionEndTitleUnknown
	}
}
