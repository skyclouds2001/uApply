<!-- 面试官配置页面，面试官项组件 -->
<template>
    <div class="interviewer">
        <!-- 面试官信息 -->
        <div class="interviewerInfo">
            <h1>{{ interviewer.name }}</h1>
            <h3> ID：{{ interviewer.id }} <br> UID：{{ interviewer.uid }} </h3>
        </div>
        <!-- 操作图标 -->
        <div class="icon">
            <img :src="deleteIcon" @click="deleteInt">
            <img :src="editIcon" @click="editInt">
        </div>
    </div>
</template>

<script>
import { onMounted } from '@vue/runtime-core';
import { getCurrentInstance } from 'vue';

export default {
    props: {
        // 面试官基本信息
        interviewer: Object
    },
    setup(props, context) {
        const deleteIcon = require("@/assets/imgs/错误.png");
        const editIcon = require("@/assets/imgs/编辑.png");

        // 实例对象
        const { proxy } = getCurrentInstance();

        const deleteInt = function () {
            storeInterviewer();
            context.emit("toast", "DeleteInterviewer");
        }

        const editInt = function () {
            storeInterviewer();
            context.emit("toast", "EditInterviewer");
        }

        // 将面试官信息存入缓存
        const storeInterviewer = function () {
            proxy.$sessionData("set", "interviewer", props.interviewer);
        }

        return {
            deleteIcon,
            editIcon,
            deleteInt,
            editInt
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($direction: row, $ver: flex-start) {
    display: flex;
    justify-content: space-around;
    align-items: $ver;
    flex-direction: $direction;
}

.interviewer {
    background-color: white;
    @include flex(row, center);
    justify-content: center;
    width: 22rem;
    height: 10rem;
    border-radius: 1rem;
    margin-right: 3%;
    margin-bottom: 2%;
    box-shadow: 1px 1px 5px 0px #c5a0ff;

    @media (max-width:1440px) {
        width: 17.5rem;
        height: 7.5rem;
    }

    @media (max-width:1280px) {
        width: 15.5rem;
        height: 7.5rem;
    }

    .interviewerInfo {
        @include flex(column);
        height: 70%;
        width: 47%;
    }

    .icon {
        @include flex(column);
        align-items: center;
        width: 25%;
        height: 80%;
        position: relative;
        top: 2%;
        left: 2%;

        img {
            width: 35%;
            cursor: pointer;
        }
    }
}
</style>