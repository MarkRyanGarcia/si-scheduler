import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'

if (import.meta.env.VITE_ENV === 'dev' || import.meta.env.VERCEL_GIT_COMMIT_REF === 'dev') {
    import('./dev-style.css')
}


createApp(App).use(router).mount('#app')