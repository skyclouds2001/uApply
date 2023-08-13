<template>
    <div class="delete_dep">
        <h1>是否确认删除</h1>
        <h2>删除后该部门信息及其相关数据将被删除清空，请确认是否进行删除</h2>
    </div>
</template>

<script>
import { onMounted } from '@vue/runtime-core'
import { getCurrentInstance } from 'vue';
import { DeleteDepartment } from "../../../../utils/apis/api.js"
import { message } from '../../../../utils/message/message.js';

export default {
    props: {
        index: Number
    },
    setup(props, context) {
        // 实例对象
        const { proxy } = getCurrentInstance();

        // 删除部门
        const confirm = async function () {
            const res = await DeleteDepartment(proxy.$sessionData("get", "ID"));
            if (res.data.code == 1004 || res.data.code == 1005){
                message.error(res.data.msg);
            }
            console.log(res);
            return 0;
        }

        return {
            confirm,
        }
    }
}
</script>

<style lang="scss" scoped>
.delete_dep {
    height: 60%;
    width: 100%;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    flex-direction: column;
}
</style>