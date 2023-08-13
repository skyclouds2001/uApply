<template>
    <div class="SuperMaster">
        <side-nav :navs="SuperMasterNavs"></side-nav>
        <div class="main">
            <!-- 蓝色背景，绝对定位 -->
            <div class="background"></div>
            <page-head :title="page_title" role="超级管理员"></page-head>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
import SideNav from '../../components/generic/side_nav/side_nav.vue';
import PageHead from '../../components/generic/page_head/page_head.vue'
import { isLogin } from '../../utils/verification/isLogin';
import { onMounted } from '@vue/runtime-core';
import SuperConfig from "./super_master_pages/super_config.vue"

export default {
    components: {
        SideNav,
        PageHead,
        SuperConfig,
    },
    setup() {
        // 侧边导航数据
        const SuperMasterNavs = [
            {
                name: "超级配置",
                label: require("@/assets/imgs/配置.png"),
                index: 0,
                nav: "config"
            },
            {
                name: "招新数据",
                label: require("@/assets/imgs/招新数据.png"),
                index: 1,
                nav: "data"
            }
        ];
        const page_title = "西电青年志愿者协会"

        onMounted(() => {
            isLogin();
        })

        return {
            SuperMasterNavs,
            page_title
        }
    }
}
</script>

<style lang="scss">
.SuperMaster {
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: space-between;

    .main {
        height: 100%;
        width: 83%;
        position: relative;
        border-left: 3px solid rgb(190, 190, 190);


        .background {
            height: 60%;
            width: 100%;
            position: absolute;
            z-index: -1;
            background: #717AF1;
        }
    }
}

.n-data-table-th{
    text-align: center;
    font-size: 120%;
    font-weight: 600;
    background: #f7f1ff;
}

</style>