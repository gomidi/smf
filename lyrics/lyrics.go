package lyrics

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gitlab.com/gomidi/midi/reader"
)

type Option func(*config)

func OptionJSONOutput() Option {
	return func(o *config) {
		o.jsonOutput = true
	}
}

func OptionTrackNo(no uint16) Option {
	return func(o *config) {
		o.track = int16(no)
	}
}

func OptionIncludeText() Option {
	return func(o *config) {
		o.includeText = true
	}
}

type config struct {
	jsonOutput  bool
	includeText bool
	track       int16
}

func Read(file string, opts ...Option) (output string, err error) {
	var c = &config{
		track: -1,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.jsonOutput {
		j := newJson()
		j.File = file
		j.includeText = c.includeText
		j.requestedTrack = c.track

		err = reader.ReadSMFFile(j.reader, file)
		if err != nil {
			return "", err
		}

		bt, err := json.MarshalIndent(j, "", " ")

		if err != nil {
			return "", err
		}

		return string(bt), nil
	}

	p := newPrinter()
	p.includeText = c.includeText
	p.requestedTrack = c.track

	err = reader.ReadSMFFile(p.reader, file)
	output = p.bf.String()
	return
}

type printer struct {
	includeText    bool
	requestedTrack int16 // if < 0 : all tracks
	reader         *reader.Reader
	bf             bytes.Buffer
}

func newPrinter() *printer {
	var f printer

	f.reader = reader.New(
		reader.NoLogger(),
		reader.Lyric(f.Lyric),
		reader.Text(f.Text),
		reader.TrackSequenceName(f.TrackSequenceName),
		reader.Instrument(f.Instrument),
		reader.Program(f.Program),
		reader.EndOfTrack(f.EndOfTrack),
	)

	return &f
}

func (f *printer) shouldWrite(track int16) bool {
	return f.requestedTrack < 0 || f.requestedTrack == track
}

func (f *printer) Lyric(p reader.Position, text string) {
	if f.shouldWrite(p.Track) {
		fmt.Fprint(&f.bf, text+" ")
	}
}

func (f *printer) Text(p reader.Position, text string) {
	if f.shouldWrite(p.Track) && f.includeText {
		fmt.Fprint(&f.bf, "[ "+text+" ] ")
	}
}

func (f *printer) TrackSequenceName(p reader.Position, name string) {
	if f.shouldWrite(p.Track) {
		fmt.Fprintln(&f.bf, fmt.Sprintf("[track: %v]\n", name))
	}
}

func (f *printer) Instrument(p reader.Position, name string) {
	if f.shouldWrite(p.Track) {
		fmt.Fprintln(&f.bf, fmt.Sprintf("[instrument: %v]\n", name))
	}
}

func (f *printer) Program(p reader.Position, name string) {
	if f.shouldWrite(p.Track) {
		fmt.Fprintln(&f.bf, fmt.Sprintf("[program: %v]\n", name))
	}
}

func (f *printer) EndOfTrack(p reader.Position) {
	if f.shouldWrite(p.Track) && p.Track > 0 {
		fmt.Fprintf(&f.bf, "\n\n------------------------\n\n")
	}
}

func newJson() *Json {
	var t = &Json{}
	t.Tracks = []*Track{&Track{}}

	t.reader = reader.New(
		reader.NoLogger(),
		reader.Lyric(t.Lyric),
		reader.Text(t.Text),
		reader.TrackSequenceName(t.Track),
		reader.Instrument(t.Instrument),
		reader.Program(t.Program),
		reader.EndOfTrack(t.EndOfTrack),
	)

	return t
}

type Track struct {
	Program    string   `json:"program,omitempty"`
	Name       string   `json:"name,omitempty"`
	Instrument string   `json:"instrument,omitempty"`
	Texts      []string `json:"texts,omitempty"`
	Lyrics     []string `json:"lyrics,omitempty"`
	No         int      `json:"no"`
}

func (t *Track) addLyric(l string) {
	t.Lyrics = append(t.Lyrics, l)
}

func (t *Track) addText(tx string) {
	t.Texts = append(t.Texts, tx)
}

type Json struct {
	File           string `json:"file"`
	current        int
	Tracks         []*Track `json:"tracks,omitempty"`
	includeText    bool
	requestedTrack int16 // if < 0 : all tracks
	trackNo        int
	reader         *reader.Reader
}

func (t *Json) shouldWrite(track int16) bool {
	return t.requestedTrack < 0 || t.requestedTrack == track
}

func (t *Json) Lyric(p reader.Position, text string) {
	if t.shouldWrite(p.Track) {
		t.Tracks[t.current].addLyric(text)
	}
}

func (t *Json) Text(p reader.Position, text string) {
	if t.shouldWrite(p.Track) && t.includeText {
		t.Tracks[t.current].addText(text)
	}
}

func (t *Json) Track(p reader.Position, name string) {
	if t.shouldWrite(p.Track) {
		t.Tracks[t.current].Name = name
	}
}

func (t *Json) Instrument(p reader.Position, name string) {
	if t.shouldWrite(p.Track) {
		t.Tracks[t.current].Instrument = name
	}
}

func (t *Json) Program(p reader.Position, name string) {
	if t.shouldWrite(p.Track) {
		t.Tracks[t.current].Program = name
	}
}

func (t *Json) EndOfTrack(p reader.Position) {
	t.trackNo++
	if t.shouldWrite(p.Track) {
		t.Tracks = append(t.Tracks, &Track{No: t.trackNo})
		t.current = len(t.Tracks) - 1
	}
}
