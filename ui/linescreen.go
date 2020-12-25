package ui

//	"github.com/gdamore/tcell/v2"
//	"github.com/rivo/tview"

/*
type lineScreen struct {
	*tview.Box
	currentLine int
	lines       []string
}

func (m *lineScreen) Focus(delegate func(p tview.Primitive)) {
	m.lines = getLines()
	sort.Strings(m.lines)
	m.showForm()
}

func (m *lineScreen) showForm() {
	form := lineForm(m.currentLineName())
	//changeScreen(form)
	pagesRight.RemovePage("form")
	pagesRight.AddAndSwitchToPage("form", form, true)
}

func (m *lineScreen) currentLineName() string {
	mt := ""
	if m.currentLine >= 0 {
		mt = m.lines[m.currentLine]
	}
	return mt
}

func newLineScreen() *lineScreen {
	return &lineScreen{Box: tview.NewBox().SetBorder(true).SetTitle("lines"), lines: getLines(), currentLine: -1}
}

func (m *lineScreen) Draw(screen tcell.Screen) {
	m.Box.Draw(screen)
	x, y, width, height := m.GetInnerRect()

	radioButton := "( )" // Unchecked.
	if m.currentLine == -1 {
		radioButton = "[red](X)" // Checked.
	}
	line := fmt.Sprintf(`%s %s`, radioButton, "NEW")
	tview.Print(screen, line, x+1, y, width, tview.AlignLeft, tcell.ColorYellow)

	for index, option := range m.lines {
		if index >= height-1 {
			break
		}
		radioButton := "( )" // Unchecked.
		if index == m.currentLine {
			radioButton = "[red](X)" // Checked.
		}
		line := fmt.Sprintf(`%s %s`, radioButton, option)
		tview.Print(screen, line, x+2, y+index+1, width, tview.AlignLeft, tcell.ColorYellow)
	}

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

}

// InputHandler returns the handler for this primitive.
func (m *lineScreen) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return m.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyDelete:
			name := m.currentLineName()
			if name != "" {
				//removeLine(name)
				m.lines = getLines()
				sort.Strings(m.lines)
				m.currentLine = -1
				m.showForm()
				//app.Refresh()
			}
		case tcell.KeyUp:
			m.currentLine--
			if m.currentLine < -1 {
				m.currentLine = -1
			}
			m.showForm()
		case tcell.KeyDown:
			m.currentLine++
			if m.currentLine >= len(m.lines) {
				m.currentLine = len(m.lines) - 1
			}
			m.showForm()

				case tcell.KeyTab:
					mt := ""
					if m.currentStack >= 0 {
						mt = m.stacks[m.currentStack]
					}
					_ = mt
					//			changeScreen(stackForm(mt))

		}
	})
}
*/
