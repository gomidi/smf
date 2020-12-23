module gitlab.com/gomidi/smf/ui

go 1.14

require (
	github.com/gdamore/encoding v0.0.0-20151215212835-b23993cbb635 // indirect
	github.com/gdamore/tcell v1.1.0
	github.com/lucasb-eyer/go-colorful v0.0.0-20180709185858-c7842319cf3a // indirect
	github.com/mattn/go-runewidth v0.0.3 // indirect
	github.com/rivo/tview v0.0.0-20180821142722-77bcb6c6b900
	gitlab.com/gomidi/midi v1.20.3
	gitlab.com/gomidi/rtmididrv v0.9.3
	gitlab.com/gomidi/smf v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.3.0 // indirect
)

replace gitlab.com/gomidi/smf => ../
