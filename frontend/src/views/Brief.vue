<template>
    <v-container fluid>
        <link href="https://fonts.googleapis.com/css2?family=Black+Han+Sans&display=swap" rel="stylesheet">
        <span class="project-title">맡기시오</span><span class="project-subtitle">LeaveItUs</span>
        <p>생활관별 휴대폰 반납 시스템</p>
        <a href="https://github.com/osamhack2020/WEB_LeaveItUs_797IsPalindrome">Web 깃헙</a><a href="https://github.com/osamhack2020/IoT_LeaveItUs_797IsPalindrome">IoT 깃헙</a>
        
        <div class="version">
            <h2>빌드 정보</h2>
            <h3>프론트앤드</h3>
            <ul>
                <li>빌드 날짜 : {{frontBuildTime}}</li>
                <li>빌드 버전 : {{frontBuildVersion}}</li>
                <li>빌드 커밋 : {{frontBuildCommit}}</li>
            </ul>
            <h3>백앤드</h3>
            <ul>
                <li>빌드 날짜 : {{backBuildTime}}</li>
                <li>빌드 버전 : {{backBuildVersion}}</li>
                <li>빌드 태그 : {{backBuildCommit}}</li>
            </ul>
        </div>
    </v-container>
</template>

<script>
import {getBuildInfo} from 'vue-cli-plugin-build-info/plugin'
import axios from 'axios';

export default {
    name: "Brief",
    data: () => ({
        frontBuildVersion: getBuildInfo().VERSION,
        frontBuildTime: new Date(getBuildInfo().TIMESTAMP).toISOString().replace("T", "-").replace(".000Z", " (UTC)"),
        frontBuildCommit: getBuildInfo().COMMIT,
        backBuildCommit: "?",
        backBuildTime: "?",
        backBuildVersion: "?",
    }),
    methods: {
        getVersion: function() {
            axios.get("./api/version")
                .then((res) => {
                        this.backBuildCommit = res.data.commit.substring(0, 6);
                        this.backBuildTime = res.data.date + " (UTC)";
                        this.backBuildVersion = res.data.tag;
                    })
        }
    },
    mounted() {
        this.getVersion();
    }
}
</script>

<style scoped>
    p, a {
        margin: 0.5rem;
    }

    .version {
        margin: 0.5rem;
    }

    h2 {
        margin-top: 1rem;
    }

    .project-title {
        font-size: 5rem;
        color: darkorange;
        font-family: 'Black Han Sans', sans-serif;
    }

    .project-subtitle {
        font-size: 1.5rem;
        color: orange;
    }
</style>

