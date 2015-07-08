package textbackend

import (
	"encoding/json"
)

// TextContent store a ContentLine list
type TextContent struct {
	Rows  []TextRow
	level int // the current level
}

// Reset the content data and the current level value
func (t *TextContent) Reset() {
	t.Rows = []TextRow{}
	t.level = 0
}

// Writeln append a line to the content array
func (t *TextContent) Writeln(text string) {
	t.Rows = append(t.Rows, TextRow{t.level, text})
}

// TODO: Write
// func (t *TextContent) Write() {
// }

// TODO: WriteTo writes text to a specific line.
// func (t *TextContent) WriteTo() {
// }

// PushLevel pushes into the next level.
func (t *TextContent) PushLevel() {
	t.level++
}

// PopLevel pop out from the level.
func (t *TextContent) PopLevel() {
	if t.level > 0 {
		t.level--
	}
}

// GetLevel return the current level depth as int.
func (t *TextContent) GetLevel() int {
	return t.level
}

// SetLevel
func (t *TextContent) SetLevel(newLevel int) {
	t.level = newLevel
}

// GetTotalLines returns the number of lines.
func (t *TextContent) GetTotalLines() int {
	return len(t.Rows)
}

// TODO: add masterLevel to get methods

// GetString renders out the content as string and you can set the type of linebreak and level char
func (t *TextContent) GetString(linebreak, whitespace string) string {
	contentStr := ""
	for _, v := range t.Rows {
		tmpRow := v.GetLevelWhitespace(whitespace) + v.Text
		contentStr += tmpRow + "\n"
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
