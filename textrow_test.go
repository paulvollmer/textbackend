package textbackend

import (
	"testing"
)

func TestTextRow_GetString(t *testing.T) {
	textRow := TextRow{}
	textRow.Text = "hello world"
	if textRow.GetString("\t") != "hello world" {
		t.Error("GetString broken")
	}

	textRow2 := TextRow{}
	textRow2.Level = 1
	textRow2.Text = "hello world"
	if textRow2.GetString("\t") != "\thello world" {
		t.Error("GetString broken")
	}
}
