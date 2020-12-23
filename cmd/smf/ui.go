package main

import (
	"gitlab.com/gomidi/smf/ui"
	"gitlab.com/metakeule/config"
)

var UI = _ui{}

func init() {
	UI.init()
}

type _ui struct {
	*config.Config
	file config.StringGetter
}

func (c *_ui) init() {
	c.Config = CONFIG.MustCommand("ui", "show UI for a SMF file")
	c.file = c.NewString("file", "the SMF file that is shown", config.Shortflag('f'), config.Required)
}

func (c *_ui) show() error {
	return ui.StartUI(c.file.Get())
}
