import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
import Index from '../components/Index'
import login from '../components/login'
import ServiceList from "../components/ServiceList"
import ServiceAddHTTP from "../components/Service/http"
import overview from "../components/Overview";
// 权限管理
import AuthDashBord from '../components/AuthDashBord/AuthDashBord'



// 解决ElementUI导航栏中的vue-router在3.0版本以上重复点菜单报错问题
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(Router)


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'overview',
      component: overview,
      meta:{
        keepAlive: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: login,
      meta: {
        keepAlive: false
      }
    },
    {
      path: '/ServiceList',
      name: 'ServiceList',
      component: ServiceList,
      meta: {
        keepAlive: true,
        requireAuth: true,
      }
    },
    {
      path: '/ServiceAddHTTP',
      name: 'ServiceAddHTTP',
      component: ServiceAddHTTP,
      meta: {
        keepAlive: true,
        requireAuth: true,
      }
    },
  ]
})
