package ffmpeg

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const command = "ffprobe"

func GetDuration(filePath string) (duration float64, err error) {
	kvArgs := KVArgs{
		"v":            "error",
		"show_entries": "format=duration",
		"of":           "default=noprint_wrappers=1:nokey=1",
	}

	outBuffer, err := runFFProbeCmd(filePath, kvArgs)
	if err != nil {
		return -1, err
	}

	return strconv.ParseFloat(strings.TrimSpace(outBuffer.String()), 64)
}

func GetFrameRate(filePath string) (frameRate float64, err error) {
	kvArgs := KVArgs{
		"v":              "error",
		"select_streams": "v:0",
		"show_entries":   "stream=avg_frame_rate",
		"of":             "default=noprint_wrappers=1:nokey=1",
	}

	outBuffer, err := runFFProbeCmd(filePath, kvArgs)
	if err != nil {
		return -1, err
	}

	fpsString := strings.TrimSpace(outBuffer.String())
	split := strings.Split(fpsString, "/")
	if len(split) != 2 {
		return -1, errors.New("ffprobe returned unexpected frame rate")
	}

	frames, err := strconv.ParseFloat(split[0], 64)
	fraction, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return -1, err
	}

	return frames / fraction, nil
}

func runFFProbeCmd(filePath string, kvArgs KVArgs) (output *bytes.Buffer, err error) {
	args := kvArgs.toArgs()
	args = append(args, filePath)

	cmd := exec.Command(command, args...)

	outBuffer := bytes.NewBuffer(nil)

	cmd.Stdout = outBuffer
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	return outBuffer, err
}
