package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Import struct {
	importArgs []KVArgs
}

type Formatted struct {
	Import
	filters KVArgs
}

func NewImageImport(filePath string, duration float64) (i *Import) {
	videoArgs := KVArgs{
		"framerate": 25,
		"loop":      1,
		"t":         duration,
		"i":         filePath,
	}
	audioArgs := KVArgs{
		"f": "lavfi",
		"t": duration,
		"i": "anullsrc",
	}
	return &Import{importArgs: []KVArgs{videoArgs, audioArgs}}
}

func NewVideoImport(filePath string) (i *Import) {
	return &Import{importArgs: []KVArgs{{
		"i": filePath,
	}}}
}

func (i *Import) Filter(filterOptions KVArgs) (f *Formatted) {
	return &Formatted{
		Import:  *i,
		filters: filterOptions,
	}
}

func Concat(items []*Formatted, outputArgs []string) (err error) {
	importArgs := make([]string, 0, 64)
	formatFilter := strings.Builder{}
	concatFilter := strings.Builder{}
	streamIndex := 0

	for _, item := range items {
		// Create import args list
		for _, args := range item.importArgs {
			importArgs = append(importArgs, args.toArgs()...)
		}
		// Create filter list
		formatFilter.WriteString(fmt.Sprintf("[%d:v]", streamIndex))
		formatFilter.WriteString(item.filters.toFilter())
		formatFilter.WriteString(fmt.Sprintf("[v%d];", streamIndex))
		// Add concat inputs
		if len(item.importArgs) == 2 {
			// Type image
			concatFilter.WriteString(fmt.Sprintf("[v%d][%d]", streamIndex, streamIndex+1))
		} else {
			// Type video
			concatFilter.WriteString(fmt.Sprintf("[v%d][%d:a]", streamIndex, streamIndex))
		}
		streamIndex += len(item.importArgs)
	}
	concatFilter.WriteString(fmt.Sprintf("concat=v=1:a=1:n=%d[out]", len(items)))

	filter := strings.Builder{}
	filter.WriteString(formatFilter.String())
	filter.WriteString(concatFilter.String())

	concatArgs := append(importArgs, "-filter_complex", filter.String(), "-map", "[out]")

	cmd := exec.Command("ffmpeg", append(concatArgs, outputArgs...)...)
	println("generated command:", cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
