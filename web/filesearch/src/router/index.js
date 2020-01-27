import Vue from 'vue'
import Router from 'vue-router'
import VueResource from 'vue-resource'
import HomeSearchPage from '@/components/HomeSearchPage'
import SearchResult from '@/components/SearchResult'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import VueClipboard from 'vue-clipboard2'

Vue.use(Router)
Vue.use(ElementUI)
Vue.use(VueResource)
Vue.use(VueClipboard)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HomeSearchPage',
      component: HomeSearchPage
    },
    {
      path: '/search_result',
      name: 'SearchResult',
      component: SearchResult
    }
  ]
})
