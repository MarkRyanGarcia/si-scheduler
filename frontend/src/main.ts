import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'

if (import.meta.env.MODE === 'development') {
    import('./dev-style.css')
}

createApp(App).use(router).mount('#app')