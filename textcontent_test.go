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
	if string(textContent.Get("\n", "\t")) != "hello world\n" {
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
		t.Error("GetString failed")
	}
}

var testTableWriteFile = []struct {
	level    uint8
	text     []string
	filename string
}{
	{
		0,
		[]string{"hello world"},
		"test1.txt",
	},
	{
		2,
		[]string{"hello world", "some more lines"},
		"test2.txt",
	},
}

func Test_TextContent_WriteFile(t *testing.T) {
	for k, tt := range testTableWriteFile {
		textContent := NewTextContent()
		textContent.SetLevel(tt.level)
		for _, tttext := range tt.text {
			textContent.Writeln(tttext)
		}
		err := textContent.WriteFile(tt.filename, 0666, "\n", "\t")
		if err != nil {
			t.Error(k, "Writefile failed")
		}
	}
}

func Test_TextContent_ReadFile(t *testing.T) {
	for k, tt := range testTableWriteFile {
		textContent, err := ReadFile(tt.filename, "\t", "\n")
		if err != nil {
			t.Error(k, "ReadFile failed")
		}
		if textContent.GetTotalLines() != len(tt.text)+1 { // +1 because there is an empty line at the end of the text
			t.Error(k, "ReadFile.GetTotalLines not equal")
		}
	}
}

func Test_TextContent_ReadFile_Error(t *testing.T) {
	_, err := ReadFile("no.file", "\t", "\n")
	if err == nil {
		t.Error("ReadFile missing error")
	}
}

// benchmarks

var resultTextContent []byte

func Benchmark_TextContent(b *testing.B) {
	var r []byte
	for n := 0; n < b.N; n++ {
		row := NewTextContent()
		row.Writeln("hello world")
		row.Writeln("more text here")
		row.Writeln("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
		r = row.Get("\n", "<br>")
	}
	resultTextContent = r
}
