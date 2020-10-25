<template>
    <v-container fluid>
        <toast :msgs="msgs"/>
        <data-table
            title="개폐 이벤트"
            :columns="columns"
            :rows="events"
            :dialog="dialog"
            disable_action="cud"
            @read="listEvent"
        >
        </data-table>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';
import DataTable from '../components/DataTable.vue';

export default {
    name: "DoorEvent",
    components: {
        Toast,
        DataTable,
    },
    data: () => ({
        columns: [
          {text: '라커 UID', value: 'locker_uid'},
          {text: '폐문 시간', value: 'close_time'},
          {text: '개문 기간', value: 'duration'},
        ],
        events: [],
        dialog: [],
        msgs: [],
    }),
    methods: {
        listEvent: function(delay) {
            const req = () => {
                axios.get("./api/locker/door")
                    .then((res) => {      
                        console.log(res.data);
                        this.events = res.data;
                        this.msgs.push({msg: "불러오기 성공!", kind: "success"});
                    })
                    .catch(this.handleError);
                }
            if (!delay || delay === 0) {
                req();
            } else {
                setTimeout(req, delay);
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
        this.listEvent();
    }
}
</script>