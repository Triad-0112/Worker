package color

import (
	"time"

	"github.com/fatih/color"
)

var textcolor = color.New(color.FgHiWhite, color.Bold).SprintfFunc()
var workercolor = color.New(color.FgHiCyan, color.Bold).SprintfFunc()
var directorycolor = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintfFunc()
var filenamecolor = color.New(color.FgHiGreen, color.BlinkRapid, color.Bold).SprintfFunc()
var timecolor = color.New(color.FgHiMagenta, color.Bold).SprintfFunc()
var notificationcolor = color.New(color.FgHiRed, color.Bold).SprintfFunc()
var now = time.Now()
