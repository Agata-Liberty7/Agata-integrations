import { createApp } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import easySpinner from 'vue-easy-spinner'

import App from './App.vue'
import HelloWorld from './components/HelloWorld.vue'


const routes = [
  { path: '/', component: HelloWorld },
  // { path: '/about', component: About },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})


let app = createApp(App)

// // LOCAL
app.config.globalProperties.REGION = 'WORLD'
// app.config.globalProperties.REGION = 'RUS'
app.config.globalProperties.HOST = 'http://localhost:8080/'

// // // WORLD
// app.config.globalProperties.REGION = 'WORLD'
// app.config.globalProperties.HOST = 'http://178.62.240.212:80/'

// // RUS
// app.config.globalProperties.REGION = 'RUS'
// app.config.globalProperties.HOST = 'http://80.249.144.41:80/'

app
  .use(router)
  .use(easySpinner, { prefix: 'easy' })
  .mount('#app')