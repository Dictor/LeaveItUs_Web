<template>
    <v-container fluid>
        <link href="https://fonts.googleapis.com/css2?family=Abel&display=swap" rel="stylesheet">
        <toast :msgs="msgs"/>
        <v-container class="locker-wrapper" fluid> 
            <div class="locker" v-for="(locker, i) in lockers" :key="i" :class="getClass(locker)">
                <p><v-icon>mdi-briefcase</v-icon> {{locker.id}} ({{locker.room_id}} 생활관)</p>
                <p class="locker-status"><v-icon>mdi-cellphone-android</v-icon>{{(lockerTag[locker.uid]) ? unmarshal(lockerTag[locker.uid].tag_uids).length : "?"}}/{{locker.tags.length}}</p>
                <p class="locker-status"><v-icon>mdi-door-open</v-icon>{{lockerDoor[locker.uid] ? lockerDoor[locker.uid].length : 0}}건</p>
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
    }),
    methods: {
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
            return axios.get("./api/locker/door").then(response => response.data);
        },
        prepareData: async function() {
            this.lockers = await this.getLocker();
            this.lockerTag = await this.getTagRecord();

            let des = await this.getLockerDoorEvent();
            this.lockerDoor = {};
            for (let de of des) {
                console.log(de);
                if (!this.lockerDoor[de.locker_uid]) {
                    this.lockerDoor[de.locker_uid] = {};
                }
                this.lockerDoor[de.locker_uid].push(de);
            }
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
        this.prepareData();
    }
}
</script>

<style scoped>
    .locker {
        border: solid black 2px;
        border-radius: 10px;
        padding: 0.3rem;
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