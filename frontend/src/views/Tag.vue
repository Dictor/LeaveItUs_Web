<template>
    <v-container fluid>
        <toast :msgs="msgs"/>
        <data-table
            title="태그"
            :columns="columns"
            :rows="tags"
            :dialog="dialog"
            ukey="uid"
            @create="addTag"
            @read="listTag"
            @update="updateTag"
            @delete="deleteTag"
        >
        </data-table>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';
import DataTable from '../components/DataTable.vue';

export default {
    name: "Tag",
    data: () => ({
        search: '',
        columns: [
          { text: '고유 ID', value: 'uid' },
          { text: '관리 ID', value: 'id' },
          { text: '사용자 ID', value: 'assignee_id' },
          { text: '장치 ID', value: 'device_id' },
          { text: '라커 ID', value: 'locker_uid' },
        ],
        tags: [],
        dialog: [
            [
                {key: "uid", label: "고유 ID", required: true, update_disable: true},
                {key: "id", label: "관리 ID", required: true},
            ],
            [
                {key: "assignee_id", label: "사용자 ID", required: true},
                {key: "device_id", label: "장치 ID", required: true},
                {key: "locker_uid", label: "라커 ID", required: true},
            ]
        ],
        msgs: [],
      }),
    components: {
        Toast,
        DataTable,
    },
    methods: {
        addTag: function(dialogData) {
            axios.post("./api/tag", dialogData)
                .then(() => {
                    this.msgs.push({msg: "추가 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listTag(300);
        },
        listTag: function(delay) {
            const req = () => {
                axios.get("./api/tag")
                    .then((res) => {
                        this.tags = res.data;
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
        updateTag: function(dialogData) {
             axios.put("./api/tag", dialogData)
                .then(() => {
                    this.msgs.push({msg: "갱신 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listTag(300);
        },
        deleteTag: function(selected) {
            const t = selected.map(x => x.id);
            if (t.length < 1) {
                this.msgs.push({msg: "삭제할 태그를 지정하세요.", kind: "error"});
                return;
            }
            //in axios 0.20, axios.delete not working with data option
            axios.request({data: {id: t}, url: "./api/tag", method: 'delete'})
                .then(() => {
                    this.msgs.push({msg: "삭제 성공!", kind: "success"});
                })
                .catch(this.handleError);
            this.listTag(300);
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
        mounted() { 
            this.listTag();
        }
    },
    mounted() {
        this.listTag();
    }
}
</script>

<style scoped>
</style>