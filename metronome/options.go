package metronome

import (
	"gitlab.com/gomidi/midi/midimessage/channel"
)

type Option func(*file)

// OptionTrack sets the track no, where the metronome resides, starting by 0
func OptionTrack(trackNo int16) Option {
	return func(f *file) {
		f.metronomeTrackno = trackNo
	}
}

// OptionFilter sets the filter for the metronome messages
func OptionFilter(md Filter) Option {
	return func(f *file) {
		f.metronomeDetector = md
	}
}

// OrFilter returns a filter that is true if any of the given filters is true
func OrFilter(filters ...Filter) Filter {
	return func(msg channel.Message) bool {
		for _, f := range filters {
			if f(msg) {
				return true
			}
		}
		return false
	}
}

// AndFilter returns a filter that is true if all of the given filters are true
func AndFilter(filters ...Filter) Filter {
	return func(msg channel.Message) bool {
		for _, f := range filters {
			if !f(msg) {
				return false
			}
		}
		return true
	}
}

// ChannelFilter returns a Filter that triggers only for messages on the given MIDI channel.
// If ch is < 0, any channel will trigger.
func ChannelFilter(ch int8) Filter {
	return func(msg channel.Message) bool {
		return (ch < 0) || uint8(ch) == msg.Channel()
	}
}

// Filter detects  metronome beats on the metronome track.
// It is a function that returns true when a metronome beat was detected.
type Filter func(msg channel.Message) bool

// NoteOnFilter returns a Filter that triggers only for the given key.
// If key is < 0, any key will trigger.
func NoteOnFilter(key int8) Filter {
	return func(msg channel.Message) bool {
		switch nt := msg.(type) {
		case channel.NoteOn:
			if nt.Velocity() == 0 {
				return false
			}
			return key < 0 || uint8(key) == nt.Key()
		default:
			return false
		}
	}
}

// CCFilter returns a Filter that triggers only for the given controller.
// If controller is < 0, any controller will trigger.
func CCFilter(controller int8) Filter {
	return func(msg channel.Message) bool {
		switch c := msg.(type) {
		case channel.ControlChange:
			if c.Value() == 0 {
				return false
			}
			return controller < 0 || uint8(controller) == c.Controller()
		default:
			return false
		}
	}
}

// AftertouchFilter returns a Filter that triggers only for the aftertouch messages.
func AftertouchFilter() Filter {
	return func(msg channel.Message) bool {
		switch at := msg.(type) {
		case channel.Aftertouch:
			return at.Pressure() != 0
		default:
			return false
		}
	}
}

// PolyAftertouchFilter returns a Filter that triggers only for the polyaftertouch messages of the given key.
// If key is < 0, any key will trigger.
func PolyAftertouchFilter(key int8) Filter {
	return func(msg channel.Message) bool {
		switch pa := msg.(type) {
		case channel.PolyAftertouch:
			if pa.Pressure() == 0 {
				return false
			}
			return key < 0 || uint8(key) == pa.Key()
		default:
			return false
		}
	}
}

// PitchbendFilter returns a Filter that triggers only for the pitchbend messages > 0.
func PitchbendFilter() Filter {
	return func(msg channel.Message) bool {
		switch pb := msg.(type) {
		case channel.Pitchbend:
			return pb.Value() > 0
		default:
			return false
		}
	}
}
