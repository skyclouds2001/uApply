<template>
    <page-head class="statusAnimate" :title="'我的状态'" @goBack="goBack"></page-head>
    <div class="status statusAnimate">
        <div class="blocks">
            <!-- 用户信息 -->
            <div class="user">
                <p>学号：{{ userInfo.userNumb }}</p>
                <p>UID :&nbsp; {{ userInfo.userUid }}</p>
                <img :src="userImg" alt="">
            </div>
            <!-- 如果没有申报记录 -->
            <div class="goApply" v-if="isEmpyt" @click="goApply">
                <span class="title">申报心仪部门</span>
                <div class="img">
                    Go
                </div>
            </div>
            <div class="emptyImg" v-if="isEmpyt">
                <img :src="emptyImg" alt="">
                <p>流程状态暂时为空哦</p>
                <p>快去申报心仪部门吧</p>
            </div>
            <!-- 如果有申报记录 -->
            <div class="applyHistory" v-for="(item, index) in applys"
                :style="{ background: (index + 1) % 3 == 2 ? '#8972FF' : (index + 1) % 3 == 0 ? '#62bdfe' : '#FFA2A2' }">
                <div class="details">
                    <p>组织 : {{ item.org_name }}</p>
                    <p>部门 : {{ item.dep_name }}</p>
                    <p style="width:100%">流程状态 : {{ item.status }}</p>
                </div>
                <div class="index"
                    :style="{ background: (index + 1) % 3 == 2 ? '#7960f6' : (index + 1) % 3 == 0 ? '#0093FA' : '#FC8080' }">
                    {{ index + 1 }}
                </div>
            </div>
            <!-- 背景 -->
            <div class="bg"></div>
        </div>
    </div>
</template>

<script>
import { onMounted, reactive, ref } from '@vue/runtime-core';
import PageHead from '../head/page_head.vue';
import cookies from "vue-cookies"
import router from '../../router';
import { GetStatus } from '../../utils/apis/api';

export default {
    components: {
        PageHead
    },
    setup() {
        onMounted(async () => {
            $(".statusAnimate").animate({
                opacity: 1,
            })
            userInfo.userNumb = cookies.get("userNumb");
            userInfo.userUid = cookies.get("userUid");

            const res = await GetStatus();
            if (res.data.rsp == null || undefined) {
                isEmpyt.value = true;
                return;
            }
            const rsp = res.data.rsp
            applys.value = rsp;
            for (let i = 0; i < rsp.length; i++) {
                applys.value[i].status = initStatus(rsp[i]);
            }
            isEmpyt.value = false;
        })

        const initStatus = function (e) {
            let text = "";
            text = statusTranslate(e.first, 1);
            if (e.second != 0) {
                text = statusTranslate(e.second, 2);
            }
            if (e.final != 0) {
                text = statusTranslate(e.final, 3);
            }
            return text;
        }

        const statusTranslate = function (code, index) {
            let statusText = "筛选简历";
            if (code == 0) {
                statusText = "筛选简历";
            } else if (code == 1) {
                statusText = "已安排面试";
            } else if (code == 2) {
                statusText = "已面试";
            } else if (code == 3) {
                statusText = "已通过";
            } else if (code == 4) {
                statusText = "已淘汰";
            }
            return "第" + index + "轮-" + statusText;
        }

        const goBack = function () {
            router.push('/');
        }

        // 用户学号及ID
        let userInfo = reactive({
            userNumb: 0,
            userUid: 0
        })

        // 是否为空
        let isEmpyt = ref(true);

        // 申报列表
        let applys = ref([]);

        const userImg = require("@/static/imgs/个人信息.png");
        const emptyImg = require("@/static/imgs/背景2.png");

        // 前往报名页面
        const goApply = function () {
            router.push("/apply");
        }

        return {
            userImg,
            emptyImg,
            isEmpyt,
            applys,
            userInfo,
            goApply,
            goBack
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($row: row, $isCerten: true) {
    display: flex;
    flex-direction: $row;

    @if ($isCerten) {
        justify-content: center;
        align-items: center;
    }
}

.statusAnimate {
    opacity: 0;
}

.blocks {
    position: relative;
    top: 0.1vh;
    width: 100%;
    height: 94.5vh;
    overflow: auto;
    // background: black;

    .bg {
        position: absolute;
        top: -2vh;
        // height: 38vh;
        height: 40vh;
        width: 100%;
        background: #6C7BFF;
        z-index: -1;
    }

    @mixin block {
        position: relative;
        height: 17vh;
        width: 100%;
        z-index: 3;
        margin-top: 2vh;
        border-radius: 1.2vh;
    }

    .user {
        position: relative;
        @include block();
        @include flex(column, );
        background: #8CD742;

        img {
            position: absolute;
            right: 0;
            height: 90%;
        }

        p {
            font-size: 0.6rem;
            width: 80%;
            height: 40%;
            display: flex;
            align-items: center;
            color: white;
            padding-left: 0vw;

            @media (min-width:1280px) {
                font-size: 0.13rem;
            }
        }
    }

    .goApply {
        position: relative;
        @include block();
        background: #F3AE84;
        display: flex;
        align-items: center;
        font-size: 0.8rem;

        @media (min-width:1280px) {
            font-size: 0.2rem;
        }

        .title {
            position: absolute;
            left: 10%;
            color: white;

            @media (min-width:1280px) {}
        }

        .img {
            @include flex(row, );
            color: white;
            position: absolute;
            right: 10%;
            height: 9vh;
            width: 9vh;
            background: #FC985C;
            border-radius: 50%;
        }
    }

    .emptyImg {
        @include flex(column, );
        height: 50%;
        width: 100%;

        img {
            width: 90%;
            opacity: 0.85;
            padding-bottom: 5%;
        }

        p {
            font-size: 0.5rem;
            color: #e1d4f2;

            @media (min-width:1280px) {
                font-size: 0.1rem;
            }
        }
    }

    .applyHistory {
        @include block();
        @include flex(row, );
        position: relative;

        .details {
            height: 70%;
            width: 65%;
            position: absolute;
            left: 10%;
            @include flex(row, );
            flex-wrap: wrap;

            p {
                display: flex;
                align-items: center;
                height: 50%;
                width: 50%;
                font-size: 0.5rem;
                color: rgb(255, 255, 255);

                @media (min-width:1280px) {
                    font-size: 0.12rem;
                }
            }

            p:hover {
                color: #6cffda;
            }
        }

        .index {
            height: 7vh;
            width: 7vh;
            font-size: 0.7rem;
            position: absolute;
            border-radius: 50%;
            right: 7%;
            color: white;
            @include flex(row, );

            @media (min-width:1280px) {
                font-size: 0.2rem;
            }
        }
    }

}
</style>