package ui

import (
	"github.com/rivo/tview"
	"gitlab.com/gomidi/midi"
)

type editlayout struct {
	*tview.Flex
	//ms *messageSelect
	//left        *tview.Flex
	left        *tview.Form
	right       *tview.Flex
	messageForm *EditForm
}

func newEditLayout(ch uint8, saveCb func(m midi.Message), cancelCb func(), setFocus func(p tview.Primitive)) *editlayout {
	el := &editlayout{
		Flex: tview.NewFlex(),
	}

	//el.ms = newMessageSelect()

	el.left = tview.NewForm()

	/*
		fm := tview.NewForm()
		fm.AddDropDown("new message", _messagesSelect, 0, func(opt string, idx int) {
			*pmsg = newMessageInTrack(uint8(ch), opt)
		})
	*/
	el.AddItem(el.left, 50, 1, true)
	el.right = tview.NewFlex()
	el.AddItem(el.right, 0, 2, false)

	el.left.AddDropDown("message type ", _messagesSelect, 0, func(opt string, idx int) {
		msg := newMessageInTrack(uint8(ch), opt)
		el.setMessageForm(msg, saveCb, cancelCb)
		setFocus(el.right)
	})

	/*
		el.ms.AddItems(func() {
			typ, _ := el.ms.GetItemText(el.ms.GetCurrentItem())
			msg := newMessageInTrack(ch, typ)
			el.setMessageForm(msg, saveCb, cancelCb)
		})
	*/

	//el.left.AddItem(el.ms.List, 0, 1, true)

	return el
}

func (e *editlayout) setMessageForm(msg midi.Message, saveCb func(m midi.Message), cancelCb func()) {
	if e.messageForm != nil {
		e.right.RemoveItem(e.messageForm)
	}

	e.messageForm = NewEditForm(msg, saveCb, cancelCb)
	e.right.AddItem(e.messageForm, 0, 1, true)
}
