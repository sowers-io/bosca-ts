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
	"strings"
)

type VerseText struct {
	Text string
}

func (v *VerseText) AddText(text string) error {
	v.Text += text
	return nil
}

type Verse struct {
	Number    string
	AltNumber string
	Sid       string
	Children  []Node
}

func (v *Verse) GetUsfm() string {
	return strings.Replace(strings.Replace(v.Sid, " ", ".", -1), ":", ".", -1)
}

func (v *Verse) GetText() string {
	buf := &strings.Builder{}
	buildText(buf, v)
	return buf.String()
}

func buildText(buf *strings.Builder, container NodeContainer) {
	for _, child := range container.GetChildren() {
		switch node := child.(type) {
		case *VerseText:
			buf.WriteString(node.Text)
		case NodeContainer:
			buildText(buf, node)
		}
	}
}

func (v *Verse) AddText(text string) error {
	if len(v.Children) == 0 {
		v.Children = append(v.Children, &VerseText{Text: text})
		return nil
	}
	last := v.Children[len(v.Children)-1]
	return last.AddText(text)
}

func (v *Verse) AddNode(node Node) {
	v.Children = append(v.Children, node)
}

func (v *Verse) GetChildren() []Node {
	return v.Children
}

func (p *Parser) onVerse(element xml.StartElement, current NodeContainer) (bool, error) {
	verse := &Verse{}
	end := false
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "number":
			verse.Number = attr.Value
		case "altnumber":
			verse.AltNumber = attr.Value
		case "sid":
			verse.Sid = attr.Value
		case "eid":
			end = true
		}
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.StartElement:
			return false, errors.New("unexpected start tag after verse")
		case xml.EndElement:
			if element.Name.Local != "verse" {
				return false, errors.New("unexpected state")
			}
			if end {
				return true, nil
			}
			current.AddNode(verse)
			return true, nil
		}
	}
}
