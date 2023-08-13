<!-- 管理员界面，新投递数据页 -->
<template>
    <div class="page">
        <!-- 数据表 -->
        <div class="new">
            <div class="title">
                <h1>本部门已面试数据</h1>
            </div>
            <div class="buttons">
                <n-button type="warning" @click="DeleteAdmission">删除成员</n-button>
                <n-button type="info" id="add" @click="AddAdmission">新增成员</n-button>
                <n-button type="success" @click="ExportExcel">导出成员名单</n-button>
            </div>
            <div class="table">
                <n-data-table :columns="columns" :data="data" size="large" :row-key="(row) => row.uid"
                    :on-update:checked-row-keys="select"></n-data-table>
            </div>
        </div>
    </div>
    <!-- 弹窗 -->
    <toast v-if="isShow" :toast="component" @isToast="isToast" @confirm="confirm"></toast>
</template>

<script>
import { NButton, NDataTable } from 'naive-ui'
import { GetAdmission } from '../../../utils/apis/api'
import { onMounted, ref } from '@vue/runtime-core'
import Toast from '../../../components/generic/toast/toast.vue'
import { getCurrentInstance } from 'vue'
import XLSX from "xlsx";

export default {
    components: {
        NButton,
        NDataTable,
        Toast
    },
    setup() {
        // 表头
        const columns = [
            {
                type: 'selection'
            },
            {
              title: "UID",
              key: "uid",
              width: 100,
            },
            {
                title: "姓名",
                key: "name",
                width: 150,
            },
            {
                title: "性别",
                key: "sex",
                width: 100,
            },
            {
                title: "联系电话",
                key: "phone",
                width: 150,
            },
            {
                title: "邮箱",
                key: "email",
                width: 200,
            },
            {
                title: "学号",
                key: "student_num",
                width: 150,
            },
            {
                title: "个人介绍",
                key: "intro",
                width: 400
            },
            {
                title: "楼号",
                key: "address_num",
                width: 150,
            },
            {
                title: "大类",
                key: "major",
                width: 150,
            }
        ]

        // 录取人员信息
        let data = ref([]);
        // 选中人员uid
        let uids = ref([]);

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 是否显示toast
        let isShow = ref(false);

        // toast的种类
        let component = ref("");

        // 需要显示或者不显示toast的时候，都需要调用此函数，不允许出现功能重复的函数
        const isToast = function () {
            // 切换toast状态
            isShow.value = !isShow.value;
        }

        // 选择人员
        const select = function (e) {
            uids.value = e;
        }

        // 删除成员
        const DeleteAdmission = function () {
            proxy.$sessionData("set", "uids", uids.value);
            component.value = "DeleteAdmission";
            isToast();
        }

        // 新增成员
        const AddAdmission = function () {
            proxy.$sessionData("set", "uids", uids.value);
            component.value = "AddAdmission";
            isToast();
        }

        // 导出成员名单
        const ExportExcel = function () {
            const temp = XLSX.utils.json_to_sheet(data.value);
            const wb = XLSX.utils.book_new();
            XLSX.utils.book_append_sheet(wb, temp, '录取名单');
            XLSX.writeFile(wb, '录取名单.xlsx');
        }

        const confirm = function () {
            init();
            // console.log("confirm");
        }

        // 获取已录取人员名单
        const init = async function () {
            const res = await GetAdmission();
            if (res.data.sum === 0) {
                data.value = [];
                return;
            }
            data.value = res.data.users;
            console.log(data.value);
        }

        // 初始化
        onMounted(async () => {
            init();
        })

        return {
            columns,
            data,
            isShow,
            component,
            select,
            DeleteAdmission,
            AddAdmission,
            ExportExcel,
            isToast,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($direction: row, $row: space-around) {
    display: flex;
    flex-direction: $direction;
    justify-content: $row;
    align-items: center;
}

.new {
    height: 95%;
    width: 95%;
    margin-top: 2%;
    background: white;
    border-radius: 1rem;
    box-shadow: 1px 1px 5px 0 #c5a0ff;
    overflow: auto;
    @include flex(column, center);

    .title {
        height: 10%;
        width: 100%;
        @include flex(row, center);
    }

    .buttons {
        height: 12%;
        width: 95%;
        @include flex(row, space-between);
        align-items: flex-start;

        #add {
            position: relative;
            right: 30%;
        }

        :deep(.n-button) {
            padding: 2% 2%;
            border-radius: 5rem;
            font-size: large;
            letter-spacing: 1.5px;
        }
    }

    .table {
        height: 78%;
        width: 100%;
        overflow: auto;
    }
}

// 深度追查
:deep(.n-data-table-th) {
    text-align: center;
    font-size: 120%;
    font-weight: 600;
    background: #f7f1ff;
}

:deep(.n-data-table-td) {
    text-align: center;
}
</style>
