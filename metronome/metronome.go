package metronome

import (
	"fmt"

	"gitlab.com/gomidi/midi/smf"
)

/*

This package allows to find the tempo changes based on a metronome track.
All events are converted from tick based time positions to milliseconds based time position.
Then the tempo changes are detected via the time position difference between the beats
in the metronome track.
Then the tick based positions are calculated for all events based on the tempo changes.
The result is written to a new file

*/

/*
SetTempo reads the given srcFile and takes the note on messages inside the nth track
(with the numer metronomeTrackNo, counting by 0) as beats of a metronome, calculates
tempo changes based on the beats of the metronome and removes
all old tempo changes, and writes the new tempo changes into the metronome track.
The resulting smf file is written to destFile.
*/
func SetTempo(srcFile, destFile string, opts ...Option) error {
	f := newFile(srcFile, opts...)
	err := f.read()

	if err != nil {
		return fmt.Errorf("can't read %q: %s", srcFile, err.Error())
	}

	f.convertTicksToMsec()
	f.calcNewTempi()
	f.convertMsecToTicks()

	err = f.writeTo(destFile)

	if err != nil && err != smf.ErrFinished {
		return fmt.Errorf("can't write %q: %s %T", destFile, err.Error(), err)
	}

	return nil
}
