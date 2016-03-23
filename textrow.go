package textbackend

// TextRow to store the level and text of one row.
type TextRow struct {
	Level int    `json:"level"`
	Text  string `json:"text"`
}

// NewTextRow initialize and return a TextRow object
func NewTextRow(l int, t string) *TextRow {
	return &TextRow{l, t}
}

// GetLevelWhitespace returns the whitespace for the row.
func (t *TextRow) GetLevelWhitespace(whitespace string) string {
	tmp := ""
	if t.Level > 0 {
		for i := 0; i < t.Level; i++ {
			tmp += whitespace
		}
	}
	return tmp
}

// GetString returns the row with whitespace and text
func (t *TextRow) GetString(whitespace string) string {
	return t.GetLevelWhitespace(whitespace) + t.Text
}
