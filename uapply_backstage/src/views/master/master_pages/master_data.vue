<template>
    <div class="page">
        <div class="data_block">
            <total v-for="item in data" :key="item.text" :data="item"></total>
        </div>
        <div class="useless">
            <img :src="img" alt="">
            <br>
            <h1>辛苦管理员大大，这个学年还要一起努力嗷~</h1>
        </div>
    </div>
</template>

<script>
import { onMounted, ref } from 'vue'
import { GetDepartmentData } from '../../../utils/apis/api'
import Total from '../../../components/super_master/super_data/total.vue';

export default {
    components: {
        Total
    },
    setup() {
        const img = require("@/assets/imgs/没啥用.png");

        // 招新数据
        let data = ref([
            {
                Id: 0,
                text: "部门招新投递总数",
                data: 0
            },
            {
                Id: 1,
                text: "部门招新通过总数",
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

        onMounted(async () => {
            const res = (await GetDepartmentData()).data;
            data.value[0].data = res.reg_sum;
            data.value[1].data = res.pass_sum;
            data.value[2].data = res.pass_male;
            data.value[3].data = res.pass_female;
        })

        return {
            data,
            img
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex() {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.data_block {
    width: 95%;
    height: 25%;
    @include flex();
}

.useless {
    height: 70%;
    width: 95%;
    border-radius: 1rem;
    background: white;
    box-shadow: 1px 1px 5px 0 #c5a0ff;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    h1{
        color: #bd94fc;
    }

    img {
        height: 60%;
    }
}
</style>
