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
	"fmt"
)

type Char struct {
	Text string
}

func (c *Char) AddText(text string) error {
	c.Text += text
	return nil
}

func (p *Parser) onChar(element xml.StartElement, current NodeContainer) (bool, error) {
	char := &Char{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "style":
			//char.Style = parseCharStyle(attr)
		}
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.CharData:
			char.Text = string(element)
		case xml.EndElement:
			if element.Name.Local != "char" {
				return false, errors.New("onChar: unexpected end")
			}
			current.AddNode(char)
			return true, nil
		default:
			return false, fmt.Errorf("onChar: unexpected element: %v", element)
		}
	}
}
