import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueCookies from 'vue-cookies'
import { localData, sessionData } from '../src/store/localStorage'

const app = createApp(App);

app.use(store);
app.use(router);
app.use(VueCookies);

app.config.globalProperties.$localData = localData;
app.config.globalProperties.$sessionData = sessionData;

app.mount("#app");

// createApp(App).use(store).use(router).mount('#app')
