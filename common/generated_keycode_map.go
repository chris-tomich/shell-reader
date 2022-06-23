package common

const (
	KeyEsc KeyCode = "Esc"
	KeyF1 KeyCode = "F1"
	KeyF2 KeyCode = "F2"
	KeyF3 KeyCode = "F3"
	KeyF4 KeyCode = "F4"
	KeyF5 KeyCode = "F5"
	KeyF6 KeyCode = "F6"
	KeyF7 KeyCode = "F7"
	KeyF8 KeyCode = "F8"
	KeyF9 KeyCode = "F9"
	KeyF10 KeyCode = "F10"
	KeyF11 KeyCode = "F11"
	KeyF12 KeyCode = "F12"
	KeyBacktick KeyCode = "Backtick"
	KeyNumber1 KeyCode = "Number1"
	KeyNumber2 KeyCode = "Number2"
	KeyNumber3 KeyCode = "Number3"
	KeyNumber4 KeyCode = "Number4"
	KeyNumber5 KeyCode = "Number5"
	KeyNumber6 KeyCode = "Number6"
	KeyNumber7 KeyCode = "Number7"
	KeyNumber8 KeyCode = "Number8"
	KeyNumber9 KeyCode = "Number9"
	KeyNumber0 KeyCode = "Number0"
	KeyMinus KeyCode = "Minus"
	KeyEquals KeyCode = "Equals"
	KeyBackspace KeyCode = "Backspace"
	KeyTilde KeyCode = "Tilde"
	KeyShift1 KeyCode = "Shift1"
	KeyShift2 KeyCode = "Shift2"
	KeyShift3 KeyCode = "Shift3"
	KeyShift4 KeyCode = "Shift4"
	KeyShift5 KeyCode = "Shift5"
	KeyShift6 KeyCode = "Shift6"
	KeyShift7 KeyCode = "Shift7"
	KeyShift8 KeyCode = "Shift8"
	KeyShift9 KeyCode = "Shift9"
	KeyShift0 KeyCode = "Shift0"
	KeyUnderscore KeyCode = "Underscore"
	KeyPlus KeyCode = "Plus"
	KeyTab KeyCode = "Tab"
	Keyq KeyCode = "q"
	Keyw KeyCode = "w"
	Keye KeyCode = "e"
	Keyr KeyCode = "r"
	Keyt KeyCode = "t"
	Keyy KeyCode = "y"
	Keyu KeyCode = "u"
	Keyi KeyCode = "i"
	Keyo KeyCode = "o"
	Keyp KeyCode = "p"
	KeyLeftSquareBracket KeyCode = "LeftSquareBracket"
	KeyRightSquareBracket KeyCode = "RightSquareBracket"
	KeyBackslash KeyCode = "Backslash"
	KeyQ KeyCode = "Q"
	KeyW KeyCode = "W"
	KeyE KeyCode = "E"
	KeyR KeyCode = "R"
	KeyT KeyCode = "T"
	KeyY KeyCode = "Y"
	KeyU KeyCode = "U"
	KeyI KeyCode = "I"
	KeyO KeyCode = "O"
	KeyP KeyCode = "P"
	KeyLeftCurlyBracket KeyCode = "LeftCurlyBracket"
	KeyRightCurlyBracket KeyCode = "RightCurlyBracket"
	KeyPipe KeyCode = "Pipe"
	Keya KeyCode = "a"
	Keys KeyCode = "s"
	Keyd KeyCode = "d"
	Keyf KeyCode = "f"
	Keyg KeyCode = "g"
	Keyh KeyCode = "h"
	Keyj KeyCode = "j"
	Keyk KeyCode = "k"
	Keyl KeyCode = "l"
	KeySemiColon KeyCode = "SemiColon"
	KeySingleQuote KeyCode = "SingleQuote"
	KeyEnter KeyCode = "Enter"
	KeyA KeyCode = "A"
	KeyS KeyCode = "S"
	KeyD KeyCode = "D"
	KeyF KeyCode = "F"
	KeyG KeyCode = "G"
	KeyH KeyCode = "H"
	KeyJ KeyCode = "J"
	KeyK KeyCode = "K"
	KeyL KeyCode = "L"
	KeyColon KeyCode = "Colon"
	KeyDoubleQuote KeyCode = "DoubleQuote"
	Keyz KeyCode = "z"
	Keyx KeyCode = "x"
	Keyc KeyCode = "c"
	Keyv KeyCode = "v"
	Keyb KeyCode = "b"
	Keyn KeyCode = "n"
	Keym KeyCode = "m"
	KeyComma KeyCode = "Comma"
	KeyPeriod KeyCode = "Period"
	KeyForwardSlash KeyCode = "ForwardSlash"
	KeyZ KeyCode = "Z"
	KeyX KeyCode = "X"
	KeyC KeyCode = "C"
	KeyV KeyCode = "V"
	KeyB KeyCode = "B"
	KeyN KeyCode = "N"
	KeyM KeyCode = "M"
	KeyLeftArrowBracket KeyCode = "LeftArrowBracket"
	KeyRightArrowBracket KeyCode = "RightArrowBracket"
	KeyQuestionMark KeyCode = "QuestionMark"
	KeySpace KeyCode = "Space"
	KeyInsert KeyCode = "Insert"
	KeyHome KeyCode = "Home"
	KeyPageUp KeyCode = "PageUp"
	KeyDelete KeyCode = "Delete"
	KeyEnd KeyCode = "End"
	KeyPageDown KeyCode = "PageDown"
	KeyArrowUp KeyCode = "ArrowUp"
	KeyArrowLeft KeyCode = "ArrowLeft"
	KeyArrowDown KeyCode = "ArrowDown"
	KeyArrowRight KeyCode = "ArrowRight"
)

var KeyCodes KeyCodesMap = KeyCodesMap{
	"Esc": []byte{0x1b},
	"F1": []byte{0x1b, 0x4f, 0x50},
	"F2": []byte{0x1b, 0x4f, 0x51},
	"F3": []byte{0x1b, 0x4f, 0x52},
	"F4": []byte{0x1b, 0x4f, 0x53},
	"F5": []byte{0x1b, 0x5b, 0x31, 0x35, 0x7e},
	"F6": []byte{0x1b, 0x5b, 0x31, 0x37, 0x7e},
	"F7": []byte{0x1b, 0x5b, 0x31, 0x38, 0x7e},
	"F8": []byte{0x1b, 0x5b, 0x31, 0x39, 0x7e},
	"F9": []byte{0x1b, 0x5b, 0x32, 0x30, 0x7e},
	"F10": []byte{0x1b, 0x5b, 0x32, 0x31, 0x7e},
	"F11": []byte{},
	"F12": []byte{0x1b, 0x5b, 0x32, 0x34, 0x7e},
	"Backtick": []byte{0x60},
	"Number1": []byte{0x31},
	"Number2": []byte{0x32},
	"Number3": []byte{0x33},
	"Number4": []byte{0x34},
	"Number5": []byte{0x35},
	"Number6": []byte{0x36},
	"Number7": []byte{0x37},
	"Number8": []byte{0x38},
	"Number9": []byte{0x39},
	"Number0": []byte{0x30},
	"Minus": []byte{0x2d},
	"Equals": []byte{0x3d},
	"Backspace": []byte{0x7f},
	"Tilde": []byte{0x7e},
	"Shift1": []byte{0x21},
	"Shift2": []byte{0x40},
	"Shift3": []byte{0x23},
	"Shift4": []byte{0x24},
	"Shift5": []byte{0x25},
	"Shift6": []byte{0x5e},
	"Shift7": []byte{0x26},
	"Shift8": []byte{0x2a},
	"Shift9": []byte{0x28},
	"Shift0": []byte{0x29},
	"Underscore": []byte{0x5f},
	"Plus": []byte{0x2b},
	"Tab": []byte{0x9},
	"q": []byte{0x71},
	"w": []byte{0x77},
	"e": []byte{0x65},
	"r": []byte{0x72},
	"t": []byte{0x74},
	"y": []byte{0x79},
	"u": []byte{0x75},
	"i": []byte{0x69},
	"o": []byte{0x6f},
	"p": []byte{0x70},
	"LeftSquareBracket": []byte{0x5b},
	"RightSquareBracket": []byte{0x5d},
	"Backslash": []byte{0x5c},
	"Q": []byte{0x51},
	"W": []byte{0x57},
	"E": []byte{0x45},
	"R": []byte{0x52},
	"T": []byte{0x54},
	"Y": []byte{0x59},
	"U": []byte{0x55},
	"I": []byte{0x49},
	"O": []byte{0x4f},
	"P": []byte{0x50},
	"LeftCurlyBracket": []byte{0x7b},
	"RightCurlyBracket": []byte{0x7d},
	"Pipe": []byte{0x7c},
	"a": []byte{0x61},
	"s": []byte{0x73},
	"d": []byte{0x64},
	"f": []byte{0x66},
	"g": []byte{0x67},
	"h": []byte{0x68},
	"j": []byte{0x6a},
	"k": []byte{0x6b},
	"l": []byte{0x6c},
	"SemiColon": []byte{0x3b},
	"SingleQuote": []byte{0x27},
	"Enter": []byte{0xd},
	"A": []byte{0x41},
	"S": []byte{0x53},
	"D": []byte{0x44},
	"F": []byte{0x46},
	"G": []byte{0x47},
	"H": []byte{0x48},
	"J": []byte{0x4a},
	"K": []byte{0x4b},
	"L": []byte{0x4c},
	"Colon": []byte{0x3a},
	"DoubleQuote": []byte{0x22},
	"z": []byte{0x7a},
	"x": []byte{0x78},
	"c": []byte{0x63},
	"v": []byte{0x76},
	"b": []byte{0x62},
	"n": []byte{0x6e},
	"m": []byte{0x6d},
	"Comma": []byte{0x2c},
	"Period": []byte{0x2e},
	"ForwardSlash": []byte{0x2f},
	"Z": []byte{0x5a},
	"X": []byte{0x58},
	"C": []byte{0x43},
	"V": []byte{0x56},
	"B": []byte{0x42},
	"N": []byte{0x4e},
	"M": []byte{0x4d},
	"LeftArrowBracket": []byte{0x3c},
	"RightArrowBracket": []byte{0x3e},
	"QuestionMark": []byte{0x3f},
	"Space": []byte{0x20},
	"Insert": []byte{0x1b, 0x5b, 0x32, 0x7e},
	"Home": []byte{0x1b, 0x5b, 0x48},
	"PageUp": []byte{0x1b, 0x5b, 0x35, 0x7e},
	"Delete": []byte{0x1b, 0x5b, 0x33, 0x7e},
	"End": []byte{0x1b, 0x5b, 0x46},
	"PageDown": []byte{0x1b, 0x5b, 0x36, 0x7e},
	"ArrowUp": []byte{0x1b, 0x5b, 0x41},
	"ArrowLeft": []byte{0x1b, 0x5b, 0x44},
	"ArrowDown": []byte{0x1b, 0x5b, 0x42},
	"ArrowRight": []byte{0x1b, 0x5b, 0x43},
}
