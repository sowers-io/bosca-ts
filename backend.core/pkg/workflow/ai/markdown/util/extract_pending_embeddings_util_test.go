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

package util

import "testing"

func TestFromTable(t *testing.T) {
	values, err := ExtractPendingEmbeddingsFromMarkdown([]byte("|A|B|C|\n|---|---|---|\n|1|2|3\n|4|5|6|\n|7|8|9\n"), "A", "B")
	if err != nil {
		t.Error(err)
	}

	if len(values) != 3 {
		t.Errorf("Wrong number of values returned from ExtractPendingEmbeddings: %d", len(values))
	}

	if values[0].Id != "1" && *values[0].Content != "2" {
		t.Errorf("Wrong value returned from ExtractPendingEmbeddings: %s", values[1])
	}
	if values[1].Id != "4" && *values[1].Content != "5" {
		t.Errorf("Wrong value returned from ExtractPendingEmbeddings: %s", values[1])
	}
	if values[2].Id != "7" && *values[2].Content != "8" {
		t.Errorf("Wrong value returned from ExtractPendingEmbeddings: %s", values[1])
	}
}
