import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import installElementPlus from './plugins/element'
import './assets/css/icon.css'
import mock from './mock'

if(process.env.NODE_ENV === 'development') {
    mock.bootstrap()
}

const app = createApp(App)
installElementPlus(app)
app
    .use(store)
    .use(router)
    .mount('#app')