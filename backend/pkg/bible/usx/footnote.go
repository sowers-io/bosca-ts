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

type FootnoteType int

type Footnote struct {
	Type     FootnoteType
	Caller   string
	Category *string
	Children []Node
}

type FootnoteVerse struct {
	Style string
	Text  string
}

func (f *Footnote) AddText(text string) error {
	if len(f.Children) == 0 {
		f.Children = append(f.Children, &Char{Text: text})
		return nil
	}
	last := f.Children[len(f.Children)-1]
	return last.AddText(text)
}

func (f *Footnote) AddNode(node Node) {
	f.Children = append(f.Children, node)
}

func (f *Footnote) GetChildren() []Node {
	return f.Children
}

func (p *Parser) onFootnote(footnoteTag xml.StartElement, footnoteType FootnoteType, current NodeContainer) (bool, error) {
	footnote := &Footnote{
		Type: footnoteType,
	}
	current.AddNode(footnote)
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.StartElement:
			if handled, err := p.onChapterContent(element, footnote); err != nil {
				return false, err
			} else if !handled {
				return false, nil
			}
		case xml.CharData:
			if err = footnote.AddText(string(element)); err != nil {
				return false, err
			}
		case xml.EndElement:
			if element.Name.Local != "note" {
				return false, fmt.Errorf("onFootnote: unexpected tag: %s", element.Name.Local)
			}
			return true, nil
		default:
			return false, fmt.Errorf("onFootnote: unexpected element: %v", element)
		}
	}
}
