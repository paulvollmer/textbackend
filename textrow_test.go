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
	if string(row.GetText()) != "hello world" {
		t.Error("TextRow.GetText() not equal")
	}
}

func Test_TextRow_GetLevelWhitespace(t *testing.T) {
	textRow := TextRow{}
	textRow.Level = 1
	if textRow.GetLevelWhitespace("\t") != "\t" {
		t.Error("GetLevelWhitespace failed")
	}
}

func Test_TextRow_GetString(t *testing.T) {
	textRow := TextRow{}
	textRow.Text = "hello world"
	if textRow.GetString("\t") != "hello world" {
		t.Error("GetString failed")
	}

	textRow2 := TextRow{}
	textRow2.Level = 1
	textRow2.Text = "hello world"
	if textRow2.GetString("\t") != "\thello world" {
		t.Error("GetString failed")
	}
}

// benchmarks

var resultTextRow string

func Benchmark_TextRow(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		row := NewTextRow(2, "hello world")
		r = row.GetString("<br>")
	}
	resultTextRow = r
}
