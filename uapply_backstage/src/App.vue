<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

export default {
}

</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
}

#app {
  height: 100vh;
  width: 100%;
  min-width: 1280px;
  min-height: 767px;

  @media (max-width:1680px) {

  }

  @media (max-width:1440px) {}

  @media (max-width:1366px) {
    min-width: 1024px;
    min-height: 600px;
  }

}



::-webkit-scrollbar {
  /*整个滚动条*/
  width: 6px;
  background-color: #dddddd;
}

::-webkit-scrollbar-track {
  /*滚动条轨道*/
  border-radius: 10px;
}

::-webkit-scrollbar-thumb {
  /*滚动滑块*/
  background: #4ec2a3;
}

</style>
