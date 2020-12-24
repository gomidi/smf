package metronome

import (
	"time"

	"gitlab.com/gomidi/midi/smf"
)

// absPosToMsec transforms the absolute ticks to microseconds based on the tempi
func absPosToMsec(metricTicks smf.MetricTicks, temps tempi, absPos uint64) (msec int64) {

	/*
		calculate the abstime in msec for every tempo position up to the last tempo position before absPos
		the abstime of a tempo position is calculated the following way:

		absTime = absTimePrevTempo + metricTicks.FractionalDuration(lastTempo, uint32(absPosCurrent - absPosPrevious)).Microseconds()

		the abstime of the ticks is

		absTime = absTimePrevTempo + metricTicks.FractionalDuration(lastTempo, uint32(absPos - absPosPrevious)).Microseconds()
	*/

	var absTimeLastTempo int64
	var absTicksLastTempo uint64
	var lastTempo float64 = 120.0

	for _, tm := range temps {
		if tm.absPos > absPos {
			break
		}

		if tm.absPos == 0 {
			lastTempo = tm.bpm
			continue
		}

		absTime := absTimeLastTempo + metricTicks.FractionalDuration(lastTempo, uint32(tm.absPos-absTicksLastTempo)).Microseconds()

		absTimeLastTempo = absTime
		absTicksLastTempo = tm.absPos
		lastTempo = tm.bpm
	}

	if absPos == absTicksLastTempo {
		return absTimeLastTempo
	}

	msec = absTimeLastTempo + metricTicks.FractionalDuration(lastTempo, uint32(absPos-absTicksLastTempo)).Microseconds()
	//fmt.Printf("converted tick at %v to microsecs %v\n", absPos, msec)
	return
}

// msecToAbsPos calculates the ticks based on the microseconds and the tempi
func msecToAbsPos(metricTicks smf.MetricTicks, temps tempi, msec int64) (absPos uint64) {

	/*
		calculate the abstick for every tempo absTime up to the last tempo time before msec
		the abstick of a tempo time is calculated the following way:

		abstick = absTickPrevTempo +  metricTicks.FractionalTicks(lastTempo, (absTimeCurrent-absTimePrevious) *time.Microsecond )

		the abstime of the ticks is

		abstick = absTickPrevTempo +  metricTicks.FractionalTicks(lastTempo, (msec-absTimePrevious) *time.Microsecond )
	*/

	var absTickLastTempo uint64
	var absTimeLastTempo int64
	var lastTempo float64 = 120.0

	for _, tm := range temps {
		if tm.msec > msec {
			break
		}

		if tm.msec == 0 {
			lastTempo = tm.bpm
			continue
		}

		abstick := absTickLastTempo + uint64(metricTicks.FractionalTicks(lastTempo, time.Duration(tm.msec-absTimeLastTempo)*time.Microsecond))

		absTickLastTempo = abstick
		absTimeLastTempo = tm.msec
		tm.absPos = abstick
		lastTempo = tm.bpm
	}

	if msec == absTimeLastTempo {
		return absTickLastTempo
	}

	absPos = absTickLastTempo + uint64(metricTicks.FractionalTicks(lastTempo, time.Duration(msec-absTimeLastTempo)*time.Microsecond))
	//fmt.Printf("converted microsec %v to abstick %v\n", msec, absPos)
	return

}

// timeDistanceToTempo calculates the tempo based on the distance in microseconds
func timeDistaneToTempo(msecA, msecB int64) (bpm float64) {
	return float64(60000000) / float64(msecB-msecA)
}
