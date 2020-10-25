<template>
    <v-navigation-drawer app mobile-breakpoint=768>
        <v-list-item>
            <router-link to="/" tag="v-list-item-content">
                <v-list-item-title class="title">
                    맡기시오
                </v-list-item-title>
                <v-list-item-subtitle>
                    휴대폰 통합관리 시스템
                </v-list-item-subtitle>
            </router-link>
        </v-list-item>
        <v-divider></v-divider>
        <v-list dense nav>
            <router-link v-for="item in items" :key="item.title" :to="item.link" tag="v-list-item" link>
                <v-list-item-icon>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                </v-list-item-content>
            </router-link>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
    export default {
        data () {
            return {
                items: [
                    { title: '현재 상태', icon: 'mdi-chart-box', link: "now"},
                    
                    { title: '휴대폰 반납 기록 관리', icon: 'mdi-cellphone-arrow-down', link: "tag-record"},
                    { title: '개폐 기록 관리', icon: 'mdi-door-open', link: "door-event"},

                    { title: '태그(필증) 관리', icon: 'mdi-bookmark-multiple-outline', link: "tag"},
                    { title: '보관함 관리', icon: 'mdi-briefcase', link: "locker"},
                    { title: '병력 관리', icon: 'mdi-account-details', link: "personnel"}
                ],
                right: null,
            }
        },
    }
</script>

<style scoped>

</style>