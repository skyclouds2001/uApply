<template>
    <div class="Master">
        <side-nav :navs="MasterNavs"></side-nav>
        <div class="main">
            <!-- 蓝色背景，绝对定位 -->
            <div class="background"></div>
            <page-head :title="page_title" role="管理员"></page-head>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
import SideNav from '../../components/generic/side_nav/side_nav.vue'
import PageHead from '../../components/generic/page_head/page_head.vue';

export default{
    components:{
        SideNav,
        PageHead,
    },
    setup(){    
        // 侧边导航数据
        const MasterNavs = [
            {
                name: "面试官配置",
                label: require("@/assets/imgs/配置.png"),
                index: 0,
                nav: "config"
            },
            {
                name: "新投递数据",
                label: require("@/assets/imgs/新投递数据.png"),
                index: 1,
                nav: "new"
            },
            {
                name: "已面数据",
                label: require("@/assets/imgs/已面试数据.png"),
                index: 2,
                nav: "interviewed"
            },
            {
                name: "淘汰数据",
                label: require("@/assets/imgs/淘汰数据.png"),
                index: 3,
                nav: "eliminate"
            },
            {
                name: "录取名单",
                label: require("@/assets/imgs/录为部员数据.png"),
                index: 4,
                nav: "admission"
            },
            {
                name: "招新数据",
                label: require("@/assets/imgs/招新数据.png"),
                index: 5,
                nav: "data"
            }
        ];

        const page_title = "西电青年志愿者协会"

        return {
            MasterNavs,
            page_title
        }
    }
}
</script>

<style lang="scss" scoped>
.Master {
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

:deep(.page){
    height: 88%;
    width: 100%;
    flex-wrap: wrap;
    display: flex;
    justify-content: center;
    align-items: flex-start;
    flex-wrap: wrap;
    overflow: auto;
}
</style>