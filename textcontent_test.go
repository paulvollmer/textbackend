package textbackend

import (
	"testing"
)

// tests

func Test_TextContent_Reset(t *testing.T) {
	textContent := NewTextContent()
	textContent.CurrentLevel = 5
	textContent.Reset()
	if textContent.GetLevel() != 0 {
		t.Error("Reset failed")
	}
}

func Test_TextContent_Writeln(t *testing.T) {
	textContent := NewTextContent()
	textContent.Writeln("hello")
	if len(textContent.Rows) != 1 || textContent.Rows[0].Text != "hello" {
		t.Error("Writeln failed")
	}
}

func Test_TextContent_Write(t *testing.T) {
	textContent := NewTextContent()
	textContent.Write("hello")
	textContent.Write("world")
	if len(textContent.Rows) != 1 || textContent.Rows[0].Text != "helloworld" {
		t.Error("Writeln failed")
	}
}

func Test_TextContent_WriteTo(t *testing.T) {
	textContent := NewTextContent()
	textContent.Writeln("hello")
	textContent.WriteTo(1, "-world")
	if len(textContent.Rows) != 1 || textContent.Rows[0].Text != "hello-world" {
		t.Error("WriteTo failed")
	}
	textContent.WriteTo(100, "testing")
	textContent.WriteTo(-100, "testing")
}

func Test_TextContent_PushLevel(t *testing.T) {
	textContent := NewTextContent()
	textContent.PushLevel()
	if textContent.CurrentLevel != 1 {
		t.Error("PushLevel failed")
	}
}

func Test_TextContent_PopLevel(t *testing.T) {
	textContent := NewTextContent()
	textContent.PushLevel()
	textContent.PushLevel()
	textContent.PopLevel()
	if textContent.CurrentLevel != 1 {
		t.Error("PopLevel failed")
	}
}

func Test_TextContent_GetLevel(t *testing.T) {
	textContent := NewTextContent()
	textContent.PushLevel()
	if textContent.GetLevel() != 1 {
		t.Error("GetLevel failed")
	}
}

func Test_TextContent_SetLevel(t *testing.T) {
	textContent := NewTextContent()
	textContent.SetLevel(5)
	if textContent.CurrentLevel != 5 {
		t.Error("SetLevel failed")
	}
}

func Test_TextContent_GetTotalLines(t *testing.T) {
	textContent := NewTextContent()
	textContent.Writeln("hello")
	textContent.Writeln("world")
	if textContent.GetTotalLines() != 2 {
		t.Error("GetTotalLines failed")
	}
}

func Test_TextContent_GetString(t *testing.T) {
	textContent := NewTextContent()
	textContent.Writeln("hello world")
	if textContent.GetString("\n", "\t") != "hello world\n" {
		t.Error("GetString failed")
	}
}

func Test_TextContent_GetStringArray(t *testing.T) {
	textContent := NewTextContent()
	textContent.Writeln("hello world")
	tmpStrArr := textContent.GetStringArray("\t")
	if len(tmpStrArr) != 1 || tmpStrArr[0] != "hello world" {
		t.Error("GetStringArray failed")
	}

}

func Test_TextContent_GetJSON(t *testing.T) {
	textContent := NewTextContent()
	textContent.Writeln("hello world")
	tmpJSON, err := textContent.GetJSON()
	if err != nil || string(tmpJSON) != `[{"level":0,"text":"hello world"}]` {
		t.Error("GetString broken")
	}
}

// benchmarks

var resultTextContent string

func Benchmark_TextContent(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		row := NewTextRow(2, "hello world")
		r = row.GetString("<br>")
	}
	resultTextContent = r
}
