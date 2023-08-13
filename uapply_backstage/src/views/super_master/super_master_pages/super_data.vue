<template>
    <div class="super_data">
        <div class="data_block">
            <data-total v-for="item in org" :key="item.text" :data="item"></data-total>
        </div>
        <div class="data_table">
            <div class="table_title">
                组织内各部门招新数据
            </div>
            <div class="table">
                <n-data-table :columns="columns" :data="dep" size="large" :bordered="false" striped
                    :row-class-name="rowClassName" :pagination="pagination"></n-data-table>
            </div>
        </div>
    </div>
</template>

<script>
import { onMounted, ref } from 'vue';
import { GetOrganizationRecruitmentDate } from "../../../utils/apis/api"
import DataTotal from "../../../components/super_master/super_data/total.vue"
import { NDataTable } from 'naive-ui';

export default {
    components: {
        DataTotal,
        NDataTable
    },
    setup() {
        let org = ref([
            {
                Id: 0,
                text: "组织招新投递总数",
                data: 0
            },
            {
                Id: 1,
                text: "组织招新通过总数",
                data: 0
            },
            {
                Id: 2,
                text: "招新通过男生总数",
                data: 0
            },
            {
                Id: 3,
                text: "招新通过女生总数",
                data: 0
            }
        ]);
        let dep = ref([]);

        const columns = [
            {
                title: "部门",
                key: "name",
                width: 80,
                className: "tableStyle"
            },
            {
                title: "投递总数",
                key: "sum",
                width: 100
            },
            {
                title: "通过总数",
                key: "pass",
                width: 100
            },
            {
                title: "通过男生总数",
                key: "male",
                width: 100
            },
            {
                title: "通过女生总数",
                key: "female",
                width: 100
            }
        ]

        // 拿到组织下所有的招新数据
        const getData = function () {
            return GetOrganizationRecruitmentDate();
        }

        // 初始化数据
        onMounted(async () => {
            // 发送请求解析数据
            const allData = (await getData()).data;
            dep.value = allData.dep;
            // 多了一个部门ID，先删了
            delete allData.org.Id;
            delete allData.org.name;
            // 创建 total data 遍历数组
            Object.keys(allData.org).forEach((key, index) => {
                org.value[index].data = allData.org[key];
            })
        })

        return {
            org,
            dep,
            columns,
            pagination: { pageSize: 10 },
            rowClassName() {
                return "tableStyle"
            }
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($row: space-between, $ver: center) {
    display: flex;
    justify-content: $row;
    align-items: $ver;
}

.super_data {
    @include flex(center);
    height: 88.5%;
    width: 100%;
    margin-bottom: 5%;
    flex-wrap: wrap;

    .data_block {
        width: 95%;
        height: 25%;
        @include flex();
    }

    .data_table {
        width: 95%;
        height: 75%;
        @include flex(center, column);
        flex-wrap: wrap;
        background: white;
        box-shadow: 2px 2px 5px 0 #c5a0ff;
        border-radius: 1rem;

        .table_title {
            width: 100%;
            height: 10%;
            @include flex(center);
            font-size: 1.8rem;
            font-weight: 600;
        }

        .table {
            height: 85%;
            width: 100%;

            .n-data-table-th {
                text-align: center;
            }

        }
    }

}


// 深度追查
:deep(.n-data-table-th){
    text-align: center;
    font-size: 120%;
    font-weight: 600;
    background: #f7f1ff;
}

:deep(.n-data-table-td){
    text-align: center;
}

</style>
