package yb

import "github.com/morris-nan/go-yiban/ybc"

func Client(options ...ybc.Options) (ybc.Client, error) {
	return ybc.New(options...)
}
