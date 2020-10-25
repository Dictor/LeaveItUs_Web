<template>
    <v-container fluid>
        <v-card class="pa-2 mb-2">
            <v-btn key="btn-create" color="green" @click="actionCreate" v-if="!isDisabled('c')" dark large class="mr-1"><v-icon>mdi-plus</v-icon>{{title}} 추가</v-btn>
            <v-btn key="btn-delete" color="error" @click="actionDelete" v-if="!isDisabled('d')" large class="mr-1"><v-icon>mdi-delete</v-icon>{{title}} 삭제</v-btn>
            <v-btn key="btn-update" color="orange" @click="actionUpdate"  v-if="!isDisabled('u')" :disabled="selected.length!==1" large class="mr-1"><v-icon>mdi-pencil</v-icon>{{title}} 변경</v-btn>
            <v-btn key="btn-read" color="blue"  @click="actionRead"  v-if="!isDisabled('r')" large><v-icon>mdi-refresh</v-icon>새로고침</v-btn>
        </v-card>
        <v-card>
            <v-card-title>
                <v-text-field
                    v-model="search"
                    append-icon="mdi-magnify"
                    :label="title + ' 검색'"
                    single-line
                    hide-details
                ></v-text-field>
            </v-card-title>
            <v-data-table
                v-model="selected"
                :headers="columns"
                :items="rows"
                :search="search"
                :item-key="ukey"
                :show-select="ukey"
            >
                <template v-if="slot" v-slot:[slot.name]="props">
                    <component v-for="(c, i) in props[slot.item_id]" :key="i" :is="slot.type">
                        {{ slot.item_renderer(c) }}
                    </component>
                </template>
            </v-data-table>
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
                                    <v-text-field v-model="Dialog[col.key]" :label="col.label + (col.required ? '*':'')" :disabled="DialogStatus==='update' && col.update_disable"></v-text-field>
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
properties : title, columns, rows, dialog, unique-key
events: create, read, update, delete
*/
export default {
    props: ["title", "columns", "rows", "dialog", "ukey", "slot", "disable_action"],
    data: () => ({
        search: "",
        selected: [],
        visibleDialog: false,
        Dialog: {},
        DialogStatus: ""
    }),
    methods: {
        isDisabled: function(this_action) {
            if (!this.disable_action) return false;
            return this.disable_action.includes(this_action);
        },
        applyDialog: function() {
            if (this.DialogStatus == "create") {
                this.$emit("create", this.Dialog);
            } else {
                this.$emit("update", this.Dialog);
            }
        },
        actionCreate: function() {
            this.Dialog = {};
            this.DialogStatus = "create";
            this.visibleDialog = true;
        },
        actionUpdate: function() {
            this.Dialog = this.selected[0];
            this.DialogStatus = "update";
            this.visibleDialog = true;
        },
        actionDelete: function() {
            this.$emit("delete", this.selected);
        },
        actionRead: function() {
            this.$emit("read");
        },
    }
}
</script>