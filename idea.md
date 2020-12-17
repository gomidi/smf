
# First Implementation / Priority

- main screen
- midi in/out config
- edit lines and cells
- show events
- track config
- beats -> tempo mode


# Idea

- a CLI gui tool to inspect and edit smf files
- if smf type 0: columns are MIDI channels
- if smf type 1: columns are track/channel combinations
- time is running top to bottom, may scroll while playing
- menu on top
- view switch / special F1-F9 keys on bottom (like midnight explorer)
- also a player: configurable MIDI-out connectors
- also a recorder: configurable MIDI-in connectors (that also let the MIDI pass through to outs)
- also a metronome (for recording)
- also a graphical MIDI keyboards
- ability to show/hide tracks
- can have "attached", non-editable and watched MIDI files that are not merged into the current MIDI file but played together and that can be merged and exported per command
- position: beat.decimal part of beat
- Comments column indicates, if there is a comment. that could be shown/edited via command shortcut or displayed while playing,
  comments are numbered by appearance, so they can by shown by number, they will be saved as text in a separate track
- ENTER creates a new line, shift+ENTER removes a line
- CTRL+left/right next/previous column
- CTRL+up/down next/previous marker
- CTRL+T new track-column
- CTRL+P play from current position
- tab/shift+tab push current cell and following cells down/up a bar while keeping the cursor in the cell
- CTRL-tab/CTRL+shift+tab push current cell and following cells down/up to the next marker while keeping the cursor in the cell
- ALT-tab/shift+alt+tab push current cell and following cells down/up a beat while keeping the cursor in the cell
- CTRL-ALT-tab/shift+ctrl+alt+tab push current cell and following cells down/up half a beat while keeping the cursor in the cell
- CTRL+shift+up/down move the line up or down while keeping the cursor in the line
- CTRL+shift+left/right move the position of the line one beat forward/back
- CTRL+Alt+left/right move the position of the line one half beat forward/back
- CTRL+N create a new line on the next bar
- CTRL+B create a new line on the next beat
- CTRL+H create a new line on the next half beat
- export of the selected column to muskel file
- have swing and delay settings for a column
- impossible actions (like pushing post the end) will be indicated by a short screen flash
- simultaneous events in empty line below
- allow sending sysex and other initialization MIDI to external gear on start
- general configs for shortcuts per instruments (like drumnotes, program changes etc.)
- ESC always gets you to the main screen
- main screen can have 5 different views that show/hide different tracks and keep their cursor position independant from each other
- CTRL+ESC hide current track from current view
- ALT+right scroll one column to the right
- ALT+left scroll one column to the left
- ALT+down scroll one column-page (=all displayed columns) to the right
- ALT+up scroll one column-page (=all displayed columns) to the left
- when doing horizontal/column scrolling the columns Comment Mark Bar and Beat stay visible
- a track may not have more than 10 characters and may only have ascii characters
- comments and text are UTF-8
- there can be one config per directory that is valid for all midi files within the directory
- editing is only possible while not playing/recording, however keyboard play and midi through is always possible while editing
- the currently selected line is highlighted while in editing mode
- while playing / recording the current position line(s) are highlighted
- ALT+S Solo current track (toggle)
- ALT+M mute current track (toggle)
- ALT+R record arm current track (toggle)
- ALT+O edit options for current track
- there is also a play line that is separate from the cursor; F1 plays/stops the play line. when stopping with F1 the play line is set to the beginning of the bar of the stopping position
- SPACE plays/stops from the cursor line. when stopping, the play line is not affected. one could start playing with F1 and stopping with SPACE or vice versa
- CTRL+SPACE sets the play line to the start of the bar of the cursor line
- CTRL+F1 jumps the screen to the playline without moving the cursor.
- ALT+F1 jumps the screen to the playline and sets the cursor to the playline while keeping it in the same column
- import of a track from another midi file at the current cursor position inside the current track
- Transpose selected notes
- shift+arrow-keys select: grow-shrink per column (left/right) and line/position (up/down)
- CTRL+ENTER: edit cell value (dialog) or simply write to override all
- CTRL++/CTRL+- change value / velocity (notes) one step at a time
- CTRL+ALT+/CTRL+ALT- change value / velocity (notes) ten steps at a time
- CTRL+#/CTRL+b change note by halftone up/down and play the result via midi ouput port
- CTRL+ALT+#/CTRL+ALT+b change note by octave up/down and play the result via midi ouput port
- allow to show / edit notes/cc/program names via descriptive textshortcuts that are resoved on a per track basis; the syntax is 'shortcut/value' e.g.
  'kick/106' (note) or 'chorus/90' (cc) or 'rockorgan' (program)
- when importing midi have intelligence to detect track delay
- allow beat-track to help retrofit tempo changes

## Main Screen

-----------------------------------------------------------------------
File | Edit | View | Config         (the menu, open the first with ALT+SPACE and then navigate with arrow keys and select with ENTER)
-----------------------------------------------------------------------
Comment | Mark  | Bar  | Beat || Drums[10] | Bass[9]  | Vocal[1] | Piano[1]  | (piano track on channel 1 etc)
        |       |      |      || S M R     | S M R    | S M R    | S M R     | (Solo/Mute/Record indicators)
----------------------------------------------------------------------- (everything above this line is static/non scrollable)
1       | Intro | 4/4  | 1.0  || C3/100    | C5_/120  |          |           | (drum note is just a 32ths, bass is note on)
        |       | 144  | 1.0  ||           |          |          |           | tempo change
                  #2                                                           (bar change)
        |       |      | 2.25 || C5/60     | _C5      | "hiho"   | CC123/100 |
=====I====>===V====C=== position indicator, always the pre-last line of the screen (each = is a bar, each letter is the first letter of a Marker) 
F1 Play | F2 Rec | F3 Metro | F4 Keyb | F5 V1 | F6 V2 | F7 V3 | F8 V4 | F9 V5 | F10 Track Properties | F11 Song Properties
(play, record, metronome, Keyboard are switches that indicate if it is active)
(views are a selector; only one view can be active at a time)


## Song Properties Screen

- meta-data like title, copyright etc.



## Track Properties Screen

|           | Drums[10] | Bass[9]   | Vocal[1] | Piano[1] | (piano track on channel 1 etc)
|           | S M R     | S M R     | S M R    | S M R    | (Solo/Mute/Record indicators)
=====================================================
| Program   |           | 39        | 20       | 1        | 
| Volume    | 100       | 80        | 60       | 110      |
| Out-Port  | 2         | 2         | 2        | 1        |
| In-Port   |           |           |          | 1        |
| Transpose |           |           |          | 12       |
| Delay     |           |           |          | -1/8     |
| Views     | 1 2 3     | 1 2 3 4 5 | 2        | 4 5      | (in which views the tracks are shown)        

