package main

import (
	"time"

	"gorm.io/gorm"
)

type (
	// Tag is definition of NFC tag which is attached on behind of device like smartphone
	Tag struct {
		UID        string `json:"uid" validate:"required,printascii" gorm:"primaryKey,unique"` // tag's unique id
		ID         string `json:"id" validate:"required,printascii"`                           // tag's managing id
		AssigneeID string `json:"assignee_id" validate:"required,printascii"`                  // id of person who assigned this tag (same with device owner)
		DeviceID   string `json:"device_id" validate:"required,printascii"`                    // id of device which this tag is attached
		LockerUID  string `json:"locker_uid" validate:"required,printascii"`
		gorm.Model        // model for managing record's crud datetime
	}

	// Locker is definition of Locker device which store devices with tag.
	Locker struct {
		UID        string         `json:"uid" validate:"required,printascii" gorm:"primaryKey,unique"` // locker's unique id
		ID         string         `json:"id" validate:"required,printascii"`                           // locker's managing id
		RoomID     string         `json:"room_id" validate:"required,printascii"`                      // id of room where locker exist in
		Security   LockerSecurity `json:"-"`                                                           // security data
		Status     LockerStatus   `json:"status"`                                  // locker's status
		Tags       []Tag          `json:"tags" validate:"required" gorm:"foreignkey:LockerUID"`        // Slice of tags which are stored in locker
		gorm.Model                // model for managing record's crud datetime
	}

	// LockerSecurity is security data like credential, certificate for communicating
	LockerSecurity struct {
		Secret string
	}

	// LockerStatus is status information of locker
	LockerStatus struct {
		LastRecieveDate time.Time `json:"last_recieve_date"`
	}

	// LockerDoorEvent is definition of locker door closing event. Detail is in API doc.
	LockerDoorEvent struct {
		ClosedTime time.Time     `json:"close_time" validate:"numeric"`
		Duration   time.Duration `json:"duration" validate:"numeric"`
		gorm.Model               // model for managing record's crud datetime
	}

	// LockerPhoneRecord is log record from locker's storing information.  Detail is in API doc.
	LockerPhoneRecord struct {
		LockerUID  string    `json:"locker_id"`
		TagUIDs    *[]string `json:"tag_uids"`
		Weight     float32   `json:"weight"`
		gorm.Model           // model for managing record's crud datetime
	}

	// Person is each human's data
	Person struct {
		ID         string `json:"id" validate:"required,printascii" gorm:"primaryKey,unique"` // person id, it's used to tag's assignee id
		Name       string `json:"name" validate:"required"`
		Department string `json:"department" validate:"required"`
		RoomID     string `json:"room_id" validate:"required,printascii"`
	}

	// Room is room data. One room has one locker.
	Room struct {
		ID      string `json:"id" validate:"required,printascii" gorm:"primaryKey,unique"` // room id, it's used to locker's room id
		Name    string `json:"name" validate:"required"`
		Persons Person `json:"persons" gorm:"foreignkey:RoomID"`
	}
)

// Tag model's CrudModel interface implementing
func (Tag) InstancePointer() interface{} {
	return &Tag{}
}

func (Tag) SlicePointer() interface{} {
	return &[]Tag{}
}

func (Tag) DeleteInstancePointer() interface{} {
	return &TagDeleteRequest{}
}

func (Tag) DeleteKey(i interface{}) interface{} {
	return i.(*TagDeleteRequest).UIDs
}

// Person model's CrudModel interface implementing
func (Person) InstancePointer() interface{} {
	return &Person{}
}

func (Person) SlicePointer() interface{} {
	return &[]Person{}
}

func (Person) DeleteInstancePointer() interface{} {
	return &PersonDeleteRequest{}
}

func (Person) DeleteKey(i interface{}) interface{} {
	return i.(*PersonDeleteRequest).IDs
}
