package textbackend

// TextRow to store the level and string of one textline
type TextRow struct {
	Level int
	Text  string
}

func (t *TextRow) GetLevelWhitespace(whitespace string) string {
	tmp := ""
	if t.Level > 0 {
		for i := 0; i < t.Level; i++ {
			tmp += whitespace
		}
	}
	return tmp
}

func (t *TextRow) GetString(whitespace string) string {
	return t.GetLevelWhitespace(whitespace) + t.Text
}
