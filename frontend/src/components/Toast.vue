<template>
    <v-container class="pa-2 mb-2 toast-parent">
        <v-container fluid v-for="(msg, i) in msgs" :key="i" class="pa-0 ma-0">
            <v-alert :type="msg.kind ? msg.kind : 'info'" class="ma-1">
                {{ !(msg.msg) ? msg: msg.msg }}
            </v-alert>
        </v-container>
    </v-container>
</template>

<script>
export default {
    name: "Toast",
    data: () => ({
    }),
    props: ["msgs"],
    mounted() {
        setInterval(() => {
            const val = this.msgs;
            if (!val) return;
            const now = Date.now();
            for (let i = 0; i < val.length; i++) {
                if (val[i].created) {
                    if (now - val[i].created > 2000) {
                        val.splice(i, 1);
                    }
                } else {
                    val[i].created = now;
                }
            }
        }, 500);
    }
}
</script>
        
<style scoped>
    .toast-parent {
        max-width: 50%;
        min-width: 300px;
        z-index: 999;
        position: fixed;
        bottom: 0;
        right: 0;
    }
</style>