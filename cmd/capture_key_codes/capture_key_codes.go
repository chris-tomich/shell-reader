package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/chris-tomich/shell-reader/common"
	"github.com/pkg/term"
)

var keys = []string{
	"Esc",
	"F1",
	"F2",
	"F3",
	"F4",
	"F5",
	"F6",
	"F7",
	"F8",
	"F9",
	"F10",
	"F11",
	"F12",
	"Backtick",
	"Number1",
	"Number2",
	"Number3",
	"Number4",
	"Number5",
	"Number6",
	"Number7",
	"Number8",
	"Number9",
	"Number0",
	"Minus",
	"Equals",
	"Backspace",
	"Tilde",
	"Shift1",
	"Shift2",
	"Shift3",
	"Shift4",
	"Shift5",
	"Shift6",
	"Shift7",
	"Shift8",
	"Shift9",
	"Shift0",
	"Underscore",
	"Plus",
	"Tab",
	"q",
	"w",
	"e",
	"r",
	"t",
	"y",
	"u",
	"i",
	"o",
	"p",
	"LeftSquareBracket",
	"RightSquareBracket",
	"Backslash",
	"Q",
	"W",
	"E",
	"R",
	"T",
	"Y",
	"U",
	"I",
	"O",
	"P",
	"LeftCurlyBracket",
	"RightCurlyBracket",
	"Pipe",
	"a",
	"s",
	"d",
	"f",
	"g",
	"h",
	"j",
	"k",
	"l",
	"SemiColon",
	"SingleQuote",
	"Enter",
	"A",
	"S",
	"D",
	"F",
	"G",
	"H",
	"J",
	"K",
	"L",
	"Colon",
	"DoubleQuote",
	"z",
	"x",
	"c",
	"v",
	"b",
	"n",
	"m",
	"Comma",
	"Period",
	"ForwardSlash",
	"Z",
	"X",
	"C",
	"V",
	"B",
	"N",
	"M",
	"LeftArrowBracket",
	"RightArrowBracket",
	"QuestionMark",
	"Space",
	"Insert",
	"Home",
	"PageUp",
	"Delete",
	"End",
	"PageDown",
	"ArrowUp",
	"ArrowLeft",
	"ArrowDown",
	"ArrowRight",
}

func convertToHexBytes(buffer []byte) string {
	byteArray := "[]byte {"

	for i, b := range buffer {
		if i == len(buffer)-1 {
			byteArray = byteArray + fmt.Sprintf("0x%x", b)
		} else {
			byteArray = byteArray + fmt.Sprintf("0x%x, ", b)
		}
	}

	byteArray = byteArray + "}"

	return byteArray
}

func convertToHexStrings(buffer []byte) []string {
	byteArray := make([]string, 0, len(buffer))

	for _, b := range buffer {
		byteArray = append(byteArray, fmt.Sprintf("0x%x", b))
	}

	return byteArray
}

func main() {
	terminalController, err := term.Open("/dev/tty")

	if err != nil {
		panic(fmt.Errorf("issue creating terminal controls\n%w", err))
	}

	err = term.RawMode(terminalController)

	if err != nil {
		panic(fmt.Errorf("issue entering raw mode\n%w", err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		terminalController.Restore()
		fmt.Println()
		os.Exit(0)
	}()

	defer func() {
		terminalController.Restore()
		fmt.Println()
	}()

	var lastBuffer = make([]byte, 0, 10)
	var buffer []byte = make([]byte, 10)
	var lastk string

	keyMap := make(common.KeyCodeList, 0, 200)

	for _, k := range keys {
		if err != io.EOF {
			if len(lastBuffer) == 1 && lastBuffer[0] == 'q' && lastk != "q" {
				break
			}

			fmt.Printf("Press %v: ", k)

			_, err = os.Stdin.Read(buffer)

			lastBuffer = lastBuffer[:0]

			for i, c := range buffer {
				if c == 0 {
					break
				}

				lastBuffer = append(lastBuffer, c)
				buffer[i] = 0
			}

			fmt.Printf("%c%c", 0x1b, 'E')

			if lastBuffer[0] == 'w' && k[0] != 'w' {
				keyMap = append(keyMap, common.KeyCodePair{
					Key:   k,
					Bytes: make([]string, 0),
				})
			} else {
				keyMap = append(keyMap, common.KeyCodePair{
					Key:   k,
					Bytes: convertToHexStrings(lastBuffer),
				})
			}

			lastk = k
		}
	}

	bytes, err := json.Marshal(keyMap)

	if err != nil {
		panic(fmt.Errorf("issue marshalling key map array to JSON\n%w", err))
	}

	err = os.WriteFile("../../common/keymap.json", bytes, 0664)

	if err != nil {
		panic(fmt.Errorf("issue writing the key map to a JSON file\n%w", err))
	}
}
