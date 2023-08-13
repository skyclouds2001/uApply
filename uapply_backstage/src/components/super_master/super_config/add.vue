<template>
    <div class="add">
        <div class="title">
            <div class="func_title">
                {{ add_type.title }}
            </div>
            <div class="func_contain">
                <p>{{ start_time }}</p>
                <p>{{ end_time }}</p>
            </div>
        </div>
        <img :src="img" @click="add">
    </div>
</template>

<script>
import { ref } from '@vue/reactivity'
import { getCurrentInstance, onMounted } from '@vue/runtime-core';
import { GetRecruitmentDate } from '../../../utils/apis/api';
import toTime from '../../../utils/dateHandle/dateHandle';

export default {
    props: {
        add_type: Object,
    },
    setup(props, context) {
        // 标题和日期
        // let { text } = toRefs(props.add_type);
        let start_time = ref(props.add_type.type == "date" ? "配置组织招新日期" : "配置部门信息" );
        let end_time = ref("");

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 是配置招新日期的图标还是新增部门的图标呢
        const img = props.add_type.type == "date" ? require("@/assets/imgs/配置招新日期.png") : require("@/assets/imgs/新增部门.png")

        const add = function () {
            // 根据type呼出不同的toast
            props.add_type.type == "date" ? context.emit("toast", "AddDate") : context.emit("toast", "AddDep");
        }

        // 重置时间
        const reset = async function () {
            const { data } = await GetRecruitmentDate(proxy);
            // console.log(data);
            if (data) {
                start_time.value = toTime(data.start * 1000);
                end_time.value = toTime(data.end * 1000);
                return;
            }
        }

        onMounted(() => {
            if (props.add_type.type == "date") {
                reset();
            }
        })

        return {
            img,
            start_time,
            end_time,
            reset,
            add
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
        justify-content: space-around;
        // align-items: center;
    }
}

.add {
    border-radius: 0.5rem;
    margin-right: 5%;
    background: white;
    display: flex;
    justify-content: space-around;
    align-items: center;

    width: 22rem;
    height: 10rem;

    @media (max-width:1440px) {
        width: 17.5rem;
        height: 7.5rem;
    }

    @media (max-width:1280px) {
        width: 15.5rem;
        height: 7.5rem;
    }

    .title {
        height: 60%;
        @include flex(column, false);

        .func_title {
            font-weight: 600;
            font-size: 2rem;

            @media (max-width:1680px) {
                font-size: 1.7rem;
            }

            @media (max-width:1440px) {
                font-size: 1.5rem;
            }

            @media (max-width:1366px) {
                font-size: 1.4rem;
            }
        }

        .func_contain {
            font-size: 1.2rem;

            @media (max-width:1680px) {
                font-size: 1.05rem;
            }

            @media (max-width:1440px) {
                font-size: 1rem;
            }

            @media (max-width:1366px) {
                font-size: 0.95rem;
            }
        }
    }

    img {
        height: 50%;
        cursor: pointer;
    }
}
</style>