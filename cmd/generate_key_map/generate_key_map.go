package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/chris-tomich/shell-reader/common"
)

func main() {
	data, err := os.ReadFile("../../common/keymap.json")

	if err != nil {
		panic(fmt.Errorf("issue with loading keymap.json file\n%w", err))
	}

	keyList := make(common.KeyCodeList, 0)

	err = json.Unmarshal(data, &keyList)

	if err != nil {
		panic(fmt.Errorf("issue unmarshalling from JSON\n%w", err))
	}

	packageName := "common"

	escapeRoot := common.NewEscapeRoot()
	keyCodeBuilder := NewKeyCodeGoCodeBuilder(packageName)

	for _, k := range keyList {
		keyCodeBuilder.WriteKeyCodeMapping(&k)

		numBytes := len(k.Bytes)
		if numBytes <= 1 || k.Bytes[0] != "0x1b" {
			continue
		}

		currentLeaf := escapeRoot

		for i := 1; i < numBytes; i++ {
			currentBytes, err := hex.DecodeString(strings.TrimPrefix(k.Bytes[i], "0x"))

			if err != nil {
				panic(fmt.Errorf("unexpected byte value for key, key: %v, bytes: %v\n%w", k.Key, k.Bytes[i], err))
			}

			if len(currentBytes) != 1 {
				panic(fmt.Errorf("unexpected number of bytes decoded for key, key: %v, bytes: %v, bytes decoded: %v", k.Key, k.Bytes[i], currentBytes))
			}

			currentByte := currentBytes[0]
			var nextLeaf *common.EscapeLeaf

			if currentLeaf.NextByte[currentByte] == nil {
				nextLeaf = common.NewEscapeLeaf(currentByte)
				currentLeaf.NextByte[currentByte] = nextLeaf
			} else {
				nextLeaf = currentLeaf.NextByte[currentByte]
			}

			currentLeaf = nextLeaf
		}

		currentLeaf.Key = k.Key
	}

	builder := &strings.Builder{}

	BuildEscapeTreeGoCode(packageName, builder, 0, escapeRoot)

	os.WriteFile("../../common/generated_escape_tree.go", []byte(builder.String()), 0664)
	os.WriteFile("../../common/generated_keycode_map.go", []byte(keyCodeBuilder.String()), 0664)
}

func PrintLeaf(depth int, leaf *common.EscapeLeaf) {
	if leaf.Key != "" {
		fmt.Printf("%vByte: %x, Key: %v\n", strings.Repeat("\t", depth), leaf.ByteValue, leaf.Key)
	} else {
		fmt.Printf("%vByte: %x\n", strings.Repeat("\t", depth), leaf.ByteValue)
	}

	for _, v := range leaf.NextByte {
		PrintLeaf(depth+1, v)
	}
}

func BuildEscapeTreeGoCode(packageName string, builder *strings.Builder, depth int, leaf *common.EscapeLeaf) {
	if depth == 0 {
		builder.WriteString(fmt.Sprintf("package %v\n\n", packageName))
		builder.WriteString("var EscapeTreeRoot *EscapeLeaf = &EscapeLeaf")
	}

	builder.WriteString("{\n")
	builder.WriteString(fmt.Sprintf("%vByteValue: 0x%x,\n", strings.Repeat("\t", depth+1), leaf.ByteValue))
	builder.WriteString(fmt.Sprintf("%vKey:       \"%v\",\n", strings.Repeat("\t", depth+1), leaf.Key))

	if leaf.NextByte == nil {
		builder.WriteString(fmt.Sprintf("%vNextByte:  nil,\n", strings.Repeat("\t", depth+1)))
	} else {
		builder.WriteString(fmt.Sprintf("%vNextByte:  map[byte]*EscapeLeaf{\n", strings.Repeat("\t", depth+1)))

		for k, v := range leaf.NextByte {
			builder.WriteString(fmt.Sprintf("%v0x%x: ", strings.Repeat("\t", depth+2), k))
			BuildEscapeTreeGoCode(packageName, builder, depth+2, v)
		}

		builder.WriteString(fmt.Sprintf("%v},\n", strings.Repeat("\t", depth+1)))
	}

	if depth == 0 {
		builder.WriteString("}\n")
	} else {
		builder.WriteString(fmt.Sprintf("%v},\n", strings.Repeat("\t", depth)))
	}
}

type KeyCodeGoCodeBuilder struct {
	packageName          string
	constantsListBuilder *strings.Builder
	keyCodeBytesBuilder  *strings.Builder
}

func NewKeyCodeGoCodeBuilder(packageName string) *KeyCodeGoCodeBuilder {
	builder := &KeyCodeGoCodeBuilder{
		packageName,
		&strings.Builder{},
		&strings.Builder{},
	}

	builder.constantsListBuilder.WriteString("const (\n")

	builder.keyCodeBytesBuilder.WriteString("var KeyCodes KeyCodesMap = KeyCodesMap{\n")

	return builder
}

func (b *KeyCodeGoCodeBuilder) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("package %v\n\n", b.packageName))

	builder.WriteString(b.constantsListBuilder.String())
	builder.WriteString(")\n\n")

	builder.WriteString(b.keyCodeBytesBuilder.String())
	builder.WriteString("}\n")

	return builder.String()
}

func (b *KeyCodeGoCodeBuilder) WriteKeyCodeMapping(mapping *common.KeyCodePair) {
	b.constantsListBuilder.WriteString(fmt.Sprintf("\tKey%v KeyCode = \"%v\"\n", mapping.Key, mapping.Key))

	bytesBuilder := strings.Builder{}
	bytesBuilder.WriteString("[]byte{")

	for i, hexString := range mapping.Bytes {
		bytesBuilder.WriteString(hexString)

		if i < len(mapping.Bytes)-1 {
			bytesBuilder.WriteString(", ")
		}
	}

	bytesBuilder.WriteString("}")

	b.keyCodeBytesBuilder.WriteString(fmt.Sprintf("\t\"%v\": %v,\n", mapping.Key, bytesBuilder.String()))
}
