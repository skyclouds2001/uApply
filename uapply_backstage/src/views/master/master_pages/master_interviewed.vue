<!-- 管理员界面，新投递数据页 -->
<template>
    <div class="page">
        <!-- 数据表 -->
        <div class="new">
            <div class="title">
                <h1>本部门已面试数据</h1>
            </div>
            <div class="buttons">
                <n-button type="info" @click="eliminate">淘汰且通知</n-button>
                <n-button type="success" @click="interview" style="position:relative;left:30%;">发送面试通知</n-button>
                <n-button type="info" @click="pass">录取通过</n-button>
            </div>
            <div class="table">
                <n-data-table :columns="columns" :data="data" size="large" :row-key="(row) => row.uid"
                    :on-update:checked-row-keys="select"></n-data-table>
            </div>
        </div>
    </div>
    <!-- 弹窗 -->
    <toast v-if="isShow" :toast="component" @isToast="isToast" @confirm="confirm"></toast>
    <n-modal v-model:show="showModal" title="确认通过" content="点击确认将选中人员通过面试" preset="dialog" positive-text="确认"
        negative-text="算了" @positive-click="onPositiveClick" />
</template>

<script>
import { NButton, NDataTable, NModal } from 'naive-ui'
import { GetAllInterviewedOnce, PassAndAdmission } from '../../../utils/apis/api'
import { onMounted, ref } from '@vue/runtime-core'
import Toast from '../../../components/generic/toast/toast.vue'
import { getCurrentInstance } from 'vue'
import { message } from '../../../utils/message/message'

export default {
    components: {
        NButton,
        NDataTable,
        Toast,
        NModal
    },
    setup() {
        // 表头
        const columns = [
            {
                type: 'selection',
                disabled(row) {
                    return row.second_status != "新投递" && row.second_status != "已通知";
                }
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
                title: "面评评价",
                key: "eva",
                width: 300,
            },
            {
                title: "个人介绍",
                key: "intro",
                width: 300,
            },
            {
                title: "评分",
                key: "mark",
                width: 70,
            },
            {
                title: "状态字段",
                key: "second_status",
                width: 150,
            }
        ]

        let showModal = ref(false);

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

        // 淘汰且通知
        const eliminate = function () {
            proxy.$sessionData("set", "uids", uids.value);
            proxy.$sessionData("set", "eliminate_index", 2);
            component.value = "Eliminate";
            isToast();
        }

        // 发送面试信息
        const interview = function () {
            proxy.$sessionData("set", "uids", uids.value);
            // 通过二面
            proxy.$sessionData("set", "inter_index", 2);
            component.value = "Interview";
            isToast();
        }

        const confirm = function () {
            setTimeout(()=>{
                init();
            },2000);
        }

        const init = async function () {
            const res = await GetAllInterviewedOnce();
            if (res.data.users.length == 0) {
                return;
            }
            // console.log(temp)
            res.data.users.forEach((item, index, arr) => {
                if (item.second_status == 0) arr[index].second_status = "新投递";
                if (item.second_status == 1) arr[index].second_status = "已通知";
                if (item.second_status == 2) arr[index].second_status = "已面试";
                if (item.second_status == 3) arr[index].second_status = "已通过";
                if (item.second_status == 4) arr[index].second_status = "已淘汰";
            });
            data.value = res.data.users;
        }

        // 批量通过面试
        const pass = function () {
            if (uids.value.length == 0) {
                message.warning("请先选中！");
                return;
            }
            showModal.value = true;
        }

        const onPositiveClick = async function () {
            const params = {
                "uid": uids.value
            }
            const res = await PassAndAdmission(params);
            if (res.data.code > 1000) {
                message.error(res.data.msg);
                init();
                return;
            }
            message.success(res.data.msg);
            init();
        }


        // 不是很需要刷新界面，所以把获取数据直接放在钩子里面
        onMounted(() => {
            init();
        })

        return {
            columns,
            data,
            isShow,
            component,
            showModal,
            select,
            eliminate,
            interview,
            isToast,
            confirm,
            pass,
            onPositiveClick
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
    box-shadow: 1px 1px 5px 0px #c5a0ff;
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
