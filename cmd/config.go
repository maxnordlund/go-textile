package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var errMissingReplacement = errors.New("missing replacement value")
var errMissingPath = errors.New("missing key path")

func init() {
	register(&configCmd{})
}

type configCmd struct {
	Client ClientOptions `group:"Client Options"`
}

func (x *configCmd) Name() string {
	return "config"
}

func (x *configCmd) Short() string {
	return "Get and set config values"
}

func (x *configCmd) Long() string {
	return `
The config command controls configuration variables. It works much like 'git config'.
The configuration values are stored in a config file inside your Textile repository.

When changing values, valid JSON types must be used. For example, a string should be
escaped or wrapped in single quotes (e.g., \"127.0.0.1:40600\") and arrays and objects
work fine (e.g. '{"API": "127.0.0.1:40600"}') but should be wrapped in single quotes.

Examples:

Get the value of the 'Addresses.API' key (can also use '/' instead of '.'):

  $ textile config Addresses.API

Print the entire Textile config file to console:

  $ textile config

Set the value of the 'Addresses.API' key:

  $ textile config Addresses.API \"127.0.0.1:40600\"
`
}

func (x *configCmd) Execute(args []string) error {
	setApi(x.Client)

	var path string
	if len(args) > 0 {
		path = "/" + strings.Replace(args[0], ".", "/", -1)
	}
	if len(args) > 1 {
		patch := []byte(fmt.Sprintf(`[
  {"op": "replace", "path": "%s", "value": %s}
]`, path, args[1]))
		res, err := request(PATCH, "config", params{
			payload: bytes.NewBuffer(patch),
		})
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode >= 400 {
			res, err := unmarshalString(res.Body)
			if err != nil {
				return err
			}
			return errors.New(res)
		}
		return nil
	}
	var info interface{}
	res, err := executeJsonCmd(GET, "config"+path, params{}, &info)
	if err != nil {
		return err
	}
	output(res)
	return nil
}
