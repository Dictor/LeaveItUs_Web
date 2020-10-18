<template>
    <v-container fluid>
        <v-card class="pa-2 mb-2">
            <!--
            <v-btn v-for="(button, i) in buttons" :key="i" @click="button.onclick" :color="button.color" large dark class="mr-1">
                <v-icon>{{button.icon}}</v-icon>
            </v-btn>
            -->
            <v-btn key="btn-create" color="green" @click="actionCreate" dark large class="mr-1"><v-icon>mdi-plus</v-icon>{{title}} 추가</v-btn>
            <v-btn key="btn-delete" color="error" @click="actionDelete" large class="mr-1"><v-icon>mdi-delete</v-icon>{{title}} 삭제</v-btn>
            <v-btn key="btn-update" color="orange" @click="actionUpdate" :disabled="selected.length===1" large class="mr-1"><v-icon>mdi-pencil</v-icon>{{title}} 변경</v-btn>
            <v-btn key="btn-read" color="blue"  @click="actionRead" large><v-icon>mdi-refresh</v-icon>새로고침</v-btn>
        </v-card>
        <v-card>
            <v-card-title>
                <v-text-field
                    v-model="search"
                    append-icon="mdi-magnify"
                    :label="title + 검색"
                    single-line
                    hide-details
                ></v-text-field>
            </v-card-title>
            <v-data-table
                v-model="selected"
                :headers="columns"
                :items="rows"
                :search="search"
                item-key="uid"
                show-select
            ></v-data-table>
        </v-card>
        <v-row justify="center">
            <v-dialog v-model="visibleDialog" persistent max-width="600px">
                <v-card>
                    <v-card-title>
                        <span class="headline">{{title}} {{DialogStatus == "create" ? "추가" : "변경"}}</span>
                    </v-card-title>
                    <v-card-text>
                        <v-container>
                            <v-row v-for="(row, ri) in dialog" :key="ri">
                                <v-col v-for="(col, ci) in row" :key="ci" :cols="24 / row.length" :sm="12 / row.length" :md="8 / row.length">
                                    <v-text-field v-model="Dialog[col.key]" :label="col.label" :required="col.required" :disabled="DialogStatus==='update'"></v-text-field>
                                </v-col>
                            </v-row>
                        </v-container>
                        <small>*표시된 항목은 필수 입력</small>
                    </v-card-text>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" text @click="visibleDialog = false;">취소</v-btn>
                        <v-btn color="blue darken-1" text @click="applyDialog(); visibleDialog = false;">적용</v-btn>
                    </v-card-actions>
                </v-card>
        </v-dialog>
    </v-row>
    </v-container>
</template>

<script>
/*
properties : title, columns, rows, dialog
events: create, read, update, delete
*/
export default {
    props: ["title", "columns", "rows", "dialog"],
    data: () => ({
        search: "",
        selected: [],
        visibleDialog: false,
        Dialog: {},
        DialogStatus: ""
    }),
    method: {
        applyDialog: () => {
            if (this.DialogStatus == "create") {
                this.$emit("create", this.Dialog);
            } else {
                this.$emit("update", this.Dialog);
            }
        },
        actionCreate: () => {
            this.Dialog = {};
            this.DialogStatus = "create";
            this.visibleDialog = true;
        },
        actionUpdate: () => {
            this.Dialog = this.selected[0];
            this.DialogStatus = "update";
            this.visibleDialog = true;
        },
        actionDelete: () => {
            this.$emit("delete", this.selected);
        },
        actionRead: () => {
            this.$emit("read");
        },
    }
}
</script>