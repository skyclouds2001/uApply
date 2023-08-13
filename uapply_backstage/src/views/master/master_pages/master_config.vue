<template>
    <div class="page">
        <!-- 新增面试官 -->
        <div class="add">
            <div class="add_interviewer">
                <div class="text">
                    <!-- <div class="text_title">新增面试官</div> -->
                    <h1>新增面试官</h1>
                    <h3>配置面试官UID</h3>
                </div>
                <img alt="" :src="addButton" @click="add_interviewer">
            </div>
        </div>
        <!-- 面试官列表 -->
        <div class="interviewers">
            <!-- 面试官们 -->
            <interviewer v-for="item in interviewers" :key="item.id" :interviewer="item" @toast="toast"></interviewer>
            <!-- 防止触底时遮盖盒子阴影 -->
            <div style="height:0.1%;width: 100%;"></div>
        </div>
    </div>
    <toast v-if="isShow" :toast="component" @isToast="isToast" @confirm="confirm"></toast>
</template>

<script>
import { ref } from 'vue';
import Toast from '../../../components/generic/toast/toast.vue';
import { onMounted } from 'vue';
import { GetAllInterviewers } from "../../../utils/apis/api"
import Interviewer from '../../../components/master/master_config/interviewer.vue';
import { message } from '../../../utils/message/message';

export default {
    components: {
        Toast,
        Interviewer
    },
    setup() {
        // 显示什么toast呢
        let component = ref("");

        // 是否显示toast
        let isShow = ref(false);

        // 需要显示或者不显示toast的时候，都需要调用此函数，不允许出现功能重复的函数
        const isToast = function () {
            // 切换toast状态
            isShow.value = !isShow.value;
        }

        const confirm = function () {
            init();
        }

        const add_interviewer = function () {
            component.value = "AddInterviewer";
            isToast();
        }

        // 根据点击的按钮展示对应的操作弹窗
        const toast = function (e) {
            component.value = e;
            isToast();
        }

        // 面试官列表
        let interviewers = ref([]);

        // 初始化
        const init = async function () {
            const res = await GetAllInterviewers();
            if (res.data.code > 1000) {
                message.error(res.data.msg);
            }
            interviewers.value = res.data.infos;
        }

        onMounted(() => {
            init();
        })

        return {
            addButton: require("@/assets/imgs/配置招新日期.png"),
            component,
            isShow,
            interviewers,
            isToast,
            confirm,
            add_interviewer,
            toast
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($direction: row, $row: center, $ver: center) {
    display: flex;
    justify-content: $row;
    align-items: $ver;
    flex-direction: $direction;
}

.interviewers {
    width: 100%;
    height: 70%;
    padding: 0 0 0 4%;
    display: flex;
    flex-wrap: wrap;
}

.add {
    width: 100%;
    height: 30%;
    display: flex;
    align-items: center;
    padding: 0 0 0 4%;

    .add_interviewer {
        background: white;
        border-radius: 1rem;
        width: 22rem;
        height: 10rem;
        @include flex();

        @media (max-width:1440px) {
            width: 17.5rem;
            height: 7.5rem;
        }

        @media (max-width:1280px) {
            width: 15.5rem;
            height: 7.5rem;
        }

        .text {
            height: 60%;
            width: 50%;
            @include flex(column, space-around, flex-start);
        }

        img {
            width: 20%;
            cursor: pointer;
        }

    }

}
</style>
