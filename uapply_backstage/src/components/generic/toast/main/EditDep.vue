<template>
    <div class="add_dep">
        <div class="dep_item">
            <h2> &nbsp; 部门全称</h2>
            <n-input size="large" placeholder="请输入10位以内中英文字符的部门全称" clearable v-model:value="params.name"></n-input>
        </div>
        <div class="dep_item">
            <h2> &nbsp; 配置部门管理账号</h2>
            <n-input size="large" placeholder="可包含12~18位英文字符/数字/@" clearable v-model:value="params.account"></n-input>
        </div>
        <div class="dep_item">
            <h2> &nbsp; 配置部门管理密码</h2>
            <n-input size="large" placeholder="可包含12~18位英文字符/数字/@" clearable v-model:value="params.password"></n-input>
        </div>
    </div>
</template>

<script>
import { NInput } from 'naive-ui'
import { reactive } from '@vue/reactivity'
import { message } from '../../../../utils/message/message'
import { EditDepartment } from "../../../../utils/apis/api"
import { getCurrentInstance, onMounted } from 'vue'

export default {
    components: {
        NInput
    },
    setup() {
        // 保存配置信息
        const params = reactive({
            name: "",
            account: "",
            password: ""
        })

        // 检测函数
        const isNull = function () {
            if (params.name == "") {
                message.warning("请填写部门名称！");
            } else if (params.account == "") {
                message.warning("请填写部门账号！");
            } else if (params.password == "") {
                message.warning("请填写部门密码!");
            } else {
                return 0;
            }
            return -1;
        }

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 当前点击部门的信息
        const department = proxy.$sessionData("get", "department");

        // 编辑部门
        const edit = async function () {
            console.log(params)
            const res = await EditDepartment(department.id, params);
            // console.log(res);
            if (res.data.code == 1005) {
                res.data.msg == "部门名字非法，请更换" ? message.warning(res.data.msg) : message.warning("账号重复，请更换");
            } else if (res.data.code == 1002) {
                message.warning(res.data.msg);
            } else {
                message.success("成功修改" + params.name);
                return 0;
            }
            return -1;
        }

        // 确认函数
        const confirm = async function () {
            // // 如果不是空则
            // if (isNull() == -1) {
            //     return -1;
            // }
            let res = 0;
            // 如果是编辑部门
            res = await edit();
            console.log(res);
            if (res != 0) {
                return -1;
            } else {
                return 0;
            }
        }

        // 初始化数据
        onMounted(() => {
            params.name = department.name;
            params.account = department.account;
            params.password = department.password;
        })

        return {
            params,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
.add_dep {
    height: 85%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;

    .dep_item {
        width: 55%;
        height: 25%;
    }
}
</style>