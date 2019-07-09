package deep

import (
	"bytes"
	"encoding/gob"
)

func Copy(dst, src interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(src); err != nil {
		return err
	}
	return dec.Decode(dst)
}
