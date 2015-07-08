package textbackend

import (
	"testing"
)

func TestTextContent_Reset(t *testing.T) {
	textContent := TextContent{}
	textContent.CurrentLevel = 5
	textContent.Reset()
	if textContent.GetLevel() != 0 {
		t.Error("Reset failed")
	}
}

func TestTextContent_Writeln(t *testing.T) {
	textContent := TextContent{}
	textContent.Writeln("hello")
	if len(textContent.Rows) != 1 || textContent.Rows[0].Text != "hello" {
		t.Error("Writeln failed")
	}
}

func TestTextContent_Write(t *testing.T) {
	textContent := TextContent{}
	textContent.Write("hello")
	textContent.Write("world")
	if len(textContent.Rows) != 1 || textContent.Rows[0].Text != "helloworld" {
		t.Error("Writeln failed")
	}
}

func TestTextContent_WriteTo(t *testing.T) {
	textContent := TextContent{}
	textContent.Writeln("hello")
	textContent.WriteTo(1, "-world")
	if len(textContent.Rows) != 1 || textContent.Rows[0].Text != "hello-world" {
		t.Error("WriteTo failed")
	}
	textContent.WriteTo(100, "testing")
	textContent.WriteTo(-100, "testing")
}

func TestTextContent_PushLevel(t *testing.T) {
	textContent := TextContent{}
	textContent.PushLevel()
	if textContent.CurrentLevel != 1 {
		t.Error("PushLevel failed")
	}
}

func TestTextContent_PopLevel(t *testing.T) {
	textContent := TextContent{}
	textContent.PushLevel()
	textContent.PushLevel()
	textContent.PopLevel()
	if textContent.CurrentLevel != 1 {
		t.Error("PopLevel failed")
	}
}

func TestTextContent_GetLevel(t *testing.T) {
	textContent := TextContent{}
	textContent.PushLevel()
	if textContent.GetLevel() != 1 {
		t.Error("GetLevel failed")
	}
}

func TestTextContent_SetLevel(t *testing.T) {
	textContent := TextContent{}
	textContent.SetLevel(5)
	if textContent.CurrentLevel != 5 {
		t.Error("SetLevel failed")
	}
}

func TestTextContent_GetTotalLines(t *testing.T) {
	textContent := TextContent{}
	textContent.Writeln("hello")
	textContent.Writeln("world")
	if textContent.GetTotalLines() != 2 {
		t.Error("GetTotalLines failed")
	}
}

func TestTextContent_GetString(t *testing.T) {
	textContent := TextContent{}
	textContent.Writeln("hello world")
	if textContent.GetString("\n", "\t") != "hello world\n" {
		t.Error("GetString failed")
	}
}

func TestTextContent_GetStringArray(t *testing.T) {
	textContent := TextContent{}
	textContent.Writeln("hello world")
	tmpStrArr := textContent.GetStringArray("\t")
	if len(tmpStrArr) != 1 || tmpStrArr[0] != "hello world" {
		t.Error("GetStringArray failed")
	}

}

func TestTextContent_GetJSON(t *testing.T) {
	textContent := TextContent{}
	textContent.Writeln("hello world")
	tmpJSON, err := textContent.GetJSON()
	if err != nil || string(tmpJSON) != `[{"level":0,"text":"hello world"}]` {
		t.Error("GetString broken")
	}
}
