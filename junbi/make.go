package junbi

import (
	"fmt"
	"log"
	"os"
)

type Junbi []string

func (j Junbi) MkDir() {
	for _, name := range j {
		if err := os.Mkdir(name, 0755); err != nil {
			log.Println(err)
		}
	}
}

func (j Junbi) MkFile() {
	for _, name := range j {
		if _, err := os.Create(name + "/1." + name); err != nil {
			log.Println(err)
		}
	}
}

type MainF struct {
	Name, Expand string
}

var expand = map[string]string{
	"rb": "ruby",
	"py": "python",
	"pl": "perl",
}

const ENCODE_TYPE = "# coding: utf-8"

func (m *MainF) Create() {
	s := ENCODE_TYPE
	f, err := os.OpenFile(m.Name+"."+m.Expand, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	e := expand[m.Expand]
	if e != "" {
		s = fmt.Sprintf("#! /usr/bin/env %s\n", e) + ENCODE_TYPE
	}
	f.WriteString(s)
}
