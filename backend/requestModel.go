package main

type (
	// TagDeleteRequest is definition of DELETE tag request's body
	TagDeleteRequest struct {
		UIDs []string `json:"uids" validate:"required,printascii"`
	}

	// PersonDeleteRequest is definition of DELETE tag request's body
	PersonDeleteRequest struct {
		IDs []string `json:"ids" validate:"required,printascii"`
	}

	// LockerDeleteRequest is definition of DELETE locker request's body
	LockerDeleteRequest struct {
		UIDs []string `json:"uids" validate:"required,printascii"`
	}

	// LockerRecordCreateRequest is definition of POST locker recored request's body
	LockerRecordCreateRequest struct {
		TagUIDs *[]string `json:"tag_uids" validate:"required"`
		Weight  float32   `json:"weight" validate:"required,numeric"`
	}

	// LockerDoorEventCreateRequest is definition of POST locker door event request's body
	LockerDoorEventCreateRequest struct {
		ClosedTime int `json:"close_time" validate:"numeric"`
		Duration   int `json:"duration" validate:"numeric"`
	}
)
