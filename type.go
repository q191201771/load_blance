package load_blance

import "errors"

var (
	ErrNoNode           = errors.New("no node exist")
	ErrNodeAlreadyExist = errors.New("node already exist")
	ErrNodeNotFound     = errors.New("node not found")
)
