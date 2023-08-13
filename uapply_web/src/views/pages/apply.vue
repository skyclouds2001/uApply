<template>
    <page-head :title="pageHead" class="applyAnimate" @goBack="goBack"></page-head>
    <div class="apply applyAnimate">
        <div class="blocks">
            <!-- 用户信息 -->
            <div class="user">
                <p>学号：{{ userInfo.userNumb }}</p>
                <p>UID :&nbsp; {{ userInfo.userUid }}</p>
                <img :src="userImg" alt="">
            </div>
            <!-- 申报须知 -->
            <div class="shouldKnow" @click="toastNotice">
                <img :src="shouldKnowImg1" class="ornaments">
                <span>申报须知</span>
                <div class="arrow"><img :src="shouldKnowImg2" alt=""></div>
            </div>
            <!-- 组织 -->
            <div class="org">
                <div class="select">
                    <p>选择意向组织</p>
                    <n-select size="large" :options="orgs" v-model:value="chosenOrgIndex"></n-select>
                </div>
                <div class="icon">
                    组
                </div>
            </div>
            <!-- 部门 -->
            <div class="dep">
                <div class="select">
                    <p>选择意向部门</p>
                    <n-select size="large" :options="deps" v-model:value="chosenDepIndex"></n-select>
                </div>
                <div class="icon">
                    部
                </div>
            </div>
            <!-- 背景 -->
            <div class="bg"></div>
        </div>
        <div class="button">
            <n-button size="large" type="success" circle strong @click="submit">提交申报</n-button>
        </div>
    </div>

    <n-modal v-model:show="showNotice">
        <n-card style="width:500px" title="申报须知">
            <p v-for="item in noticeText">{{ item }} <br> <br> </p>
        </n-card>
    </n-modal>

</template>

<script>
import PageHead from '../head/page_head.vue';
import { NButton, NModal, NCard, NSelect } from 'naive-ui';
import { onMounted, reactive, ref, watch } from '@vue/runtime-core';
import { message } from '../../utils/message/message';
import { GetDepartment, GetOrganization, GetResume, SubmitRegistration } from '../../utils/apis/api';
import cookies from "vue-cookies";
import router from '../../router';

export default {
    components: {
        PageHead,
        NButton,
        NModal,
        NCard,
        NSelect
    },
    watch() {

    },
    setup() {
        const goBack = function () {
            router.push('/');
        }

        onMounted(async () => {
            // 进入动画
            $(".applyAnimate").animate({
                opacity: 1,
            })
            // 获取正在招新的组织列表
            const resOrg = await GetOrganization();
            let tmpDeps = [];
            resOrg.data.organizations.forEach(async (item1, index1) => {
                // 更新组织列表
                orgs.value[index1] = {
                    label: item1.org_name,
                    value: item1.org_id,
                }
                // console.log(deps.value);
            });
            // 初始默认展示第一个组织
            chosenOrgIndex.value = orgs.value[0].value;
            // 默认展示第一个部门
            // chosenDepIndex.value = deps.value[0].value;
            userInfo.userNumb = cookies.get("userNumb");
            userInfo.userUid = cookies.get("userUid");
        })

        // 用户学号及ID
        let userInfo = reactive({
            userNumb: 0,
            userUid: 0
        })

        // 招新的组织以及组织的部门列表
        let orgs = ref([]);
        // 所有组织的部门都存放在这里，第一个部门存放在二维数组的[0]...
        let deps = ref([]);

        // 是否要请求
        let isGetDeps = 1;

        // 选中的组织的索引
        let chosenOrgIndex = ref(0);
        let chosenDepIndex = ref(0);

        // 暂时存放部门
        let tmp = []

        watch(chosenOrgIndex, async (newV, oldV) => {
            // 如果是还未发送过请求
            if (isGetDeps == 1) {
                for (let i = 0; i < orgs.value.length; i++) {
                    tmp[i] = (await GetDepartment(orgs.value[i].value)).data.deps;
                    tmp[i].forEach((item, index, arr) => {
                        arr[index] = {
                            label: item.dep_name,
                            value: item.dep_id
                        }
                    })
                }
                // console.log(tmp);
                deps.value = tmp[0];
                chosenDepIndex.value = deps.value[0].value;
                isGetDeps = 0;
                return;
            }
            for (let i = 0; i < orgs.value.length; i++) {
                if (orgs.value[i].value == newV) {
                    if (tmp[i].length == 0) {
                        console.log(i);
                        deps.value = [
                            {
                                label: "暂无部门可选",
                                value: -1
                            }
                        ]
                        chosenDepIndex.value = -1;
                        return;
                    }
                    deps.value = tmp[i];
                    chosenDepIndex.value = deps.value[i].value;
                }
            }
            // for (let i = 0; i < orgs.value.length; i++) {
            //     if (orgs.value[i].value == newV) {
            //         // 获取当前组织下正在招新的部门
            //         const resDep = await GetDepartment(item1.org_id);
            //         // console.log(resDep.data.deps);
            //         tmpDeps[index1] = resDep.data.deps;
            //         if (tmpDeps[index1].length != 0) {
            //             tmpDeps[index1].forEach((item2, index2, arr2) => {
            //                 arr2[index2] = {
            //                     label: item2.dep_name,
            //                     value: item2.dep_id,
            //                 }
            //             })
            //         }
            //         deps.value[index1] = tmpDeps[index1];
            //     }
            // }
        });

        const pageHead = "我要申报";

        const userImg = require("@/static/imgs/个人信息.png");

        const shouldKnowImg1 = require("@/static/imgs/我的状态--装饰.png");
        const shouldKnowImg2 = require("@/static/imgs/我的状态.png");

        const noticeText = [
            "①各组织招新时段不同，处于其招新时段内的组织及部门可被查询到。若意向组织/部门未出现在列表，请耐心等待其招新启动/联系意向组织负责人确认",
            "②每个组织可申报至多两个志愿部门，已申报部门不可变更/取消",
            "③申报成功后请耐心等待面试通知，通知将通过短信/邮件发送，请确认当前简历中联系方式正确。",
            "④如因简历联系方式有误等原因导致无法成功收到面试通知，请联系电话：15209934027解决"
        ]

        let showNotice = ref(false);
        const toastNotice = function () {
            showNotice.value = true;
        }

        const submit = async function () {
            // 查看是否存在简历
            const resume = await GetResume();
            if (!resume.data.exist) {
                message.warning("请先填写简历！");
                return;
            }
            const params = {
                org_id: chosenOrgIndex.value,
                dep_id: chosenDepIndex.value
            }
            const res = await SubmitRegistration(params);
            message.success(res.data.msg);
        }

        return {
            pageHead,
            shouldKnowImg1,
            shouldKnowImg2,
            showNotice,
            noticeText,
            userImg,
            orgs,
            deps,
            chosenOrgIndex,
            chosenDepIndex,
            userInfo,
            toastNotice,
            submit,
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

.applyAnimate {
    opacity: 0;
}

.blocks {
    position: relative;
    top: 0.1vh;
    width: 100%;
    height: 70vh;

    .bg {
        position: absolute;
        top: -2vh;
        // height: 38vh;
        height: 30vh;
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

    .shouldKnow {
        @include block();
        background: #FFA2A2;
        @include flex(row, );
        cursor: pointer;

        span {
            font-size: 0.7rem;
            color: white;

            @media (min-width:1280px) {
                font-size: 0.18rem;
            }
        }

        .ornaments {
            height: 60%;
            position: absolute;
            left: 0;
            bottom: 20%;
            transform: rotate(190deg);
        }

        .arrow {
            @include flex(row, );
            position: absolute;
            right: 10%;
            height: 7vh;
            width: 7vh;
            background: #FC8080;
            border-radius: 50%;
            transform: rotate(45deg);

            img {
                height: 60%;
            }
        }
    }

    @mixin icon() {
        @include flex(row, );
        font-size: 0.8rem;
        color: white;
        position: absolute;
        right: 10%;
        height: 7vh;
        width: 7vh;
        border-radius: 50%;

        @media (min-width:1280px) {
            font-size: 0.18rem;
        }
    }

    @mixin select() {
        height: auto;
        width: 50%;
        position: absolute;
        left: 15%;

        p {
            color: white;
            font-size: 0.5rem;

            @media (min-width:1280px) {
                font-size: 0.1rem;
            }
        }
    }

    .org {
        @include block();
        @include flex(row, true);
        justify-content: space-around;
        background: #9d89ff;

        .select {
            @include select();
        }

        .icon {
            @include icon;
            background: #7960f6;
        }
    }

    .dep {
        @include block();
        @include flex(row, true);
        justify-content: space-around;
        background: #62bdfe;

        .select {
            @include select();
        }

        .icon {
            @include icon;
            background: #0093FA;
        }
    }
}

.button {
    height: 25vh;
    width: 100%;
    @include flex(row, true);

    button {
        height: 30%;
        width: 40%;
        font-size: 2vh;
        letter-spacing: 0.2px;
    }
}
</style>