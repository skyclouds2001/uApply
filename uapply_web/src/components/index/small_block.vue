<template>
    <div id="block" :class="type == 'resume' ? 'bg1' : 'bg2'">
        <span>{{ buttonText }}</span>
        <div id="image" class="image">
            <img :src="src">
        </div>
    </div>
</template>

<script>
export default {
    props: {
        type: String
    },
    setup(props) {
        const src = props.type == "resume" ? require("@/static/imgs/我的简历.png") : require("@/static/imgs/我要申报.png");

        return {
            buttonText: props.type == "resume" ? "我的简历" : "我要申报",
            src,
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
    height: 28%;
    width: 49%;
    border-radius: 1.8vh;

    span {
        color: white;
        font-size: 2.3vh;
        margin-right: 5%;
    }

    #image{
        height: 40%;
        padding: 0 5%;
        border-radius: 50%;
        @include flex(row, true );

        img{
            height: 60%;
        }
    }
}

.bg1 {
    background: #8972FF;

    .image{
        background: #5f35ea;
    }
}

.bg2 {
    background: #0093FA;

    .image{
        background: #0353c2;
    }
}
</style>