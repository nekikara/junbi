package junbi

import (
	"log"
	"os"
)

type FormatMaker interface {
	MkDir()
	MkFile()
	CreateScriptFile()
}

type Ruby struct {
	dataAnswer
}

var NewRubyFormat = func(name string) FormatMaker {
	var rb Ruby
	rb.dataAnswer = newDataAnswer()
	rb.name = name
	rb.extension = "rb"
	rb.shebang = "#! /usr/bin/env ruby"
	return rb
}

type Perl struct {
	dataAnswer
}

var NewPerlFormat = func(name string) FormatMaker {
	var pl Perl
	pl.dataAnswer = newDataAnswer()
	pl.name = name
	pl.extension = "pl"
	pl.shebang = "#! /usr/bin/perl"
	pl.encodeType = "use strict;\nuse utf8;\nuse warnings;"
	return pl
}

type Python struct {
	dataAnswer
}

var NewPythonFormat = func(name string) FormatMaker {
	var py Python
	py.dataAnswer = newDataAnswer()
	py.name = name
	py.extension = "py"
	py.shebang = "#! /usr/bin/env python"
	return py
}

var format = map[string]func(string) FormatMaker{
	"rb": NewRubyFormat,
	"py": NewPythonFormat,
	"pl": NewPerlFormat,
}

func NewFormatMaker(name, extension string) FormatMaker {
	return format[extension](name)
}

type dataAnswer struct {
	dataAnswer []string
	name       string
	extension  string
	shebang    string
	encodeType string
}

func newDataAnswer() dataAnswer {
	return dataAnswer{[]string{"data", "answer"}, "", "", "", "# coding: utf-8"}
}

func (da dataAnswer) MkDir() {
	for _, name := range da.dataAnswer {
		if err := os.Mkdir(name, 0755); err != nil {
			log.Println(err)
		}
	}
}

func (da dataAnswer) MkFile() {
	for _, name := range da.dataAnswer {
		if _, err := os.Create(name + "/1." + name); err != nil {
			log.Println(err)
		}
	}
}

func (da dataAnswer) CreateScriptFile() {
	f, err := os.OpenFile(da.name+"."+da.extension, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	var s string
	e := da.extension
	if e != "" {
		s = da.shebang + "\n" + da.encodeType
	}
	f.WriteString(s)
}
