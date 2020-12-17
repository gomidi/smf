package metronome

import (
	"gitlab.com/gomidi/midi/filter"
)

type Option func(*file)

// OptionTrack sets the track no, where the metronome resides, starting by 0
func OptionTrack(trackNo int16) Option {
	return func(f *file) {
		f.metronomeTrackno = trackNo
	}
}

// OptionFilter sets the filter for the metronome messages
func OptionFilter(md filter.Filter) Option {
	return func(f *file) {
		f.metronomeDetector = md
	}
}
