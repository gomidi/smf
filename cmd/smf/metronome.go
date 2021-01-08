package main

import (
	"fmt"
	"strings"

	"gitlab.com/gomidi/midi/filter"
	. "gitlab.com/gomidi/smf/metronome"
	"gitlab.com/metakeule/config"
)

var METRONOME = metro{}

func init() {
	METRONOME.init()
}

type metro struct {
	*config.Config
	srcFile  config.StringGetter
	destFile config.StringGetter
	trackNo  config.Int32Getter
	typ      config.StringGetter
	channel  config.Int32Getter
	value    config.Int32Getter
}

func (s *metro) init() {
	s.Config = CONFIG.MustCommand("metro", "set the tempo by analysing the beats inside a metronome track").Skip("midifile")
	s.srcFile = s.LastString("midifile", "source file", config.Required)
	s.destFile = s.NewString("out", "output file", config.Shortflag('o'))
	s.trackNo = s.NewInt32("track", "track no where only the metronome beats reside", config.Shortflag('t'), config.Default(int32(0)), config.Required)
	s.typ = s.NewString("mtype", "type of the metronome beats; available types are: 'no' (NoteOn), 'cc' (ControlChange), 'at' (Aftertouch), 'pa' (Polyaftertouch), 'pb' (PitchBend)", config.Default("no"))
	s.channel = s.NewInt32("mchan", "channel of the metronome beats (-1=all)", config.Default(int32(-1)))
	s.value = s.NewInt32("mval", "value of the metronome beats (-1=all); values the meaning for the different types is: NoteOn:  key, ControlChange: controller, Polyaftertouch: key", config.Default(int32(-1)))

}

func (s metro) setTempo() error {
	destFile := s.destFile.Get()
	if !s.destFile.IsSet() {
		destFile = s.srcFile.Get()
	}

	var opts []Option
	opts = append(opts, Track(int16(s.trackNo.Get())))

	// default: all channels
	var chFilter = filter.Channel(-1)

	if s.channel.IsSet() {
		chFilter = filter.Channel(int8(s.channel.Get()))
	}

	// default: all notes
	var fTyp = filter.NoteOn(-1)

	if s.typ.IsSet() {
		v := strings.ToLower(strings.TrimSpace(s.typ.Get()))
		switch v {
		case "no":
			fTyp = filter.NoteOn(int8(s.value.Get()))
		case "cc":
			fTyp = filter.CC(int8(s.value.Get()))
		case "at":
			fTyp = filter.Aftertouch()
		case "pa":
			fTyp = filter.PolyAftertouch(int8(s.value.Get()))
		case "pb":
			fTyp = filter.Pitchbend()
		default:
			return fmt.Errorf(`unknown type of metronome: %v, known types are: 
'no' (NoteOn), 
'cc' (ControlChange), 
'at' (Aftertouch), 
'pa' (Polyaftertouch), 
'pb' (PitchBend)
`, v)
		}
	}

	opts = append(opts, Filter(filter.And(chFilter, fTyp)))
	return SetTempo(s.srcFile.Get(), destFile, opts...)
}
