package config

import (
	"encoding/json"
	"io/ioutil"
)

type unmarshaller struct {
	err error

	bytes []byte
}

func (u *unmarshaller) readFile(path string) {

	bs, be := ioutil.ReadFile(path)

	u.err = be

	if u.err == nil {

		u.bytes = bs
	}
}

func (u *unmarshaller) unmarshal(target interface{}) {

	if u.err == nil {

		err := json.Unmarshal(u.bytes, target)
		u.err = err
	}
}
