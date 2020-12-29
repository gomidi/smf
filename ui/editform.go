package ui

import (
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/midimessage/meta"
	"gitlab.com/gomidi/midi/midimessage/sysex"

	//	"github.com/gomidi/connect/rtmidiadapter"
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func (n *EditForm) addProgramFields(v meta.Program) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Program(text)
	})
}

func (n *EditForm) addInstrumentFields(v meta.Instrument) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Instrument(text)
	})
}

func (n *EditForm) addTrackSequenceNameFields(v meta.TrackSequenceName) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.TrackSequenceName(text)
	})
}

func (n *EditForm) addCopyrightFields(v meta.Copyright) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Copyright(text)
	})
}

func (n *EditForm) addCuepointFields(v meta.Cuepoint) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Cuepoint(text)
	})
}

func (n *EditForm) addTextFields(v meta.Text) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Text(text)
	})
}

func (n *EditForm) addLyricFields(v meta.Lyric) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Lyric(text)
	})
}

func (n *EditForm) addMarkerFields(v meta.Marker) {
	n.AddInputField("text", printv(v.Text()), 300, func(textToCheck string, lastChar rune) bool {
		return true
	}, func(text string) {
		n.Message = meta.Marker(text)
	})
}

func (n *EditForm) addPitchbendFields(v channel.Pitchbend) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).Aftertouch(v.Pressure())
		})
	*/
	n.AddInputField("value", printv(v.Value()), 3, func(textToCheck string, lastChar rune) bool {
		i, err := strconv.Atoi(textToCheck)
		if err != nil {
			return false
		}
		if i < channel.PitchLowest || i > channel.PitchHighest {
			return false
		}
		return true
	}, func(text string) {
		i, err := strconv.Atoi(text)
		if err != nil {
			return
		}
		if i < channel.PitchLowest || i > channel.PitchHighest {
			return
		}
		n.Message = channel.Channel(v.Channel()).Pitchbend(int16(i))
	})
}

func (n *EditForm) addAftertouchFields(v channel.Aftertouch) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).Aftertouch(v.Pressure())
		})
	*/
	n.AddInputField("pressure", printv(v.Pressure()), 3, func(textToCheck string, lastChar rune) bool {
		i, ok := getUint8(textToCheck)
		if i < 1 {
			ok = false
		}
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if i < 1 || !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).Aftertouch(i)
	})
}

func (n *EditForm) addProgramChangeFields(v channel.ProgramChange) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).Aftertouch(v.Pressure())
		})
	*/
	n.AddInputField("program", printv(v.Program()), 3, func(textToCheck string, lastChar rune) bool {
		i, ok := getUint8(textToCheck)
		if i < 1 {
			ok = false
		}
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if i < 1 || !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).ProgramChange(i)
	})
}

//addPolyAftertouchFields
func (n *EditForm) addPolyAftertouchFields(v channel.PolyAftertouch) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).PolyAftertouch(v.Key(), v.Pressure())
		})
	*/
	n.AddInputField("key", printv(v.Key()), 3, func(textToCheck string, lastChar rune) bool {
		_, ok := getUint8(textToCheck)
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).PolyAftertouch(i, v.Pressure())
	})

	n.AddInputField("pressure", printv(v.Pressure()), 3, func(textToCheck string, lastChar rune) bool {
		i, ok := getUint8(textToCheck)
		if i < 1 {
			ok = false
		}
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if i < 1 || !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).PolyAftertouch(v.Key(), i)
	})
}

func (n *EditForm) addNoteOnFields(v channel.NoteOn) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).NoteOn(v.Key(), v.Velocity())
		})
	*/
	n.AddInputField("key", printv(v.Key()), 3, func(textToCheck string, lastChar rune) bool {
		_, ok := getUint8(textToCheck)
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).NoteOn(i, v.Velocity())
	})

	n.AddInputField("velocity", printv(v.Velocity()), 3, func(textToCheck string, lastChar rune) bool {
		i, ok := getUint8(textToCheck)
		if i < 1 {
			ok = false
		}
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if i < 1 || !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).NoteOn(v.Key(), i)
	})
}

func (n *EditForm) addCCFields(v channel.ControlChange) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).ControlChange(v.Controller(), v.Value())
		})
	*/
	n.AddInputField("controller", printv(v.Controller()), 3, func(textToCheck string, lastChar rune) bool {
		_, ok := getUint8(textToCheck)
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).ControlChange(i, v.Value())
	})

	n.AddInputField("value", printv(v.Value()), 3, func(textToCheck string, lastChar rune) bool {
		_, ok := getUint8(textToCheck)
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).ControlChange(v.Controller(), i)
	})
}

func (n *EditForm) addNoteOffFields(v channel.NoteOff) {
	/*
		n.AddInputField("channel", printv(v.Channel()), 2, func(textToCheck string, lastChar rune) bool {
			i, ok := getUint8(textToCheck)
			if i < 1 || i > 15 {
				ok = false
			}
			return ok
		}, func(text string) {
			i, ok := getUint8(text)
			if i < 1 || i > 15 || !ok {
				return
			}
			n.Message = channel.Channel(i).NoteOff(v.Key())
		})
	*/
	n.AddInputField("key", printv(v.Key()), 3, func(textToCheck string, lastChar rune) bool {
		_, ok := getUint8(textToCheck)
		return ok
	}, func(text string) {
		i, ok := getUint8(text)
		if !ok {
			return
		}
		n.Message = channel.Channel(v.Channel()).NoteOff(i)
	})
}

func newMessageInTrack(ch uint8, typ string) midi.Message {
	switch typ {
	case "noteon":
		return channel.Channel(ch).NoteOn(60, 120)
	case "noteoff":
		return channel.Channel(ch).NoteOff(60)
	case "controlchange":
		return channel.Channel(ch).ControlChange(0, 100)
	case "aftertouch":
		return channel.Channel(ch).Aftertouch(100)
	case "pitchbend":
		return channel.Channel(ch).Pitchbend(0)
	case "polyaftertouch":
		return channel.Channel(ch).PolyAftertouch(60, 100)
	case "programchange":
		return channel.Channel(ch).ProgramChange(0)
	case "copyright":
		return meta.Copyright("")
	case "cuepoint":
		return meta.Cuepoint("")
	case "lyric":
		return meta.Lyric("")
	case "marker":
		return meta.Marker("")
	case "text":
		return meta.Text("")
	default:
		panic("unknown typ " + typ)
	}
}

type messageSelect struct {
	*tview.List
}

var _messagesSelect = []string{
	"noteon",
	"noteoff",
	"controlchange",
	"aftertouch",
	"pitchbend",
	"polyaftertouch",
	"programchange",
	"copyright",
	"cuepoint",
	"lyric",
	"marker",
	"text",
}

func newMessageSelect() *messageSelect {
	var s messageSelect
	s.List = tview.NewList()
	return &s

	//l.GetCurrentItem()
}

func (ms *messageSelect) AddItems(selected func()) {
	for _, m := range _messagesSelect {
		ms.AddItem(m, "", rune(m[0]+m[1]), selected)
	}

	//l.GetCurrentItem()
}

func NewEditForm(msg midi.Message, saveCb func(m midi.Message), cancelCb func()) *EditForm {
	if msg == nil {
		panic("msg must not be nil")
	}
	var n EditForm
	n.Form = tview.NewForm()
	n.Message = msg
	n.Form.SetTitle(fmt.Sprintf("%T", n.Message))
	n.Form.SetTitleAlign(tview.AlignCenter)
	n.Form.SetBorder(true)

	switch v := n.Message.(type) {
	case channel.NoteOn:
		n.addNoteOnFields(v)
	case channel.NoteOff:
		n.addNoteOffFields(v)
	case channel.NoteOffVelocity:
		n.addNoteOffFields(channel.Channel(v.Channel()).NoteOff(v.Key()))
	case channel.ControlChange:
		n.addCCFields(v)
	case channel.Aftertouch:
		n.addAftertouchFields(v)
	case channel.Pitchbend:
		n.addPitchbendFields(v)
	case channel.PolyAftertouch:
		n.addPolyAftertouchFields(v)
	case channel.ProgramChange:
		n.addProgramChangeFields(v)
	case meta.Copyright:
		n.addCopyrightFields(v)
	case meta.Cuepoint:
		n.addCuepointFields(v)
	case meta.Lyric:
		n.addLyricFields(v)
	case meta.Marker:
		n.addMarkerFields(v)
	case meta.Text:
		n.addTextFields(v)
	case meta.Program:
		n.addProgramFields(v)
	case meta.Instrument:
		n.addInstrumentFields(v)
	case meta.TrackSequenceName:
		n.addTrackSequenceNameFields(v)
	case meta.Tempo:
	case meta.TimeSig:
	case meta.Key:
	case meta.SequenceNo:
	case meta.SequencerData:
	case meta.Channel:
	case meta.Device:
	case meta.Port:
	case sysex.SysEx:
	default:
		panic(fmt.Sprintf("%T not implemented yet", msg))

	}

	n.AddButton("save", func() {
		saveCb(n.Message)
	})
	n.AddButton("cancel", func() {
		cancelCb()
	})

	return &n
}

type EditForm struct {
	*tview.Form
	Message midi.Message
}
