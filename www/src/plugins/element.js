import ElementPlus from 'element-plus'
import { createI18n } from 'vue-i18n'
import 'element-plus/theme-chalk/index.css'
import * as Icons from '@element-plus/icons-vue'
import localeZH from 'element-plus/lib/locale/lang/zh-cn'
import localeEN from 'element-plus/lib/locale/lang/en'
import messages from '../utils/i18n'

const i18n = createI18n({
  locale: localeZH.name,
  fallbackLocale: localeEN.name,
  messages,
})

function transElconName(iconName) {
  return 'el-icon' + iconName.replace(/[A-Z]/g, (match) => '-' + match.toLowerCase())
}

export default (app) => {
  for (const key in Icons) {
    app.component(transElconName(key), Icons[key])
  }
  app.use(ElementPlus, { locale:localeZH })
  app.use(i18n)
}
