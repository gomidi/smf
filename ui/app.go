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

type App struct {
	file         string
	inModal      bool
	song         *smf.Song
	app          *tview.Application
	layout       *tview.Flex
	pages        *tview.Pages
	info         *tview.TextView
	mainScreen   *mainScreen
	clipBoard    interface{}
	activeScreen activeScreen
}

type activeScreen interface {
	InputHandler() func(event *tcell.EventKey, setFocus func(tview.Primitive))
}

const (
	infoText = "[red]F10 [white]help [yellow]| [red]F6 [white]connection [yellow]| [red]F7[white] conditions [yellow]| [red]F8[white] actions [yellow]| [red]F9[white] lines [yellow]| [red]CTRL+S [white]save [yellow]| [red]CTRL+Q [white]quit"
	help     = `HELP
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
)

func New(file string) *App {
	return &App{
		file: file,
	}
}

func (u *App) Run() (err error) {
	encoding.Register()

	u.song, err = smf.ReadSMF(u.file)

	if err != nil {
		return err
	}

	u.app = tview.NewApplication()
	u.newRunnerScreen()

	u.pages = tview.NewPages()
	u.pages.AddAndSwitchToPage("main", u.mainScreen, true)
	u.info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(tview.AlignCenter)

	u.info.SetText(infoText)

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
	u.backToTableScreen()

	/*
		app.SetAfterDrawFunc(func(sc tcell.Screen) {
			_, runScreen.height = sc.Size()
		})
	*/

	return u.app.Run()
}

func (u *App) showHelp() {
	m := tview.NewTextView()
	//m.InputHandle
	m.SetBorder(true)
	m.SetTitleAlign(tview.AlignCenter)
	m.SetTitle("HELP")
	m.SetText(help)
	m.SetTextAlign(tview.AlignLeft)
	m.SetInputCapture(func(*tcell.EventKey) *tcell.EventKey {
		u.changeScreen(u.layout)
		u.app.SetFocus(u.pages)
		return nil
	})
	u.changeScreen(m)
}

func (u *App) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	//fmt.Printf("key %v %v", event.Key(), event.Rune())

	switch event.Key() {
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
		return nil
	case tcell.KeyCtrlQ:
		u.app.Stop()
		return nil
	case tcell.KeyCtrlC:
		u.app.Stop()
		return nil
	case tcell.KeyEscape:
		u.backToTableScreen()
		return nil
	}

	if u.activeScreen != nil {
		fn := u.activeScreen.InputHandler()
		fn(event, func(p tview.Primitive) {
			u.app.SetFocus(p)
		})
		return nil
	}

	switch event.Key() {
	//case tcell.Keyctr

	// TODO find a way to track CTRL+left (page column wise), CTRL+right (page column wise),
	// CTRL+up and CTRL+down (change key for noteon/noteoff and polyaftertouch and value for all other channel messages)

	case tcell.KeyDEL:
		u.mainScreen.deleteMessage()
	case tcell.KeyDelete:
		u.mainScreen.deleteMessage()
	case tcell.KeyRune: // letters
		//fmt.Printf("rune: %v", event.Rune())
		switch event.Rune() {
		case ' ': // play/stop starting from currently selected cell (solo or not depending on solo state)
		case '.': // stop playing and move selection/cursor to stopped position within current column
		case '?': // show the one-letter shortcuts (or all?)
			u.showHelp()
			return nil
		case 't': // add a new track (to edit a track, go to the first line, select the track and press ENTER, to delete a track press then delete)
		case '0': // go to the table header with the cursor
			u.mainScreen.Table.Select(0, u.mainScreen.selectedCol)
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
			u.mainScreen.copyCell()
		case 'p': // p: properties; edit properties of current track
			u.mainScreen.pasteCell()
		case 'v': // paste
			u.mainScreen.pasteCell()
		case 'x': // cut
			u.mainScreen.copyCell()
			u.mainScreen.deleteMessage()
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
				u.mainScreen.replaceMessage()
				return nil
			} else {
				return event
			}
		default:
			u.mainScreen.Table.InputHandler()(event, nil)
		}
	case tcell.KeyEnter:
		if !u.inModal {
			u.mainScreen.enterCell(func(p tview.Primitive) {
				u.app.SetFocus(p)
			})
			return nil
		} else {
			return event
		}
	case tcell.KeyDown:
		u.mainScreen.selectNextLine()
	case tcell.KeyUp:
		u.mainScreen.selectPrevLine()
	case tcell.KeyLeft:
		u.mainScreen.selectPrevCol()
	case tcell.KeyRight:
		u.mainScreen.selectNextCol()
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

	//case tcell.KeyCtrlT:
	case tcell.KeyF8:
		//pagesRight.RemovePage("form")
		u.pages.SwitchToPage("action")
		u.app.SetFocus(u.pages)
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
	return nil
	//return nil
}

func (u *App) backToTableScreen() {
	u.changeScreen(u.layout)
	u.app.SetFocus(u.pages)
	u.inModal = false
	u.activeScreen = nil
}

func (u *App) changeScreen(p tview.Primitive) {
	u.app.SetRoot(p, true).SetFocus(p)
	u.activeScreen = p
}

func (u *App) showError(err error) {
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
