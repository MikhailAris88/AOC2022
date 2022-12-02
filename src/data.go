package src

import (
	"github.com/echojc/aocutil"
)

func GetData() (*aocutil.Input, error) {
	data, err := aocutil.NewInputFromFile(".session_id")
	if err != nil {
		return nil, err
	}
	return data, err
}
