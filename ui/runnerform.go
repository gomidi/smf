package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/cc"
	"gitlab.com/gomidi/midi/gm"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/midimessage/meta"
	"gitlab.com/gomidi/smf"
)

func (s *runnerScreen) addRunButton(form *tview.Form) {
	form.AddButton("Start", func() {
		if s.chosenInport == nil {
			showError(fmt.Errorf("choose the MIDI in port first"))
			return
		}
		if s.chosenOutport == nil {
			showError(fmt.Errorf("choose the MIDI out port first"))
			return
		}
		if s.chosenLine == "" {
			showError(fmt.Errorf("choose a line first"))
			return
		}
		v := tview.NewTextView()
		v.SetWrap(true)
		v.SetChangedFunc(func() {
			app.Draw()
		})
		v.SetBorder(true).SetTitle("connection log")
		pagesRight.AddAndSwitchToPage("form", v, true)
		fmt.Fprintf(v, "STARTING LINE %#v BETWEEN PORT %v AND %v\n", s.chosenLine, s.chosenInport, s.chosenOutport)
		//var err error
		/*
			s.proxy, err = newProxy(v, s.chosenLine, s.chosenInport, s.chosenOutport)
			if err != nil {
				showError(err)
				return
			}
		*/
		stop := make(chan bool)
		stopDone := make(chan bool)
		go func() {
			/*
				err := s.proxy.Start()
				if err != nil {
					showError(err)
				}
			*/
			<-stop
			/*
				s.proxy.Stop()
				s.proxy = nil
			*/
			stopDone <- true
			fmt.Fprintf(v, "CONNECTION CLOSED\n")
		}()

		form.RemoveButton(0)
		form.AddButton("Stop", func() {
			stop <- true
			form.RemoveButton(0)
			<-stopDone
			s.addRunButton(form)
		})
	})
}

func findInPort(name string) midi.In {
	/*
		for _, port := range ins {
			if port.String() == name {
				return port
			}
		}
	*/
	return nil
}

func findOutPort(name string) midi.Out {
	/*
		for _, port := range outs {
			if port.String() == name {
				return port
			}
		}
	*/
	return nil
}

func showMessage(msg midi.Message) string {
	switch v := msg.(type) {
	case channel.NoteOn:
		return fmt.Sprintf("[red]%s[white]/%v", smf.KeyToNote(v.Key()), v.Velocity())
	case channel.NoteOff:
		return fmt.Sprintf("[white]/%s", smf.KeyToNote(v.Key()))
	case channel.NoteOffVelocity:
		return fmt.Sprintf("[white]/%s", smf.KeyToNote(v.Key()))
	case channel.Aftertouch:
		return fmt.Sprintf("[green]AT%v", v.Pressure())
	case channel.ControlChange:
		name := cc.Name[v.Controller()]
		if name == "" {
			name = fmt.Sprintf("%v", v.Controller())
		} else {
			name = fmt.Sprintf("%v(%s)", v.Controller(), name)
		}
		return fmt.Sprintf("[green]CC%s[white]/%v", name, v.Value())
	case channel.Pitchbend:
		return fmt.Sprintf("[green]PB%v", v.Value())
	case channel.PolyAftertouch:
		return fmt.Sprintf("[green]PA%v[white]/%v", v.Key(), v.Pressure())
	case channel.ProgramChange:
		name := gm.Instr(v.Program()).String()
		if name == "" {
			name = fmt.Sprintf("%v", v.Program())
		} else {
			name = fmt.Sprintf("%v(%s)", v.Program(), name)
		}
		return fmt.Sprintf("[green]PC%s", name)
	case meta.Lyric:
		return fmt.Sprintf("[cyan]%q", v.Text())
	case meta.Text:
		return fmt.Sprintf("[cyan]'%s'", v.Text())
	default:
		return fmt.Sprintf("[white]%s", msg.String())
	}
}

func (s *runnerScreen) setTableHeader() {

	s.Table.SetCell(0, 0, tview.NewTableCell("Bar").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 1, tview.NewTableCell("Meter").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 2, tview.NewTableCell("Comment").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 3, tview.NewTableCell("Mark").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 4, tview.NewTableCell("Tempo").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 5, tview.NewTableCell("Beat").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))

	s.cols = 6

	for i, t := range s.song.Tracks {
		//if t.WithContent {
		//fmt.Fprintf(&bf, " %s[%v] | ", t.Name, t.Channel)
		var str = t.Name
		if t.Channel >= 0 {
			str += fmt.Sprintf(" [red][%v]", t.Channel)
		}
		s.Table.SetCell(0, 6+i, tview.NewTableCell(str).SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
		s.cols++
		//}
	}

	/*
		for l := 1; l < s.height; l++ {
			for cc := 0; cc < s.cols; cc++ {
				s.Table.SetCell(l, cc, tview.NewTableCell(" X ").SetTextColor(tcell.ColorGreen))
			}
		}
	*/
}

// run a configuration
//func (s *runnerScreen) runnerForm() *tview.Form {
func (s *runnerScreen) refreshNotes() {

	for l := 1; l < s.lines; l++ {
		for cc := 0; cc < s.cols; cc++ {
			s.Table.SetCell(l, cc, tview.NewTableCell(" "))
		}
	}

	/*
		outports := getOutPorts()
		inports := getInPorts()

		form := tview.NewForm()
		form.
			AddDropDown("MIDI in port", inports, -1, func(option string, optionIndex int) {
				s.chosenInport = findInPort(inports[optionIndex])

			}).
			AddDropDown("MIDI out port", outports, -1, func(option string, optionIndex int) {
				s.chosenOutport = findOutPort(outports[optionIndex])
			}).
			AddFormItem(s.linesDropDown)

		s.addRunButton(form)
		form.SetBorder(true).SetTitle("connection").SetTitleAlign(tview.AlignCenter)
		return form
	*/

	//table.SetBorders(true)
	//lorem := strings.Split("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", " ")
	//cols, rows := 10, 40
	//word := 0

	/*
		| Comment | Mark | Tempo  | Beat |  metronome[0] |  melody[0] |

	*/

	var line = 1

	for _, b := range s.song.Bars {
		_ = b
		bt := fmt.Sprintf(" %v", b.No+1)
		mt := fmt.Sprintf("%v/%v", b.TimeSig[0], b.TimeSig[1])

		s.Table.SetCell(line, 0, tview.NewTableCell(bt).SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
		s.Table.SetCell(line, 1, tview.NewTableCell(mt).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))

		for cc := 2; cc < s.cols; cc++ {
			s.Table.SetCell(line, cc, tview.NewTableCell("--").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorGrey))
		}

		line++

		for _, p := range b.Positions {
			tempo := " "
			if p.Tempo != 0 {
				tempo = fmt.Sprintf("%0.2f", tempo)
			}

			var frac float64

			if p.Fraction[1] > 0 {
				frac = p.Fraction[0] / p.Fraction[1]
			}

			beat := fmt.Sprintf("%0.4f", float64(p.Beat)+float64(1)+frac)

			s.Table.SetCell(line, 2, tview.NewTableCell(p.Comment).SetAlign(tview.AlignCenter))
			s.Table.SetCell(line, 3, tview.NewTableCell(p.Mark).SetAlign(tview.AlignCenter))
			s.Table.SetCell(line, 4, tview.NewTableCell(tempo).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft).SetAttributes(tcell.AttrBold))
			s.Table.SetCell(line, 5, tview.NewTableCell(beat).SetTextColor(tcell.ColorGrey).SetAlign(tview.AlignLeft).SetAttributes(tcell.AttrBold))

			//fmt.Fprintf(&bf, "| %s | %s | %s | %s | ", p.Comment, p.Mark, tempo, beat)

			for n, t := range s.song.Tracks {
				//if t.WithContent {
				//var printed bool
				for _, m := range p.Messages {
					if m.TrackNo == t.No {
						//fmt.Fprintf(&bf, " %s | ", showMessage(m.Message))
						//printed = true
						s.Table.SetCell(line, 6+n,
							tview.NewTableCell(showMessage(m.Message)).
								SetAlign(tview.AlignCenter).
								SetAttributes(tcell.AttrBold).
								SetReference(m),
						)
					}
				}
				/*
					if !printed {
						fmt.Fprintf(&bf, "  | ")
					}
				*/
				//}
			}

			line++

			/*
				if line == s.height-4 {
					break
				}
			*/
			//fmt.Fprintf(&bf, "\n")
		}
	}

	s.lines = line

	/*
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				color := tcell.ColorWhite
				if c < 1 || r < 1 {
					color = tcell.ColorYellow
				}
				table.SetCell(r, c,
					tview.NewTableCell(lorem[word]).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
				word = (word + 1) % len(lorem)
			}
		}
		table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		}).SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		})
	*/

	//return table
}
