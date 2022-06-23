package vt

import "fmt"

const (
	esc = 0x1b
)

func MoveBackward(distance int) string {
	return fmt.Sprintf("%v%d%c", string([]byte{esc, '['}), distance, 'D')
}

func MoveForward(distance int) string {
	return fmt.Sprintf("%v%d%c", string([]byte{esc, '['}), distance, 'C')
}

func CursorBackward() string {
	return string([]byte{esc, '[', '0', 'D'})
}

func CursorDown() string {
	return string([]byte{esc, '[', '0', 'B'})
}

func CursorForward() string {
	return string([]byte{esc, '[', '0', 'C'})
}

func CursorUp() string {
	return string([]byte{esc, '[', '0', 'A'})
}

func Backspace() string {
	return string([]byte{0x08, ' ', 0x08})
}

func SaveCursorPosition() string {
	return string([]byte{esc, '7'})
}

func RestoreCursorPosition() string {
	return string([]byte{esc, '8'})
}

func EraseToEndOfLine() string {
	return string([]byte{esc, '[', 'K'})
}

func EraseEntireLine() string {
	return string([]byte{esc, '[', '2', 'K'})
}
