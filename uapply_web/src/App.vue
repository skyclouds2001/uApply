<template>
  <div class="app">
    <router-view></router-view>
  </div>
</template>

<script>
import indexVue from './views/pages/index.vue';

export default{
  components:{
    indexVue
  },
  setup(){

  }
}
</script>

<style lang="scss">
body{
  display: flex;
  justify-content: center;
}

*{
  margin: 0;
  padding: 0;
}

#app {
  height: 100vh;
  width: 100vw;
  box-shadow: 0px 0.1px 3px 0px #c0bfbf;
  position: fixed;
  top: 0;

  
  @media (min-width:1280px) {
    max-width: 2.5rem;
  }
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
