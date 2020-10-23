<template>
    <v-container fluid>
        <toast :msgs="msgs"/>
        <data-table
            title="라커"
            :columns="columns"
            :rows="lockers"
            :dialog="dialog"
            ukey="id"
            @create="addLocker"
            @read="listLocker"
            @update="updateLocker"
            @delete="deleteLocker"
        >
        </data-table>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';
import DataTable from '../components/DataTable.vue';

export default {
    name: "Locker",
    components: {
        Toast,
        DataTable,
    },
    data: () => ({
        columns: [
          {text: '하드웨어 ID', value: 'uid'},
          {text: '관리 ID', value: 'id'},
          {text: '생활관 ID', value: 'room_id'},
          {text: '할당된 태그', value: 'tags'},
        ],
        lockers: [],
        dialog: [
            [
                {key: "uid", label: "하드웨어 ID", required: true, update_disable: true},
                {key: "id", label: "관리 ID", required: true},
            ],
            [
                {key: "room_id", label: "생활관 ID", required: true},
                {key: "tags", label: "할당된 태그", required: true},,
            ]
        ],
        msgs: [],
    }),
    methods: {
        addLocker: function(dialogData) {
            axios.post("./api/locker", dialogData)
                .then(() => {
                    this.msgs.push({msg: "추가 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listLocker(300);
        },
        listLocker: function(delay) {
            const req = () => {
                axios.get("./api/locker")
                    .then((res) => {
                        this.persons = res.data;
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
        updateLocker: function(dialogData) {
             axios.put("./api/locker", dialogData)
                .then(() => {
                    this.msgs.push({msg: "갱신 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listLocker(300);
        },
        deleteLocker: function(selected) {
            const t = selected.map(x => x.id);
            if (t.length < 1) {
                this.msgs.push({msg: "삭제할 라커를 지정하세요.", kind: "error"});
                return;
            }
            //in axios 0.20, axios.delete not working with data option
            axios.request({data: {id: t}, url: "./api/locker", method: 'delete'})
                .then(() => {
                    this.msgs.push({msg: "삭제 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listLocker(300);
        },
        handleError: function(error) {
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
    }
}
</script>