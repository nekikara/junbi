package main

import (
	"flag"
	"os"

	"junbi/junbi"
)

var (
	initFlag bool
	langFlag string
	nameFlag string
)

func init() {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flg.BoolVar(&initFlag, "init", false, "true or false")
	flg.StringVar(&langFlag, "lang", "", "言語の拡張子を指定")
	flg.StringVar(&nameFlag, "name", "", "編集するメインのファイル名を指定")
	flg.Parse(os.Args[1:])
	if 0 < flg.NArg() {
		flg.Parse(flg.Args()[1:])
	}
}

func main() {
	if initFlag {
		j := junbi.Junbi{"data", "answer"}
		j.MkDir()
		j.MkFile()
	}
	if langFlag != "" && nameFlag != "" {
		m := junbi.MainF{nameFlag, langFlag}
		m.Create()
	}
}
