<template>
    <v-container fluid>
        <toast :msgs="msgs"/>
        <data-table
            title="병력"
            :columns="columns"
            :rows="persons"
            :dialog="dialog"
            ukey="id"
            @create="addPerson"
            @read="listPerson"
            @update="updatePerson"
            @delete="deletePerson"
        >
        </data-table>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';
import DataTable from '../components/DataTable.vue';

export default {
    name: "Personnel",
    components: {
        Toast,
        DataTable,
    },
    data: () => ({
        columns: [
          { text: '군번', value: 'id'},
          { text: '이름', value: 'name'},
          { text: '부서', value: 'department'},
          { text: '생활관 ID', value: 'room_id'},
        ],
        persons: [],
        dialog: [
            [
                {key: "id", label: "군번", required: true, update_disable: true},
                {key: "name", label: "이름", required: true},
            ],
            [
                {key: "department", label: "부서", required: true},
                {key: "room_id", label: "생활관 ID", required: true},
            ]
        ],
        msgs: [],
    }),
    methods: {
        addPerson: function(dialogData) {
            axios.post("./api/person", dialogData)
                .then(() => {
                    this.msgs.push({msg: "추가 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listPerson(300);
        },
        listPerson: function(delay) {
            const req = () => {
                axios.get("./api/person")
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
        updatePerson: function(dialogData) {
             axios.put("./api/person", dialogData)
                .then(() => {
                    this.msgs.push({msg: "갱신 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listPerson(300);
        },
        deletePerson: function(selected) {
            const t = selected.map(x => x.id);
            if (t.length < 1) {
                this.msgs.push({msg: "삭제할 병력를 지정하세요.", kind: "error"});
                return;
            }
            //in axios 0.20, axios.delete not working with data option
            axios.request({data: {id: t}, url: "./api/person", method: 'delete'})
                .then(() => {
                    this.msgs.push({msg: "삭제 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listPerson(300);
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
    },
    mounted() {
        this.listPerson();
    }
}
</script>