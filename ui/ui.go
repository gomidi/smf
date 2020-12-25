package ui

import (
	"fmt"

	"github.com/gdamore/tcell/encoding"
	"github.com/gdamore/tcell/v2"

	//	"github.com/gomidi/connect/rtmidiadapter"
	"github.com/rivo/tview"
	"gitlab.com/gomidi/smf"
	//lineconfig "gitlab.com/gomidi/midiline/config"
)

type UI struct {
	file      string
	inModal   bool
	song      *smf.Song
	app       *tview.Application
	layout    *tview.Flex
	pages     *tview.Pages
	info      *tview.TextView
	runScreen *mainScreen
	clipBoard interface{}
}

func New(file string) *UI {
	return &UI{
		file: file,
	}
}

func (u *UI) Run() (err error) {
	encoding.Register()

	u.song, err = smf.ReadSMF(u.file)

	if err != nil {
		return err
	}

	u.app = tview.NewApplication()
	u.newRunnerScreen()

	u.pages = tview.NewPages()
	u.pages.AddAndSwitchToPage("runner", u.runScreen, true)
	u.info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(tview.AlignCenter)

	u.info.SetText("[red]F10 [white]help [yellow]| [red]F6 [white]connection [yellow]| [red]F7[white] conditions [yellow]| [red]F8[white] actions [yellow]| [red]F9[white] lines [yellow]| [red]CTRL+S [white]save [yellow]| [red]CTRL+Q [white]quit")

	// Create the main layout.
	u.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexColumn).
				AddItem(u.pages, 0, 1, true),
			/*.
			AddItem(pagesRight, 0, 2, true), */
			0, 2, true).
		AddItem(u.info, 1, 1, false)

	u.app.SetInputCapture(u.inputCapture)
	u.changeScreen(u.layout)
	u.app.SetFocus(u.pages)

	/*
		app.SetAfterDrawFunc(func(sc tcell.Screen) {
			_, runScreen.height = sc.Size()
		})
	*/

	return u.app.Run()
}

func (u *UI) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	//fmt.Printf("key %v %v", event.Key(), event.Rune())

	switch event.Key() {
	//case tcell.Keyctr

	// TODO find a way to track CTRL+left (page column wise), CTRL+right (page column wise),
	// CTRL+up and CTRL+down (change key for noteon/noteoff and polyaftertouch and value for all other channel messages)

	case tcell.KeyDEL:
		row, col := u.runScreen.Table.GetSelection()
		u.runScreen.deleteMessage(row, col)
	case tcell.KeyDelete:
		row, col := u.runScreen.Table.GetSelection()
		u.runScreen.deleteMessage(row, col)
	case tcell.KeyRune: // letters
		//fmt.Printf("rune: %v", event.Rune())
		switch event.Rune() {
		case ' ': // play/stop starting from currently selected cell (solo or not depending on solo state)
		case '.': // stop playing and move selection/cursor to stopped position within current column
		case '?': // show the one-letter shortcuts (or all?)
			//m := tview.NewModal()
			m := tview.NewTextView()
			m.SetBorder(true)
			m.SetTitleAlign(tview.AlignCenter)
			m.SetTitle("HELP")
			//m.SetText("HELP. Nothing to see here, work in progress")
			//m.SetText(fmt.Sprintf("height: %v lines", runScreen.height))
			help := `HELP
? show help screen

SPACEBAR  play/stop starting from currently selected cell
DELETE    remove message of selected cell
ENTER     edit message of selected cell or insert a new message


0 select header of current column
1 insert message on the next beat 1 in the current column
2 insert message on the next beat 2 in the current column
3 insert message on the next beat 3 in the current column
4 insert message on the next beat 4 in the current column
5 insert message on the next beat 5 in the current column
6 insert message on the next beat 6 in the current column
7 insert message on the next beat 7 in the current column
8 insert message on the next beat 8 in the current column
9 insert message on the next beat 9 in the current column

c copy message of the selected cell
x cut message the selected cell
v paste message into the selected cell (replacing the any message there)
a add message on same position in the current column
b add message from clipboard on same position in the current column
r replace the current message type

n new line within current bar
k kill the current beat line

t add a new track
o change trackorder, i.e. move track position
p edit properties of current track
s solo the current track

b b: new bar and beat line within that bar
g go to bar: enter bar number to navigate to
m move bar
d delete the current bar

l insert lyrics into current column
i insert a series of messages
`

			m.SetText(help)
			m.SetTextAlign(tview.AlignLeft)
			m.SetInputCapture(func(*tcell.EventKey) *tcell.EventKey {
				u.changeScreen(u.layout)
				u.app.SetFocus(u.pages)
				return nil
			})

			/*
				m.AddButtons([]string{"Ok"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Ok" {
							changeScreen(layout)
							app.SetFocus(pages)
						}
					})
			*/
			u.changeScreen(m)
			return nil
		case 't': // add a new track (to edit a track, go to the first line, select the track and press ENTER, to delete a track press then delete)
		case '0': // go to the table header with the cursor
			u.runScreen.Table.Select(0, u.runScreen.selectedCol)
		case '1': // insert line/event on the next beat 1 (this bar or the next bar) of the current track
		case '2': // insert line/event on the next beat 2 (this bar or the next bar) of the current track
		case '3': // insert line/event on the next beat 3 (this bar or the next bar) of the current track
		case '4': // insert line/event on the next beat 4 (this bar or the next bar) of the current track
		case '5': // insert line/event on the next beat 5 (this bar or the next bar) of the current track
		case '6': // insert line/event on the next beat 6 (this bar or the next bar) of the current track
		case '7': // insert line/event on the next beat 7 (this bar or the next bar) of the current track
		case '8': // insert line/event on the next beat 8 (this bar or the next bar) of the current track
		case '9': // insert line/event on the next beat 9 (this bar or the next bar) of the current track
		case 'g': // go to bar: enter bar number to navigate to
		case 'c': // copy
			row, col := u.runScreen.Table.GetSelection()
			u.runScreen.copyCell(row, col)
		case 'p': // p: properties; edit properties of current track
			row, col := u.runScreen.Table.GetSelection()
			u.runScreen.pasteCell(row, col)
		case 'v': // paste
			row, col := u.runScreen.Table.GetSelection()
			u.runScreen.pasteCell(row, col)
		case 'x': // cut
			row, col := u.runScreen.Table.GetSelection()
			u.runScreen.copyCell(row, col)
			u.runScreen.deleteMessage(row, col)
		case 'o': // change trackorder, i.e. move track position
		case 'i':
			// insert a series of events: type (linear/exponential), startpoint (current cell), startvalue, targetvalue, timestep, valuestep (when linear)
		case 's': // solo the current track
		case 'd': // delete the current bar
		case 'k': // kill the current beat line
		case 'l': // insert lyrics into current track (distribute it on the following notes, separated by empty space and dashes -)
		case 'n': // n: new line within current bar
		case 'b': // b: new bar and beat line within that bar
		case 'm': /// m: move bar (to move beat line simply edit the beat
			// move the line
		case 'r': // r: replace the current message type
			//fmt.Printf("new")
			if !u.inModal {
				row, col := u.runScreen.Table.GetSelection()
				u.runScreen.insertMessageFunc(row, col)
				return nil
			} else {
				return event
			}
		default:
			u.runScreen.Table.InputHandler()(event, nil)
		}
	case tcell.KeyEnter:
		if !u.inModal {
			row, col := u.runScreen.Table.GetSelection()
			u.runScreen.selectedFunc(row, col)
			return nil
		} else {
			return event
		}
	case tcell.KeyDown:
		u.runScreen.selectNextLine()
	case tcell.KeyUp:
		u.runScreen.selectPrevLine()
	case tcell.KeyLeft:
		u.runScreen.selectPrevCol()
	case tcell.KeyRight:
		u.runScreen.selectNextCol()
	case tcell.KeyF10:
		// print help
		m := tview.NewModal()
		//m.SetText("HELP. Nothing to see here, work in progress")
		//m.SetText(fmt.Sprintf("height: %v lines", u.runScreen.height))
		m.AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Ok" {
					u.changeScreen(u.layout)
					u.app.SetFocus(u.pages)
				}
			})
		u.changeScreen(m)
	case tcell.KeyF7:
		//pagesRight.RemovePage("form")
		//			pages.RemovePage("items")
		//			changeScreen(newMatcherScreen())
		//			ms := newMatcherScreen()
		//			pages.AddAndSwitchToPage("items", ms, true)
		//			app.SetFocus(ms)
		//				app.SetFocus(pagesRight)
		//				app.SetFocus(pages)
		u.pages.SwitchToPage("Condition")
		u.app.SetFocus(u.pages)

	case tcell.KeyCtrlS:
		//err := saveConfig()
		m := tview.NewModal()
		/*
			if err == nil {
				m.SetText("File saved.")

			} else {
				//m.SetText(fmt.Sprintf("ERROR: could not save file %#v:%v", CONFIG_FILE, err))
			}
		*/
		m.AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Ok" {
					u.changeScreen(u.layout)
					u.app.SetFocus(u.pages)
				}
			})
		u.changeScreen(m)

	//case tcell.KeyCtrlT:
	case tcell.KeyF8:
		//pagesRight.RemovePage("form")
		u.pages.SwitchToPage("action")
		u.app.SetFocus(u.pages)
	case tcell.KeyCtrlQ:
		/*
			if runScreen.proxy != nil {
				runScreen.proxy.Stop()
				runScreen.proxy = nil
			}
		*/
		u.app.Stop()
	case tcell.KeyCtrlC:
		/*
			if runScreen.proxy != nil {
				runScreen.proxy.Stop()
				runScreen.proxy = nil
			}
		*/
		u.app.Stop()
	case tcell.KeyF9:
		//pagesRight.RemovePage("form")
		u.pages.SwitchToPage("line")
		u.app.SetFocus(u.pages)
	case tcell.KeyF6:
		//pagesRight.RemovePage("form")
		u.pages.SwitchToPage("runner")
		u.app.SetFocus(u.pages)
	//case tcell.KeyEscape, tcell.KeyTab:
	//	return event
	default:
		//			panic(fmt.Sprintf("key: %#v", event.Key()))
	}
	return event
	//return nil
}

func (u *UI) backToTableScreen() {
	u.changeScreen(u.layout)
	u.app.SetFocus(u.pages)
	u.inModal = false
}

func (u *UI) changeScreen(p tview.Primitive) {
	u.app.SetRoot(p, true).SetFocus(p)
}

func (u *UI) showError(err error) {
	if err == nil {
		return
	}
	m := tview.NewModal()
	m.SetText(fmt.Sprintf("ERROR: %#v", err.Error()))
	m.AddButtons([]string{"Ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Ok" {
				u.changeScreen(u.layout)
				u.app.SetFocus(u.pages)
			}
		})
	u.changeScreen(m)
}
