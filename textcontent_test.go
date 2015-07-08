package textbackend

import (
	"testing"
)

func TestTextContent_GetString(t *testing.T) {
	textContent := TextContent{}

	textContent.Writeln("hello world")
	if textContent.GetString("\n", "\t") != "hello world\n" {
		t.Error("GetString broken")
	}

	tmpStrArr := textContent.GetStringArray("\t")
	if tmpStrArr == nil && len(tmpStrArr) == 1 && tmpStrArr[0] == "hello world" {
		t.Error("GetString broken")
	}

	tmpJSON, err := textContent.GetJSON()
	if err != nil && string(tmpJSON) != `[{"level":0,"text":"hello world"}]` {
		t.Error("GetString broken")
	}

	if textContent.GetTotalLines() != 1 {
		t.Error("GetTotalLines broken")
	}

	textContent.PushLevel()
	if textContent.GetLevel() != 1 {
		t.Error("GetLevel broken")
	}

	textContent.SetLevel(7)
	if textContent.GetLevel() != 7 {
		t.Error("SetLevel broken")
	}

	textContent.PopLevel()
	if textContent.GetLevel() != 6 {
		t.Error("PopLevel broken")
	}

	textContent.Reset()
	if textContent.GetLevel() != 0 {
		t.Error("PopLevel broken")
	}

}
