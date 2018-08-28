package gotg

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	resultSuccess = "SUCCESS"
)

// checkSuccess return an error if the return buffer is not according to pytg expected value:
// {"result": "SUCCESS"}
func checkSuccess(buf []byte) error {
	var result Result
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return err
	}

	if result.Result != resultSuccess {
		return errors.New(fmt.Sprintf("Unsuccessful result: %s", buf))
	}

	return nil
}
