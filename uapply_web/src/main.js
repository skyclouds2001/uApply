import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'amfe-flexible'
import $ from 'jquery'
import VueCookies from 'vue-cookies'
// import 'default-passive-events';

// import VueCookies from 'vue-cookies'
import { localData,sessionData } from "./store/localstorage"

const app = createApp(App);

app.use(store);
app.use(router);
app.use(VueCookies);

app.config.globalProperties.$localData = localData;
app.config.globalProperties.$sessionData = sessionData;

app.mount("#app");