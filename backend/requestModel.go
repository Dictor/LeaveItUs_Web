package main

type (
	TagDeleteRequest struct {
		UIDs []string `json:"uids" validate:"required,printascii"`
	}
)
