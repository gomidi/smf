package main

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/metakeule/config"
)

var CAT = cat{}

func init() {
	CAT.init()
}

type cat struct {
	*config.Config
	file config.StringGetter
}

func (c *cat) init() {
	c.Config = CONFIG.MustCommand("cat", "cat shows the content of an SMF (MIDI) file")
	c.file = c.NewString("file", "path of the midi file", config.Shortflag('f'), config.Required)
}

func (c *cat) print() error {
	rd := reader.New(reader.NoLogger(),
		reader.Each(func(pos *reader.Position, msg midi.Message) {
			fmt.Fprintf(os.Stdout, "T%v [%v] %s\n", pos.Track, pos.AbsoluteTicks, msg.String())
		}),
	)

	return reader.ReadSMFFile(rd, c.file.Get())
}
