package main

import (
	"github.com/BurntSushi/toml"

	"bytes"
	"io/ioutil"
)

type Config struct {
	Baseurl string

	Mailname string
	Mailfrom string
	Smtpuser string
	Smtppass string
	Smtpserv string
}

var (
	cfg Config
)

func TomlDecodeFile(fpath string, v interface{}) (toml.MetaData, error) {
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		return toml.MetaData{}, err
	}
	// skip BOM (berucht op Windows)
	if bytes.HasPrefix(bs, []byte{239, 187, 191}) {
		bs = bs[3:]
	}
	return toml.Decode(string(bs), v)
}
