<template>
    <div class="eliminate">
        <n-input type="textarea" class="input" placeholder="请输入通知内容" :disabled="true" v-model:value="value"></n-input>
    </div>
</template>

<script>
import { NInput } from 'naive-ui'
import { ref } from '@vue/reactivity'
import { getCurrentInstance, onMounted } from 'vue';
import { EliminateInterviewee } from '../../../../utils/apis/api';
import { message } from '../../../../utils/message/message';

export default {
    components: {
        NInput
    },
    setup() {
        // 通知内容
        const value = ref(`xx同学你好：

感谢同学投递【xxxx部门】，经评估认为您暂时与本部门需求不匹配，

也预祝同学成功加入其他心仪部门 ♥

（如非本人请忽略本条信息）

(功能禁用！目前淘汰不发送面试通知！)`);

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 第几轮
        let turn = 0;

        const confirm = async function () {
            const uids = proxy.$sessionData("get", "uids");
            const param = {
                uid: uids
            }
            if (uids.length == 0) {
                message.warning("请选择后操作！");
                return 0;
            }
            // console.log(param)
            const res = await EliminateInterviewee(turn,param);
            if (res.data.msg == "淘汰成功") {
                message.success("淘汰成功！");
                return 0;
            }
            message.error(res.data.msg);
            return -1;
        }

        onMounted(() => {
            turn = proxy.$sessionData("get", "eliminate_index");
            // console.log(turn);
        })

        return {
            value,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.eliminate {
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;

    .input {
        height: 70%;
        width: 90%;
        position: relative;
        bottom: 10%;
        font-size: large;
        border-radius: 1rem;
    }
}
</style>