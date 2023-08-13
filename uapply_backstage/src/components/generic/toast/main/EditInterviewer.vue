<template>
    <div class="editInterviewer">
        <div class="input">
            <h3>面试官小程序用户姓名</h3>
            <n-input size="large" placeholder="请输入目标面试官姓名" v-model:value="value"></n-input>
        </div>
    </div>
</template>

<script>
import { NInput } from 'naive-ui'
import { ref } from '@vue/reactivity'
import { EditInterviewer } from '../../../../utils/apis/api';
import { getCurrentInstance } from 'vue';
import { message } from '../../../../utils/message/message';

export default {
    components: {
        NInput
    },
    setup() {
        let value = ref();

        // 实例对象
        const { proxy } = getCurrentInstance();

        const confirm = async function () {
            // 没填怎么办
            if(value.value == undefined){
                message.warning("请填写面试官姓名");
                return -1;
            }
            // 拿到当前点击的UID
            const interviewer = proxy.$sessionData("get", "interviewer");
            const param = {
                "uid":interviewer.uid,
                "name": value.value
            }
            const res = await EditInterviewer(param);
            console.log(res);
            if(res.data.msg == "修改成功"){
                message.success("成功修改面试官信息！");
            } else {
                message.error(res.data.msg);
                return -1;
            }
            return 0;
        }

        return {
            value,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.editInterviewer {
    height: 60%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;

    .input {
        width: 60%;
    }
}
</style>