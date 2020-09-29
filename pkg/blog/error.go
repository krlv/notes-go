package blog

import "errors"

// ErrNotFound is returned when page object not found
var ErrNotFound = errors.New("page not found")
