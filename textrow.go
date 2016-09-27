package textbackend

// TextRow to store the level and text of one row.
type TextRow struct {
	Level uint8  `json:"level"`
	Text  string `json:"text"`
}

// NewTextRow initialize and return a TextRow object
func NewTextRow(l uint8, t string) *TextRow {
	return &TextRow{l, t}
}

// GetLevel return the level of the TextRow
func (t *TextRow) GetLevel() uint8 {
	return t.Level
}

// GetText return the text of the TextRow
func (t *TextRow) GetText() string {
	return t.Text
}

// GetLevelWhitespace returns the whitespace for the row
func (t *TextRow) GetLevelWhitespace(whitespace string) string {
	if t.Level > 0 {
		buf := []byte(whitespace)
		var i uint8
		for i = 1; i < t.Level; i++ {
			buf = append(buf, whitespace...)
		}
		return string(buf)
	}
	return ""
}

// GetString returns the row with whitespace and text
func (t *TextRow) GetString(whitespace string) string {
	return t.GetLevelWhitespace(whitespace) + t.Text
}

// AppendText add the given string to the text row
func (t *TextRow) AppendText(text string) {
	t.Text = t.Text + text
}
