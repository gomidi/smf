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

func (s *mainScreen) showMessage(msg midi.Message) string {
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

func (s *mainScreen) setTableHeader() {

	s.Table.SetCell(0, 0, tview.NewTableCell("Bar").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 1, tview.NewTableCell("Meter").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 2, tview.NewTableCell("Comment").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 3, tview.NewTableCell("Mark").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 4, tview.NewTableCell("Tempo").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))
	s.Table.SetCell(0, 5, tview.NewTableCell("Beat").SetBackgroundColor(tcell.ColorBlue).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter).SetAttributes(tcell.AttrBold))

	s.cols = 6

	for i, t := range s.ui.song.Tracks {
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
func (s *mainScreen) refreshNotes() {

	/*
		for l := 1; l < s.lines; l++ {
			for cc := 0; cc < s.cols; cc++ {
				if cc > 5 {
					var tm smf.TrackMessage
					tm.TrackNo = uint16(cc - 6)
					s.Table.SetCell(l, cc, tview.NewTableCell("").SetReference(tm))
				} else {
					s.Table.SetCell(l, cc, tview.NewTableCell(" "))
				}
			}
		}
	*/

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

	for _, b := range s.ui.song.Bars {
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

			for cc := 6; cc < s.cols; cc++ {
				var tm smf.TrackMessage
				tm.TrackNo = uint16(cc - 6)
				tm.Position = p
				tm.AbsPos = p.AbsTicks()
				s.Table.SetCell(line, cc, tview.NewTableCell("").SetReference(&tm).
					SetAlign(tview.AlignCenter).
					SetAttributes(tcell.AttrBold))
			}

			for n, t := range s.ui.song.Tracks {
				//if t.WithContent {
				//var printed bool
				for _, m := range p.Messages {
					if m.TrackNo == t.No {
						//fmt.Fprintf(&bf, " %s | ", showMessage(m.Message))
						//printed = true
						s.Table.SetCell(line, 6+n,
							tview.NewTableCell(s.showMessage(m.Message)).
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
}

type mainScreen struct {
	//*tview.Box
	//	currentTransformer int
	//	transformers       []string
	//*tview.Form
	*tview.Table
	//lines         []string
	linesDropDown *tview.DropDown
	chosenLine    string
	chosenInport  midi.In
	chosenOutport midi.Out
	// height        int
	currentBar    int
	topPosition   int
	selectedLine  int
	selectedCol   int
	cols          int
	colLeftOffset int
	lines         int
	ui            *UI
	//proxy         *midiproxy.Proxy
}

func (sc *mainScreen) insertMessageFunc(row int, column int) {
	/*
		for c := 0; c < sc.cols; c++ {
			sc.Table.GetCell(sc.selectedLine, c).SetBackgroundColor(tcell.ColorBlack)
		}
		for c := 0; c < sc.cols; c++ {
			sc.Table.GetCell(row, c).SetBackgroundColor(tcell.ColorRed)
		}
	*/

	//println("selectedFuncCalled")
	//fmt.Print("getting ref")
	cell := sc.Table.GetCell(row, column)
	ref := cell.GetReference()
	//panic(ref)

	if ref != nil {
		//fmt.Print("have ref")
		tm, ok := ref.(*smf.TrackMessage)

		if ok {

			//fmt.Print("have trackmessage")

			//tm.TrackNo
			ch := sc.ui.song.Tracks[tm.TrackNo].Channel
			if ch >= 0 {
				var _msg midi.Message
				var pmsg = &_msg

				fm := tview.NewForm()
				fm.AddDropDown("new message", _messagesSelect, 0, func(opt string, idx int) {
					*pmsg = newMessageInTrack(uint8(ch), opt)
				})
				fm.AddButton("ok", func() {
					_fm := NewEditForm(*pmsg, func(msg midi.Message) {
						tm.Message = msg
						cell.SetReference(tm)
						cell.SetText(sc.showMessage(msg))
						//sc.refresh()

						sc.ui.backToTableScreen()
						sc.Select(row, column)
					}, func() {
						sc.ui.backToTableScreen()
						sc.Select(row, column)
					})
					fm.Blur()
					sc.ui.changeScreen(_fm)
				})

				fm.AddButton("cancel", func() {
					sc.ui.backToTableScreen()
					sc.Select(row, column)
				})

				sc.ui.inModal = true
				sc.ui.changeScreen(fm)
			}

		}
	}

}

func (sc *mainScreen) deleteMessage(row int, column int) {
	cell := sc.Table.GetCell(row, column)
	ref := cell.GetReference()
	//panic(ref)

	if ref != nil {
		ms, ok := ref.(*smf.TrackMessage)
		if ok && ms != nil {
		}

		ms.Message = nil
		cell.SetReference(ms)
		cell.SetText("")
	}
	sc.Select(row, column)
}

func (sc *mainScreen) copyCell(row int, column int) {
	cell := sc.Table.GetCell(row, column)
	ref := cell.GetReference()
	//panic(ref)

	if ref != nil {
		ms, ok := ref.(*smf.TrackMessage)
		if ok && ms != nil && ms.Message != nil {
			sc.ui.clipBoard = ms.Message
		}
	}

	sc.Select(row, column)
}

func (sc *mainScreen) pasteCell(row int, column int) {
	cell := sc.Table.GetCell(row, column)
	ref := cell.GetReference()
	//panic(ref)

	if ref != nil {
		ms, ok := ref.(*smf.TrackMessage)
		if ok && ms != nil && sc.ui.clipBoard != nil {
			msg, isMsg := sc.ui.clipBoard.(midi.Message)

			if isMsg {

				_, isChMsg := msg.(channel.Message)

				if isChMsg {
					ch := sc.ui.song.Tracks[ms.TrackNo].Channel
					if ch >= 0 {
						c := channel.Channel(ch)
						switch v := msg.(type) {
						case channel.Aftertouch:
							msg = c.Aftertouch(v.Pressure())
						case channel.NoteOn:
							msg = c.NoteOn(v.Key(), v.Velocity())
						case channel.NoteOff:
							msg = c.NoteOff(v.Key())
						case channel.Pitchbend:
							msg = c.Pitchbend(v.Value())
						case channel.PolyAftertouch:
							msg = c.PolyAftertouch(v.Key(), v.Pressure())
						case channel.ControlChange:
							msg = c.ControlChange(v.Controller(), v.Value())
						case channel.ProgramChange:
							msg = c.ProgramChange(v.Program())
						default:
							panic("unsupported channel message")
						}
					}
				}
				//ms.Message = msg
			}

			ms.Message = msg

			cell.SetReference(ms)
			cell.SetText(sc.showMessage(msg))
		}
	}

	sc.Select(row, column)
}

func (sc *mainScreen) selectedFunc(row int, column int) {
	/*
		for c := 0; c < sc.cols; c++ {
			sc.Table.GetCell(sc.selectedLine, c).SetBackgroundColor(tcell.ColorBlack)
		}
		for c := 0; c < sc.cols; c++ {
			sc.Table.GetCell(row, c).SetBackgroundColor(tcell.ColorRed)
		}
	*/

	//println("selectedFuncCalled")

	cell := sc.Table.GetCell(row, column)
	ref := cell.GetReference()
	//panic(ref)

	if ref != nil {
		ms, ok := ref.(*smf.TrackMessage)
		if ok && ms != nil {
			//panic(ms.Message.String())
			sc.ui.inModal = true

			if ms.Message == nil {

				ch := sc.ui.song.Tracks[ms.TrackNo].Channel
				if ch >= 0 {
					var msg midi.Message
					var pmsg = &msg

					fm := tview.NewForm()
					fm.AddDropDown("new message", _messagesSelect, 0, func(opt string, idx int) {
						*pmsg = newMessageInTrack(uint8(ch), opt)
					})
					fm.AddButton("ok", func() {
						//fmt.Print("[red]" + (*pmsg).String())
						//tm.Message = *pmsg
						//changeScreen(layout)
						//app.SetFocus(pages)
						//inModal = false
						//cell.SetReference(tm)
						//cell.SetText(showMessage(*pmsg))

						_fm := NewEditForm(*pmsg, func(msg midi.Message) {
							ms.Message = msg
							cell.SetReference(ms)
							cell.SetText(sc.showMessage(msg))
							//sc.refresh()

							sc.ui.backToTableScreen()
							sc.Select(row, column)
						}, func() {
							sc.ui.backToTableScreen()
							sc.Select(row, column)
						})
						fm.Blur()
						sc.ui.changeScreen(_fm)
					})

					fm.AddButton("cancel", func() {
						sc.ui.backToTableScreen()
						sc.Select(row, column)
					})

					sc.ui.inModal = true
					sc.ui.changeScreen(fm)
				}
			} else {
				fm := NewEditForm(ms.Message, func(msg midi.Message) {
					ms.Message = msg
					cell.SetReference(ms)
					cell.SetText(sc.showMessage(msg))
					//sc.refresh()

					sc.ui.backToTableScreen()
					sc.Select(row, column)
				}, func() {
					sc.ui.backToTableScreen()
					sc.Select(row, column)
				})
				sc.ui.changeScreen(fm)
			}

			/*
				m := tview.NewModal()
				m.SetTitle(fmt.Sprintf("track: %v position: %v", ms.TrackNo, ms.AbsPos))
				//m.SetText("HELP. Nothing to see here, work in progress")
				m.SetTitleColor(tcell.ColorWhite)
				m.SetText(ms.Message.String())
				m.SetBorder(true)
				m.AddButtons([]string{"Ok"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Ok" {
							changeScreen(layout)
							app.SetFocus(pages)
							inModal = false
						}
					})
				changeScreen(m)
			*/

		}
	}

	sc.selectedLine = row
	//table.SetSelectable(false, false)
	//sc.Table.SetSelectable(false, false)
}

func (sc *mainScreen) Focus(delegate func(p tview.Primitive)) {
	//	sc.refresh()
	//	sc.Form.Focus(delegate)
}

/*
func (sc *runnerScreen) lineIndex() int {
	for i, st := range sc.lines {
		if sc.chosenLine == st {
			return i
		}
	}
	return -1
}
*/

func (sc *mainScreen) refresh() {
	//_, _, _, sc.height = sc.GetRect()
	sc.refreshNotes()
	/*
		sc.lines = getLines()
		sort.Strings(sc.lines)
		sc.linesDropDown.SetOptions(sc.lines, func(option string, optionIndex int) {
			sc.chosenLine = option
		})
		sc.linesDropDown.SetCurrentOption(sc.lineIndex())
	*/
}

func (sc *mainScreen) selectNextCol() {
	if sc.selectedCol > sc.cols-1 {
	} else {
		sc.selectedCol++
		sc.selectLine()
	}
}

func (sc *mainScreen) selectPrevCol() {
	if sc.selectedCol <= 0 {
		// do nothing
	} else {
		sc.selectedCol--
		sc.selectLine()
	}
}

func (sc *mainScreen) selectNextLine() {
	//if sc.selectedLine == sc.height-2 {
	// TODO: scroll one line
	//} else {
	sc.selectedLine++
	sc.selectLine()
	//}
}

func (sc *mainScreen) selectPrevLine() {
	if sc.selectedLine <= 0 {
		// do nothing
	} else {
		sc.selectedLine--
		sc.selectLine()
	}
}

func (sc *mainScreen) selectLine() {
	/*
		for c := 0; c < sc.cols; c++ {
			if c == 5 {
				sc.Table.GetCell(sc.selectedLine, c).SetTextColor(tcell.ColorGrey).SetBackgroundColor(tcell.ColorBlack)
				//	txt := cl.Text
				//	cl.SetTextColor(tcell.ColorGrey).SetBackgroundColor(tcell.ColorBlack)
				//	cl.SetText(txt)
			} else {
				sc.Table.GetCell(sc.selectedLine, c).SetBackgroundColor(tcell.ColorBlack)
			}
		}
		for c := 0; c < sc.cols; c++ {
			switch c {
			case 5:
				sc.Table.GetCell(no, c).SetBackgroundColor(tcell.ColorGrey).SetTextColor(tcell.ColorWhite)
			case sc.selectedCol:
				sc.Table.GetCell(no, c).SetBackgroundColor(tcell.ColorGrey)
			default:
			}
		}
		sc.selectedLine = no
	*/

	sc.Table.Select(sc.selectedLine, sc.selectedCol)
	/*
		.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				sc.Table.SetSelectable(true, true)
			}
		})
	*/
}

func (u *UI) newRunnerScreen() *mainScreen {
	sc := &mainScreen{}
	u.runScreen = sc
	sc.ui = u
	sc.chosenInport = nil
	sc.chosenOutport = nil
	sc.Table = tview.NewTable()
	sc.Table.SetSelectable(true, true)
	sc.Table.SetFixed(1, 6)
	sc.Table.SetBorders(false)
	sc.Table.SetBordersColor(tcell.ColorGrey)
	sc.Table.SetSeparator('|')
	//sc.Table.SetEvaluateAllRows(true)
	sc.Table.WrapInputHandler(func(key *tcell.EventKey, cb func(tview.Primitive)) {
		panic(key.Name())
	})
	//sc.Table.
	var sty tcell.Style

	sc.Table.SetSelectedStyle(sty.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
	//sc.selectLine()
	sc.setTableHeader()
	//sc.selectedCol = 6
	//var selectedRow = 0

	//sc.Table.SetSelectedFunc()

	//sc.linesDropDown = tview.NewDropDown()
	//sc.linesDropDown.SetLabel("Stack")
	sc.refresh()
	//sc.Form = sc.runnerForm()
	return sc
}

//func (m *runnerScreen) Draw(screen tcell.Screen) {

//	m.Box.Draw(screen)
/*	x, y, width, height := m.GetInnerRect() */
/*
	radioButton := "\u25ef" // Unchecked.
	if m.currentTransformer == -1 {
		radioButton = "\u25c9" // Checked.
	}
	line := fmt.Sprintf(`%s[white]  %s`, radioButton, "*new")
	tview.Print(screen, line, x+2, y, width, tview.AlignLeft, tcell.ColorYellow)

	for index, option := range m.transformers {
		if index >= height-1 {
			break
		}
		radioButton := "\u25ef" // Unchecked.
		if index == m.currentTransformer {
			radioButton = "\u25c9" // Checked.
		}
		line := fmt.Sprintf(`%s[white]  %s`, radioButton, option)
		tview.Print(screen, line, x+2, y+index+1, width, tview.AlignLeft, tcell.ColorYellow)
	}
	/*
		for index, option := range r.options {
			if index >= height {
				break
			}
			radioButton := "\u25ef" // Unchecked.
			if index == r.currentOption {
				radioButton = "\u25c9" // Checked.
			}
			line := fmt.Sprintf(`%s[white]  %s`, radioButton, option)
			tview.Print(screen, line, x, y+index, width, tview.AlignLeft, tcell.ColorYellow)
		}
*/

//}

// InputHandler returns the handler for this primitive.
//func (m *runnerScreen) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
//return m.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
/*
	return func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyUp:
			m.currentTransformer--
			if m.currentTransformer < -1 {
				m.currentTransformer = -1
			}
		case tcell.KeyDown:
			m.currentTransformer++
			if m.currentTransformer >= len(m.transformers) {
				m.currentTransformer = len(m.transformers) - 1
			}
		case tcell.KeyTab:
			mt := ""
			if m.currentTransformer >= 0 {
				mt = m.transformers[m.currentTransformer]
			}
			form := transformerForm(mt)
			//changeScreen(form)
			pagesRight.RemovePage("form")
			pagesRight.AddAndSwitchToPage("form", form, true)
			setFocus(form)
		default:

		}
	}
*/
//})
//}
