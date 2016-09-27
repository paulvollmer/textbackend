package textbackend

import (
	"testing"
)

// tests

func Test_TextRow_NewTextRow(t *testing.T) {
	row := NewTextRow(1, "hello world")
	if row.GetLevel() != 1 {
		t.Error("TextRow.GetLevel() not equal")
	}
	if row.GetText() != "hello world" {
		t.Error("TextRow.GetText() not equal")
	}
	if row.GetTextLength() != 11 {
		t.Error("TextRow.GetTextLength() not equal")
	}
}

var testTableNewTextRowFromString = []struct {
	text          string
	ws            string
	expectedLevel uint8
	expectedText  string
}{
	{"hello", "\t", 0, "hello"},
	{"\t\thello", "\t", 2, "hello"},
	{"\t\thello world", "\t", 2, "hello world"},
	{"\t\thello\tworld", "\t", 2, "hello\tworld"},
	{"\t\t hello\tcrazy\tworld\t", "\t", 2, " hello\tcrazy\tworld\t"},
}

func Test_TextRow_NewTextRowFromString(t *testing.T) {
	for k, tt := range testTableNewTextRowFromString {
		row := NewTextRowFromString(tt.text, tt.ws)
		if row.Level != tt.expectedLevel {
			t.Error(k, "NewTextRowFromString Level not equal")
		}
		if row.Text != tt.expectedText {
			t.Error(k, "NewTextRowFromString Text not equal")
		}
	}
}

func Test_TextRow_GetLevelWhitespace(t *testing.T) {
	textRow := TextRow{}
	textRow.Level = 1
	if textRow.GetLevelWhitespace("\t") != "\t" {
		t.Error("GetLevelWhitespace failed")
	}
	textRow.Level = 3
	if textRow.GetLevelWhitespace("\t") != "\t\t\t" {
		t.Error("GetLevelWhitespace failed")
	}
}

func Test_TextRow_GetString(t *testing.T) {
	textRow := NewTextRow(0, "hello world")
	if textRow.GetString("\t") != "hello world" {
		t.Error("GetString failed")
	}
	textRow2 := NewTextRow(1, "hello world")
	if textRow2.GetString("\t") != "\thello world" {
		t.Error("GetString failed")
	}
}

func Test_TextRow_AppendText(t *testing.T) {
	row := NewTextRow(0, "hello")
	row.AppendText(" world")
	if row.GetText() != "hello world" {
		t.Error("AppendText failed")
	}
}

func Test_TextRow_EqualTo(t *testing.T) {
	row := NewTextRow(0, "hello")
	if row.EqualTo("hello") == false {
		t.Error("TextRow.TextEqualTo() not equal, expected true")
	}
	if row.EqualTo("hello broken") == true {
		t.Error("TextRow.TextEqualTo() not equal, expected false")
	}
}

// benchmarks

var resultTextRow string

func benchmarkTextRow(l uint8, c string, b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		row := NewTextRow(l, c)
		r = row.GetString("<br>")
	}
	resultTextRow = r
}

func Benchmark_TextRow_level4(b *testing.B) {
	benchmarkTextRow(4, "hello world", b)
}
func Benchmark_TextRow_level8(b *testing.B) {
	benchmarkTextRow(8, "hello world", b)
}
func Benchmark_TextRow_level16(b *testing.B) {
	benchmarkTextRow(16, "hello world", b)
}
func Benchmark_TextRow_small(b *testing.B) {
	benchmarkTextRow(2, "hello world", b)
}
func Benchmark_TextRow_middle(b *testing.B) {
	benchmarkTextRow(2, "hello world, this is just a test", b)
}
func Benchmark_TextRow_large(b *testing.B) {
	benchmarkTextRow(2, "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", b)
}
