<template>
    <div class="error">
        <h3>404</h3>
        <p>{{  errorMsg  }}</p>
    </div>
</template>

<script>
export default {
    setup() {
        return {
            errorMsg: "请求资源错误"
        }
    }
}
</script>

<style lang="scss" scoped>
.error {
    width: 100%;
    height: 70vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    h3 {
        font-size: 2.5rem;

        @media (min-width:1280px) {
            font-size: 0.6rem;
        }
    }

    p {
        font-size: 0.6rem;

        @media (min-width:1280px) {
            font-size: 0.15rem;
        }
    }
}
</style>