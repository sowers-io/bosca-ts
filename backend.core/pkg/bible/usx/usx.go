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

type USX struct {
	BookIdentification        BookIdentification
	BookIntroductions         []BookIntroduction
	BookHeaders               []BookHeader
	BookTitles                []BookTitle
	BookIntroductionEndTitles []BookIntroductionEndTitle
	BookChapterLabel          *BookChapterLabel
	Chapters                  []*Chapter
}

type ChapterVerses struct {
	Chapter *Chapter
	Verses  []*Verse
}

func (u *USX) GetUsfm() string {
	return u.BookIdentification.Code.ToString()
}

func (u *USX) FindChapterVerses() []*ChapterVerses {
	chapters := make([]*ChapterVerses, 0)
	for _, chapter := range u.Chapters {
		chapters = append(chapters, &ChapterVerses{
			Chapter: chapter,
			Verses:  findVerses(chapter),
		})
	}
	return chapters
}

func findVerses(container NodeContainer) []*Verse {
	verses := make([]*Verse, 0)
	for _, node := range container.GetChildren() {
		if verse, ok := node.(*Verse); ok {
			verses = append(verses, verse)
		} else if container, ok := node.(NodeContainer); ok {
			verses = append(verses, findVerses(container)...)
		}
	}
	return verses
}
