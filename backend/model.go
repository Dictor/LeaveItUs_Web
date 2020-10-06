package main

import (
	"gorm.io/gorm"
)

type (
	// Tag is definition of NFC tag which is attached on behind of device like smartphone
	Tag struct {
		UID        string // tag's unique id
		ID         string // tag's managing id
		AssigneeID string // id of person who assigned this tag (same with device owner)
		DeviceID   string // id of device which this tag is attached
		gorm.Model        // model for managing record's crud datetime
	}

	// Locker is definition of Locker device which store devices with tag.
	Locker struct {
		UID        string         // locker's unique id
		ID         string         // locker's managing id
		RoomID     string         // id of room where locker exist in
		Security   LockerSecurity // security data
		Status     LockerStatus   // locker's status
		Tags       *[]Tag         // Slice of tags which are stored in locker
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
		ID         string // person id, it's used to tag's assignee id
		Name       string
		Department string
	}

	// Room is room data. One room has one locker.
	Room struct {
		ID      string // room id, it's used to locker's room id
		Name    string
		Persons []Person // slice of persons belong in this room
	}
)
