<template>
    <div class="delete_admission">
        <h1>是否确认删除选中成员？</h1>
        <h2>删除后所选成员将从部员名单中移除，请自行通知目标同学</h2>

    </div>
</template>

<script>
import { getCurrentInstance } from 'vue';
import { DeleteAdmission } from "../../../../utils/apis/api.js"
import { message } from '../../../../utils/message/message.js';

export default {
    props: {
        index: Number
    },
    setup(props, context) {
        // 实例对象
        const { proxy } = getCurrentInstance();

        // 删除已录取人员
        const confirm = async function () {
            const uids = proxy.$sessionData("get", "uids");
            if (uids.length == 0) {
                message.warning("请选择后进行操作");
            } else {
                // 这个时候已经选中了
                const params = {
                    "uids": uids
                }
                // 执行删除操作
                const res = await DeleteAdmission(params);
                // console.log(params);
                if (res.data.msg > 1000) {
                    message.error(res.data.msg);
                } else {
                    message.success(res.data.msg);
                }
                return 0;
            }
            return -1;
        }

        return {
            confirm,
        }
    }
}
</script>

<style lang="scss" scoped>
.delete_admission {
    height: 60%;
    width: 100%;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    flex-direction: column;
}
</style>