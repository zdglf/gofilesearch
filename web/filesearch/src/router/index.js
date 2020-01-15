import Vue from 'vue'
import Router from 'vue-router'
import HomeSearchPage from '@/components/HomeSearchPage'
import SearchResult from '@/components/SearchResult'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(Router)
Vue.use(ElementUI)

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
