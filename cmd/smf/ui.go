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
	c.Config = CONFIG.MustCommand("ui", "show UI for a SMF file").Skip("midifile")
	c.file = c.LastString("midifile", "the SMF file that is edited", config.Required)
}

func (c *_ui) show() error {
	return ui.New(c.file.Get()).Run()
}
