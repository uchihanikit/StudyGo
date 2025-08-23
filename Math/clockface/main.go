package main

import (
	"os"
	"time"

	svg "github.com/uchihanikit/StudyGo/Math/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
