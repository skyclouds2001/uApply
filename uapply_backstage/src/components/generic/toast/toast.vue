<template>
    <div class="toast">
        <toast-head :toast="toast" @hidden="hidden" @confirm="confirm"></toast-head>
        <component :is="toast" ref="son"></component>
    </div>
</template>

<script>
import ToastHead from './toast_head.vue'
import AddDate from './main/AddDate.vue'
import AddDep from './main/AddDep.vue'
import DeleteDep from './main/DeleteDep.vue'
import EditDep from './main/EditDep.vue'

import AddInterviewer from './main/AddInterviewer.vue'
import DeleteInterviewer from './main/DeleteInterviewer.vue'
import EditInterviewer from './main/EditInterviewer.vue'
import Eliminate from './main/Eliminate.vue'
import Interview from './main/Inter.vue'
import DeleteAdmission from './main/DeleteAdmission.vue'
import AddAdmission from './main/AddAdmission.vue'

import { ref } from '@vue/runtime-core'

export default {
    components: {
        ToastHead,
        AddDate,
        AddDep,
        DeleteDep,
        EditDep,
        AddInterviewer,
        DeleteInterviewer,
        EditInterviewer,
        Eliminate,
        Interview,
        DeleteAdmission,
        AddAdmission
    },
    props: {
        // 当前toast的种类
        toast: String
    },
    setup(props, context) {
        // 隐藏toast
        const hidden = function () {
            context.emit("isToast");
        }

        // 子组件实例
        let son = ref(null);

        // 当确认按钮被点击,执行确认操作
        // 由头部子组件触发
        const confirm = async function () {
            // 调用主要内容组件，发送请求、保存数据
            const res = await son.value.confirm();
            // 如果子组件返回-1，说明发生了错误
            if (res == -1) {
                return;
            }
            console.log(props.toast);
            // 请求成功，触发父组件确认功能
            context.emit("confirm", props.toast);
            // 并关闭弹窗
            hidden();
        }

        return {
            hidden,
            son,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.toast {
    background: rgb(247, 247, 247);
    border-radius: 2vh;
    width: 50rem;
    height: 28rem;
    position: fixed;
    top: 22%;
    left: 35%;
    z-index: 5;
    overflow: hidden;
    box-shadow: 1px 1px 5px 1px rgb(216, 216, 216);

    @media (max-width:1440px) {
        width: 42rem;
        height: 25rem;
    }

    @media (max-width:1366px) {
        width: 40rem;
        height: 25rem;
    }

}
</style>