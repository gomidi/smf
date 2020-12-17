# smf

Various tools to deal with Standard MIDI Files (SMF).

## Installation

    go get -u github.com/gomidi/smf/cmd/smf

## Usage

    smf help
    
### extract lyrics from an SMF file

    smf lyrics -f='your-midi-file.mid'
    
documentation:

    smf help lyrics

### set tempo based on a metronome track

set the tempo based on the metronome beats on the second track 

    smf metro -f='input.mid' -o='ouput.mid' -t=1
    
documentation:

    smf help metro

### see content of SMF file

    smf cat -f='your-midi-file.mid'
    
documentation:

    smf help cat
 
