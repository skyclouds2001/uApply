<template>
    <div class="page_head">
        <div class="org_title">
            {{ title }}
        </div>
        <div class="user_title">
            {{ user_title }}
            <n-button round size="large" @click="sign_out">退 出</n-button>
        </div>
    </div>
</template>

<script>
import { NButton } from 'naive-ui'
import cookies from "vue-cookies";
import router from '../../../router';

export default {
    components: {
        NButton
    },
    props: {
        title: String,
        role: String,
    },
    setup(props, context) {
        const user_title = "欢迎回来，" + props.role + "~"

        const sign_out = function(){
            cookies.remove("Token");
            router.replace('/');
        }

        return {
            user_title,
            sign_out
        }
    }
}
</script>

<style lang="scss" scoped>
.page_head {
    height: 10%;
    width: 90%;
    padding: 0 5%;
    border-bottom: 1px solid lightgray;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: white;

    font-size: 1.7rem;

    @media (max-width:1680px) {
        font-size: 1.5rem;
    }

    @media (max-width:1440px) {
        font-size: 1.3rem;
    }

    @media (max-width:1280px) {
        font-size: 1.2rem;
    }

    .user_title {
        width: 30%;
        display: flex;
        align-items: center;
        justify-content: space-between;

        @media (max-width:1680px) {
            width: 33%;
        }

        @media (max-width:1440px) {
            width: 35%;
        }

        @media (max-width:1280px) {
            width: 37%;
        }

        .n-button {
            color: white;
        }
    }
}
</style>