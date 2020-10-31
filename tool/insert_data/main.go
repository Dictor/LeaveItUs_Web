package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	resty "github.com/go-resty/resty/v2"
)

type (
	// Tag is definition of NFC tag which is attached on behind of device like smartphone
	Tag struct {
		UID        string `json:"uid" validate:"required,printascii" gorm:"primaryKey,unique"` // tag's unique id
		ID         string `json:"id" validate:"required,printascii"`                           // tag's managing id
		AssigneeID string `json:"assignee_id" validate:"required,printascii"`                  // id of person who assigned this tag (same with device owner)
		DeviceID   string `json:"device_id" validate:"required,printascii"`                    // id of device which this tag is attached
		LockerUID  string `json:"locker_uid" validate:"required,printascii"`
	}

	// Locker is definition of Locker device which store devices with tag.
	Locker struct {
		UID    string `json:"uid" validate:"required,printascii" gorm:"primaryKey,unique"` // locker's unique id
		ID     string `json:"id" validate:"required,printascii"`                           // locker's managing id
		RoomID string `json:"room_id" validate:"required,printascii"`                      // id of room where locker exist in
		Tags   []Tag  `json:"tags" gorm:"foreignkey:LockerUID;references:UID"`             // Slice of tags which are stored in locker
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

var (
	roomAmount int = 15
	tagAmount  int = 150
	host           = "http://gcp.kimdictor.kr"
)

func main() {
	tStart := int(time.Now().UnixNano())
	for i := 1; i <= roomAmount; i++ {
		request("/api/locker", makeLocker(i))
		//request("/api/room", makeRoom(i))
	}
	for i := 1; i <= tagAmount; i++ {
		request("/api/tag", makeTag(i, roomAmount))
		request("/api/person", makePerson(i, roomAmount))
	}
	tEnd := int(time.Now().UnixNano())
	var eTime float32 = float32(tEnd-tStart) / float32(1000000000)
	var reqCount int = roomAmount + tagAmount*2
	fmt.Printf("========\n%d requests are done\ntotal elapsed time is %f second.\n%f req/s\n========\n",
		reqCount,
		eTime,
		float32(reqCount)/eTime)
}

func request(path string, s interface{}) {
	client := resty.New()
	resp, err := client.R().
		SetBody(s).
		Post(host + path)
	if err != nil {
		log.Printf("[%s] %s", path, err)
	}
	log.Printf("[%s] %d", path, resp.StatusCode())
}

func makeLocker(number int) *Locker {
	return &Locker{
		UID:    fmt.Sprintf("l%d", number),
		ID:     fmt.Sprintf("lm%d", number),
		RoomID: fmt.Sprintf("l%d", number),
	}
}

func makeRoom(number int) *Room {
	return &Room{
		ID:   fmt.Sprintf("r%d", number),
		Name: fmt.Sprintf("%d생활관", number),
	}
}

func makeTag(number int, lockerAmount int) *Tag {
	return &Tag{
		UID:        fmt.Sprintf("t%d", number),
		ID:         fmt.Sprintf("tm%d", number),
		AssigneeID: fmt.Sprintf("p%d", number),
		DeviceID:   fmt.Sprintf("d%d", number),
		LockerUID:  fmt.Sprintf("l%d", (number%lockerAmount)+1),
	}
}

func makePerson(number int, roomAmount int) *Person {
	return &Person{
		ID:         fmt.Sprintf("p%d", number),
		Name:       faker.Name(),
		Department: faker.CCType(),
		RoomID:     fmt.Sprintf("r%d", (number%roomAmount)+1),
	}
}
