import 'bootstrap/dist/css/bootstrap.min.css';
import 'element-plus/dist/index.css';

import {Component, createApp} from 'vue';

import App from './App.vue';
import {router} from './router';

const app = createApp(App as Component);

app.use(router);
app.mount('#app');
