package puzzle

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// State ...
type State struct {
	Dim    uint      `json:"dim"`
	Puzzle [][]uint8 `json:"puzzle"`
}

func LoadStateFromFile(path string) (s State, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&s)
	return
}

func (s State) String() string {
	boxStride := s.Dim
	stride := boxStride * boxStride
	sb := new(strings.Builder)
	sb.WriteString("\n--------------------------------\n")
	for y := uint(0); y < stride; y += boxStride {
		for by := y; by < y+boxStride; by++ {
			row := s.Puzzle[by]
			sb.WriteString("| ")
			for x := uint(0); x < stride; x += boxStride {
				for bx := x; bx < x+boxStride; bx++ {
					fmt.Fprintf(sb, "%0.2X ", row[bx])
				}
				if x+boxStride < stride {
					sb.WriteByte(' ')
				}
			}
			sb.WriteString("|\n")
		}
		if y+boxStride < stride {
			sb.WriteString("|                              |\n")
		}
	}
	sb.WriteString("--------------------------------\n")

	return sb.String()
}
