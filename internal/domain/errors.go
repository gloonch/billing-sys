package domain

import "errors"

var ErrUnitNotFound = errors.New("unit not found")
var ErrInsufficientData = errors.New("insufficient data for calculating charge")
