package metronome

import (
	"io"
	"sort"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/filter"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/midimessage/meta"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/writer"
)

type file struct {
	filePath          string
	metronomeTrackno  int16
	metricTicks       smf.MetricTicks
	beats             []uint64
	beatsMsec         beatsMsec
	tracks            []*track
	originalTempi     tempi
	realTempi         tempi
	metronomeDetector filter.Filter
}

// newFile returns a new file.
// the default metronome track no is 0, the default metronome detector is NoteOnMetronome
// on any channel for any key.
func newFile(filePath string, opts ...Option) *file {
	f := &file{
		filePath:          filePath,
		metronomeTrackno:  0,
		metronomeDetector: filter.NoteOn(-1),
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

func (f *file) convertTicksToMsec() {
	for _, t := range f.tracks {
		for _, ev := range *t {
			ev.msec = absPosToMsec(f.metricTicks, f.originalTempi, ev.absPos)
		}
	}
}

func (f *file) convertMsecToTicks() {
	for i, rt := range f.realTempi {
		rt.absPos = msecToAbsPos(f.metricTicks, f.realTempi, rt.msec)
		f.realTempi[i] = rt
	}
	for _, t := range f.tracks {
		for _, ev := range *t {
			ev.absPos = msecToAbsPos(f.metricTicks, f.realTempi, ev.msec)
		}
	}
}

// read reads the smf file and finds the metronome beats
func (f *file) read() error {
	var currentTrack = &track{}
	f.tracks = append(f.tracks, currentTrack)
	r := reader.New(
		reader.NoLogger(),
		reader.Each(func(p *reader.Position, msg midi.Message) {
			if msg == meta.EndOfTrack {
				currentTrack = &track{}
				f.tracks = append(f.tracks, currentTrack)
				return
			}

			switch v := msg.(type) {
			case meta.Tempo:
				tmpo := &tempo{absPos: p.AbsoluteTicks, bpm: v.FractionalBPM()}
				f.originalTempi = append(f.originalTempi, tmpo)
				return
			case channel.Message:
				if int16(len(f.tracks)-1) == f.metronomeTrackno {
					if f.metronomeDetector(v) {
						f.beats = append(f.beats, p.AbsoluteTicks)
					}
				}
			}
			*currentTrack = append(*currentTrack, &event{absPos: p.AbsoluteTicks, msg: msg})
		}),
		reader.SMFHeader(func(hd smf.Header) {
			f.metricTicks = hd.TimeFormat.(smf.MetricTicks)
		}),
	)
	return reader.ReadSMFFile(r, f.filePath)
}

// calcNewTempi calculates the new tempi
func (f *file) calcNewTempi() {
	for _, b := range f.beats {
		f.beatsMsec = append(f.beatsMsec, absPosToMsec(f.metricTicks, f.originalTempi, b))
	}
	f.realTempi = f.originalTempi._detectTempi(f.metricTicks, f.beatsMsec)
	for _, tempo := range f.realTempi {
		te := event{msec: tempo.msec, msg: meta.FractionalBPM(tempo.bpm)}
		*f.tracks[f.metronomeTrackno] = append(*f.tracks[f.metronomeTrackno], &te)
	}
	sort.Sort(f.tracks[f.metronomeTrackno])
}

// writeNew writes the new smf file
func (f *file) writeTo(filePath string) error {
	return writer.WriteSMF(filePath, uint16(len(f.tracks)), func(wr *writer.SMF) error {
		for _, tr := range f.tracks {
			sort.Sort(tr)
			lastPos := uint64(0)
			for _, ev := range *tr {
				wr.SetDelta(uint32(ev.absPos - lastPos))
				err := wr.Write(ev.msg)
				if err != nil {
					return err
				}
				lastPos = ev.absPos
			}
			err := wr.Write(meta.EndOfTrack)
			if err != nil && err != io.EOF {
				return err
			}
		}
		return nil
	})
}
