<template>
    <div class="page">
        <div class="eliminate">
            <div class="title">
                <h1>本部门已淘汰数据</h1>
            </div>
            <div class="table">
                <n-data-table :columns="columns" :data="data" size="large"></n-data-table>
            </div>
        </div>
    </div>
</template>

<script>
import { NDataTable } from 'naive-ui';
import { GetAllEliminate } from '../../../utils/apis/api';
import { onMounted, ref } from 'vue';

export default {
    components: {
        NDataTable
    },
    setup() {
        const columns = [
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
                width:150
            },
            {
                title: "面评评价",
                key: "eva",
                width: 350,
            },
            {
                title: "个人介绍",
                key: "intro",
                width: 350,
            },
            {
                title: "评分",
                key: "mark",
                width: 80,
            },
            // {
            //     title: "状态字段",
            //     key: "status",
            //     width: 150,
            // }
        ];

        let data = ref([]);

        onMounted(async () => {
            const res = await GetAllEliminate();
            data.value = res.data.users;
        })

        return {
            columns,
            data
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex() {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
}

.eliminate {
    height: 95%;
    width: 95%;
    margin-top: 2%;
    background: white;
    box-shadow: 1px 1px 5px 0px #c5a0ff;
    border-radius: 1rem;
    padding: 0.1%;
    @include flex();

    .title {
        height: 10%;
        width: 100%;
        @include flex();
    }

    .table {
        height: 90%;
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
