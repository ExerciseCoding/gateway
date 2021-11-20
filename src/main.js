// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

// 引入 axios
import axios from 'axios'

Vue.prototype.$axios = axios
// 引入echarts
import echarts from 'echarts'
Vue.prototype.$echarts = echarts
import { Bus } from './utils/bus'
// import 'v-charts/lib/style.css'


router.beforeEach((to, from, next) => {
  /* 判断该路由是否需要登录权限 */
  // requireAuth从router下的index.js取出
  console.log('===')
  console.log(to.matched)
  console.log('===')
  if (to.matched.some(record => record.meta.requireAuth)) {
    // axios.get('/api/user/userinfo').then(rsp => {
    //   rsp = rsp.data
    //   if (rsp.code === 401) {
    //     // 未登录
    //     next({
    //       path: '/login',
    //     })
    //   } else {
    //     // 已登录
    //     this.$userinfo = rsp.data[0]
    //     localStorage.setItem("userinfo", JSON.stringify(rsp.data[0]));
    //     Bus.$emit('change')
    //     next()
    //   }
    // })
  }
  next();
})

// 引入element ui
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);

Vue.config.productionTip = false

// utc2local
Vue.prototype.utc2local = function(utc_datetime) {
  // 转为正常的时间格式 年-月-日 时:分:秒
  let T_pos = utc_datetime.indexOf('T');
  let Z_pos = utc_datetime.indexOf('Z');
  let year_month_day = utc_datetime.substr(0, T_pos);
  let hour_minute_second = utc_datetime.substr(T_pos + 1, Z_pos - T_pos - 1);
  let new_datetime = year_month_day + " " + hour_minute_second;

  // 处理成为时间戳
  timestamp = new Date(Date.parse(new_datetime));
  timestamp = timestamp.getTime();
  timestamp = timestamp / 1000;

  // 增加8个小时，北京时间比utc时间多八个时区
  let timestamp = timestamp + 8 * 60 * 60;

  // 时间戳转为时间
  let date = new Date(parseInt(timestamp) * 1000)
  let year = formatNumber(date.getFullYear())
  let month = formatNumber(date.getMonth() + 1)
  let day = formatNumber(date.getDate())
  let hour = formatNumber(date.getHours())
  let min = formatNumber(date.getMinutes())
  let sec = formatNumber(date.getSeconds())
  let time = year + '/' + month + '/' + day + ' ' + hour + ':' + min + ':' + sec
  //let beijing_datetime = new Date(parseInt(timestamp) * 1000).toLocaleString('chinese', { hour12: false }).replace(/24:00:00/g, "00:00:00")
  return time;
}

function formatNumber(value) {
  if (value < 10) {
    value = "0" + value
  }
  return value
}

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
