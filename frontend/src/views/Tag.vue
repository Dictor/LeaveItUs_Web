<template>
    <v-container fluid>
        <toast :msgs="msgs"/>
        <v-card class="pa-2 mb-2">
            <v-btn key="" color="green" v-on:click="showAddDialog" dark large class="mr-1"><v-icon>mdi-plus</v-icon>태그 추가</v-btn>
            <v-btn key="" color="error" v-on:click="deleteTag(); listTag(300);" large class="mr-1"><v-icon>mdi-delete</v-icon>태그 삭제</v-btn>
            <v-btn key="" color="blue"  v-on:click="listTag" large><v-icon>mdi-refresh</v-icon>새로고침</v-btn>
        </v-card>
        <v-card>
            <v-card-title>
                <v-text-field
                    v-model="search"
                    append-icon="mdi-magnify"
                    label="태그 검색"
                    single-line
                    hide-details
                ></v-text-field>
            </v-card-title>
            <v-data-table
                v-model="selected"
                :headers="headers"
                :items="tags"
                :search="search"
                item-key="uid"
                show-select
            ></v-data-table>
        </v-card>
        <v-row justify="center">
            <v-dialog v-model="visibleAddDialog" persistent max-width="600px">
                <v-card>
                    <v-card-title>
                        <span class="headline">태그 추가</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row>
                                <v-col cols="12" sm="6" md="4">
                                    <v-text-field v-model="addDialog.uid" label="태그 고유 ID" required></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="6" md="4">
                                    <v-text-field v-model="addDialog.id" label="태그 관리 ID" required></v-text-field>
                                </v-col>
                            </v-row>
                            <v-row>
                                <v-col cols="12" sm="6" md="4">
                                    <v-text-field v-model="addDialog.assignee_id" label="태그 사용자 ID" required></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="6" md="4">
                                    <v-text-field v-model="addDialog.device_id" label="태그 장치 ID" required></v-text-field>
                                </v-col>
                            </v-row>
                        </v-container>
                        <small>*indicates required field</small>
                    </v-card-text>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" text @click="closeAddDialog">취소</v-btn>
                        <v-btn color="blue darken-1" text @click="addTag(); closeAddDialog(); listTag();">추가</v-btn>
                    </v-card-actions>
                </v-card>
        </v-dialog>
    </v-row>
    </v-container>
</template>

<script>
import axios from 'axios';
import Toast from '../components/Toast.vue';

export default {
    name: "Tag",
    data: () => ({
        search: '',
        headers: [
          { text: '고유 ID', value: 'uid' },
          { text: '관리 ID', value: 'id' },
          { text: '사용자 ID', value: 'assignee_id' },
          { text: '장치 ID', value: 'device_id' },
        ],
        tags: [],
        selected: [],
        visibleAddDialog: false,
        addDialog: {},
        msgs: [],
      }),
    components: {
        Toast,
    },
    methods: {
        showAddDialog: function() {
            this.addDialog = {
                uid: "",
                id: "",
                assignee_id: "",
                device_id: ""
            };
            this.visibleAddDialog = true;
        },
        closeAddDialog: function() {
            this.visibleAddDialog = false;
        },
        addTag: function() {
            axios.post("./api/tag", this.addDialog)
                .then(() => {
                    this.msgs.push({msg: "추가 성공!", kind: "success"});
                })
                .catch((err) => {
                    this.msgs.push({msg: err, kind: "error"});
                });
        },
        deleteTag: function() {
            const t = this.selected.map(x => x.uid);
            this.selected = [];
            if (t.length < 1) {
                this.msgs.push({msg: "삭제할 태그를 지정하세요.", kind: "error"});
                return;
            }
            //in axios 0.20, axios.delete not working with data option
            axios.request({data: {uids: t}, url: "./api/tag", method: 'delete'})
                .then(() => {
                    this.msgs.push({msg: "삭제 성공!", kind: "success"});
                })
                .catch((err) => {
                    this.msgs.push({msg: err, kind: "error"});
                });
        },
        listTag: function(delay) {
            const req = () => {
                axios.get("./api/tag")
                    .then((res) => {
                        this.tags = res.data;
                        this.msgs.push({msg: "불러오기 성공!", kind: "success"});
                    })
                    .catch((err) => {
                        this.msgs.push({msg: err, kind: "error"});
                    });
            }
            if (!delay || delay === 0) {
                req();
            } else {
                setTimeout(req, delay);
            }
        },
        mounted() { 
            this.listTag();
        }
    } 
}
</script>

<style scoped>
</style>