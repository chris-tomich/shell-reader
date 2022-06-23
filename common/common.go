package common

type KeyCodeList []KeyCodePair

type KeyCodePair struct {
	Key   string
	Bytes []string
}

type EscapeLeaf struct {
	ByteValue byte
	Key       string
	NextByte  map[byte]*EscapeLeaf
}

func NewEscapeRoot() *EscapeLeaf {
	return &EscapeLeaf{
		ByteValue: 0x1b,
		Key:       "",
		NextByte:  make(map[byte]*EscapeLeaf),
	}
}

func NewEscapeLeaf(byteValue byte) *EscapeLeaf {
	return &EscapeLeaf{
		ByteValue: byteValue,
		Key:       "",
		NextByte:  make(map[byte]*EscapeLeaf),
	}
}

type KeyCode string

const NUL KeyCode = ""

type KeyCodesMap map[KeyCode][]byte
