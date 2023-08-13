<template>
    <div class="add_date">
        <div class="data_picker">
            <n-date-picker clearable size="large" type="datetimerange" start-placeholder="请选择开始日期"
                end-placeholder="请选择开始日期" v-model:value="date"></n-date-picker>
        </div>
    </div>
</template>

<script>
import { NDatePicker } from 'naive-ui'
import { getCurrentScope, ref } from '@vue/reactivity'
import { SetRecruitmentDate } from '../../../../utils/apis/api';
import { message } from '../../../../utils/message/message';
import { getCurrentInstance } from '@vue/runtime-core';
import toTime from "../../../../utils/dateHandle/dateHandle"

export default {
    components: {
        NDatePicker
    },
    setup(props, context) {
        // 选中的时间
        let date = ref(null);

        // 保存VUE实例对象
        const { proxy } = getCurrentInstance();


        // toast确认后的需要进行的操作等等
        const confirm = async function () {
            // 如果未选择日期，提示一下后直接返回
            if (date.value == null) {
                message.info("请选择日期！")
                return -1
            };
            // 配置请求参数
            const param = {
                start: date.value[0] / 1000,
                end: date.value[1] / 1000
            }
            // console.log(param);
            // 发送请求
            const res = await SetRecruitmentDate(param);

            if (res.data.msg == "修改招新时间成功") {
                // 通知
                message.success("配置招新时间成功!");
                // 保存人看的懂的时间
                const temp = {
                    start: toTime(param.start),
                    end: toTime(param.end),
                }
                // 将日期保存到本地浏览器
                proxy.$localData("set", "date", temp);
                return;
            }
            message.error("出错啦，请重新配置！");
            return -1;
        }

        return {
            date,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.add_date {
    height: 85%;
    width: 100%;

    .data_picker {
        width: 70%;
        position: relative;
        left: 15%;
        top: 20%;
    }
}
</style>