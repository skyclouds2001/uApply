<template>
    <div class="super_config">
        <div class="add_config">
            <add-config v-for="item in add_type" :key="item.title" :add_type="item" @toast="toast" ref="config"></add-config>
        </div>
        <div class="department_config">
            <department v-for="item in departments" :key="item.Id" :info="item" @toast="toast"></department>
        </div>
    </div>
    <toast v-if="isShow" :toast="component" @isToast="isToast" @confirm="confirm"></toast>
</template>

<script>
import { onMounted, ref } from 'vue';
import AddConfig from "../../../components/super_master/super_config/add.vue";
import Department from "../../../components/super_master/super_config/department.vue";
import { GetAllDepartmentInfo } from "../../../utils/apis/api";
import Toast from '../../../components/generic/toast/toast.vue';

export default {
    components: {
        AddConfig,
        Department,
        Toast
    },
    setup() {
        // 最上面两个配置栏的循环信息
        const add_type = [
            {
                type: "date",
                title: "招新日期设置",
                text: "配置组织招新日期"
            },
            {
                type: "dep",
                title: "新增部门",
                text: "配置部门信息"
            }
        ];
        // 部门信息
        // 可先存放在localstorage里面，等到下次加载时，首先从内存中加载这一块，然后和请求做检验，如果不一样再渲染
        let departments = ref([]);

        // 是否显示toast
        let isShow = ref(false);

        // toast的种类
        let component = ref("");

        // config实例
        let config = ref(null);

        // 获取并返回全部部门信息，初始化/更新页面
        const set_departments = async function () {
            const res = await GetAllDepartmentInfo();
            departments.value = JSON.parse(res.request.response).departments;
        }

        // 需要显示或者不显示toast的时候，都需要调用此函数，不允许出现功能重复的函数
        const isToast = function () {
            // 切换toast状态
            isShow.value = !isShow.value;
        }

        // 当日期、部门按钮被点击后，根据点击种类显示toast
        const toast = function (e) {
            isToast();
            // 设置toast中的内容
            component.value = e;
        }

        // 点击确认按钮之后，触发页面更新
        // 为什么要用[0]取出？因为ref部署在v-for循环出的对象身上
        const confirm = function (e) {
            if (e === "AddDate") {
                config.value[0].reset();
            }
            else if (e === "AddDep" || e === "DeleteDep" || e === "EditDep") {
                set_departments();
            }
        }

        onMounted(() => {
            set_departments();
        })

        return {
            add_type,
            departments,
            isShow,
            component,
            config,
            confirm,
            isToast,
            toast,
            set_departments
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($direction: row, $isAround: true) {
    display: flex;
    flex-direction: $direction;

    @if ($isAround) {
        justify-content: center;
        align-items: center;
    }

    @else {
        justify-content: left;
        align-items: center;
    }

    flex-wrap: wrap;
}

.super_config {
    height: 88.5%;
    width: 100%;
    margin-bottom: 5%;
    @include flex(column, false);
    flex-wrap: nowrap;
    overflow: auto;

    .add_config {
        width: 94%;
        padding-top: 3%;
        @include flex(row, false);
    }

    .department_config {
        width: 94%;
        margin-bottom: 2%;
        @include flex(row, false)
    }

}
</style>
