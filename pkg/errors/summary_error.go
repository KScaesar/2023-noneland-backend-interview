package errors

import "net/http"

var (
	ErrInvalidToken = NewCustomError(
		"invalid token",
		700,
		http.StatusUnauthorized,
	)
	ErrTokenExpired = NewCustomError(
		"token expired",
		710,
		http.StatusUnauthorized,
	)

	ErrInvalidParams = NewCustomError(
		"invalid params",
		800,
		http.StatusBadRequest,
	)
	ErrNotFound = NewCustomError(
		"not found",
		810,
		http.StatusNotFound,
	)
	ErrKeyDuplicated = NewCustomError(
		"key or name is duplicated",
		820,
		http.StatusConflict,
	)

	ErrSystem = NewCustomError(
		"system fail",
		900,
		http.StatusInternalServerError,
	)
	ErrTimeout = NewCustomError(
		"system time out",
		910,
		http.StatusInternalServerError,
	)

	ErrUnknown3rdParty = NewCustomError(
		"unknown third party",
		-1,
		http.StatusInternalServerError,
	)
)
