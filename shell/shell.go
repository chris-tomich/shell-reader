package shell

import (
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/chris-tomich/shell-reader/common"
	"github.com/chris-tomich/shell-reader/common/vt"
	"github.com/pkg/term"
)

type historyEntry struct {
	Previous *historyEntry
	Next     *historyEntry
	Line     string
}

type Reader struct {
	reader             io.Reader
	terminalController *term.Term
	lastEntry          *historyEntry
}

func NewReader() (*Reader, error) {
	terminalController, err := term.Open("/dev/tty")

	if err != nil {
		return nil, fmt.Errorf("failed to open a connection to the terminal /dev/tty\n%w", err)
	}

	err = term.CBreakMode(terminalController)

	if err != nil {
		return nil, fmt.Errorf("failed to set the terminal to cbreak mode\n%w", err)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	go func() {
		s := <-signalChannel

		if s == os.Interrupt {
			terminalController.Restore()
			fmt.Println("\n\n(signal interrupt received)")
			os.Exit(0)
		}
	}()

	return &Reader{
		os.Stdin,
		terminalController,
		nil,
	}, nil
}

func (r *Reader) Close() {
	if r.terminalController != nil {
		r.terminalController.Restore()
	}
}

func (r *Reader) ReadLine() (string, error) {
	line := make([]byte, 0, 80)
	buffer := make([]byte, 12)
	cursorPosition := 0
	historyCursor := (*historyEntry)(nil)

	_, err := r.reader.Read(buffer)

	if err == io.EOF {
		return r.returnLine("", err)
	}

	if err != nil {
		return r.returnLine("", fmt.Errorf("issue reading from reader\n%w", err))
	}

	for buffer[0] != '\n' {
		isEscaped := false

		var escapedCode common.KeyCode

		if buffer[0] == common.KeyCodes[common.KeyEsc][0] && buffer[1] != 0x00 {
			isEscaped = true
			escapedCode, err = readEscapedKey(buffer)

			if err != nil {
				_, err = r.reader.Read(buffer)

				if err == io.EOF {
					return r.returnLine(string(line), err)
				}

				if err != nil {
					return r.returnLine("", fmt.Errorf("issue reading from reader\n%w", err))
				}

				continue
			}
		}

		switch {
		case isEscaped && (escapedCode == common.KeyF1 ||
			escapedCode == common.KeyF2 ||
			escapedCode == common.KeyF3 ||
			escapedCode == common.KeyF4 ||
			escapedCode == common.KeyF5 ||
			escapedCode == common.KeyF6 ||
			escapedCode == common.KeyF7 ||
			escapedCode == common.KeyF8 ||
			escapedCode == common.KeyF9 ||
			escapedCode == common.KeyF10 ||
			escapedCode == common.KeyF11 ||
			escapedCode == common.KeyF12):
			break

		case isEscaped && (escapedCode == common.KeyInsert || escapedCode == common.KeyPageDown || escapedCode == common.KeyPageUp):
			break

		case isEscaped && escapedCode == common.KeyHome:
			if cursorPosition > 0 {
				fmt.Printf(vt.MoveBackward(cursorPosition))
				cursorPosition = 0
			}

		case isEscaped && escapedCode == common.KeyEnd:
			distance := len(line) - cursorPosition

			if distance > 0 {
				fmt.Printf(vt.MoveForward(distance))
				cursorPosition = len(line)
			}

		case isEscaped && escapedCode == common.KeyArrowUp:
			if historyCursor == nil {
				historyCursor = r.lastEntry
			} else {
				if historyCursor.Previous != nil {
					historyCursor = historyCursor.Previous
				}
			}

			if historyCursor != nil {
				if cursorPosition > 0 {
					fmt.Printf(vt.MoveBackward(cursorPosition))
					fmt.Printf(vt.EraseToEndOfLine())
				}
				line = make([]byte, len(historyCursor.Line))
				copy(line, []byte(historyCursor.Line))
				cursorPosition = len(line)
				fmt.Printf(string(line))
			}

		case isEscaped && escapedCode == common.KeyArrowDown:
			if historyCursor != nil {
				if historyCursor.Next != nil {
					historyCursor = historyCursor.Next
				}
			}

			if historyCursor != nil {
				if cursorPosition > 0 {
					fmt.Printf(vt.MoveBackward(cursorPosition))
					fmt.Printf(vt.EraseToEndOfLine())
				}
				line = make([]byte, len(historyCursor.Line))
				copy(line, []byte(historyCursor.Line))
				cursorPosition = len(line)
				fmt.Printf(string(line))
			}

		case isEscaped && escapedCode == common.KeyArrowLeft:
			if cursorPosition > 0 {
				fmt.Printf(vt.CursorBackward())
				cursorPosition--
			}

		case isEscaped && escapedCode == common.KeyArrowRight:
			if cursorPosition < len(line) {
				fmt.Printf(vt.CursorForward())
				cursorPosition++
			}

		case isEscaped && escapedCode == common.KeyDelete:
			if cursorPosition < len(line) {
				fmt.Printf(vt.SaveCursorPosition())
				fmt.Printf("%v%c", string(line[cursorPosition+1:]), ' ')
				fmt.Printf(vt.RestoreCursorPosition())
				line = append(line[0:cursorPosition], line[cursorPosition+1:]...)
			}

		case isEscaped:
			return r.returnLine(string(escapedCode), nil)

		case buffer[0] == common.KeyCodes[common.KeyEsc][0]:
			break

		case buffer[0] == common.KeyCodes[common.KeyTab][0]:
			tab := "    "
			fmt.Printf(tab)
			line = append(line, []byte(tab)...)
			cursorPosition += 4
			break

		case buffer[0] == common.KeyCodes[common.KeyBackspace][0]:
			if cursorPosition == 0 {
				break
			} else if cursorPosition == len(line) {
				fmt.Printf(vt.Backspace())
				cursorPosition--

				line = line[0 : len(line)-1]
			} else {
				cursorPosition--
				newLine := make([]byte, len(line)-1)

				copy(newLine[0:cursorPosition], line[0:cursorPosition])
				copy(newLine[cursorPosition:], line[cursorPosition+1:])

				distance := len(newLine) - cursorPosition + 1

				fmt.Printf(vt.CursorBackward())
				fmt.Printf("%s ", string(newLine[cursorPosition:]))
				fmt.Printf(vt.MoveBackward(distance))

				line = newLine
			}

		default:
			if cursorPosition == len(line) {
				fmt.Printf("%c", buffer[0])
				cursorPosition++

				line = append(line, buffer[0])
			} else {
				newLine := make([]byte, len(line)+1)

				copy(newLine[0:cursorPosition], line[0:cursorPosition])
				copy(newLine[cursorPosition+1:], line[cursorPosition:])
				newLine[cursorPosition] = buffer[0]

				fmt.Printf("%c%s", buffer[0], string(line[cursorPosition:]))
				fmt.Printf(vt.MoveBackward(len(line[cursorPosition:])))

				line = newLine

				cursorPosition++
			}
		}

		_, err := r.reader.Read(buffer)

		if err == io.EOF {
			return r.returnLine(string(line), err)
		}

		if err != nil {
			return r.returnLine("", fmt.Errorf("issue reading from reader\n%w", err))
		}
	}

	fmt.Printf("\n")

	return r.returnLine(string(line), nil)
}

func readEscapedKey(buffer []byte) (common.KeyCode, error) {
	i := 1
	currentLeaf := common.EscapeTreeRoot.NextByte[buffer[i]]

	for currentLeaf != nil && currentLeaf.Key == "" && buffer[i] != 0x00 {
		i++
		currentLeaf = currentLeaf.NextByte[buffer[i]]
	}

	if currentLeaf == nil {
		return common.NUL, fmt.Errorf("unexpected key")
	}

	if currentLeaf.Key == "" {
		return common.NUL, fmt.Errorf("unexpected key")
	}

	return common.KeyCode(currentLeaf.Key), nil
}

func (r *Reader) returnLine(line string, err error) (string, error) {
	if line == "" {
		return line, err
	}

	newEntry := &historyEntry{}

	if r.lastEntry != nil {
		r.lastEntry.Next = newEntry
		newEntry.Previous = r.lastEntry
	}

	r.lastEntry = newEntry

	newEntry.Line = line

	return line, err
}
