<template>
    <div class="pageHead">
        <span>{{ title }}</span>
        <span class="returnIcon" v-if="isShowReturen" @click="goBack">{{ returnIcon }}</span>
    </div>
</template>

<script>
/**
 * 页头
 * @title 父组件传参，作为页头标题
 */
import { getCurrentInstance, onMounted } from '@vue/runtime-core'

export default {
    props: {
        title: {
            type: String,
            require: true,
        },
        isShowReturen: {
            type: Boolean,
            default: true
        }
    },
    setup(props, context) {
        onMounted(() => {
            console.log(props.title);
        })

        // 实例对象
        const { proxy } = getCurrentInstance();

        const goBack = function () {
            // router.push("/");
            proxy.$emit("goBack");
        }

        return {
            returnIcon: "<",
            goBack
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($row: row, $isCerten: true) {
    display: flex;
    flex-direction: $row;

    @if ($isCerten) {
        justify-content: center;
        align-items: center;
    }
}

.pageHead {
    position: relative;
    box-sizing: border-box;
    height: 4.9vh;
    width: 100%;
    background: #6C7BFF;
    @include flex(row, true);

    span {
        color: white;
        font-size: 1.7vh;
    }

    .returnIcon {
        position: absolute;
        left: 5%;
        font-size: large;
        cursor: pointer;
    }

    .returnIcon:hover {
        transform: scale(1.2);
    }
}
</style>