<template>
  <!-- 登录页面 -->
  <div class="login">
    <!-- 设置离谱的背景 -->
    <div class="threeJs" ref="threeJs"></div>
    <!-- 登录框,使用登录框组件 -->
    <div class="login_block">
      <login-block @login="login"></login-block>
    </div>
    <div class="website_text">
      <span>- 西安电子科技大学青鸢工作室 -</span>
      <span>备案号</span>
    </div>
  </div>
</template>

<script>
import { ref } from '@vue/reactivity';
import LoginBlock from '../../components/login/login_block.vue';
import { OrganizationLogin, DepartmentLogin } from '../../utils/apis/api';
import { message } from '../../utils/message/message';
import cookies from 'vue-cookies';
import router from '../../router';
import { getCurrentInstance } from '@vue/runtime-core';
import * as THREE from 'three';
import Stats from '../../utils/stats';

export default {
  components: {
    LoginBlock,
  },
  setup(props, context) {
    const title = ref('Welcome!');
    const white_meteor = require('@/assets/imgs/白色流星.png');
    const blue_meteor = require('@/assets/imgs/蓝色流星.png');
    let stats;
    let camera, scene, renderer;
    let objects;
    let mouseX = 0,
      mouseY = 0;
    let windowHalfX = window.innerWidth / 2;
    let windowHalfY = window.innerHeight / 2;

    // 实例对象
    const { proxy } = getCurrentInstance();

    // 发起登录请求
    const login = async function (e) {
      // 登录请求体
      const params = {
        account: e.account,
        password: e.password,
      };
      // 请求返回
      let res = {};

      res = await (e.role == '超级管理员'
        ? OrganizationLogin(params)
        : DepartmentLogin(params));
      // console.log(res);
      // 如果登录失败
      if (res.data.code == 1004) {
        message.info('账号不存在，请检查角色/账号/密码是否有误！');
        return;
      }
      // 将用户信息保存到本地
      proxy.$localData('set', 'userInfo', res.data);
      // 保存Token
      cookies.set('Token', res.data.accessToken);
      if (e.role == '超级管理员') {
        router.push('/SuperMaster');
      } else if (e.role == '管理员') {
        router.push('/Master');
      }
    };

    return {
      title,
      login,
      blue_meteor,
      white_meteor,
      stats,
      camera,
      scene,
      renderer,
      objects,
      mouseX,
      mouseY,
      windowHalfX,
      windowHalfY,
    };
  },

  mounted() {
    document.addEventListener('mousemove', this.onDocumentMouseMove);
    this.$nextTick(this.init);
  },

  methods: {
    init() {
      const container = this.$refs.threeJs;

      this.camera = new THREE.PerspectiveCamera(
        60,
        window.innerWidth / window.innerHeight,
        1,
        10000,
      );
      this.camera.position.z = 3200;
      this.scene = new THREE.Scene();
      const groups = new THREE.Group();
      this.scene.background = new THREE.Color(0xffffff);

      this.objects = [];

      const material = new THREE.MeshNormalMaterial();

      const loader = new THREE.BufferGeometryLoader();
      const geometry = new THREE.BoxGeometry(1, 1, 1);

      for (let i = 0; i < 5000; i++) {
        const mesh = new THREE.Mesh(geometry, material);

        mesh.position.x = Math.random() * 8000 - 4000;
        mesh.position.y = Math.random() * 8000 - 4000;
        mesh.position.z = Math.random() * 8000 - 4000;
        mesh.rotation.x = Math.random() * 2 * Math.PI;
        mesh.rotation.y = Math.random() * 2 * Math.PI;
        mesh.scale.x = mesh.scale.y = mesh.scale.z = Math.random() * 50 + 100;

        this.objects.push(mesh);

        groups.add(mesh);
      }
      this.scene.add(groups);

      this.renderer = new THREE.WebGLRenderer();
      this.renderer.setPixelRatio(window.devicePixelRatio);
      this.renderer.setSize(window.innerWidth, window.innerHeight);
      container.appendChild(this.renderer.domElement);

      this.stats = new Stats();
      container.appendChild(this.stats.dom);

      window.addEventListener('resize', this.onWindowResize);

      this.animate();
    },
    onWindowResize() {
      this.windowHalfX = window.innerWidth / 2;
      this.windowHalfY = window.innerHeight / 2;

      this.camera.aspect = window.innerWidth / window.innerHeight;
      this.camera.updateProjectionMatrix();

      this.renderer.setSize(window.innerWidth, window.innerHeight);
    },
    onDocumentMouseMove(event) {
      this.mouseX = (event.clientX - this.windowHalfX) * 10;
      this.mouseY = (event.clientY - this.windowHalfY) * 10;
    },
    animate() {
      requestAnimationFrame(this.animate);
      this.render();
      this.stats.update();
    },
    render() {
      this.camera.position.x += (this.mouseX - this.camera.position.x) * 0.05;
      this.camera.position.y += (-this.mouseY - this.camera.position.y) * 0.05;
      this.camera.lookAt(this.scene.position);
      this.renderer.render(this.scene, this.camera);
    },
  },
};
</script>

<style lang="scss" scoped>
/*
* @封装flex布局
* column: 主轴方向
* isCenter: 是否垂直水平居中
*/
@mixin flex($direction: column, $isCenter: true) {
  display: flex;
  flex-direction: $direction;

  @if $isCenter {
    align-items: center;
    justify-content: center;
  }
}

.login {
  height: 100%;
  width: 100%;
  @include flex(column, true);

  .login_block {
    position: absolute;
    margin-top: 4%;
    @include flex(column, true);
  }

  .website_text {
    height: 10%;
    width: 100%;
    position: absolute;
    bottom: 2.5%;
    @include flex(column, true);
    color: #8e8c90;
    z-index: -1;

    font-size: 1.5rem;
  }

  .threeJs {
    opacity: 0.9;
  }

  .background {
    height: 100%;
    width: 100%;
    z-index: -1;

    // 配置流星图案
    img {
      height: 12rem;
      position: absolute;
    }

    .meteor1 {
      left: 20%;
      top: 10%;
    }

    .meteor2 {
      left: 6%;
      bottom: 13%;
    }

    .meteor3 {
      right: 30%;
      bottom: 7%;
    }

    .meteor4 {
      right: 0%;
      bottom: 33%;
    }

    .floor1 {
      width: 100%;
      height: 20%;
      background: linear-gradient(to right, #717af1, #c081ff);
      @include flex(column, true);
      color: white;
      font-weight: 900;

      font-size: 6rem;

      @media (max-width: 1680px) {
        font-size: 5rem;
      }

      @media (max-width: 1440px) {
        font-size: 4rem;
      }

      @media (max-width: 1366px) {
        font-size: 3.5rem;
      }
    }

    .floor2 {
      width: 100%;
      height: 32%;
      background: rgba(51, 59, 200, 1);
    }

    .floor3 {
      width: 100%;
      height: 48%;
      background: rgba(23, 43, 77, 1);
    }
  }
}
</style>
