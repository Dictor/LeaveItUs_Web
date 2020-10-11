<template>
    <v-container fluid>
        <v-container class="pa-2 mb-2 toast">
            <v-container v-for="(error, i) in errors" v-bind:key="i" class="pa-0 ma-0">
                <v-alert :type="error.kind" class="ma-1" v-if="error.visible">
                    {{ error.msg }}
                </v-alert>
            </v-container>
        </v-container>
        <v-card class="pa-2 mb-2">
            <v-btn key="" color="green" v-on:click="shoWAddDialog" dark large class="mr-1"><v-icon>mdi-plus</v-icon>태그 추가</v-btn>
            <v-btn key="" color="error" v-on:click="deleteTag(); listTag();" large class="mr-1"><v-icon>mdi-delete</v-icon>태그 삭제</v-btn>
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
                        <v-btn color="blue darken-1" text @click="shoWAddDialog">취소</v-btn>
                        <v-btn color="blue darken-1" text @click="addTag(); visibleAddDialog = false; listTag();">추가</v-btn>
                    </v-card-actions>
                </v-card>
        </v-dialog>
    </v-row>
    </v-container>
</template>

<script>
import axios from 'axios';

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
        errors: []
      }),
    methods: {
        shoWAddDialog: function() {
            this.addDialog = {
                uid: "",
                id: "",
                assignee_id: "",
                device_id: ""
            };
            this.visibleAddDialog = true;
        },
        addTag: function() {
            axios.post("./api/tag", this.addDialog)
                .then(() => {
                    this.addToast("추가 성공!", "success");
                })
                .catch((err) => {
                    this.addToast(err, "error");
                });
        },
        deleteTag: function() {
            const t = this.selected.map(x => x.uid);
            this.selected = [];
            if (t.length < 1) {
                this.addToast("삭제할 태그를 지정하세요.", "error");
                return;
            }
            //in axios 0.20, axios.delete not working with data option
            axios.request({data: {uids: t}, url: "./api/tag", method: 'delete'})
                .then(() => {
                    this.addToast("삭제 성공!", "success");
                })
                .catch((err) => {
                    this.addToast(err, "error");
                });
        },
        listTag: function() {
            axios.get("./api/tag")
                .then((res) => {
                    this.tags = res.data;
                    this.addToast("불러오기 성공!", "success");
                })
                .catch((err) => {
                    this.addToast(err, "error");
                });
        },
        addToast: function(msg, kind) {
            let i = this.errors.push({msg: msg, visible: true, kind: kind});
            setTimeout(() => {
                this.errors[i-1].visible = false;
            }, 2000);
        },
        mounted() { 
            this.listTag();
        }
    } 
}
</script>

<style scoped>
    .toast {
        position: fixed;
        bottom: 0;
        right: 0;
    }
</style>