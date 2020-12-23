package smf

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/midimessage/meta"
	"gitlab.com/gomidi/midi/midimessage/meta/meter"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"
)

func ReadSMF(file string) (*Song, error) {
	s := New()
	// add the first track
	s.AddTrack(false, -1)

	var rd = reader.New(
		reader.NoLogger(),
		reader.Each(s.scanMessage),
	)

	err := reader.ReadSMFFile(rd, file)

	if err != nil {
		return nil, err
	}

	if rd.Header().Type() != 1 {
		return nil, fmt.Errorf("wrong SMF type %v, currently only supports type 1, please convert your midi file %s", rd.Header().Type(), file)
	}

	tpq, isMetric := rd.Header().TimeFormat.(smf.MetricTicks)

	if !isMetric {
		return nil, fmt.Errorf("wrong time format type %s, currently only supports metric time, please convert your midi file %s", rd.Header().TimeFormat.String(), file)
	}

	s.ticksPerQN = tpq.Ticks4th()

	err = s.finishScan()

	if err != nil {
		return nil, err
	}

	return s, nil
}

type TempoChange struct {
	AbsPos   uint64
	TempoBPM float64
}

type TempoChanges []*TempoChange

func (p TempoChanges) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

func (p TempoChanges) Len() int {
	return len(p)
}

func (p TempoChanges) Less(a, b int) bool {
	return p[a].AbsPos < p[b].AbsPos
}

type TimeSigs []*TimeSig

func (p TimeSigs) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

func (p TimeSigs) Len() int {
	return len(p)
}

func (p TimeSigs) Less(a, b int) bool {
	return p[a].AbsPos < p[b].AbsPos
}

type TimeSig struct {
	AbsPos uint64
	Num    uint8
	Denom  uint8
}

type Song struct {
	CopyRight           string
	Properties          map[string]string
	Bars                []*Bar
	Tracks              []*Track
	scannedMessages     []*TrackMessage
	scannedTempoChanges TempoChanges
	scannedTimeSig      TimeSigs
	ticksPerQN          uint32
	lastPos             uint64
}

func (s *Song) AddTrack(withContent bool, channel int8) *Track {
	t := &Track{
		Song:        s,
		WithContent: withContent,
		Channel:     channel,
	}

	s.Tracks = append(s.Tracks, t)
	s.RenumberTracks()
	return t
}

func (s *Song) createBarsUntil(from, to uint64, num, denom uint8) {
	pos := from

	for pos < to {
		b := s.AddBar(pos, num, denom)
		pos = b.EndPos()
	}
}

func (s *Song) createBars(firstTimeSig [2]uint8, changes TimeSigs) {
	num := firstTimeSig[0]
	denom := firstTimeSig[1]
	var pos uint64

	for _, change := range changes {
		if pos < change.AbsPos {
			s.createBarsUntil(pos, change.AbsPos, num, denom)
		}
		num, denom = change.Num, change.Denom
		b := s.AddBar(change.AbsPos, num, denom)
		pos = b.EndPos()
	}

	s.createBarsUntil(pos, s.lastPos, num, denom)
	s.RenumberBars()
}

func (s *Song) findBar(pos uint64) (bar *Bar) {
	for _, b := range s.Bars {
		if pos >= b.AbsPos {
			bar = b
		}
	}
	return
}

func (s *Song) NoOfContentTracks() (no uint16) {
	for _, tr := range s.Tracks {
		if tr.WithContent {
			no++
		}
	}
	return
}

func (s *Song) finishScan() (err error) {
	sort.Sort(s.scannedTempoChanges)
	sort.Sort(s.scannedTimeSig)

	if len(s.scannedTimeSig) > 0 && s.scannedTimeSig[0].AbsPos == 0 {
		var rest TimeSigs
		if len(s.scannedTimeSig) > 1 {
			rest = s.scannedTimeSig[1:]
		}
		s.createBars([2]uint8{s.scannedTimeSig[0].Num, s.scannedTimeSig[0].Denom}, rest)
	} else {
		s.createBars([2]uint8{4, 4}, s.scannedTimeSig)
	}

	for _, msg := range s.scannedMessages {
		b := s.findBar(msg.AbsPos)
		if b == nil {
			return fmt.Errorf("can't find bar for message: %v at position %v", msg.Message, msg.AbsPos)
		}
		b.SetMessageByRelTicks(msg.AbsPos-b.AbsPos, msg.TrackNo, msg.Message)
		b.SortPositions()
	}

	return nil
}

func (s *Song) LastTrack() *Track {
	return s.Tracks[len(s.Tracks)-1]
}

func (s *Song) scanMessage(p *reader.Position, msg midi.Message) {
	if p.AbsoluteTicks > s.lastPos {
		s.lastPos = p.AbsoluteTicks
	}

	if msg == meta.EndOfTrack {
		s.AddTrack(false, -1)
		return
	}

	t := s.LastTrack()
	switch m := msg.(type) {
	case meta.Copyright:
		s.CopyRight = m.Text()
	case meta.TrackSequenceName:
		t.Name = m.Text()
	case meta.Instrument:
		t.Instrument = m.Text()
	case meta.TimeSig:
		ts := &TimeSig{
			AbsPos: p.AbsoluteTicks,
			Num:    m.Numerator,
			Denom:  m.Denominator,
		}
		s.scannedTimeSig = append(s.scannedTimeSig, ts)
	case meta.Tempo:
		tc := &TempoChange{
			AbsPos:   p.AbsoluteTicks,
			TempoBPM: m.FractionalBPM(),
		}
		s.scannedTempoChanges = append(s.scannedTempoChanges, tc)
	default:
		if msg != nil {
			tm := &TrackMessage{}
			tm.Message = msg
			tm.TrackNo = t.No
			tm.AbsPos = p.AbsoluteTicks
			s.scannedMessages = append(s.scannedMessages, tm)
			t.WithContent = true
			if chMsg, is := msg.(channel.Message); is {
				if t.Channel >= 0 && uint8(t.Channel) != chMsg.Channel() {
					panic(fmt.Sprintf("track no %v (%s) has mixed channel messages for channel %v and %v - not supported", t.No, t.Name, t.Channel, chMsg.Channel()))
				}

				if t.Channel < 0 {
					t.Channel = int8(chMsg.Channel())
				}
			}
		}
	}
}

func (s *Song) AddBar(pos uint64, num, denom uint8) *Bar {
	b := &Bar{
		AbsPos:  pos,
		Song:    s,
		TimeSig: [2]uint8{num, denom},
	}

	s.Bars = append(s.Bars, b)

	return b
}

func (s *Song) Save(file string) error {
	return writer.WriteSMF(file, s.NoOfContentTracks()+2, s.writeSMF, smfwriter.Format(smf.SMF1), smfwriter.TimeFormat(smf.MetricTicks(s.ticksPerQN)))
}

func (s *Song) writeTimeSigTrack(w *writer.SMF) error {
	timesig := [2]uint8{4, 4}
	var pos uint64

	for _, b := range s.Bars {
		if b.TimeSig != timesig {
			delta := uint32(b.AbsPos - pos)
			w.SetDelta(delta)
			w.Write(meter.Meter(b.TimeSig[0], b.TimeSig[1]))
			timesig = b.TimeSig
			pos = b.AbsPos
		}
	}
	return nil
}

func (s *Song) writeTempoTrack(w *writer.SMF) error {
	tempo := float32(120.0)

	var pos uint64
	for _, b := range s.Bars {
		for _, p := range b.Positions {
			if p.Tempo != 0 && p.Tempo != tempo {
				absPos := p.AbsTicks()
				delta := uint32(absPos - pos)
				w.SetDelta(delta)
				w.Write(meta.Tempo(p.Tempo))
				tempo = p.Tempo
				pos = absPos
			}
		}
	}

	return nil
}

func (s *Song) writeSMF(w *writer.SMF) (err error) {
	err = s.writeTimeSigTrack(w)
	if err != nil {
		return
	}

	err = writer.EndOfTrack(w)
	if err != nil {
		return
	}

	err = s.writeTempoTrack(w)
	if err != nil {
		return
	}

	err = writer.EndOfTrack(w)
	if err != nil {
		return
	}

	for _, tr := range s.Tracks {
		if tr.WithContent {
			var lastTick uint64
			for _, b := range s.Bars {
				for _, p := range b.Positions {
					ticks := p.AbsTicks()

					for _, m := range p.Messages {
						if m.TrackNo == tr.No && m.Message != nil {
							delta := ticks - lastTick
							if tr.Channel < 0 {
								panic(fmt.Sprintf("channel for content track no %v (%s) is -1, but content tracks must have channels", tr.No, tr.Name))
							}
							w.SetChannel(uint8(tr.Channel))
							w.SetDelta(uint32(delta))
							w.Write(m.Message)
							lastTick = ticks
						}
					}
				}
			}

			err = writer.EndOfTrack(w)
			if err != nil {
				return
			}
		}
	}

	return nil
}

func (s *Song) TrackWidth(i int) uint8 {
	// TODO calculate the track width
	return 0
}

func KeyToNote(key uint8) string {
	nt := key % 12
	oct := key / 12

	notes := map[uint8]string{
		0:  "C",
		1:  "C#",
		2:  "D",
		3:  "D#",
		4:  "E",
		5:  "F",
		6:  "F#",
		7:  "G",
		8:  "G#",
		9:  "A",
		10: "A#",
		11: "B",
	}

	return fmt.Sprintf("%s%v", notes[nt], oct)
}

func ShowMessage(msg midi.Message) string {
	switch v := msg.(type) {
	case channel.NoteOn:
		return fmt.Sprintf("%s/%v_", KeyToNote(v.Key()), v.Velocity())
	case channel.NoteOff:
		return fmt.Sprintf("_%s", KeyToNote(v.Key()))
	case channel.NoteOffVelocity:
		return fmt.Sprintf("_%s", KeyToNote(v.Key()))
	/*
		case channel.Aftertouch:
		case channel.ControlChange:
		case channel.Pitchbend:
		case channel.PolyAftertouch:
		case channel.ProgramChange:
	*/
	case meta.Lyric:
		return fmt.Sprintf("%q", v.Text())
	case meta.Text:
		return fmt.Sprintf("'%s'", v.Text())
	default:
		return msg.String()
	}
}

func (s *Song) BarLines() string {
	var bf bytes.Buffer

	fmt.Fprintf(&bf, "| Comment | Mark | Tempo  | Beat | ")

	for _, t := range s.Tracks {
		if t.WithContent {
			fmt.Fprintf(&bf, " %s[%v] | ", t.Name, t.Channel)
		}
	}

	fmt.Fprintf(&bf, "\n")

	for _, b := range s.Bars {
		_ = b
		fmt.Fprintf(&bf, "----------- #%v %v/%v --------------\n", b.No, b.TimeSig[0], b.TimeSig[1])
		for _, p := range b.Positions {
			tempo := ""
			if p.Tempo != 0 {
				tempo = fmt.Sprintf("%0.2f", tempo)
			}

			var frac float64

			if p.Fraction[1] > 0 {
				frac = p.Fraction[0] / p.Fraction[1]
			}

			beat := fmt.Sprintf("%0.4f", float64(p.Beat)+float64(1)+frac)

			fmt.Fprintf(&bf, "| %s | %s | %s | %s | ", p.Comment, p.Mark, tempo, beat)

			for _, t := range s.Tracks {
				if t.WithContent {
					var printed bool
					for _, m := range p.Messages {
						if m.TrackNo == t.No {
							fmt.Fprintf(&bf, " %s | ", ShowMessage(m.Message))
							printed = true
						}
					}
					if !printed {
						fmt.Fprintf(&bf, "  | ")
					}
				}
			}

			fmt.Fprintf(&bf, "\n")
		}
	}

	return bf.String()
}

func (s *Song) RenumberBars() {
	for i := range s.Bars {
		s.Bars[i].No = uint16(i)
	}
}

func (s *Song) RenumberTracks() {
	for i := range s.Tracks {
		s.Tracks[i].No = uint16(i)
	}
}

type TrackMessage struct {
	TrackNo  uint16
	AbsPos   uint64
	Message  midi.Message
	Position *Position
}

type Positions []*Position

func (p Positions) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

func (p Positions) Len() int {
	return len(p)
}

func (p Positions) Less(a, b int) bool {
	if p[a].Bar.No < p[b].Bar.No {
		return true
	}

	if p[a].Bar.No > p[b].Bar.No {
		return false
	}

	if p[a].Beat < p[b].Beat {
		return true
	}

	if p[a].Beat > p[b].Beat {
		return false
	}

	var frac_a float64
	var frac_b float64

	if p[a].Fraction[1] > 0 {
		frac_a = float64(p[a].Fraction[0]) / float64(p[a].Fraction[1])
	}

	if p[b].Fraction[1] > 0 {
		frac_b = float64(p[b].Fraction[0]) / float64(p[b].Fraction[1])
	}

	return frac_a < frac_b
}

type Bar struct {
	Song      *Song
	No        uint16
	TimeSig   [2]uint8
	Positions Positions
	AbsPos    uint64
}

func (b *Bar) EndPos() uint64 {
	return b.AbsPos + b.Length()
}

func (b *Bar) Length() uint64 {
	l := float64(b.Song.ticksPerQN*4*uint32(b.TimeSig[0])) / float64(b.TimeSig[1])
	return uint64(math.Round(l))
}

func (b *Bar) SetMessageByRelTicks(ticks uint64, trackNo uint16, msg midi.Message) {
	beat := uint8(ticks / uint64(b.Song.ticksPerQN))
	ticksRest := ticks % uint64(b.Song.ticksPerQN)

	var pos *Position

	for _, p := range b.Positions {
		if p.Beat == beat && p.WithinFraction(ticksRest) {
			pos = p
			break
		}
	}

	if pos == nil {
		pos = b.AddPosition()
		pos.Beat = beat
		pos.Fraction[0] = float64(ticksRest)
		pos.Fraction[1] = float64(b.Song.ticksPerQN)
	}

	pos.SetMessage(trackNo, msg)
}

func (b *Bar) AddPosition() *Position {
	p := &Position{
		Bar: b,
	}

	b.Positions = append(b.Positions, p)
	return p
}

func (b *Bar) Columns() []string {
	//cols := make([]string)
	return nil
}

func (b *Bar) SortPositions() {
	sort.Sort(b.Positions)
}

type Position struct {
	Bar      *Bar
	Comment  string
	Mark     string
	Beat     uint8
	Tempo    float32
	Fraction [2]float64
	Messages []*TrackMessage
}

/*
WithinFraction determines, if the given ticks are within the fraction of the position.
The given ticks must be less than a quarternote (Songs ticks per quarternote).
The fraction is a fraction of a quarternote. So we first have to check, to which fraction
of the qn the given ticks correspond and then to check, if the difference between this fraction
and the fraction of the Position lies within the tolerance
*/

func (p *Position) WithinFraction(ticks uint64) bool {
	//tolerance := float64(0.0000001)
	tolerance := float64(0.001)
	fracTicks := float64(ticks) / float64(p.Bar.Song.ticksPerQN)
	if fracTicks >= 1 {
		panic("must not happen, we are on the wrong beat")
	}
	fracPos := p.Fraction[0] / p.Fraction[1]

	//fmt.Printf("\nwithin fraction %v vs %v (ticks: %v perQN: %v)\n", fracPos, fracTicks, ticks, p.Bar.Song.ticksPerQN)

	return math.Abs(fracPos-fracTicks) < tolerance
}

func (p *Position) AbsTicks() uint64 {
	beatTicks := p.Bar.Song.ticksPerQN * uint32(p.Beat)
	fracTicks := math.Round((float64(p.Bar.Song.ticksPerQN) * p.Fraction[0]) / p.Fraction[1])
	return p.Bar.AbsPos + uint64(beatTicks) + uint64(fracTicks)
}

func (p *Position) AddMessage(track uint16, msg midi.Message) {
	tm := &TrackMessage{
		TrackNo:  track,
		Message:  msg,
		Position: p,
	}

	p.Messages = append(p.Messages, tm)
}

func (p *Position) GetMessage(track uint16) *TrackMessage {
	for _, m := range p.Messages {
		if m.TrackNo == track {
			return m
		}
	}

	return nil
}

func (p *Position) SetMessage(track uint16, msg midi.Message) {
	var tm *TrackMessage
	for _, m := range p.Messages {
		if m.TrackNo == track {
			tm = m
			break
		}
	}

	if tm == nil {
		p.AddMessage(track, msg)
		return
	}

	tm.Message = msg
}

type Track struct {
	Song        *Song
	No          uint16
	Channel     int8 // -1 == not defined
	Name        string
	Instrument  string
	Solo        bool
	Mute        bool
	RecordArm   bool
	WithContent bool
	External    bool // for non editable track
}

/*

-----------------------------------------------------------------------
File | Edit | View | Config         (the menu, open the first with ALT+SPACE and then navigate with arrow keys and select with ENTER)
-----------------------------------------------------------------------
Comment | Mark  | Bar  | Beat || Drums[10] | Bass[9]  | Vocal[1] | Piano[1]  | (piano track on channel 1 etc)
        |       |      |      || S M R     | S M R    | S M R    | S M R     | (Solo/Mute/Record indicators)
----------------------------------------------------------------------- (everything above this line is static/non scrollable)
1       | Intro | 4/4  | 1.0  || C3/100    | C5_/120  |          |           | (drum note is just a 32ths, bass is note on)
        |       | 144  | 1.0  ||           |          |          |           | tempo change
                  #2                                                           (bar change)
        |       |      | 2.25 || C5/60     | _C5      | "hiho"   | CC123/100 |
=====I====>===V====C=== position indicator, always the pre-last line of the screen (each = is a bar, each letter is the first letter of a Marker)
F1 Play | F2 Rec | F3 Metro | F4 Keyb | F5 V1 | F6 V2 | F7 V3 | F8 V4 | F9 V5 | F10 Track Properties | F11 Song Properties
(play, record, metronome, Keyboard are switches that indicate if it is active)
(views are a selector; only one view can be active at a time)

*/

func New() *Song {
	return &Song{
		Properties: map[string]string{},
	}
}
