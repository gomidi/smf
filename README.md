# smf

Various tools to deal with Standard MIDI Files (SMF).

Note: If you are reading this on Github, please note that the repo has moved to Gitlab (gitlab.com/gomidi/smf) and this is only a mirror.


## Installation

    go get -u gitlab.com/gomidi/smf/cmd/smf

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
 
