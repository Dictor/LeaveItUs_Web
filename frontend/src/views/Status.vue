<template>
    <v-container fluid>
        <link href="https://fonts.googleapis.com/css2?family=Abel&display=swap" rel="stylesheet">
        <toast :msgs="msgs"/>
        <v-card class="pa-2 mb-2">
            <span class="ml-2">{{valueTimeAgo}} 불러옴</span>
            <v-select
                v-model="doorlimit"
                :items="doorlimitDropdown"
                item-text="label"
                item-value="value"
                solo
                style="display: inline-block; height: 50px; width: 100px;"
                class="ml-4"
            ></v-select> 이내 기록 불러옴
        </v-card>
        <v-container class="locker-wrapper" fluid> 
            <div class="locker" v-for="(locker, i) in lockers" :key="i" :class="getClass(locker)">
                <p><v-icon>mdi-briefcase</v-icon> {{locker.id}} ({{locker.room_id}} 생활관)</p>
                <p class="locker-status"><v-icon>mdi-cellphone-android</v-icon>{{(lockerTag[locker.uid]) ? unmarshal(lockerTag[locker.uid].tag_uids).length : "?"}}/{{locker.tags.length}}</p>
                <p class="locker-status"><v-icon>mdi-door-open</v-icon>{{isArray(lockerDoor[locker.uid]) ? lockerDoor[locker.uid].length : 0}}건</p>
            </div>
        </v-container>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';

export default {
    name: "Status",
    components: {
        Toast,
    },
    data: () => ({
        lockers: [],
        lockerDoor: {},
        lockerTag: {},
        msgs: [],
        valueTimeAgo: "",
        valueTime: Date.now(),
        doorlimitDropdown: [
            {label: "1분", value: 60},
            {label: "10분", value: 600},
            {label: "30분", value: 1800},
            {label: "1시간", value: 3600},
            {label: "3시간", value: 10800},
            {label: "6시간", value: 21600},
            {label: "하루", value: 86400},
            {label: "전체", value: ""},
        ],
        doorlimit: 3600
    }),
    methods: {
        isArray: function(obj) {
            return Array.isArray(obj);
        },
        getClass: function(l) {
            if (!this.lockerTag[l.uid]) {
                return ["locker-disable"];
            } 
            let total = l.tags.length;
            let current = this.unmarshal(this.lockerTag[l.uid].tag_uids).length;
            if (total === current) {
                return ["locker-good"]
            } else {
                return ["locker-warning"]
            }
        },
        unmarshal: function(d) {
            return JSON.parse(d);
        },
        getLocker: function() {
            return axios.get("./api/locker").then(response => response.data);
        },
        getTagRecord: function() {
            return axios.get("./api/locker/latest/tag").then(response => response.data);
        },
        getLockerDoorEvent: function() {
            return axios.get("./api/locker/door?limit=" + this.doorlimit).then(response => response.data);
        },
        prepareData: async function() {
            this.lockers = await this.getLocker();
            this.lockerTag = await this.getTagRecord();

            let des = await this.getLockerDoorEvent();
            this.lockerDoor = {};
            for (let de of des) {
                if (!this.lockerDoor[de.locker_uid]) {
                    this.lockerDoor[de.locker_uid] = [];
                }
                this.lockerDoor[de.locker_uid].push(de);
            }
            this.valueTime = Date.now();
        },
        handleError: function(error) {
            console.log(error);
            if (error.response) {
                if (error.response.data.msg) {
                    this.msgs.push({msg: "서버의 실패 응답 : " + error.response.data.msg + "(" + error.response.status + ")", kind: "error"});
                } else {
                    this.msgs.push({msg: error + " (" + error.response.status + ")", kind: "error"});
                }
            }
            else if (error.request) {
                this.msgs.push({msg: "응답 수신 실패", kind: "error"});
            }
            else {
                this.msgs.push({msg: "요청 실패", kind: "error"});
            }
        },
    },
    mounted() {
        setInterval(this.prepareData, 3000);
        setInterval(() => {
            this.valueTimeAgo = ((Date.now() - this.valueTime) / 1000).toFixed(1) + "초 전";
        }, 200);
    }
}
</script>

<style scoped>
    .locker {
        border: solid black 2px;
        border-radius: 10px;
        padding: 0.3rem;
        margin: 0.3rem;
    }

    .locker > p:nth-child(1) {
        font-weight: bold;
        font-size: 1.3rem;
    }

    .locker > p {
        margin-block-end: 0.2rem;
        text-align: center;
    }

    .locker-status {
        display: inline-block;
        box-sizing: border-box;
        font-size: 1.8rem;
        margin-left: 0.3rem;
        font-family: 'Abel', sans-serif;
    }

    .locker-wrapper {
        display: flex;
    }

    .locker-disable {
        background-color: gray;
    }

    .locker-warning {
        background-color: darkgoldenrod;
    }

    .locker-good {
        background-color: limegreen;
    }
</style>