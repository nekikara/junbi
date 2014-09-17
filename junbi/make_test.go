package junbi

import (
	"fmt"
	"testing"
)

/*
func estMkDir(t *testing.T) {
	j = beforeProcess()
	j.MkDir()

	for _, name := range j {
		f, err := os.Stat(name)
		if err != nil {
			t.Errorf("%v", err)
		}
		if !f.IsDir() {
			t.Errorf("%s はディレクトリではありません", f.Name())
		}
	}

	afterProcess(j)
}

func estMkFile(t *testing.T) {
	j = beforeProcess()
	j.MkDir()
	j.MkFile()

	for _, name := range j {
		if _, err := os.Stat(name + "/1." + name); err != nil {
			t.Errorf("%v", err)
		}
	}

	afterProcess(j)
}

func beforeProcess() Junbi {
	return Junbi{"data", "answer"}
}

func afterProcess(j Junbi) {
	for _, name := range j {
		os.RemoveAll(name)
	}
}

func estCreate(t *testing.T) {
	cases := []string{
		"rb",
		"py",
		"pl",
		"",
		"test",
	}
	for _, tc := range cases {
		name := "hoge"
		file := name + "." + tc
		m := MainF{name, tc}
		m.Create()
		if _, err := os.Stat(file); err != nil {
			t.Errorf("ファイルがうまく作成できていません。\n%v", err)
		}
		os.Remove(file)
	}
}
*/

func TestNewRubyFormat(t *testing.T) {
	r := NewRubyFormat("hogehoge")
	fmt.Println(r.name)
}
