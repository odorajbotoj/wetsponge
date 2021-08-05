package WetSpongeCore

import (
	"fmt"
	"time"
)

type Ver_struct struct {
	Main        uint16
	Sub         uint16
	Revision    uint16
	Time        time.Time
	Description string
}

func (v Ver_struct) String() string {
	ret := fmt.Sprintf("%d.%d.%d", v.Main, v.Sub, v.Revision)
	return ret
}

func (v Ver_struct) GetInfo() string {
	ret := fmt.Sprintf("Version:%d.%d.%d\nTime:%s\nDescription:%s\n", v.Main, v.Sub, v.Revision, v.Time, v.Description)
	return ret
}

var cstZone = time.FixedZone("CST", 8*3600)
var VERSION = Ver_struct{1, 0, 0, time.Date(2021, 8, 5, 14, 30, 0, 0, cstZone), "The First!"}

func init() {
	const AUTHOR string = "OdorajBotoj"
	fmt.Printf("Welcome!WetSponge(Core)Version:%s,By:%s\n", VERSION, AUTHOR)
}
