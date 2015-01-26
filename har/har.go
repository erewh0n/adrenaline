package har

import (
	"encoding/json"
	"io/ioutil"
)

func FromFile(filename *string) (*HARLog, error) {
	harContent, err := ioutil.ReadFile(*filename)

	if err != nil {
		return nil, err
	}
	var har HARLog
	json.Unmarshal(harContent, &har)

	return &har, nil
}
