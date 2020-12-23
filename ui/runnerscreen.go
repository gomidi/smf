package ui

import (
	"sort"

	"github.com/rivo/tview"
	"gitlab.com/gomidi/midi"
)

type runnerScreen struct {
	//*tview.Box
	//	currentTransformer int
	//	transformers       []string
	*tview.Form
	lines         []string
	linesDropDown *tview.DropDown
	chosenLine    string
	chosenInport  midi.In
	chosenOutport midi.Out
	//proxy         *midiproxy.Proxy
}

func (sc *runnerScreen) Focus(delegate func(p tview.Primitive)) {
	sc.refresh()
	sc.Form.Focus(delegate)
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

func newRunnerScreen() *runnerScreen {
	sc := &runnerScreen{}
	sc.chosenInport = nil
	sc.chosenOutport = nil
	sc.linesDropDown = tview.NewDropDown()
	sc.linesDropDown.SetLabel("Stack")
	sc.refresh()
	sc.Form = sc.runnerForm()
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
