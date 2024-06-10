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

type BookIdentification struct {
	Style string
	Code  BookCode
	Text  string
}

type BookHeader struct {
	Style BookHeaderType
	Text  string
}

type BookTitle struct {
	Style BookTitleType
	Text  string
}

type BookIntroduction struct {
	Style BookIntroductionType
}

type BookIntroductionEndTitle struct {
	Style BookIntroductionEndTitleType
}

type BookChapterLabel struct {
	Text string
}
