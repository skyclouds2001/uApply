<template>
    <div class="deleteInterviewer">
        <h1>是否确认删除</h1>
        <h2>删除后该用户将不能使用面试官端-小程序功能，请确认是否删除</h2>
    </div>
</template>

<script>
import { DeleteInterviewer } from '../../../../utils/apis/api';
import { getCurrentInstance } from 'vue';
import { message } from '../../../../utils/message/message';

export default {
    setup(props, context) {
        // 实例对象
        const { proxy } = getCurrentInstance();

        const confirm = async function () {
            const interviewer = proxy.$sessionData("get", "interviewer");
            const res = await DeleteInterviewer(interviewer.uid);
            message.info(res.data.msg);
            return 0;
        }

        return {
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.deleteInterviewer {
    height: 60%;
    width: 100%;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    flex-direction: column;
}
</style>