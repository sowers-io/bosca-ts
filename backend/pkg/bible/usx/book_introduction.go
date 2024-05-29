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

type BookIntroductionType int

const (
	BookIntroductionUnknown BookIntroductionType = iota
	// BookIntroductionIMT Introduction major title, level 1 - (if single level)
	BookIntroductionIMT
	// BookIntroductionIMT1  Introduction major title, level 1 (if multiple levels)
	BookIntroductionIMT1
	// BookIntroductionIMT2 Introduction major title, level 2
	BookIntroductionIMT2
	// BookIntroductionIMT3 Introduction major title, level 3
	BookIntroductionIMT3
	// BookIntroductionIMT4 Introduction major title, level 4 (usually within parenthesis)
	BookIntroductionIMT4
	// BookIntroductionIB Introduction blank line
	BookIntroductionIB
	// BookIntroductionIE Introduction ending marker
	BookIntroductionIE
	// BookIntroductionILI A list entry, level 1 (if single level)
	BookIntroductionILI
	// BookIntroductionILI1 A list entry, level 1 (if multiple levels)
	BookIntroductionILI1
	// BookIntroductionILI2 A list entry, level 2
	BookIntroductionILI2
	// BookIntroductionIMI Introduction prose paragraph, with no first line indent (may occur after poetry)
	BookIntroductionIMI
	// BookIntroductionIMQ Introduction prose paragraph text, indented, with no first line indent
	BookIntroductionIMQ
	// BookIntroductionIO Introduction prose paragraph, quote from the body text, with no first line indent
	BookIntroductionIO
	// BookIntroductionIO1 Introduction outline text, level 1 (if single level)
	BookIntroductionIO1
	// BookIntroductionIO2 Introduction outline text, level 1 (if multiple levels)
	BookIntroductionIO2
	// BookIntroductionIO3 Introduction outline text, level 2
	BookIntroductionIO3
	// BookIntroductionIO4 Introduction outline text, level 3
	BookIntroductionIO4
	// BookIntroductionIOT Introduction outline text, level 4
	BookIntroductionIOT
	// BookIntroductionIP Introduction outline title
	BookIntroductionIP
	// BookIntroductionIPI Introduction prose paragraph
	BookIntroductionIPI
	// BookIntroductionIPQ Introduction prose paragraph, indented, with first line indent
	BookIntroductionIPQ
	// BookIntroductionIPR Introduction prose paragraph, quote from the body text
	BookIntroductionIPR
	// BookIntroductionIQ Introduction prose paragraph, right aligned
	BookIntroductionIQ
	// BookIntroductionIQ1 Introduction poetry text, level 1 (if single level)
	BookIntroductionIQ1
	// BookIntroductionIQ2 Introduction poetry text, level 1 (if multiple levels)
	BookIntroductionIQ2
	// BookIntroductionIQ3 Introduction poetry text, level 2
	BookIntroductionIQ3
	// BookIntroductionIS Introduction poetry text, level 3
	BookIntroductionIS
	// BookIntroductionIS1 Introduction section heading, level 1 (if single level)
	BookIntroductionIS1
	// BookIntroductionIS2 Introduction section heading, level 1 (if multiple levels)
	BookIntroductionIS2
	// BookIntroductionIMTE Introduction major title at introduction end, level 1 (if single level)
	BookIntroductionIMTE
	// BookIntroductionIMTE1 Introduction major title at introduction end, level 1 (if multiple levels)
	BookIntroductionIMTE1
	// BookIntroductionIMTE2 Introduction major title at introduction end, level 2
	BookIntroductionIMTE2
	// BookIntroductionIEX Introduction explanatory or bridge text (e.g. explanation of missing book in Short Old Testament)
	BookIntroductionIEX
	// BookIntroductionREM Remark
	BookIntroductionREM
)

func parseBookIntroduction(attr xml.Attr) BookIntroductionType {
	switch attr.Value {
	case "imt":
		return BookIntroductionIMT
	case "imt1":
		return BookIntroductionIMT1
	case "imt2":
		return BookIntroductionIMT2
	case "imt3":
		return BookIntroductionIMT3
	case "imt4":
		return BookIntroductionIMT4
	case "ib":
		return BookIntroductionIB
	case "ie":
		return BookIntroductionIE
	case "ili":
		return BookIntroductionILI
	case "ili1":
		return BookIntroductionILI1
	case "ili2":
		return BookIntroductionILI2
	case "imi":
		return BookIntroductionIMI
	case "imq":
		return BookIntroductionIMQ
	case "io":
		return BookIntroductionIO
	case "io1":
		return BookIntroductionIO1
	case "io2":
		return BookIntroductionIO2
	case "io3":
		return BookIntroductionIO3
	case "io4":
		return BookIntroductionIO4
	case "iot":
		return BookIntroductionIOT
	case "ip":
		return BookIntroductionIP
	case "ipi":
		return BookIntroductionIPI
	case "ipq":
		return BookIntroductionIPQ
	case "ipr":
		return BookIntroductionIPR
	case "iq":
		return BookIntroductionIQ
	case "iq1":
		return BookIntroductionIQ1
	case "iq2":
		return BookIntroductionIQ2
	case "iq3":
		return BookIntroductionIQ3
	case "is":
		return BookIntroductionIS
	case "is1":
		return BookIntroductionIS1
	case "is2":
		return BookIntroductionIS2
	case "imte":
		return BookIntroductionIMTE
	case "imte1":
		return BookIntroductionIMTE1
	case "imte2":
		return BookIntroductionIMTE2
	case "iex":
		return BookIntroductionIEX
	case "rem":
		return BookIntroductionREM
	default:
		return BookIntroductionUnknown
	}
}
