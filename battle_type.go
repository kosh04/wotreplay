package wotreplay

import (
	"encoding/json"
	"fmt"
)

type battleType int

// http://wiki.wargaming.net/en/Game_Modes
func (x battleType) String() string {
	msg := []string{
		0:  "???",
		1:  "Random",
		2:  "Training",
		3:  "Tank Company",
		4:  "Special Battle",    // wotreplays say "Tournament"
		5:  "Clan Wars",         //
		6:  "???",               //
		7:  "7vs7",              // Team Battle ?
		8:  "Historical Battle", //
		9:  "AprilFool 2014",    // wotreplays say "Soccer"
		10: "Strongholds",
		11: "Battle for Stronghold",
		12: "Ladder",
		13: "Global map",
		17: "Proving Ground",
		22: "Ranked Battle",
		23: "Boot Camp",
		24: "Grand Battle", // a.k.a 30vs30
		26: "Halloween 2017",
		27: "Frontline",
	}
	i := int(x)
	if !(0 <= i && i < len(msg)) {
		return fmt.Sprintf("%d:???", i)
	}
	return fmt.Sprintf("%d:%s", i, msg[i])
}

func (x battleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
