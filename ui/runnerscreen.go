package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/smf"
)

/*
// Demo code for the Table primitive.
package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	lorem := strings.Split("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", " ")
	cols, rows := 10, 40
	word := 0
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
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

*/

var inModal bool

type runnerScreen struct {
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
	song          *smf.Song
	height        int
	currentBar    int
	topPosition   int
	selectedLine  int
	selectedCol   int
	cols          int
	colLeftOffset int
	lines         int
	//proxy         *midiproxy.Proxy
}

func (sc *runnerScreen) insertMessageFunc(row int, column int) {
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
			ch := sc.song.Tracks[tm.TrackNo].Channel
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
						cell.SetText(showMessage(msg))
						//sc.refresh()

						changeScreen(layout)
						app.SetFocus(pages)
						inModal = false
						runScreen.Select(row, column)
					}, func() {
						changeScreen(layout)
						app.SetFocus(pages)
						inModal = false
						runScreen.Select(row, column)
					})
					fm.Blur()
					changeScreen(_fm)
				})

				fm.AddButton("cancel", func() {
					changeScreen(layout)
					app.SetFocus(pages)
					inModal = false
					runScreen.Select(row, column)
				})

				inModal = true
				changeScreen(fm)
			}

		}
	}

}

func (sc *runnerScreen) deleteMessage(row int, column int) {
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
	runScreen.Select(row, column)
}

func (sc *runnerScreen) selectedFunc(row int, column int) {
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
			inModal = true

			if ms.Message == nil {

				ch := sc.song.Tracks[ms.TrackNo].Channel
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
							cell.SetText(showMessage(msg))
							//sc.refresh()

							changeScreen(layout)
							app.SetFocus(pages)
							inModal = false
							runScreen.Select(row, column)
						}, func() {
							changeScreen(layout)
							app.SetFocus(pages)
							inModal = false
							runScreen.Select(row, column)
						})
						fm.Blur()
						changeScreen(_fm)
					})

					fm.AddButton("cancel", func() {
						changeScreen(layout)
						app.SetFocus(pages)
						inModal = false
						runScreen.Select(row, column)
					})

					inModal = true
					changeScreen(fm)
				}
			} else {
				fm := NewEditForm(ms.Message, func(msg midi.Message) {
					ms.Message = msg
					cell.SetReference(ms)
					cell.SetText(showMessage(msg))
					//sc.refresh()

					changeScreen(layout)
					app.SetFocus(pages)
					inModal = false
					runScreen.Select(row, column)
				}, func() {
					changeScreen(layout)
					app.SetFocus(pages)
					inModal = false
					runScreen.Select(row, column)
				})
				changeScreen(fm)
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

func (sc *runnerScreen) Focus(delegate func(p tview.Primitive)) {
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

func (sc *runnerScreen) refresh() {
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

func (sc *runnerScreen) selectNextCol() {
	if sc.selectedCol > sc.cols-1 {
	} else {
		sc.selectedCol++
		sc.selectLine()
	}
}

func (sc *runnerScreen) selectPrevCol() {
	if sc.selectedCol <= 0 {
		// do nothing
	} else {
		sc.selectedCol--
		sc.selectLine()
	}
}

func (sc *runnerScreen) selectNextLine() {
	//if sc.selectedLine == sc.height-2 {
	// TODO: scroll one line
	//} else {
	sc.selectedLine++
	sc.selectLine()
	//}
}

func (sc *runnerScreen) selectPrevLine() {
	if sc.selectedLine <= 0 {
		// do nothing
	} else {
		sc.selectedLine--
		sc.selectLine()
	}
}

func (sc *runnerScreen) selectLine() {
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

func newRunnerScreen(s *smf.Song) *runnerScreen {
	sc := &runnerScreen{}
	sc.song = s
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
