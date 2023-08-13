<template>
    <div class="interview">
        <div class="item">
            <h3>选择面试时间</h3>
            <n-date-picker type="datetimerange" start-placeholder="选择面试开始时间" end-placeholder="选择面试结束时间" size="large"
                v-model:value="config.time"></n-date-picker>
        </div>
        <div class="item">
            <h3>面试地点说明</h3>
            <n-input size="large" placeholder="通知面试地点" v-model:value="config.inter_address"></n-input>
        </div>
        <div class="item">
            <h3>负责人联系方式</h3>
            <n-input size="large" placeholder="负责人联系方式" v-model:value="config.contact_phone"></n-input>
        </div>
    </div>
</template>

<script>
import { NInput, NDatePicker } from 'naive-ui';
import { onMounted, reactive, ref, toRef } from 'vue';
import { getCurrentInstance } from 'vue';
import { PassInterview, SetRecruitmentConfig, GetRecruitmentConfig, PassAndAdmission, SendNotification } from '../../../../utils/apis/api';
import { message } from '../../../../utils/message/message';

export default {
    components: {
        NInput,
        NDatePicker
    },
    setup() {
        // 日期值
        let time = ref(null);

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 第几轮面试
        let turn = 1;

        // 面试配置
        let config = reactive({
            time: null,
            inter_address: "",
            contact_phone: ""
        })


        // 发送面试通知
        const send = async function () {
            // 先取出需要通知的uid数组
            const uids = proxy.$sessionData("get", "uids");
            // console.log(uids);
            // 未选中则提示
            if (uids.length == 0) {
                message.warning("请选中后进行操作！");
                return -1;
            }

            await configRecruitment();

            // await passRecruitment(uids);
            await notify(uids);
        }

        const notify = async function (uids) {
            const params = {
                "uids": uids
            }
            const res = await SendNotification(turn, params);
            // console.log(res);
            if (res.data.code > 1000) {
                message.error(res.data.msg);
                return;
            }
            message.success(res.data.msg);
        }

        // 通过面试现在到另一个地方去了
        // 通过面试
        const passRecruitment = async function (uids) {
            // 请求体
            const params = {
                "uid": uids
            }

            // 发送通过请求
            const res = null;
            if (turn == 1) {
                res = await PassInterview(turn, params);
            } else if (turn == 2) {
                res = await PassAndAdmission(params);
            } else {
                message.error("出错啦！请重试！");
            }

            // 错误处理
            if (res.data.code > 1000) {
                message.error(res.data.msg);
                return -1;
            }
            message.success(res.data.msg);
            return 0;
        }

        // 配置面试信息
        const configRecruitment = async function () {
            // 请求体
            const params = {
                "start": config.time[0] / 1000,
                "end": config.time[1] / 1000,
                "inter_address": config.inter_address,
                "contact_phone": config.contact_phone
            }

            const res = await SetRecruitmentConfig(turn, params);

            // 错误处理
            if (res.data.code > 1000) {
                message.error(res.data.msg);
                return -1;
            }
            message.success(res.data.msg);
            return 0;
        }

        // 获取并自动填充面试配置
        const init = async function () {
            const res = await GetRecruitmentConfig(turn);
            if (res.data.code > 1000) {
                message.success(res.data.msg);
                return;
            }
            // 更新或初始化
            config.inter_address = res.data.inter_address;
            config.contact_phone = res.data.contact_phone;
            config.time = [res.data.start * 1000, res.data.end * 1000];
        }

        onMounted(() => {
            // 先初始化第几轮面试
            turn = proxy.$sessionData("get", "inter_index");
            init();
        })

        const confirm = async function () {
            const res = send();
            if (res == -1) {
                return -1;
            }
            return 0;
        }

        return {
            time,
            config,
            confirm
        }
    }

    // // 设置面试配置
    // const setConfig = async function () {
    //     // 面试信息配置API请求体
    //     const configParam = {
    //         first_start: time1.value[0],
    //         first_end: time1.value[1],
    //         second_start: time2.value[0],
    //         second_end: time2.value[1],
    //         inter_address: place.value,
    //         contact_phone: contact.value
    //     }
    //     const configRes = await SetRecruitmentConfig(configParam);
    //     if (configRes.data.msg == "设置成功") {
    //         message.success("请求成功");
    //         // 如果设置成功，先进行缓存，下次直接使用
    //         proxy.$localData("set", "time1", time1.value);
    //         proxy.$localData("set", "time2", time2.value);
    //         proxy.$localData("set", "place", place.value);
    //         proxy.$localData("set", "contact", contact.value);
    //     } else {
    //         message.error(configRes.data.msg);
    //         return -1;
    //     }
    // }

}
</script>

<style lang="scss" scoped>
@mixin flex($ver: center) {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: $ver;
}

.interview {
    height: 70%;
    width: 100%;
    position: relative;
    top: 5%;
    @include flex();

    .item {
        @include flex(flex-start);
        width: 50%;
    }
}
</style>


