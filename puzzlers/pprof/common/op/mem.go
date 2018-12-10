package op

import (
	"bytes"
	"encoding/json"
	"math/rand"
)

type box struct {
	Str   string
	Code  rune
	Bytes []byte
}

func MemProfile() error {
	max := 5000
	var buf bytes.Buffer
	for j := 0; j < max; j++ {
		seed := rand.Intn(95) + 32
		one := createBox(seed)
		str, e := jsonStr(one)
		if e != nil {
			return e
		}
		buf.Write(str)
		buf.WriteByte('\t')
	}
	_ = buf.String()
	return nil
}

func createBox(seed int) box {
	if seed <= 0 {
		seed = 1
	}
	var array []byte
	size := seed * 8
	for i := 0; i < 8; i++ {
		array = append(array, byte(size))
	}
	return box{
		Str:   string(seed),
		Code:  rune(seed),
		Bytes: array,
	}
}

func jsonStr(one box) ([]byte, error) {
	return json.Marshal(one)
}
