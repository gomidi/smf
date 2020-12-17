package metronome

import (
	"sort"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/smf"
)

type tempo struct {
	absPos uint64
	msec   int64
	bpm    float64
}

type tempi []*tempo

func (t tempi) Swap(a, b int) {
	t[a], t[b] = t[b], t[a]
}

func (t tempi) Less(a, b int) bool {
	if t[a].absPos == 0 || t[b].absPos == 0 {
		return t[a].msec < t[b].msec
	}
	return t[a].absPos < t[b].absPos
}

func (t tempi) Len() int {
	return len(t)
}

type beatsMsec []int64

func (original tempi) getTempoForPos(pos uint64) (bpm float64) {
	bpm = 120.0

	for _, o := range original {
		if o.absPos <= pos {
			bpm = o.bpm
		}
	}

	return bpm
}

func (original tempi) _detectTempi(metricTicks smf.MetricTicks, bs beatsMsec) (realTempi tempi) {
	sort.Sort(original)

	prev := int64(0)

	for _, b := range bs {
		if b == 0.0 {
			continue
		}

		realBPM := timeDistaneToTempo(prev, b)
		realTempi = append(realTempi, &tempo{msec: prev, bpm: realBPM})
		prev = b
	}

	return realTempi
}

type event struct {
	absPos uint64
	msec   int64
	msg    midi.Message
}

type track []*event

func (t track) Swap(a, b int) {
	t[a], t[b] = t[b], t[a]
}

func (t track) Less(a, b int) bool {
	if t[a].absPos == 0 || t[b].absPos == 0 {
		return t[a].msec < t[b].msec
	}
	return t[a].absPos < t[b].absPos
}

func (t track) Len() int {
	return len(t)
}
