package yach

import "errors"

var ErrNoBindingFound = errors.New("no binding exists for the specified key")
var ErrNoValueFound = errors.New("no value was found for the specified key")
