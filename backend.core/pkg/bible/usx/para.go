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
	"fmt"
)

type ParaStyleType int

type Para struct {
	Style    ParaStyleType
	Children []Node
}

func (p *Para) AddText(text string) error {
	if p.Children == nil {
		p.Children = append(p.Children, &VerseText{Text: text})
		return nil
	}
	return p.Children[len(p.Children)-1].AddText(text)
}

func (p *Para) AddNode(node Node) {
	if p.Children == nil {
		p.Children = make([]Node, 0)
		p.Children = append(p.Children, node)
	} else {
		switch nodeToAdd := node.(type) {
		case *Verse:
			p.Children = append(p.Children, nodeToAdd)
		default:
			switch lastChild := p.Children[len(p.Children)-1].(type) {
			case *Verse:
				lastChild.AddNode(node)
			default:
				p.Children = append(p.Children, node)
			}
		}
	}
}

func (p *Para) GetChildren() []Node {
	return p.Children
}

func (p *Parser) onParaStandard(chapterTag xml.StartElement, current NodeContainer) (bool, error) {
	para := &Para{}
	current.AddNode(para)
	for _, attr := range chapterTag.Attr {
		switch attr.Name.Local {
		case "style":
			para.Style = parseParaStyle(attr)
		}
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.StartElement:
			if handled, err := p.onChapterContent(element, para); err != nil {
				return false, err
			} else if !handled {
				return false, nil
			}
		case xml.CharData:
			if err = para.AddText(string(element)); err != nil {
				return false, err
			}
		case xml.EndElement:
			if element.Name.Local != "para" {
				return false, fmt.Errorf("onParaStandard: unexpected tag: %s", element.Name.Local)
			}
			return true, nil
		default:
			return false, fmt.Errorf("onParaStandard: unexpected element: %v", element)
		}
	}
}

const (
	ParaStyleTypeUnknown ParaStyleType = iota
	// ParaStyleTypeRESTORE is a comment about when text was restored
	ParaStyleTypeRESTORE
	// ParaStyleTypeCLS is closure of an Epistle
	ParaStyleTypeCLS
	// ParaStyleTypeIEX is introduction explanatory or bridge text (e.g. explanation of missing book in Short Old Testament)
	ParaStyleTypeIEX
	// ParaStyleTypeIP is division or Section introductory paragraph (study Bible)
	ParaStyleTypeIP
	// ParaStyleTypeLIT is for a comment or note inserted for liturgical use
	ParaStyleTypeLIT
	// ParaStyleTypeM is paragraph text, with no first line indent (may occur after poetry) aka: Paragraph Continuation
	ParaStyleTypeM
	// ParaStyleTypeMI is paragraph text, indented, with no first line indent; often used for discourse
	ParaStyleTypeMI
	// ParaStyleTypeNB is paragraph text, with no break from previous paragraph text (at chapter boundary)
	ParaStyleTypeNB
	// ParaStyleTypeP is paragraph text, with first line indent
	ParaStyleTypeP
	// ParaStyleTypePB is page Break used for new reader portions and children's bibles where content is controlled by the page
	ParaStyleTypePB
	// ParaStyleTypePC is inscription (paragraph text centered)
	ParaStyleTypePC
	// ParaStyleTypePI is paragraph text, level 1 indent (if single level), with first line indent; often used for discourse
	ParaStyleTypePI
	// ParaStyleTypePI1 is paragraph text, level 1 indent (if multiple levels), with first line indent; often used for discourse
	ParaStyleTypePI1
	// ParaStyleTypePI2 is paragraph text, level 2 indent, with first line indent; often used for discourse
	ParaStyleTypePI2
	// ParaStyleTypePI3 is paragraph text, level 3 indent, with first line indent; often used for discourse
	ParaStyleTypePI3
	// ParaStyleTypePO is letter opening
	ParaStyleTypePO
	// ParaStyleTypePR is text refrain (paragraph right-aligned)
	ParaStyleTypePR
	// ParaStyleTypePMO is embedded text opening
	ParaStyleTypePMO
	// ParaStyleTypePM is embedded text paragraph
	ParaStyleTypePM
	// ParaStyleTypePMC is embedded text closing
	ParaStyleTypePMC
	// ParaStyleTypePMR is embedded text refrain
	ParaStyleTypePMR
	// ParaStyleTypePH is paragraph text, with level 1 hanging indent (if single level) (DEPRECATED - use para@style li#)
	ParaStyleTypePH
	// ParaStyleTypePH1 is paragraph text, with level 1 hanging indent (if multiple levels)
	ParaStyleTypePH1
	// ParaStyleTypePH2 is paragraph text, with level 2 hanging indent
	ParaStyleTypePH2
	// ParaStyleTypePH3 is paragraph text, with level 3 hanging indent
	ParaStyleTypePH3
	// ParaStyleTypeQ is poetry text, level 1 indent (if single level)
	ParaStyleTypeQ
	// ParaStyleTypeQ1 is poetry text, level 1 indent (if multiple levels)
	ParaStyleTypeQ1
	// ParaStyleTypeQ2 is poetry text, level 2 indent
	ParaStyleTypeQ2
	// ParaStyleTypeQ3 is poetry text, level 3 indent
	ParaStyleTypeQ3
	// ParaStyleTypeQ4 is poetry text, level 4 indent
	ParaStyleTypeQ4
	// ParaStyleTypeQA is poetry text, Acrostic marker/heading
	ParaStyleTypeQA
	// ParaStyleTypeQC is poetry text, centered
	ParaStyleTypeQC
	// ParaStyleTypeQR is poetry text, Right Aligned
	ParaStyleTypeQR
	// ParaStyleTypeQM is poetry text, embedded, level 1 indent (if single level)
	ParaStyleTypeQM
	// ParaStyleTypeQM1 is poetry text, embedded, level 1 indent (if multiple levels)
	ParaStyleTypeQM1
	// ParaStyleTypeQM2 is poetry text, embedded, level 2 indent
	ParaStyleTypeQM2
	// ParaStyleTypeQM3 is poetry text, embedded, level 3 indent
	ParaStyleTypeQM3
	// ParaStyleTypeQD is a Hebrew musical performance annotation, similar in content to Hebrew descriptive title.
	ParaStyleTypeQD
	// ParaStyleTypeB is poetry text stanza break (e.g. stanza break)
	ParaStyleTypeB
	// ParaStyleTypeD is a Hebrew text heading, to provide description (e.g. Psalms)
	ParaStyleTypeD
	// ParaStyleTypeMS is a major section division heading, level 1 (if single level)
	ParaStyleTypeMS
	// ParaStyleTypeMS1 is a major section division heading, level 1 (if multiple levels)
	ParaStyleTypeMS1
	// ParaStyleTypeMS2 is a major section division heading, level 2
	ParaStyleTypeMS2
	// ParaStyleTypeMS3 is a major section division heading, level 3
	ParaStyleTypeMS3
	// ParaStyleTypeMR is a major section division references range heading
	ParaStyleTypeMR
	// ParaStyleTypeR is parallel reference(s)
	ParaStyleTypeR
	// ParaStyleTypeS is a section heading, level 1 (if single level)
	ParaStyleTypeS
	// ParaStyleTypeS1 is a section heading, level 1 (if multiple levels)
	ParaStyleTypeS1
	// ParaStyleTypeS2 is a section heading, level 2 (e.g. Proverbs 22-24)
	ParaStyleTypeS2
	// ParaStyleTypeS3 is a section heading, level 3 (e.g. Genesis "The First Day")
	ParaStyleTypeS3
	// ParaStyleTypeS4 is a section heading, level 4
	ParaStyleTypeS4
	// ParaStyleTypeSR is a section division references range heading
	ParaStyleTypeSR
	// ParaStyleTypeSP is a heading, to identify the speaker (e.g. Job)
	ParaStyleTypeSP
	// ParaStyleTypeSD is vertical space used to divide the text into sections, level 1 (if single level)
	ParaStyleTypeSD
	// ParaStyleTypeSD1 is semantic division location (vertical space used to divide the text into sections), level 1 (if multiple levels)
	ParaStyleTypeSD1
	// ParaStyleTypeSD2 is semantic division location (vertical space used to divide the text into sections), level 2
	ParaStyleTypeSD2
	// ParaStyleTypeSD3 is semantic division location (vertical space used to divide the text into sections), level 3
	ParaStyleTypeSD3
	// ParaStyleTypeSD4 is semantic division location (vertical space used to divide the text into sections), level 4
	ParaStyleTypeSD4
	// ParaStyleTypeTS is translator's chunk (to identify chunks of text suitable for translating at one time)
	ParaStyleTypeTS
	// ParaStyleTypeCP is published chapter number
	ParaStyleTypeCP
	// ParaStyleTypeCL is chapter label used for translations that add a word such as "Chapter"
	ParaStyleTypeCL
	// ParaStyleTypeCD is chapter Description (Publishing option D, e.g. in Russian Bibles)
	ParaStyleTypeCD
	// ParaStyleTypeMTE is the main title of the book repeated at the end of the book, level 1 (if single level)
	ParaStyleTypeMTE
	// ParaStyleTypeMTE1 is the main title of the book repeat /ed at the end of the book, level 1 (if multiple levels)
	ParaStyleTypeMTE1
	// ParaStyleTypeMTE2 is a secondary title occurring before or after the 'ending' main title
	ParaStyleTypeMTE2
	// ParaStyleTypeP0 is front or back matter text paragraph, level 1
	ParaStyleTypeP0
	// ParaStyleTypeP1 is front or back matter text paragraph, level 1 (if multiple levels)
	ParaStyleTypeP1
	// ParaStyleTypeP2 is front or back matter text paragraph, level 2 (if multiple levels)
	ParaStyleTypeP2
	// ParaStyleTypeK1 is concordance main entry text or keyword, level 1
	ParaStyleTypeK1
	// ParaStyleTypeK2 is concordance main entry text or keyword, level 2
	ParaStyleTypeK2
	// ParaStyleTypeREM is a remark
	ParaStyleTypeREM
)

func parseParaStyle(attr xml.Attr) ParaStyleType {
	switch attr.Value {
	case "restore":
		return ParaStyleTypeRESTORE
	case "cls":
		return ParaStyleTypeCLS
	case "iex":
		return ParaStyleTypeIEX
	case "ip":
		return ParaStyleTypeIP
	case "lit":
		return ParaStyleTypeLIT
	case "m":
		return ParaStyleTypeM
	case "mi":
		return ParaStyleTypeMI
	case "nb":
		return ParaStyleTypeNB
	case "p":
		return ParaStyleTypeP
	case "pb":
		return ParaStyleTypePB
	case "pc":
		return ParaStyleTypePC
	case "pi":
		return ParaStyleTypePI
	case "pi1":
		return ParaStyleTypePI1
	case "pi2":
		return ParaStyleTypePI2
	case "pi3":
		return ParaStyleTypePI3
	case "po":
		return ParaStyleTypePO
	case "pr":
		return ParaStyleTypePR
	case "pmo":
		return ParaStyleTypePMO
	case "pm":
		return ParaStyleTypePM
	case "pmc":
		return ParaStyleTypePMC
	case "pmr":
		return ParaStyleTypePMR
	case "ph":
		return ParaStyleTypePH
	case "ph1":
		return ParaStyleTypePH1
	case "ph2":
		return ParaStyleTypePH2
	case "ph3":
		return ParaStyleTypePH3
	case "q":
		return ParaStyleTypeQ
	case "q1":
		return ParaStyleTypeQ1
	case "q2":
		return ParaStyleTypeQ2
	case "q3":
		return ParaStyleTypeQ3
	case "q4":
		return ParaStyleTypeQ4
	case "qa":
		return ParaStyleTypeQA
	case "qc":
		return ParaStyleTypeQC
	case "qr":
		return ParaStyleTypeQR
	case "qm":
		return ParaStyleTypeQM
	case "qm1":
		return ParaStyleTypeQM1
	case "qm2":
		return ParaStyleTypeQM2
	case "qm3":
		return ParaStyleTypeQM3
	case "qd":
		return ParaStyleTypeQD
	case "b":
		return ParaStyleTypeB
	case "d":
		return ParaStyleTypeD
	case "ms":
		return ParaStyleTypeMS
	case "ms1":
		return ParaStyleTypeMS1
	case "ms2":
		return ParaStyleTypeMS2
	case "ms3":
		return ParaStyleTypeMS3
	case "mr":
		return ParaStyleTypeMR
	case "r":
		return ParaStyleTypeR
	case "s":
		return ParaStyleTypeS
	case "s1":
		return ParaStyleTypeS1
	case "s2":
		return ParaStyleTypeS2
	case "s3":
		return ParaStyleTypeS3
	case "s4":
		return ParaStyleTypeS4
	case "sr":
		return ParaStyleTypeSR
	case "sp":
		return ParaStyleTypeSP
	case "sd":
		return ParaStyleTypeSD
	case "sd1":
		return ParaStyleTypeSD1
	case "sd2":
		return ParaStyleTypeSD2
	case "sd3":
		return ParaStyleTypeSD3
	case "sd4":
		return ParaStyleTypeSD4
	case "ts":
		return ParaStyleTypeTS
	case "cp":
		return ParaStyleTypeCP
	case "cl":
		return ParaStyleTypeCL
	case "cd":
		return ParaStyleTypeCD
	case "mte":
		return ParaStyleTypeMTE
	case "mte1":
		return ParaStyleTypeMTE1
	case "mte2":
		return ParaStyleTypeMTE2
	case "p0":
		return ParaStyleTypeP0
	case "p1":
		return ParaStyleTypeP1
	case "p2":
		return ParaStyleTypeP2
	case "k1":
		return ParaStyleTypeK1
	case "k2":
		return ParaStyleTypeK2
	case "rem":
		return ParaStyleTypeREM
	default:
		return ParaStyleTypeUnknown
	}
}
