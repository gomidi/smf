package ui

import (
	"sort"

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

type runnerScreen struct {
	//*tview.Box
	//	currentTransformer int
	//	transformers       []string
	//*tview.Form
	*tview.Table
	lines         []string
	linesDropDown *tview.DropDown
	chosenLine    string
	chosenInport  midi.In
	chosenOutport midi.Out
	song          *smf.Song
	//proxy         *midiproxy.Proxy
}

func (sc *runnerScreen) Focus(delegate func(p tview.Primitive)) {
	//	sc.refresh()
	//	sc.Form.Focus(delegate)
}

func (sc *runnerScreen) lineIndex() int {
	for i, st := range sc.lines {
		if sc.chosenLine == st {
			return i
		}
	}
	return -1
}

func (sc *runnerScreen) refresh() {
	sc.lines = getLines()
	sort.Strings(sc.lines)
	sc.linesDropDown.SetOptions(sc.lines, func(option string, optionIndex int) {
		sc.chosenLine = option
	})
	sc.linesDropDown.SetCurrentOption(sc.lineIndex())
}

func newRunnerScreen(s *smf.Song) *runnerScreen {
	sc := &runnerScreen{}
	sc.song = s
	sc.chosenInport = nil
	sc.chosenOutport = nil
	//sc.linesDropDown = tview.NewDropDown()
	//sc.linesDropDown.SetLabel("Stack")
	//sc.refresh()
	//sc.Form = sc.runnerForm()
	sc.Table = sc.runnerForm()
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
