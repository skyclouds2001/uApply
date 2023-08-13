<template>
    <page-head :title="title" class="indexAnimate" :isShowReturen="false"></page-head>
    <div class="index indexAnimate">
        <div class="background">
            <img :src="background" alt="">
        </div>
        <div class="buttons">
            <small-block :type="'resume'" @click="goResume"></small-block>
            <small-block :type="'apply'" @click="goApply"></small-block>
            <large-block :type="'status'" @click="goStatus"></large-block>
            <large-block :type="'login'" ref="loginBlock" @click="goLogin"></large-block>
        </div>
        <div class="studio">
            <span>- 青鸢工作室出品 -</span>
        </div>
    </div>

    <n-modal v-model:show="showLogin">
        <n-card style="width:400px" title="登录">
            <div class="cardTitle">使用西电统一登录账号登录：</div>
            <div>
                账号：<n-input placeholder="请输入学号" v-model:value="loginData.username"></n-input>
            </div>
            <div style="height:1vh"></div>
            <div>
                密码：<n-input placeholder="请输入密码" v-model:value="loginData.password" type="password"></n-input>
            </div>
            <div class="button">
                <n-button type="success" round strong size="large" @click="Login">登 录
                </n-button>
            </div>
        </n-card>
    </n-modal>
</template>

<script>
import PageHead from '../head/page_head.vue';
import SmallBlock from '../../components/index/small_block.vue';
import LargeBlock from '../../components/index/large_block.vue';
import { dialog, message } from '../../utils/message/message';
import router from '../../router';
import { getCurrentInstance, onMounted, reactive, ref } from '@vue/runtime-core';
import $ from 'jquery'
import { UserLogin } from "../../utils/apis/api"
import { NModal, NCard, NInput, NButton } from 'naive-ui';
import cookies from "vue-cookies";

export default {
    components: {
        PageHead,
        SmallBlock,
        LargeBlock,
        NModal,
        NCard,
        NInput,
        NButton
    },
    setup() {
        onMounted(async () => {
            $(".indexAnimate").animate({
                opacity: 1,
            })
            if (cookies.get("userUid") != null) {
                isLogin = true;
                return;
            }
        })

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 背景图片
        const background = require("@/static/imgs/背景.png");

        // 是否已经登录
        let isLogin = false;
        // 是否显示登录框
        let showLogin = ref(false);
        // 登录信息
        let loginData = reactive({
            username: "",
            password: ""
        })
        // 登录状态
        let loginStatus = ref("login");

        const goResume = function () {
            router.push("/resume");
        }
        const goApply = function () {
            router.push("/apply");
        }
        const goStatus = function () {
            router.push("/status");
        }
        const goLogin = function () {
            if (isLogin) {
                message.success("您已登录");
                console.log(loginStatus.value);
                dialog.success({
                    title: "是否退出登录？",
                    content: () => "退出之后，需要重新使用西电统一登录账号再次登录，是否确认退出？",
                    positiveText: "退出!",
                    negativeText: "算了",
                    maskClosable: true,
                    onPositiveClick: () => {
                        cookies.remove("userInfo");
                        cookies.remove("userUid");
                        cookies.remove("userToken");
                        cookies.remove("userNumb");
                        isLogin = false;
                        proxy.$refs.loginBlock.update();
                    }
                })
                return;
            }
            showLogin.value = true;
        }

        const Login = async function () {
            const res = await UserLogin(loginData);
            // console.log(res);
            if (res.data.code > 1000) {
                message.error(res.data.msg);
                return;
            }
            if (res.data.accessToken != undefined) {
                message.success("登录成功！");
                showLogin.value = false;
                cookies.set("userInfo", res.data);
                cookies.set("userUid", res.data.uid);
                cookies.set("userToken", res.data.accessToken);
                // 学号也存一下吧
                cookies.set("userNumb",loginData.username);
                isLogin = true;
                proxy.$refs.loginBlock.update();
            }
        }

        return {
            title: "U报名 - 首页",
            background,
            showLogin,
            loginData,
            loginStatus,
            goResume,
            goApply,
            goStatus,
            goLogin,
            Login
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

.indexAnimate {
    opacity: 0;
}

.index {

    @include flex(row, false);
    height: 95vh;
    width: 100%;
    flex-wrap: wrap;

    .background {
        @include flex(column, );
        justify-content: flex-end;
        overflow: hidden;
        z-index: -1;
        height: 37%;
        width: 100%;
        background: #6C7BFF;

        img {
            position: relative;
            top: 10%;
            width: 100%;
        }
    }

    .buttons {
        @include flex(row, false);
        justify-content: space-between;
        flex-wrap: wrap;
        height: 53%;
        width: 100%;
        cursor: pointer;
    }

    .studio {
        height: 10%;
        width: 100%;
        @include flex(row, true);

        span {
            color: #e1d4f2;
            font-size: large;
        }
    }
}

.cardTitle {
    position: relative;
    bottom: 1.5vh;
}

.button {
    margin-top: 3vh;
    height: 5vh;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;

    button {
        width: 35%;
        height: 4vh;
    }
}
</style>