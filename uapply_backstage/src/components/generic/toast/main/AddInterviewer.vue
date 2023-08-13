<template>
    <div class="AddInterviewer">
        <div class="input">
            <h3>面试官小程序用户UID</h3>
            <n-input size="large" placeholder="请输入目标面试官的数字UID" v-model:value="intID"></n-input>
        </div>
        <br>
        <div class="input">
            <h3>面试官姓名</h3>
            <n-input size="large" placeholder="请输入目标面试官的姓名" v-model:value="name"></n-input>
        </div>
    </div>
</template>

<script>
import { NInput } from 'naive-ui'
import { ref } from '@vue/reactivity'
import { AddInterviewer } from "../../../../utils/apis/api"
import { message } from '../../../../utils/message/message'

export default {
    components: {
        NInput
    },
    setup() {
        let intID = ref();
        let name = ref();

        const confirm = async function () {
            // 如果没有填写
            if(intID.value == undefined || name.value == undefined){
                message.warning("请填写信息后确认");
                return -1;
            }
            // 传参
            const param = {
                "uid": Number(intID.value),
                "name": name.value
            }
            // 发送请求和错误处理
            const res = await AddInterviewer(param);
            console.log(res);
            if (res.data.code == 1004 || res.data.code == 1005) {
                message.error(res.data.msg);
                return -1;
            }
            return 0;
        }

        return {
            confirm,
            intID,
            name
        }
    }
}
</script>

<style lang="scss" scoped>
.AddInterviewer {
    height: 70%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    flex-direction: column;

    .input {
        width: 60%;
    }
}
</style>