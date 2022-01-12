package color

import (
	"time"

	"github.com/fatih/color"
)

var Textcolor = color.New(color.FgHiWhite, color.Bold).SprintfFunc()
var Workercolor = color.New(color.FgHiCyan, color.Bold).SprintfFunc()
var Directorycolor = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintfFunc()
var Filenamecolor = color.New(color.FgHiGreen, color.BlinkRapid, color.Bold).SprintfFunc()
var Timecolor = color.New(color.FgHiMagenta, color.Bold).SprintfFunc()
var Notificationcolor = color.New(color.FgHiRed, color.Bold).SprintfFunc()
var Now = time.Now()
