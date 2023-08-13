<template style="position:relative">
    <!-- 页头 -->
    <page-head :title="'完善简历'" class="resumeAnimate" @goBack="goBack"></page-head>
    <!-- 功能区 -->
    <div style="position:relative" class="resumeAnimate">
        <!-- 绝对定位的背景 -->
        <div class="bg">
            <img :src="background">
        </div>
        <!-- 简历区 -->
        <div class="table">
            <!-- “须知”图标 -->
            <img :src="noticeIcon" @click="showNotice">
            <!-- 须知弹窗 -->
            <n-modal v-model:show="isShowNotice">
                <n-card title="填写须知" style="width:400px">
                    <p v-for="item in noticeText">{{ item }} <br><br></p>
                </n-card>
            </n-modal>
            <!-- 简历标题 -->
            <div class="resumeTitle">
                <span>我的简历</span>
            </div>
            <!-- 简历表单 -->
            <div class="resumeTable">
                <div class="item">
                    <span style="width:20%;">姓名：</span>
                    <n-input placeholder="请输入姓名" v-model:value="params.data.name"></n-input>
                </div>
                <div class="item">
                    <span style="width:20%;">性别：</span>
                    <n-select :options="options" v-model:value="params.data.sex" class="select"></n-select>
                </div>
                <div class="item">
                    <span style="width:20%;">学号：</span>
                    <n-input placeholder="请输入11位学号" v-model:value="params.data.student_num"></n-input>
                </div>
                <div class="item">
                    <span style="width:20%;">楼号：</span>
                    <n-input placeholder="请输入宿舍楼号" v-model:value="params.data.address_num"></n-input>
                </div>
                <div class="item">
                    <span style="width:20%;">大类：</span>
                    <n-input placeholder="请输入专业大类" v-model:value="params.data.major"></n-input>
                </div>
                <div class="item">
                    <span style="width:20%;">邮箱：</span>
                    <n-input placeholder="请输入常用邮箱" v-model:value="params.data.email"></n-input>
                </div>
                <div class="item">
                    <span style="width:30%;">手机号：</span>
                    <n-input placeholder="请输入常用手机号" v-model:value="params.data.phone_num"></n-input>
                </div>
                <div class="item_">
                    <div style="height: 1%;"></div>
                    <span style="width:100%;">个人简介：</span>
                    <n-input style="height:75%;margin-top:1%;" type="textarea"
                        placeholder="请输入个人简介，如个人爱好、性格、技能及实践经历等，字数上限200字" v-model:value="params.data.intro"></n-input>
                </div>
            </div>
        </div>
        <!-- 按钮区 -->
        <div class="button">
            <button @click="save">
                保存简历
            </button>
        </div>
    </div>
</template>

<script>
import router from '../../router';
import { NModal, NCard, NInput, NSelect } from 'naive-ui';
import PageHead from '../head/page_head.vue';
import { reactive, ref, toRefs } from '@vue/reactivity';
import { onMounted } from '@vue/runtime-core';
import { GetResume, SaveResume } from '../../utils/apis/api';
import { dialog, message } from '../../utils/message/message';

export default {
    components: {
        PageHead,
        NModal,
        NCard,
        NInput,
        NSelect
    },
    setup() {
        const goBack = function () {
            router.push('/');
        }

        onMounted(async () => {
            $(".resumeAnimate").animate({
                opacity: 1
            })

            const res = await GetResume();
            if (!res.data.exist) {
                message.success("未找到已有简历，请填写");
            } else if (res.data.exist) {
                dialog.success({
                    title: "已有简历备份",
                    content: "查询到已有保存的简历，如需自动恢复已保存内容请点击恢复，如需重新填写简历请点击重新填写。",
                    positiveText: "恢复",
                    negativeText: "重新填写",
                    onPositiveClick: () => {
                        params.data = res.data.resume;
                        message.success("成功恢复！");
                    },
                    onNegativeClick: () => {
                        params.data = {};
                    }
                })
            }
            // console.log(res.data.resume);
            // console.log(params);
        })

        const background = require("@/static/imgs/背景.png");

        const save = async function () {
            for (const key in params.data) {
                if (params.data[key] == "") {
                    message.warning("请填完后保存~");
                }
            }
            const res = await SaveResume(params.data);
            if (res.data.code > 1000) {
                message.error(res.data.msg);
                return;
            }
            message.success("保存成功！");
        }

        const noticeIcon = require("@/static/imgs/须知.png");

        let isShowNotice = ref(false);
        const showNotice = function () {
            isShowNotice.value = true;
        }

        const noticeText = [
            "①安全隐私说明：xxxxxxxx",
            "②内容填写完整提示：xxxxxxxx",
            "③申报简历规则：每次申报时提交简历以最新保存的简历为准",
            "④简历可重复编辑，请注意保存"
        ]

        const options = [
            {
                label: "男生",
                value: "男",
            },
            {
                label: "女生",
                value: "女"
            }
        ]

        let params = reactive({
            data: {
                name: "",
                sex: "男",
                student_num: "",
                address_num: "",
                major: "",
                email: "",
                phone_num: "",
                intro: ""
            }
        })

        return {
            background,
            noticeIcon,
            isShowNotice,
            options,
            params,
            noticeText,
            save,
            showNotice,
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

.resumeAnimate {
    opacity: 0;
}

.bg {
    overflow: hidden;
    position: relative;
    top: 0.1vh;
    height: 35%;
    width: 100%;
    background: #6C7BFF;
    z-index: 2;

    img {
        width: 105%;
    }
}

.table {
    height: 80vh;
    width: 90%;
    margin-left: 5%;
    background: white;
    box-shadow: 0px 0.1px 6px 0px #bcbcbc;
    border-radius: 1vh;
    position: absolute;
    top: 2vh;
    z-index: 3;

    @media (min-width:1280px) {
        box-shadow: 0px 0.1px 3px 0px #c0bfbf;
    }

    img {
        position: absolute;
        right: 5%;
        height: 5%;
        top: 3%;
        cursor: pointer;
    }

    .resumeTitle {
        @include flex(row, true);
        height: 10%;
        width: 100%;

        span {
            font-size: 0.6rem;
            font-weight: 600;
            letter-spacing: 1px;

            @media (min-width:1280px) {
                font-size: 0.15rem;
                letter-spacing: 0.5px;
            }
        }
    }

    .resumeTable {
        height: 90%;
        width: 100%;
        @include flex(row, true);
        flex-wrap: wrap;
        justify-content: center;

        .item {
            height: 7%;
            width: 82%;
            @include flex(row, false);
            align-items: center;

            .select {
                @media (max-width:768px) {
                    position: relative;
                    bottom: 25%;
                }
            }
        }

        .item_ {
            height: 50%;
            width: 82%;
        }

        span {
            font-size: 0.4rem;

            @media (min-width:1280px) {
                font-size: large;
            }
        }
    }
}

button {
    position: fixed;
    width: inherit;
    bottom: 3%;
    z-index: 3;
    width: 45%;
    height: 7%;
    color: white;
    background: #9160FF;
    border: none;
    border-radius: 10vh;
    font-size: 0.5rem;
    left: 27.5%;

    @media (min-width:1280px) {
        width: 10%;
        font-size: 0.12rem;
        left: 45%;
    }
}

button:hover {
    background: #6d3fd8;
}
</style>