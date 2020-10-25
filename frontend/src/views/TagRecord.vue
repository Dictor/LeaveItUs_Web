<template>
    <v-container fluid>
        <toast :msgs="msgs"/>
        <data-table
            title="반납 기록"
            :columns="columns"
            :rows="records"
            :dialog="dialog"
            disable_action="cud"
            @read="listRecord"
        >
        </data-table>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';
import DataTable from '../components/DataTable.vue';

export default {
    name: "TagRecord",
    components: {
        Toast,
        DataTable,
    },
    data: () => ({
        columns: [
          {text: '라커 UID', value: 'locker_uid'},
          {text: '반납된 태그', value: 'tag_uids'},
          {text: '무게', value: 'weight'},
        ],
        records: [],
        dialog: [],
        msgs: [],
    }),
    methods: {
        listRecord: function(delay) {
            const req = () => {
                axios.get("./api/locker/tag")
                    .then((res) => {      
                        console.log(res.data);
                        this.records = res.data;
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
        this.listRecord();
    }
}
</script>