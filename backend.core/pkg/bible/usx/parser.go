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
	"io"
	"strings"
)

type phase int

const (
	bookIdentification phase = iota
	bookHeaders
	bookTitles
	bookIntroduction
	bookIntroductionEndTitles
	bookChapterLabel
	chapter
	chapterContent
)

type Parser struct {
	phase   phase
	usx     *USX
	decoder *xml.Decoder
}

func NewParser(decoder *xml.Decoder) *Parser {
	return &Parser{
		phase:   bookIdentification,
		usx:     &USX{},
		decoder: decoder,
	}
}

func (p *Parser) Parse() (*USX, error) {
	for {
		token, err := p.decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == "usx" {
				if err = p.onUsx(); err != nil {
					return nil, err
				}
			}
		}
	}
	return p.usx, nil
}

func (p *Parser) onUsx() error {
	for {
		token, err := p.decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		switch element := token.(type) {
		case xml.EndElement:
			if element.Name.Local != "usx" {
				return fmt.Errorf("parse: unexpected end tag: %s", element.Name.Local)
			}
		default:
			var chapter *Chapter
			if len(p.usx.Chapters) > 0 {
				chapter = p.usx.Chapters[len(p.usx.Chapters)-1]
			}
			if handled, err := p.onUsxElement(token, chapter); err != nil {
				return err
			} else if !handled {
				return fmt.Errorf("onUsx: unhandled token: %v", token)
			}
		}
	}
}

func (p *Parser) onUsxElement(token xml.Token, current NodeContainer) (bool, error) {
	var err error
	var handled = false
	switch element := token.(type) {
	case xml.StartElement:
		switch element.Name.Local {
		case "book":
			if handled, err = p.onBook(element); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			fallthrough
		case "para":
			if handled, err = p.onParaPhase(element, current); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			fallthrough
		case "chapter":
			if handled, err = p.onChapter(element); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
		default:
			if p.phase != chapterContent {
				return false, fmt.Errorf("required phase to be chapterContent, unexpected phase: %s -> %d", element.Name.Local, p.phase)
			}
			if handled, err = p.onChapterContent(element, current); err != nil {
				return false, err
			} else {
				return handled, nil
			}
		}
	case xml.EndElement:
		if element.Name.Local != "usx" {
			return false, fmt.Errorf("unexpected state: %s", element.Name.Local)
		}
		return false, nil
	case xml.CharData:
		if strings.Replace(strings.Replace(strings.Trim(string(element), " "), "\r", "", -1), "\n", "", -1) == "" {
			return true, nil
		}
		return false, fmt.Errorf("unexpected text: %s", string(element))
	default:
		return false, fmt.Errorf("unexpected state: %v", element)
	}
	return false, fmt.Errorf("unexpected state")
}

func (p *Parser) onBook(element xml.StartElement) (bool, error) {
	switch p.phase {
	case bookIdentification:
		if handled, err := p.onBookIdentification(element); err != nil {
			return false, err
		} else {
			return handled, nil
		}
	case bookHeaders:
		if success, err := p.onBookHeaders(element); err != nil {
			return true, err
		} else if !success {
			p.phase = bookHeaders
			return false, nil
		}
		return false, nil
	default:
		return false, errors.New("unexpected state")
	}
}

func (p *Parser) onParaPhase(element xml.StartElement, current NodeContainer) (bool, error) {
	for {
		switch p.phase {
		case bookHeaders:
			if handled, err := p.onBookHeaders(element); err != nil {
				return true, err
			} else if handled {
				return true, nil
			}
			p.phase = bookTitles
		case bookTitles:
			if handled, err := p.onBookTitles(element); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			p.phase = bookIntroduction
		case bookIntroduction:
			if handled, err := p.onBookIntroduction(element); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			p.phase = bookIntroductionEndTitles
		case bookIntroductionEndTitles:
			if handled, err := p.onBookIntroductionEndTitles(element); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			p.phase = bookChapterLabel
		case bookChapterLabel:
			if handled, err := p.onBookChapterLabel(element); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			p.phase = chapter
		case chapterContent:
			if handled, err := p.onChapterContent(element, current); err != nil {
				return false, err
			} else if handled {
				return true, nil
			}
			return false, nil
		default:
			return false, fmt.Errorf("onParaPhase: unexpected phase: %s -> %d", element.Name.Local, p.phase)
		}
	}
}

func (p *Parser) onBookIdentification(element xml.StartElement) (bool, error) {
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "code":
			p.usx.BookIdentification.Code = parseBookType(attr)
		case "style":
			p.usx.BookIdentification.Style = attr.Value
		}
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.CharData:
			p.usx.BookIdentification.Text = string(element)
		case xml.EndElement:
			if element.Name.Local != "book" {
				return false, errors.New("onBookIdentification: unexpected state")
			}
			p.phase = bookHeaders
			return true, nil
		}
	}
}

func (p *Parser) onBookHeaders(element xml.StartElement) (bool, error) {
	header := BookHeader{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "style":
			header.Style = parseBookHeader(attr)
		}
	}
	if header.Style == BookHeaderTypeUnknown {
		return false, nil
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.CharData:
			header.Text = string(element)
		case xml.EndElement:
			if element.Name.Local != "para" {
				return false, errors.New("onBookHeaders: unexpected state")
			}
			p.usx.BookHeaders = append(p.usx.BookHeaders, header)
			return true, nil
		}
	}
}

func (p *Parser) onBookTitles(element xml.StartElement) (bool, error) {
	title := BookTitle{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "style":
			title.Style = parseBookTitle(attr)
		}
	}
	if title.Style == BookTitleTypeUnknown {
		return false, nil
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.CharData:
			title.Text = string(element)
		case xml.EndElement:
			if element.Name.Local != "para" {
				return false, errors.New("onBookTitles: unexpected state")
			}
			p.usx.BookTitles = append(p.usx.BookTitles, title)
			return true, nil
		}
	}
}

func (p *Parser) onBookIntroduction(element xml.StartElement) (bool, error) {
	introduction := BookIntroduction{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "style":
			introduction.Style = parseBookIntroduction(attr)
		}
	}
	if introduction.Style == BookIntroductionUnknown {
		return false, nil
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.EndElement:
			if element.Name.Local != "para" {
				return false, errors.New("onBookIntroduction: unexpected state")
			}
			p.usx.BookIntroductions = append(p.usx.BookIntroductions, introduction)
			return true, nil
		}
	}
}

func (p *Parser) onBookIntroductionEndTitles(element xml.StartElement) (bool, error) {
	endTitle := BookIntroductionEndTitle{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "style":
			endTitle.Style = parseBookIntroductionEndTitle(attr)
		}
	}
	if endTitle.Style == BookIntroductionEndTitleUnknown {
		return false, nil
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.EndElement:
			if element.Name.Local != "para" {
				return false, errors.New("onBookIntroductionEndTitles: unexpected state")
			}
			p.usx.BookIntroductionEndTitles = append(p.usx.BookIntroductionEndTitles, endTitle)
			return true, nil
		}
	}
}

func (p *Parser) onBookChapterLabel(element xml.StartElement) (bool, error) {
	label := BookChapterLabel{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "style":
			if attr.Value != "cl" {
				return false, nil
			}
		}
	}
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.CharData:
			label.Text = string(element)
		case xml.EndElement:
			if element.Name.Local != "para" {
				return false, errors.New("onBookChapterLabel: unexpected state")
			}
			p.usx.BookChapterLabel = &label
			return true, nil
		}
	}
}

func (p *Parser) onChapter(element xml.StartElement) (bool, error) {
	chapter := &Chapter{}
	for _, attr := range element.Attr {
		switch attr.Name.Local {
		case "number":
			chapter.Number = attr.Value
		case "altnumber":
			chapter.AltNumber = attr.Value
		case "pubnumber":
			chapter.PubNumber = attr.Value
		case "sid":
			chapter.Sid = attr.Value
		}
	}
	p.phase = chapterContent
	for {
		token, err := p.decoder.Token()
		if err != nil {
			return false, err
		}
		switch element := token.(type) {
		case xml.StartElement:
			return false, errors.New("unexpected start tag after chapter")
		case xml.EndElement:
			if element.Name.Local != "chapter" {
				return false, errors.New("onChapter: unexpected state")
			}
			if chapter.Number == "" {
				return true, nil
			}
			p.usx.Chapters = append(p.usx.Chapters, chapter)
			return true, nil
		}
	}
}

func (p *Parser) onChapterStart(element xml.StartElement) error {
	return errors.New("onChapterStart: unimplemented")
}

func (p *Parser) onChapterContent(element xml.StartElement, current NodeContainer) (bool, error) {
	switch element.Name.Local {
	case "chapter":
		return false, errors.New("onChapterContent: unsupported end chapter")
	case "para":
		return p.onParaStandard(element, current)
	case "char":
		return p.onChar(element, current)
	case "verse":
		return p.onVerse(element, current)
	case "list":
		return p.onList(element, current)
	case "table":
		return p.onTable(element, current)
	case "note":
		return p.onNote(element, current)
	case "ref":
		return p.onReference(element, current)
	case "sidebar":
		return false, errors.New("onChapterContent: unsupported content type: sidebar")
	default:
		return false, fmt.Errorf("onChapterContent: unexpected content tag: %s", element.Name.Local)
	}
}

func (p *Parser) onList(element xml.StartElement, current NodeContainer) (bool, error) {
	return false, errors.New("onList: unimplemented")
}

func (p *Parser) onTable(element xml.StartElement, current NodeContainer) (bool, error) {
	return false, errors.New("onTable: unimplemented")
}

func (p *Parser) onCrossReference(element xml.StartElement, crossReferenceType CrossReferenceType, current NodeContainer) (bool, error) {
	return false, errors.New("onFootnote: unimplemented")
}

func (p *Parser) onReference(element xml.StartElement, current NodeContainer) (bool, error) {
	return false, errors.New("onReference: unimplemented")
}
