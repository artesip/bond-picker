package domain

import "errors"

var (
	BadCredentialsErr = errors.New("invalid credentials")
	ConflictErr       = errors.New("conflict")
	ValidationErr     = errors.New("validation error")
)
