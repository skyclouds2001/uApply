<template>
    <div id="block" :class="type == 'status' ? 'bg1' : 'bg2'">
        <span>{{  text  }}</span>
    </div>
</template>

<script>
import cookies from 'vue-cookies';
import { ref } from '@vue/reactivity';
import { onMounted } from '@vue/runtime-core';

export default {
    props: {
        type: String,
    },
    setup(props) {
        console.log(props.type);

        const text = ref("");

        const update = function () {
            text.value = props.type == "status" ? "我的状态" : cookies.get("userUid") == null ? "登录获取ID" : "您的ID为：" + cookies.get("userUid");
        }

        onMounted(() => {
            update();
        })

        return {
            text,
            update
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

#block {
    @include flex(row, true);
    height: 32%;
    width: 100%;
    border-radius: 2vh;

    span {
        color: white;
        font-size: 3.3vh;
    }
}

.bg1 {
    background: #FFA2A2;
}

.bg2 {
    background: #8CD742;
}
</style>