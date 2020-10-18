package main

import (
	"gorm.io/gorm"
)

type (
	// Tag is definition of NFC tag which is attached on behind of device like smartphone
	Tag struct {
		UID        string `json:"uid" validate:"required,printascii" gorm:"primaryKey,unique"` // tag's unique id
		ID         string `json:"id" validate:"required,printascii"`                           // tag's managing id
		AssigneeID string `json:"assignee_id" validate:"required,printascii"`                  // id of person who assigned this tag (same with device owner)
		DeviceID   string `json:"device_id" validate:"required,printascii"`                    // id of device which this tag is attached
		gorm.Model        // model for managing record's crud datetime
	}

	// Locker is definition of Locker device which store devices with tag.
	Locker struct {
		UID        string         `json:"uid" validate:"required,printascii" gorm:"primaryKey,unique"` // locker's unique id
		ID         string         `json:"id" validate:"required,printascii"`                           // locker's managing id
		Room       Room           `json:"room" validate:"required,printascii"`                         // id of room where locker exist in
		Security   LockerSecurity `json:"security" validate:"required"`                                // security data
		Status     LockerStatus   `json:"status" validate:"required"`                                  // locker's status
		Tags       *[]Tag         `json:"tags" validate:"required"`                                    // Slice of tags which are stored in locker
		gorm.Model                // model for managing record's crud datetime
	}

	// LockerSecurity is security data like credential, certificate for communicating
	LockerSecurity struct {
	}

	// LockerStatus is status information of locker
	LockerStatus struct {
	}

	// Person is each human's data
	Person struct {
		ID         string `json:"id" validate:"required,printascii" gorm:"primaryKey,unique"` // person id, it's used to tag's assignee id
		Name       string `json:"name" validate:"required"`
		Department string `json:"department" validate:"required"`
	}

	// Room is room data. One room has one locker.
	Room struct {
		ID      string   `json:"id" validate:"required,printascii" gorm:"primaryKey,unique"` // room id, it's used to locker's room id
		Name    string   `json:"name" validate:"required"`
		Persons []Person `json:"person" validate:"required"` // slice of persons belong in this room
	}
)
