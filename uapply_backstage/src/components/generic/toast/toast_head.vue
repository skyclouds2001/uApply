<template>
    <div class="toast_head">
        <h2>{{ title }}</h2>
        <div class="buttons">
            <n-button round type="info" @click="cancel">取 消</n-button>
            <n-button round type="success" @click="confirm">确 认</n-button>
        </div>
    </div>
</template>

<script>
import { onMounted, ref } from '@vue/runtime-core'
import { NButton } from 'naive-ui'

export default {
    props: {
        toast:String
    },
    components:{
        NButton
    },
    setup(props, context) {
        // toast头部标题
        let title = ref("")

        // 取消toast，就是隐藏toast
        const cancel = function(){
            context.emit("hidden");
        }

        // 确认操作
        const confirm = function(){
            context.emit("confirm");
        }

        onMounted(() => {
            // 初始化toast的时候进行标题配置
            console.log(props.toast)
            switch (props.toast) {
                case "AddDate":
                    title.value = "招新日期配置";
                    break;
                case "AddDep":
                    title.value = "部门信息配置";
                    break;
                case "DeleteDep":
                    title.value = "是否确认删除";
                    break;
                case "EditDep" :
                    title.value = "部门信息编辑";
                    break;
                case "AddInterviewer" :
                    title.value = "面试官信息配置";
                    break;
                case "EditInterviewer" :
                    title.value = "面试官信息编辑";
                    break;
                case "DeleteInterviewer" :
                    title.value = "是否删除面试官";
                    break;
                case "Eliminate" : 
                    title.value = "面试信息配置-淘汰"
                    break;
                case "Interview" : 
                    title.value = " 面试信息配置-面试通知"
                    break;
                case "DeleteAdmission" : 
                    title.value = "是否确认删除"
                case "AddAdmission" : 
                    title.value = "成员信息配置"
            }
        })

        return {
            title,
            cancel,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.toast_head {
    height: 15%;
    width: 94%;
    background: rgb(255, 255, 255);
    border-bottom: 1px solid rgb(192, 192, 192);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 3%;

    h2{
        font-size: 1.8rem;
        
        @media (max-width:1440px) {
            font-size: 1.6rem;
        }
        @media (max-width:1280px) {
            font-size: 1.5rem;
        }
    }

    .buttons{
        height: 100%;
        display: flex;
        align-items: center;
        button{
            height: 70%;
            width: 120px;
            margin-left: 20px;
            font-size: 1.2rem;
        }
    }
}
</style>