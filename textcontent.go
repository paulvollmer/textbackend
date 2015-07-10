/*
Package textbackend simple usage example:

    textContent := TextContent{}
    textContent.Writeln("hello world")
    textContent.PushLevel()
    textContent.Writeln("one more line")
    textContent.PopLevel()
    theText := textContent.GetStringArray("\t")
*/
package textbackend

import (
	"encoding/json"
)

// TextContent store a ContentLine list
type TextContent struct {
	Rows         []TextRow
	CurrentLevel int // the current level
}

// Reset the content data and the current level value
func (t *TextContent) Reset() {
	t.Rows = []TextRow{}
	t.CurrentLevel = 0
}

// Writeln append a line to the content array
func (t *TextContent) Writeln(text string) {
	t.Rows = append(t.Rows, TextRow{t.CurrentLevel, text})
}

// Write writes text to the latest row
func (t *TextContent) Write(text string) {
	if len(t.Rows) != 0 {
		t.Rows[len(t.Rows)-1].Text += text
	} else {
		t.Writeln(text)
	}
}

// WriteTo writes text to a specific line.
func (t *TextContent) WriteTo(row int, text string) {
	if row >= 0 && row <= len(t.Rows) {
		t.Rows[row-1].Text += text
	}
}

// PushLevel pushes into the next level.
func (t *TextContent) PushLevel() {
	t.CurrentLevel++
}

// PopLevel pop out from the level.
func (t *TextContent) PopLevel() {
	if t.CurrentLevel > 0 {
		t.CurrentLevel--
	}
}

// GetLevel return the current level depth as int.
func (t *TextContent) GetLevel() int {
	return t.CurrentLevel
}

// SetLevel set the level to a specific depth
func (t *TextContent) SetLevel(depth int) {
	t.CurrentLevel = depth
}

// GetTotalLines returns the number of lines.
func (t *TextContent) GetTotalLines() int {
	return len(t.Rows)
}

// GetString renders out the content as string and you can set the type of linebreak and level char
func (t *TextContent) GetString(linebreak, whitespace string) string {
	contentStr := ""
	for _, v := range t.Rows {
		contentStr += v.GetLevelWhitespace(whitespace) + v.Text + linebreak
	}
	return contentStr
}

// GetStringArray return the content as array.
func (t *TextContent) GetStringArray(whitespace string) []string {
	tmp := []string{}
	for _, v := range t.Rows {
		tmpRow := v.GetLevelWhitespace(whitespace) + v.Text
		tmp = append(tmp, tmpRow)
	}
	return tmp
}

// GetJSON return the content as array with row objects.
func (t *TextContent) GetJSON() ([]byte, error) {
	return json.Marshal(t.Rows)
}
