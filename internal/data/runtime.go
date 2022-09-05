package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// implementing MarshalJSON function for the type Runtime so that
// when calling json.Marshal(), the return will be of the right format
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// before returning the value, we must surround it by double quotes
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}
