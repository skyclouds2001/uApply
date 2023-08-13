<template>
    <!-- 登录界面 --- 登录框组件 -->
    <div class="login_block">
        <!-- 登录组件标题 -->
        <div class="login_title">
            {{ block_title }}
        </div>
        <!-- 用户信息登录框 -->
        <div class="login_users">
            <!-- 角色选择 -->
            <div class="input_items">
                <!-- 图片 -->
                <img :src="role_img" class="head_img1">
                <!-- 标题 -->
                <div class="items_title1">角色</div>
                <!-- 账号级别选择 -->
                <n-select style="width: 30%" :options="role_options" v-model:value="account_infor.role" size="large">
                </n-select>
            </div>
            <!-- 账号 -->
            <div class="input_items input_AccountPassword">
                <!-- 图片 -->
                <img :src="account_img" class="head_img1">
                <!-- 标题 -->
                <div class="items_title1">账号</div>
                <!-- 账号输入栏 -->
                <input type="text" placeholder="请输入账号" v-model="account_infor.account">
            </div>
            <!-- 密码 -->
            <div class="input_items input_AccountPassword">
                <img :src="password_img" class="head_img2">
                <div class="items_title2">密码</div>
                <!-- 密码输入栏 -->
                <input :type="isPassword ? 'password' : 'text'" placeholder="请输入密码" v-model="account_infor.password">
                <!-- 选择是否可见 -->
                <img :src="isVisible ? visible : invisible" style="height:50%" @click="changeVisible">
            </div>
        </div>
        <!-- 登录按钮 -->
        <div class="login_button">
            <n-button color="#676EE4" @click="login" @keyup.enter="login">登 录</n-button>
        </div>
    </div>
</template>

<script>
import { reactive, ref } from '@vue/reactivity'
import { NButton, NSpace, NSelect } from "naive-ui"
import { message } from '../../utils/message/message';
import { onMounted } from '@vue/runtime-core';

export default {
    name: "LoginBlock",
    components: {
        NButton,
        NSelect,
        NSpace,
    },
    emits: [
        "login"
    ],
    setup(props, context) {
        const block_title = "U报名后台管理系统-登录";
        const role_img = require("@/assets/imgs/角色选择.png");
        const account_img = require("@/assets/imgs/账号.png");
        const password_img = require("@/assets/imgs/密码.png");
        const visible = require("@/assets/imgs/密码可见.png");
        const invisible = require("@/assets/imgs/密码不可见.png");
        let isPassword = ref(true);
        let isVisible = ref(false);
        const role_options = reactive([
            {
                label: "超级管理员",
                value: "超级管理员"
            },
            {
                label: "管理员",
                value: "管理员"
            }
        ])
        // 存放用户登录信息
        let account_infor = reactive({
            role: '超级管理员',
            account: "",
            password: ""
        })

        // 切换密码是否可见，及是否可见图标
        const changeVisible = function () {
            isVisible.value = !isVisible.value;
            isPassword.value = !isPassword.value;
        };

        // 点击登录
        // 做一个简单的账号密码是否填写的校验
        const login = function () {
            if (account_infor.account == "") {
                message.info("请输入账号后再登录！")
                return ;
            } else if (account_infor.password == "") {
                message.info("请输入密码后再登录！")
                return ;
            }
            context.emit("login", account_infor);
        }

        // 监听键盘，如果按下回车就执行登录
        onMounted(() => {
            window.addEventListener("keyup", function (e) {
                if (e.key == "Enter") {
                    login();
                }
            })
        })

        return {
            block_title,
            role_img,
            account_img,
            password_img,
            isPassword,
            visible,
            invisible,
            isVisible,
            changeVisible,
            account_infor,
            role_options,
            login
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($direction: column, $isCenter: true) {
    display: flex;
    flex-direction: $direction;

    @if $isCenter {
        align-items: center;
        justify-content: center;
    }
}

// 登录框适配
@mixin login_block {
    padding: 1%;
    background: rgb(247, 250, 250);
    border-radius: 2vh;

    @media (max-width:1440px) {
        width: 45rem;
        height: 30rem;
    }

    @media (max-width:1366px) {
        width: 40rem;
        height: 25rem;
    }
}

// 登录框
.login_block {
    width: 50rem;
    height: 34rem;
    @include login_block;

    // 登录标题
    .login_title {
        height: 15%;
        width: 90%;
        margin-bottom: 2%;
        @include flex(column, true);
        font-weight: 900;
        letter-spacing: 1px;
        font-size: 2rem;
        color: rgba(148, 148, 148, 1);
    }

    // 用户角色选项及账号密码输入框
    .login_users {
        height: 55%;
        width: 80%;
        display: flex;
        flex-direction: column;
        justify-content: space-around;

        .input_AccountPassword {
            background: #ffffff;
            box-shadow: 3px 3px 0px 2px rgb(240, 240, 240);
        }

        .input_items {
            height: 25%;
            width: 100%;
            display: flex;
            justify-content: left;
            align-items: center;

            .head_img1 {
                height: 55%;
                padding-left: 2.5%;
            }

            .head_img2 {
                height: 95%;
                padding-left: 0%;
            }

            @mixin items_title {
                font-weight: 550;
                font-size: 1.5rem;
                color: rgba(148, 148, 148, 1);
            }

            .items_title1 {
                padding: 0 5% 0 3%;
                @include items_title;
            }

            .items_title2 {
                padding: 0 5% 0 1%;
                @include items_title;
            }

            input {
                height: 50%;
                width: 60%;
                margin-right: 3%;
                border: none;
                padding: 0 1%;
                outline: none;
                font-size: 1.2rem;
                letter-spacing: 2px;
            }


        }

    }

    // 登录按钮
    .login_button {
        height: 15%;
        width: 100%;
        position: relative;
        top: 2%;
        left: 0.5%;
        @include flex(column, true);

        // 按钮样式 - Naive UI 默认类名
        .n-button {
            height: 90%;
            width: 30%;
            box-shadow: 3px 3px 5px 0px #3307CC;
            font-size: 1.7rem;
        }
    }

    input::-webkit-input-placeholder {
        /* // 针对 谷歌 内核 */
        color: rgb(184, 182, 182);
    }

    input:-moz-placeholder {
        /* // 火狐 */
        color: rgb(184, 182, 182);
    }

    input:-ms-input-placeholder {
        /* // 微软ie */
        color: rgb(184, 182, 182);
    }


}
</style>

