module gitlab.com/gomidi/smf/cmd/smf

go 1.14

require (
	gitlab.com/gomidi/midi v1.20.3
	gitlab.com/gomidi/rtmididrv v0.9.3
	gitlab.com/gomidi/smf v0.0.5
	gitlab.com/gomidi/smf/ui v0.0.5
	gitlab.com/metakeule/config v1.15.5
)

replace gitlab.com/gomidi/smf => ../../
replace gitlab.com/gomidi/smf/ui => ../../ui
