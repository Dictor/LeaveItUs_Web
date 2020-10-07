<template>
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
            :headers="headers"
            :items="tags"
            :search="search"
        ></v-data-table>
    </v-card>
</template>

<script>
export default {
    name: "Tag",
    data: () => ({
        search: '',
        headers: [
          { text: '고유 ID', value: 'uid' },
          { text: '관리 ID', value: 'id' },
          { text: '사용자 ID', value: 'user_id' },
          { text: '장치 ID', value: 'device_id' },
        ],
        tags: [
        ],
      }),
}
</script>