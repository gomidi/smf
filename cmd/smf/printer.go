package main

import (
	"fmt"

	"gitlab.com/gomidi/smf"
	"gitlab.com/metakeule/config"
)

var PRINTER = printer{}

func init() {
	PRINTER.init()
}

type printer struct {
	*config.Config
	srcFile config.StringGetter
}

func (p *printer) init() {
	p.Config = CONFIG.MustCommand("printer", "print smf (pre ui for debugging)")
	p.srcFile = p.NewString("file", "source file", config.Shortflag('f'), config.Required)
}

func (p *printer) print() error {
	s, err := smf.ReadSMF(p.srcFile.Get())
	if err != nil {
		return err
	}
	fmt.Println(s.BarLines())
	return nil
}
