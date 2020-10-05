<template>
    <v-navigation-drawer permanent>
        <v-list-item>
            <v-list-item-content>
                <v-list-item-title class="title">
                    맡기시오
                </v-list-item-title>
                <v-list-item-subtitle>
                    휴대폰 통합관리 시스템
                </v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list dense nav>
            <v-list-item v-for="item in items" :key="item.title" link>
                <v-list-item-icon>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
    export default {
        data () {
            return {
                items: [
                    { title: '태그(필증) 관리', icon: 'mdi-card-account-details' },
                    { title: '기기 관리', icon: 'mdi-home-lightbulb' },
                    { title: '병력 관리', icon: 'mdi-account-details' },
                    { title: '통계 및 기록', icon: 'mdi-database-search' },
                ],
                right: null,
            }
        },
    }
</script>

<style scoped>

</style>