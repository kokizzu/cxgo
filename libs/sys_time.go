package libs

import (
	_ "embed"

	"github.com/gotranspile/cxgo/runtime/csys"
	"github.com/gotranspile/cxgo/types"
)

const (
	sysTimeH = "sys/time.h"
)

//go:embed sys_time.h
var hsys_time string

func init() {
	RegisterLibrary(sysTimeH, func(c *Env) *Library {
		intT := types.IntT(4)
		timeLib := c.LookupLibrary(timeH)
		timevalT := timeLib.GetType("timeval")
		timespecT := timeLib.GetType("timespec")
		_ = timespecT
		return &Library{
			Imports: map[string]string{
				"libc": RuntimeLibc,
				"csys": RuntimePrefix + "csys",
			},
			Idents: map[string]*types.Ident{
				"gettimeofday": c.NewIdent("gettimeofday", "csys.GetTimeOfDay", csys.GetTimeOfDay, c.FuncTT(intT, c.PtrT(timevalT), c.PtrT(nil))),
			},
			// TODO
			Header: hsys_time,
		}
	})
}
