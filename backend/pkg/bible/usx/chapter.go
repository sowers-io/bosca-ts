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
	"errors"
	"strings"
)

type Chapter struct {
	Number    string
	AltNumber string
	PubNumber string
	Sid       string
	Children  []Node
}

func (v *Chapter) GetUsfm() string {
	return strings.Replace(strings.Replace(v.Sid, " ", ".", -1), ":", ".", -1)
}

func (c *Chapter) AddText(text string) error {
	return errors.New("unexpected text in chapter")
}

func (c *Chapter) AddNode(node Node) {
	c.Children = append(c.Children, node)
}

func (c *Chapter) GetChildren() []Node {
	return c.Children
}

type ChapterStart struct {
}

type ChapterContent struct {
}
