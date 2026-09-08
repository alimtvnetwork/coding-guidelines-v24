package applogger

import "fmt"

type DriverType byte

const (
	DriverConsole DriverType = iota
	DriverFile
	DriverSQLite
	DriverZap
	DriverComposite
	DriverRotatingFile
)

var driverNames = [...]string{
	"Console",
	"File",
	"SQLite",
	"Zap",
	"Composite",
	"RotatingFile",
}

func (d DriverType) Name() string {
	if int(d) < len(driverNames) {
		return driverNames[d]
	}

	return fmt.Sprintf("Driver(%d)", byte(d))
}

func (d DriverType) String() string {
	return d.Name()
}

func (d DriverType) IsValid() bool {
	return int(d) < len(driverNames)
}
