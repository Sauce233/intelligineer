import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import './styles.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import piniaPluginPeristedstate from 'pinia-plugin-persistedstate'
import 'element-plus/theme-chalk/dark/css-vars.css'

import zh from '@/locale/zh'
import zhTW from '@/locale/zhTW'
import en from '@/locale/en'

import App from './App.vue'
import router from './router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

const i18n = createI18n({
  locale: 'zh',
  fallbackLocale: 'zh',
  messages: {
    'zh-TW': zhTW,
    en,
    zh,
  },
  silentFallbackWarn: true,
  silentTranslationWarn: true,
})
app.use(i18n)

const pinia = createPinia()
pinia.use(piniaPluginPeristedstate)
app.use(pinia)
app.use(router)
app.use(ElementPlus, { locale: zhCn })

app.mount('#app')
