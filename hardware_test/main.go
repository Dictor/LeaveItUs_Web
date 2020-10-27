package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	resty "github.com/go-resty/resty/v2"
)

const (
	host = "http://gcp.kimdictor.kr"
)

var (
	lockerUIDs []string = []string{"l1", "l2", "l3"}
)

func main() {
	rand.Seed(time.Now().Unix())
	for {
		lockerUID := lockerUIDs[rand.Intn(len(lockerUIDs))]
		makeDoorRequest(lockerUID, int(time.Now().Unix()), rand.Intn(100000))
		makeTagRequest(lockerUID, rand.Intn(100000))
		time.Sleep(time.Second * 3)
	}
}

func makeDoorRequest(lockerUID string, closedTime int, duration int) {
	client := resty.New()
	resp, err := client.R().
		SetBody(map[string]interface{}{"close_time": closedTime, "duration": duration}).
		Post(host + "/api/locker/" + lockerUID + "/door")
	if err != nil {
		log.Printf("[makeDoorRequest] %s", err)
	}
	log.Printf("[makeDoorRequest] %d", resp.StatusCode())
}

func makeTagRequest(lockerUID string, weight int) {
	client := resty.New()
	resp, _ := client.R().Get(host + "/api/locker/" + lockerUID + "/tag")
	tags := []interface{}{}
	json.Unmarshal(resp.Body(), &tags)

	taguids := []string{}
	for _, t := range tags {
		pt := t.(map[string]interface{})
		taguids = append(taguids, pt["uid"].(string))
	}
	log.Printf("[makeTagRequest] %d, get %d tags", resp.StatusCode(), len(taguids))

	if len(taguids) == 0 {
		log.Printf("[makeTagRequest] no tags on locker %s", lockerUID)
		return
	}
	resp, err := client.R().
		SetBody(map[string]interface{}{"tag_uids": taguids[0:rand.Intn(len(taguids)+1)], "weight": weight}).
		Post(host + "/api/locker/" + lockerUID + "/tag")
	if err != nil {
		log.Printf("[makeTagRequest] %s", err)
	}
	log.Printf("[makeTagRequest] %d", resp.StatusCode())
}
