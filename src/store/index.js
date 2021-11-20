import Vue from 'vue'
import Vuex from 'vuex'

// 1. 安装插件
Vue.use(Vuex)

// 2. 创建对象
const state = {
  times: '',
  project: '',
  module: '',
  chart_member: '',
  chart_title: '工单数量',
  chart_type: 'line',
  chart_am: '',
  chart_group: '',
  chart_name: ['ticketAllCount_Group', 'ticketAllCount_Toc'],
  chart_time: [],
  chart_seriesName: 'Content',
  view_type: 'week',
  // 第三层视图类型标识
  view_groupType: 'member',
  first_title:'ToC',
  second_title:'变更数量及质量',
  // 趋势图or结果图
  trendResult_Type: 'weekTrend',
  // 跳转到下一级后视图的个数
  view_count: 0,
  // 产品线下钻类型: project 项目维度  module 模块维度
  product_type: '',
  member: '',
  first: false,
  prepath: '/'
}

const store = new Vuex.Store({
  state,
  mutations: {
    changeTimes(state, times){
      state.times = times
    },
    changeProject(state, project){
      state.project = project
    },
    changeModule(state, module){
      state.module = module
    },
    changeChartMember(state, chart_member){
      state.chart_member = chart_member
    },
    changeChartGroup(state, chart_group){
        state.chart_group = chart_group
    },
    changeChartName(state, chart_name){
        state.chart_name = chart_name
    },
    changeChartAM(state, chart_am){
      state.chart_am = chart_am
    },
    changeChartNameTwo(state, chart_name){
      state.chart_name = chart_name
    },
    changeChartTime(state, chart_time){
        state.chart_time = chart_time
    },
    changeChartTitle(state, chart_title){
      state.chart_title = chart_title
    },
    changeChartType(state, chart_type){
      state.chart_type = chart_type
    },
    changeMember(state, member){
      state.member = member
    },
    changeFirst(state, first){
      state.first = first
    },
    changePrePath(state, prepath){
      state.prepath = prepath
    },
    changeSeriesName(state,chart_seriesName){
      state.chart_seriesName = chart_seriesName
    },
    viewType(state,view_type){
      state.view_type = view_type
    },
    viewGroupType(state,view_groupType){
      state.view_groupType = view_groupType
    },
    firstTitle(state,first_title){
      state.first_title = first_title
    },
    secondTitle(state,second_title){
      state.second_title = second_title
    },
    trendResultType(state,trendResult_Type){
      state.trendResult_Type = trendResult_Type
    },
    viewCount(state, view_count){
      state.view_count = view_count
    },
    setProductType(state,product_type){
      state.product_type = product_type
    },
  }
})


// 3. 导出store对象
export default store
