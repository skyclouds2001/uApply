<template>
    <!-- 侧边导航页 -->
    <div class="side_nav">
        <!-- 导航logo -->
        <div class="page_logo">
            <img :src="page_logo" style="height:60%">
        </div>
        <!-- 导航按钮 -->
        <div class="navs">
            <router-link class="nav" :to="item.nav" :class="item.index == active_index ? 'nav_active' : ''"
                v-for="(item, index) in navs" @click="changeNav(index)">
                <img :src="item.label">
                <span>{{ item.name }}</span>
            </router-link>
        </div>
    </div>
</template>

<script>
import { ref } from '@vue/reactivity';

export default {
    props: {
        navs: Array,
    },
    setup(props, context) {
        const page_logo = require("@/assets/imgs/page_logo.png");

        let active_index = ref(0);

        const changeNav = function (index) {
            // 修改选中index修改导航选中状态
            active_index.value = index;
            // 根据点击的索引,导航到不同的组件
        }

        return {
            page_logo,
            changeNav,
            active_index
        }
    }
}
</script>

<style lang="scss" scoped>
@mixin flex($direction: row, $isCenter: true) {
    display: flex;
    flex-direction: $direction;

    @if ($isCenter) {
        justify-content: center;
        align-items: center;
    }

    @else {
        justify-content: space-between;
        align-items: center;
    }
}

.side_nav {
    height: 100%;
    width: 17%;
    // border-right: 3px solid rgb(190, 190, 190);
    @include flex(column, false);

    .page_logo {
        height: 10%;
        width: 100%;
        border-bottom: 1px solid lightgray;
        @include flex(row, true);
    }

    .navs {
        height: 89%;
        width: 100%;
        position: relative;

        .nav {
            height: 12%;
            width: 64%;
            padding: 0 18%;
            @include flex(row, false);
            cursor: pointer;
            text-decoration: none;

            img {
                height: 40%;
            }

            span {
                color: #6b6b6b;


                font-size: 1.7rem;

                @media (max-width:1680px) {
                    font-size: 1.5rem;
                }

                @media (max-width:1440px) {
                    font-size: 1.3rem;
                }

                @media (max-width:1280px) {
                    font-size: 1.2rem;
                }
            }
        }

        .nav_active {
            background: rgb(241, 241, 241);

            span {
                color: black;
            }
        }
    }
}
</style>