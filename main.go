//generate:boxy
package main

import (
	"flag"

	"github.com/OrbTools/OrbMap/devices/orbweaver"
	"github.com/OrbTools/OrbMap/emu"
	"github.com/OrbTools/OrbMap/interface/keyevents"
)

func main() {
	var orbs *string
	orbs = flag.String("orbweaver", "", "Comma seperated string of orbs for the orbweaver")
	flag.Parse()
	KeyBus := make(chan *keyevents.KeyEvent, 128)
	if *orbs != "" {
		Maps := orbweaver.ProcOrbFiles(*orbs)
		go orbweaver.OrbLoop(Maps, KeyBus)
	}
	emu.ProcKey(KeyBus)
}
