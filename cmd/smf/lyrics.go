package main

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/smf/lyrics"
	"gitlab.com/metakeule/config"
)

var LYRICS = lyr{}

func init() {
	LYRICS.init()
}

type lyr struct {
	*config.Config
	file        config.StringGetter
	track       config.Int32Getter
	includeText config.BoolGetter
	asJson      config.BoolGetter
}

func (c *lyr) init() {
	c.Config = CONFIG.MustCommand("lyrics", "extracts lyrics from a SMF file, tracks are separated by an empty line")
	c.file = c.NewString("file", "the SMF file that is read", config.Shortflag('f'), config.Required)
	c.track = c.NewInt32(
		"track",
		"the track from which the lyrics are taken. -1 means all tracks, 0 is the first, 1 the second etc",
		config.Shortflag('t'),
		config.Default(int32(-1)),
	)

	c.includeText = c.NewBool(
		"text",
		"include free text entries in the SMF file. Text is surrounded by doublequotes",
	)

	c.asJson = c.NewBool(
		"json",
		"output json format",
		config.Shortflag('j'),
	)
}

func (c *lyr) print() error {
	var options []lyrics.Option
	if c.track.Get() >= 0 {
		options = append(options, lyrics.OptionTrackNo(uint16(c.track.Get())))
	}

	if c.includeText.Get() {
		options = append(options, lyrics.OptionIncludeText())
	}

	if c.asJson.Get() {
		options = append(options, lyrics.OptionJSONOutput())
	}

	str, err := lyrics.Read(c.file.Get(), options...)

	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, str)
	return nil
}
