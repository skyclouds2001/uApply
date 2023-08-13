<template>
    <div class="data_total">
        <div class="data">
            <h2>{{ data.text }}</h2>
            <h1>{{ data.data }}</h1>
            <h4>2022-2023学年</h4>
        </div>
        <img class="img" :style="background" :src="img">
    </div>
</template>

<script>
import { onMounted, reactive, ref } from '@vue/runtime-core';
export default {
    props: {
        // 招新数据
        data: Object
    },
    setup(props,context) {
        let img = ref(null);

        let background = ref("");

        onMounted(() => {
            const ID = props.data.Id;
            switch(ID){
                case 0:
                    img.value = require("@/assets/imgs/招新总投递数.png");
                    background.value = "background:#FE4066;";
                    break;
                case 1:
                    img.value = require("@/assets/imgs/招新通过总数.png");
                    background.value = "background:#FE9940;";
                    break;
                case 2:
                    img.value = require("@/assets/imgs/招新通过总男生.png");
                    background.value = "background:#009966;";
                    break;
                case 3:
                    img.value = require("@/assets/imgs/招新通过总女生.png");
                    background.value = "background:#1017FB;";
                    break;
            }
        })

        return {
            img,
            background
        }
    }
}
</script>

<style lang="scss" scoped>
.data_total {
    height: 80%;
    width: 23%;
    background: white;
    border-radius: 1rem;
    display: flex;
    justify-content: space-around;
    align-items: center;

    .data{
        height: 80%;
        padding-left: 2%;
        display: flex;
        flex-direction: column;
        justify-content: space-around;

        h1{
            font-size: 3rem;

            @media (max-width:1440px) {
                font-size: 2rem;
            }

            @media (min-width:2048px) and (max-width:2560px) {
                font-size: 4rem;
            }

            @media (min-width:2560px) {
                font-size: 4.5rem;
            }
        }

        h2{
            @media (max-width:1440px) {
                font-size: 1rem;
            }
            @media (min-width:2048px) {
                font-size: 1.5rem;
            }
        }

        h4{
            color:rgb(146, 146, 146);

            @media (min-width:1920px) {
                font-size: 1.2rem;
            }
        }
    }

    .img{
        height: 28%;
        padding: 3%;
        border-radius: 50%;
        position: relative;
        bottom: 10%;
    }
}
</style>