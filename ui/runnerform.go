package ui

import (
	"fmt"

	"github.com/rivo/tview"
	"gitlab.com/gomidi/midi"
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

// run a configuration
func (s *runnerScreen) runnerForm() *tview.Form {
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
}
