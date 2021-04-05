import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import * as Wails from '@wailsapp/runtime';
import 'bootstrap'

document.title = "Echo-Slam"

Wails.Init(() => {
    let app = createApp(App);
    app.use(router);
    app.mount('#app');
});
