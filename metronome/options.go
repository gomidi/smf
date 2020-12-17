package metronome

import (
	"gitlab.com/gomidi/midi/filter"
)

type Option func(*file)

// Track sets the track no, where the metronome resides, starting by 0
func Track(trackNo int16) Option {
	return func(f *file) {
		f.metronomeTrackno = trackNo
	}
}

// Filter sets the filter for the metronome messages
func Filter(md filter.Filter) Option {
	return func(f *file) {
		f.metronomeDetector = md
	}
}
