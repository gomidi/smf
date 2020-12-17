package main

import (
	"fmt"
	"os"

	. "gitlab.com/metakeule/config"
)

var CONFIG = MustNew("smf", "0.0.1", "tools to deal with SMF/MIDI files")

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err.Error())
		fmt.Fprint(os.Stderr, "USAGE\n\n"+CONFIG.Usage())
		os.Exit(1)
		return
	}
	os.Exit(0)
}

func run() error {
	err := CONFIG.Run()

	if err != nil {
		return err
	}

	switch CONFIG.ActiveCommand() {
	case METRONOME.Config:
		return METRONOME.setTempo()
	case PRINTER.Config:
		return PRINTER.print()
	case CAT.Config:
		return CAT.print()
	case LYRICS.Config:
		return LYRICS.print()
	default:
		fmt.Fprint(os.Stdout, CONFIG.Usage())
	}

	return nil
}

/*
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"

	"encoding/json"
	"sort"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/mid"
	_ "gitlab.com/gomidi/midiline/actions/channelchange"
	_ "gitlab.com/gomidi/midiline/actions/channelmirror"
	_ "gitlab.com/gomidi/midiline/actions/channelrouter"
	_ "gitlab.com/gomidi/midiline/actions/copy"
	_ "gitlab.com/gomidi/midiline/actions/metronome"
	_ "gitlab.com/gomidi/midiline/actions/outbreak"
	_ "gitlab.com/gomidi/midiline/actions/to_aftertouch"
	_ "gitlab.com/gomidi/midiline/actions/to_cc"
	_ "gitlab.com/gomidi/midiline/actions/to_note"
	_ "gitlab.com/gomidi/midiline/actions/to_pitchbend"
	_ "gitlab.com/gomidi/midiline/actions/to_polyaftertouch"
	_ "gitlab.com/gomidi/midiline/actions/to_programchange"
	_ "gitlab.com/gomidi/midiline/actions/transpose"
	_ "gitlab.com/gomidi/midiline/conditions/logic"
	_ "gitlab.com/gomidi/midiline/conditions/message"
	_ "gitlab.com/gomidi/midiline/conditions/typ"
	lineconfig "gitlab.com/gomidi/midiline/config"
	_ "gitlab.com/gomidi/midiline/value"
	"gitlab.com/gomidi/midiproxy"
	"gitlab.com/gomidi/rtmididrv"
	"gitlab.com/metakeule/config"
*/

/*
func main() {
	smf.New()
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "missing argument (file)")
	}
	song, err := smf.ReadSMF(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
	}

	fmt.Println(song.BarLines())
}
*/
