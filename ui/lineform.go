package ui

import (
	"fmt"

	"github.com/rivo/tview"
	//lineconfig "gitlab.com/gomidi/midiline/config"
)

func lineFormAdd() (form *tview.Form) {
	/*
		form = tview.NewForm()
		var newFunc = ""
		var newargs []interface{}
		var newname = ""
		fns := getTransformerFuncs()
		form.
			AddInputField("Name", "", 100, nil, func(text string) {
				newname = text
			}).
			AddDropDown("Func", getTransformerFuncs(), 0, func(option string, optionIndex int) {
				for i, fn := range fns {
					if i == optionIndex {
						newFunc = fn
					}
				}
			}).
			AddInputField("Arguments (separated with comma)", "", 150, func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				var args []interface{}
				err := json.Unmarshal([]byte("["+text+"]"), &args)
				if err == nil {
					newargs = args
				}
			}).
			AddButton("add", func() {
				var m lineconfig.Transform
				m.Name = newname
				m.Func = newFunc
				m.Args = newargs
				addTransformer(m)
				pages.SwitchToPage("action")
				app.SetFocus(pages)
				//			pagesRight.RemovePage("form")
				//			app.SetFocus(pages)
				//				changeScreen(newTransformerScreen())
				//changeScreen(runnerPage())
			}).SetCancelFunc(func() {
			pages.SwitchToPage("action")
			app.SetFocus(pages)
			app.SetFocus(pagesRight)
		})
		form.SetBorder(true).SetTitle("Add Transformer").SetTitleAlign(tview.AlignLeft)
		return form
	*/

	form = tview.NewForm()
	//	var newMatch = ""
	//	var newTransform = ""
	//	var newname = ""
	//	var s lineconfig.Line
	//	var m lineconfig.Transformation
	//	fns := getMatchers()
	//	trs := getTransformers()
	form.
		AddInputField("name", "", 100, nil, func(text string) {
			//s.Name = text
		}).
		AddButton("add", func() {
			//addLine(s)
			pages.SwitchToPage("line")
			app.SetFocus(pages)
			//changeScreen(newStackScreen())
			//changeScreen(runnerPage())
		}).
		/*
			AddDropDown("Matcher", fns, -1, func(option string, optionIndex int) {
				for i, fn := range fns {
					if i == optionIndex {
						m.Match = fn
					}
				}
			}).
			AddDropDown("Transformer", trs, -1, func(option string, optionIndex int) {
				for i, fn := range trs {
					if i == optionIndex {
						m.Transform = fn
					}
				}
			}).
			AddButton("add transformation", func() {

				//			m.Name = newname
				//			m.Match = newMatch
				//			m.Transform = newTransform
				s.Transformations = append(s.Transformations, m)
				//addStack(m)
				//changeScreen(newStackScreen())
				//changeScreen(runnerPage())
			}).
		*/
		SetCancelFunc(func() {
			pages.SwitchToPage("line")
			app.SetFocus(pages)
			app.SetFocus(pagesRight)
		})
	form.SetBorder(true).SetTitle("new line").SetTitleAlign(tview.AlignLeft)
	return form
}

func lineFormEdit(selectedLine string) (form *tview.Form) {
	form = tview.NewForm()
	//var mt lineconfig.Line
	var selected int = -1
	/*
		for i, m := range data.Lines {
			if m.Name == selectedLine {
				selected = i
				mt = m
			}
		}

		oldName := mt.Name
	*/

	if selected == -1 {
		showError(fmt.Errorf("unknown line %#v", selectedLine))
		//panic("unknown line " + selectedStack)
	}

	//	var selectedMatch = -1
	Conditions := getConditions()

	//	var selectedTransformer = -1
	actions := getActions()

	/*
		var oldName = mt.Name
		_ = oldName
	*/

	var steps []string

	/*
		for i, trf := range mt.Steps {
			steps = append(steps, fmt.Sprintf("%v. If %#v then do %#v.", i+1, trf.Condition, trf.Action))
		}
	*/

	var findConditionIndex = func(name string) int {
		for i, mmt := range Conditions {
			if mmt == name {
				return i
			}
		}
		return -1
	}

	_ = findConditionIndex

	var findActionIndex = func(name string) int {
		for i, ttrf := range actions {
			if ttrf == name {
				return i
			}
		}
		return -1
	}

	_ = findActionIndex

	//var tr lineconfig.Step
	var removeActionIndex = -1

	var ConditionDropDown = tview.NewDropDown()
	ConditionDropDown.SetLabel("Condition")
	ConditionDropDown.SetOptions(Conditions, func(option string, optionIndex int) {
		for i, f := range Conditions {
			if i == optionIndex {
				//tr.Condition = f
				_ = f
			}
		}
	})
	ConditionDropDown.SetCurrentOption(-1)

	var actionDropDown = tview.NewDropDown()
	actionDropDown.SetLabel("action")
	actionDropDown.SetOptions(actions, func(option string, optionIndex int) {
		for i, f := range actions {
			if i == optionIndex {
				//tr.Action = f
				_ = f
			}
		}
	})
	actionDropDown.SetCurrentOption(-1)

	var stepDropDown = tview.NewDropDown()
	stepDropDown.SetLabel("select step")

	var refreshStepDropDown = func() {
		stepDropDown.SetOptions(steps, func(option string, optionIndex int) {
			for i := range steps {
				if i == optionIndex {
					removeActionIndex = i
					//tr.Condition = mt.Steps[i].Condition
					//ConditionDropDown.SetCurrentOption(findConditionIndex(mt.Steps[i].Condition))
					//tr.Action = mt.Steps[i].Action
					//actionDropDown.SetCurrentOption(findActionIndex(mt.Steps[i].Action))
					//					findMatchIndex
				}
			}
		})
		stepDropDown.SetCurrentOption(-1)
	}

	refreshStepDropDown()

	if len(steps) > 0 {
		form.AddFormItem(stepDropDown)
	}

	form.
		/*
			AddInputField("Name", mt.Name, 100, nil, func(text string) {
				mt.Name = text
			}).
		*/
		/*
			AddButton("save", func() {
				replaceStack(oldName, mt)
				pages.SwitchToPage("line")
				app.SetFocus(pages)
				//			changeScreen(newStackScreen())
				//			changeScreen(runnerPage())
			}).
		*/
		//		AddFormItem(stepDropDown).
		/*
			AddDropDown("select transformation", transformations, -1, func(option string, optionIndex int) {
				for i := range transformations {
					if i == optionIndex {
						removeTrafoIndex = i
					}
				}
			}).
		*/
		AddButton("remove step", func() {
			if removeActionIndex == -1 {
				return
			}

			//var newtrf []lineconfig.Step

			/*
				for i, trr := range mt.Steps {
					if i != removeActionIndex {
						newtrf = append(newtrf, trr)
					}
				}
			*/

			/*
				mt.Steps = newtrf
				replaceLine(oldName, mt)
			*/
			pages.SwitchToPage("line")
			app.SetFocus(pages)
			app.SetFocus(pagesRight)
			refreshStepDropDown()

			//mt.Transformations = append(mt.Transformations, tr)
			//replaceStack(oldName, mt)
			//			changeScreen(newStackScreen())
			//			changeScreen(runnerPage())
		}).
		AddFormItem(ConditionDropDown).
		AddFormItem(actionDropDown).
		/*
			AddDropDown("Match", fns, -1, func(option string, optionIndex int) {
				for i, f := range fns {
					if i == optionIndex {
						tr.Match = f
					}
				}
			}).
			AddDropDown("Transform", trs, -1, func(option string, optionIndex int) {
				for i, f := range trs {
					if i == optionIndex {
						tr.Transform = f
					}
				}
			}).
		*/
		AddButton("add step", func() {
			//mt.Steps = append(mt.Steps, tr)
			//replaceLine(oldName, mt)
			pages.SwitchToPage("line")
			app.SetFocus(pages)
			app.SetFocus(pagesRight)
			refreshStepDropDown()
			//replaceStack(oldName, mt)
			//			changeScreen(newStackScreen())
			//			changeScreen(runnerPage())
		}).
		SetCancelFunc(func() {
			pages.SwitchToPage("line")
			app.SetFocus(pages)
			app.SetFocus(pagesRight)
		})

	form.SetBorder(true).SetTitle(fmt.Sprintf("edit line %#v", selectedLine)).SetTitleAlign(tview.AlignLeft)
	/*
		form = tview.NewForm()
		var mt lineconfig.Transform
		var selected int = -1
		for i, m := range data.Transformers {
			if m.Name == selectedTransformer {
				selected = i
				mt = m
			}
		}

		if selected == -1 {
			panic("unknown matcher " + selectedTransformer)
		}

		var args []string

		for _, arg := range mt.Args {
			args = append(args, fmt.Sprintf("%v", arg))
		}

		var selectedFunc = -1
		fns := getTransformerFuncs()

		for i, f := range fns {
			if f == mt.Func {
				selectedFunc = i
			}
		}
		var oldName = mt.Name
		_ = oldName

		form.
			AddInputField("Name", mt.Name, 100, nil, func(text string) {
				mt.Name = text
			}).
			//		AddFormItem(dd).

			AddDropDown("Func", fns, selectedFunc, func(option string, optionIndex int) {
				for i, f := range fns {
					if i == optionIndex {
						mt.Func = f
					}
				}
			}).
			AddInputField("Arguments (separated with comma)", strings.Join(args, ","), 150, func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				var args []interface{}
				err := json.Unmarshal([]byte("["+text+"]"), &args)
				if err == nil {
					mt.Args = args
				}
			}).
			AddButton("save", func() {
				replaceTransformer(oldName, mt)
				pages.SwitchToPage("action")
				app.SetFocus(pages)
				//			pagesRight.RemovePage("form")
				//			app.SetFocus(pages)
				//				changeScreen(newTransformerScreen())
				//			changeScreen(runnerPage())
			}).SetCancelFunc(func() {
			pages.SwitchToPage("action")
			app.SetFocus(pages)
			app.SetFocus(pagesRight)
		})
		form.SetBorder(true).SetTitle("Edit Transformer").SetTitleAlign(tview.AlignLeft)
		return form
	*/
	return form
}

// edit or create a stack
func lineForm(selectedLine string) (form *tview.Form) {
	if selectedLine == "" {
		return lineFormAdd()
	}
	return lineFormEdit(selectedLine)
}
