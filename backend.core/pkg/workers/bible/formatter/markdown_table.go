package formatter

import (
	"bosca.io/pkg/bible/usx"
	"bytes"
)

func ToMarkdownTable(chapter *usx.Chapter) string {
	buf := new(bytes.Buffer)
	buf.WriteString("Chapter USFM: ")
	buf.WriteString(chapter.GetUsfm())
	buf.WriteString("\r\n")
	buf.WriteString("|Verse USFM|Verse Content|")
	buf.WriteString("\r\n")
	buf.WriteString("|---|---|---|---|")
	buf.WriteString("\r\n")
	for _, verse := range chapter.FindVerses() {
		buf.WriteString("|")
		buf.WriteString(verse.GetUsfm())
		buf.WriteString("|")
		buf.WriteString(verse.GetText())
		buf.WriteString("|")
		buf.WriteString("\r\n")
	}
	buf.WriteString("\r\n")
	return buf.String()
}
