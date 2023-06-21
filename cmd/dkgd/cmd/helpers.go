package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "strconv"
	// "strings"
	//"github.com/fairblock/dkg-core/utils"
)

func getByteCodes(file string) ([]byte, error) {
	jsonStr, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	table := make(map[string]interface{})
	err = json.Unmarshal(jsonStr, &table)
	if err != nil {
		return nil, err
	}

	str, ok := table["bytecode"].(string)
	if !ok {
		return nil, fmt.Errorf("could not retrieve bytecode from file")
	}

	return hex.DecodeString(str)
}
