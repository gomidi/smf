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

func getOutPorts() (ports []string) {
	/*
		for _, port := range outs {
			ports = append(ports, port.String())
		}
	*/
	return
}

func getInPorts() (ports []string) {
	/*
		for _, port := range ins {
			ports = append(ports, port.String())
		}
	*/
	return
}

func getLines() (res []string) {
	/*
		for _, st := range data.Lines {
			res = append(res, st.Name)
		}
	*/
	return res
}

func getActions() (res []string) {
	/*
		for _, st := range data.Actions {
			res = append(res, st.Name)
		}
	*/
	return res
}

func getConditions() (res []string) {
	/*
		for _, st := range data.Conditions {
			res = append(res, st.Name)
		}
	*/
	return res
}

func getConditionMakers() (res [][2]string) {
	/*
		names := lineconfig.RegisteredConditionMaker()
		infos := lineconfig.RegisteredConditionMakerInfos()
		sort.Strings(names)

		for _, name := range names {
			res = append(res, [2]string{name, infos[name]})
		}
	*/
	return
}

func getActionMakers() (res [][2]string) {
	/*
		names := lineconfig.RegisteredActionMaker()
		infos := lineconfig.RegisteredActionMakerInfos()
		sort.Strings(names)

		for _, name := range names {
			res = append(res, [2]string{name, infos[name]})
		}
	*/
	return
}

/*
func addLine(m lineconfig.Line) {
	data.Lines = append(data.Lines, m)
}

func addAction(m lineconfig.Action) {
	data.Actions = append(data.Actions, m)
}

func addCondition(m lineconfig.Condition) {
	data.Conditions = append(data.Conditions, m)
}

func replaceLine(oldname string, n lineconfig.Line) {
	var nm = []lineconfig.Line{n}

	for _, m := range data.Lines {
		if m.Name != oldname {
			nm = append(nm, m)
		}
	}

	data.Lines = nm
}

func replaceAction(oldname string, n lineconfig.Action) {
	var nm = []lineconfig.Action{n}

	for _, m := range data.Actions {
		if m.Name != oldname {
			nm = append(nm, m)
		}
	}

	data.Actions = nm
}

func removeCondition(name string) {
	var nm = []lineconfig.Condition{}

	for _, m := range data.Conditions {
		if m.Name != name {
			nm = append(nm, m)
		}
	}

	data.Conditions = nm
}

func removeLine(name string) {
	var nm = []lineconfig.Line{}

	for _, m := range data.Lines {
		if m.Name != name {
			nm = append(nm, m)
		}
	}

	data.Lines = nm
}

func removeAction(name string) {
	var nm = []lineconfig.Action{}

	for _, m := range data.Actions {
		if m.Name != name {
			nm = append(nm, m)
		}
	}

	data.Actions = nm
}

func replaceCondition(oldname string, n lineconfig.Condition) {
	var nm = []lineconfig.Condition{n}

	for _, m := range data.Conditions {
		if m.Name != oldname {
			nm = append(nm, m)
		}
	}

	data.Conditions = nm
}
*/

func changeScreen(p tview.Primitive) {
	app.SetRoot(p, true).SetFocus(p)
}

func showError(err error) {
	if err == nil {
		return
	}
	m := tview.NewModal()
	m.SetText(fmt.Sprintf("ERROR: %#v", err.Error()))
	m.AddButtons([]string{"Ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Ok" {
				changeScreen(layout)
				app.SetFocus(pages)
			}
		})
	changeScreen(m)
}

var layout *tview.Flex

//var runScreen *runnerScreen

func saveConfig() error {
	/*
		bt, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		return ioutil.WriteFile(CONFIG_FILE, bt, 0644)
	*/
	return nil
}

func StartUI(file string) error {
	encoding.Register()

	smf.New()
	song, err := smf.ReadSMF(file)

	if err != nil {
		return err
	}

	fmt.Println(song.BarLines())

	/*
		_, err := os.Stat(CONFIG_FILE)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
			data = &lineconfig.Config{}

		} else {
			data, err = readConfig(CONFIG_FILE)
			if err != nil {
				return err
			}
		}
	*/
	app = tview.NewApplication()
	/*
		if len(data.Stacks) > 0 {
			chosenStack = data.Stacks[0].Name
		}
	*/
	pagesRight = tview.NewPages()
	//runScreen = newRunnerScreen()

	//changeScreen(runnerForm())
	pages = tview.NewPages()
	//pages.AddAndSwitchToPage("runner", runScreen, true)
	//pages.AddPage("Condition", newConditionScreen(), true, false)
	//pages.AddPage("action", newActionScreen(), true, false)
	pages.AddPage("line", newLineScreen(), true, false)

	/*
		pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEnter {
				panic("got enter")
				return nil
			}
			return event
		})
	*/
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(tview.AlignCenter)
	_ = info

	info.SetText("[red]F10 [white]help [yellow]| [red]F6 [white]connection [yellow]| [red]F7[white] conditions [yellow]| [red]F8[white] actions [yellow]| [red]F9[white] lines [yellow]| [red]CTRL+S [white]save [yellow]| [red]CTRL+Q [white]quit")
	//	tview.Print(info, "hello world", 0, 0, 11, 1, tcell.ColorWhite)
	// Create the main layout.

	layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexColumn).
				AddItem(pages, 0, 1, true).
				AddItem(pagesRight, 0, 2, true),
			0, 2, true).
		AddItem(info, 1, 1, false)

	/*
		layout := tview.NewGrid()
		layout.AddItem(pages, 1, 1, 300, 200, 300, 200, true)
	*/
	//	layout.AddItem(info, 2, 1, 300, 200, 300, 200, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			app.SetFocus(pages)
		case tcell.KeyRight:
			app.SetFocus(pagesRight)
		case tcell.KeyF10:
			// print help
			m := tview.NewModal()
			m.SetText("HELP. Nothing to see here, work in progress")
			m.AddButtons([]string{"Ok"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Ok" {
						changeScreen(layout)
						app.SetFocus(pages)
					}
				})
			changeScreen(m)
		case tcell.KeyF7:
			pagesRight.RemovePage("form")
			//			pages.RemovePage("items")
			//			changeScreen(newMatcherScreen())
			//			ms := newMatcherScreen()
			//			pages.AddAndSwitchToPage("items", ms, true)
			//			app.SetFocus(ms)
			//				app.SetFocus(pagesRight)
			//				app.SetFocus(pages)
			pages.SwitchToPage("Condition")
			app.SetFocus(pages)

		case tcell.KeyCtrlS:
			err := saveConfig()
			m := tview.NewModal()
			if err == nil {
				m.SetText("File saved.")

			} else {
				//m.SetText(fmt.Sprintf("ERROR: could not save file %#v:%v", CONFIG_FILE, err))
			}
			m.AddButtons([]string{"Ok"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Ok" {
						changeScreen(layout)
						app.SetFocus(pages)
					}
				})
			changeScreen(m)

		//case tcell.KeyCtrlT:
		case tcell.KeyF8:
			pagesRight.RemovePage("form")
			pages.SwitchToPage("action")
			app.SetFocus(pages)
		case tcell.KeyCtrlQ:
			/*
				if runScreen.proxy != nil {
					runScreen.proxy.Stop()
					runScreen.proxy = nil
				}
			*/
			app.Stop()
		case tcell.KeyCtrlC:
			/*
				if runScreen.proxy != nil {
					runScreen.proxy.Stop()
					runScreen.proxy = nil
				}
			*/
			app.Stop()
		case tcell.KeyF9:
			pagesRight.RemovePage("form")
			pages.SwitchToPage("line")
			app.SetFocus(pages)
		case tcell.KeyF6:
			pagesRight.RemovePage("form")
			pages.SwitchToPage("runner")
			app.SetFocus(pages)
		case tcell.KeyEscape, tcell.KeyTab:
			return event
		default:
			//			panic(fmt.Sprintf("key: %#v", event.Key()))
		}
		return event
		//return nil
	})

	/*
		flex := tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(
							tview.NewFlex().SetDirection(tview.FlexColumn).
							AddItem(nil,0,1,false).
							//AddItem(runnerPage(),0,6,true), 0, 5, true).
							//AddItem(newMatcherScreen(),0,6,true), 0, 5, true).
							AddItem(newTransformerScreen(),0,6,true), 0, 5, true).
							//AddItem(nil,0,1,true), 0, 5, true).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 0, 1, false)

		changeScreen(flex)
	*/
	//changeScreen(runnerPage())
	changeScreen(layout)
	app.SetFocus(pages)
	/*
		pages.AddPage("runner", runnerPage(), false, true)
		pages.AddPage("matcher", matcherPage(), false, false)
		pages.AddPage("transformer", transformerPage(), false,false)
		pages.AddPage("stacks", stacksPage(),false,false)
	*/
	return app.Run()
}
