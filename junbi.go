package main

import (
	"flag"
	"junbi/junbi"
	"os"
)

var (
	initFlag      bool
	extensionFlag string
	nameFlag      string
)

func init() {
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flg.BoolVar(&initFlag, "init", false, "true or false")
	flg.BoolVar(&initFlag, "i", false, "true or false")
	flg.StringVar(&extensionFlag, "extension", "", "言語の拡張子を指定")
	flg.StringVar(&extensionFlag, "e", "", "言語の拡張子を指定")
	flg.StringVar(&nameFlag, "name", "", "編集するメインのファイル名を指定")
	flg.StringVar(&nameFlag, "n", "", "編集するメインのファイル名を指定")
	flg.Parse(os.Args[1:])
	if 0 < flg.NArg() {
		flg.Parse(flg.Args()[1:])
	}
}

func main() {
	var lang junbi.FormatMaker
	if extensionFlag != "" && nameFlag != "" {
		lang = junbi.NewFormatMaker(nameFlag, extensionFlag)
		lang.CreateScriptFile()
	}
	if initFlag {
		lang.MkDir()
		lang.MkFile()
	}
}
