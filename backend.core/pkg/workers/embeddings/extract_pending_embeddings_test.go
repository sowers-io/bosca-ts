package embeddings

import "testing"

func TestFromTable(t *testing.T) {
	values, err := extractPendingEmbeddingsFromMarkdown([]byte("|A|B|C|\n|---|---|---|\n|1|2|3\n|4|5|6|\n|7|8|9\n"), "A", "B")
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
