<template>
    <div class="add_admission">
        <div class="item">
            <h3>成员姓名</h3>
            <n-input v-model:value="params.name" size="large" placeholder="请输入姓名"></n-input>
        </div>
        <div class="item">
            <h3>成员性别</h3>
            <n-select :options="sexs" v-model:value="params.sex"></n-select>
        </div>
        <div class="item">
            <h3>ID</h3>
            <n-input v-model:value="params.uid" size="large" placeholder="请输入ID"></n-input>
        </div>
        <div class="item">
            <h3>联系电话</h3>
            <n-input v-model:value="params.phone_num" size="large" placeholder="请输入联系电话"></n-input>
        </div>
        <div class="item">
            <h3>邮箱</h3>
            <n-input v-model:value="params.email" size="large" placeholder="请输入邮箱地址"></n-input>
        </div>
        <div class="item">
            <h3>学号</h3>
            <n-input v-model:value="params.student_num" size="large" placeholder="请输入学号"></n-input>
        </div>
        <div class="item">
            <h3>楼号</h3>
            <n-input v-model:value="params.address_num" size="large" placeholder="请输入楼号(选填)"></n-input>
        </div>
        <div class="item">
            <h3>大类</h3>
            <n-input v-model:value="params.major" size="large" placeholder="请输入大类/专业(选填)"></n-input>
        </div>
        <div class="intro">
            <br>
            <h3>个人简介/评价</h3>
            <n-input type="textarea" placeholder="请填写个人简介/评价" v-model:value="params.intro"></n-input>
        </div>
    </div>
</template>

<script>
import { reactive, ref } from '@vue/reactivity';
import { NInput, NSelect } from 'naive-ui';
import { message } from '../../../../utils/message/message';
import { AddAdmission } from '../../../../utils/apis/api';

export default {
    components: {
        NInput,
        NSelect
    },
    setup() {
        // 参数
        let params = reactive({
            name: "",
            sex: "男",
            phone_num: "",
            email: "",
            student_num: "",
            address_num: "",
            major: "",
            uid: "",
            intro: ""
        });

        // 性别选项卡
        const sexs = [
            {
                label: "男",
                value: "男"
            },
            {
                label: "女",
                value: "女"
            }
        ]

        const confirm = async function () {
            params.uid = Number(params.uid);
            const res = await AddAdmission(params);
            console.log(res);
            if (res.data.code >= 1002){
                message.error(res.data.msg);
            } else {
                message.success(res.data.msg);
                return 0;
            }
            return -1;
        }

        return {
            params,
            sexs,
            confirm
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($row: center) {
    display: flex;
    justify-content: $row;
    align-items: center;
    flex-wrap: wrap;
}

.add_admission {
    height: 85%;
    width: 100%;
    overflow: auto;
    @include flex(space-around);

    .item {
        height: 20%;
        width: 35%;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        justify-content: center;
        // @include flex(flex-start);
        // align-items: center;
        // position:relative;
        // left: 5%;
    }

    #big {
        position: relative;
        right: 25%;
    }

    .intro {
        width: 80%;
        height: 40%;
    }
}
</style>