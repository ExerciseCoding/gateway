<template>
  <div class="overview">
      <el-row>
        <el-col :span="24">
          <el-breadcrumb separator-class="el-icon-arrow-right" style="margin-bottom: 20px; border: 1px solid #eee" >
            <!-- <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item> -->
            <el-breadcrumb-item :to="{ path: '/' }"><b>服务列表</b></el-breadcrumb-item>
          </el-breadcrumb>
        </el-col>
      </el-row>
      <el-row>
          <el-form :inline="true" class="demo-form-inline">
            <el-form-item>
              <!-- <el-date-picker v-model="queryparam.chart_time" size="small" type="datetimerange" :picker-options="pickerOptions" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间" align="right"></el-date-picker> -->
              <el-input v-model="input" placeholder="服务名/服务描述" size="small"></el-input>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" icon="el-icon-search" size="small">搜索</el-button>
              <!-- <el-button type="primary" @click="queryChart()" size="small">查询</el-button> -->
              <el-button type="primary" to="/ServiceAddHTTP" size="small" icon="el-icon-edit">添加HTTP服务</el-button>
              <el-button type="primary" @click="queryQuick('OneMonth')" size="small" icon="el-icon-edit">添加TCP服务</el-button>
              <el-button type="primary" @click="queryQuick('ThreeMonth')" size="small" icon="el-icon-edit">添加GRPC服务</el-button>
<!--              <el-button type="primary" @click="queryQuick('HalfYear')" size="small">最近半年</el-button>-->
            </el-form-item>
          </el-form>
      </el-row>
    <el-row class="数据" style="height: calc(100vh - 300px)">
      <el-scrollbar style="height: 100%">
        <el-row>
          <el-table :data="serviceList" border style="width: 100%">
            <el-table-column fixed prop="id" label="ID" width="50">
            </el-table-column>
            <el-table-column prop="service_name" label="服务名称" width="120">
            </el-table-column>
            <el-table-column prop="service_desc" label="服务描述" width="120">
            </el-table-column>
            <el-table-column prop="load_type" label="类型" width="120">
            </el-table-column>
            <el-table-column prop="service_addr" label="服务地址" width="300">
            </el-table-column>
            <el-table-column prop="qps" label="QPS" width="80">
            </el-table-column>
            <el-table-column prop="qpd" label="日请求数" width="80">
            </el-table-column>
            <el-table-column prop="total_node" label="节点数" width="80">
            </el-table-column>
            <el-table-column prop="compile" label="操作" width="300">
              <template slot-scope="scope">
                <el-button @click="handleClick(scope.row)" type="primary" size="mini">统计</el-button>
                <el-button @click="handleClick(scope.row)" type="primary" size="mini">修改</el-button>
                <el-button type="primary" size="mini">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </el-scrollbar>
    </el-row>
    <el-row class="页数" style="background-color: unset;text-align: center">
        <el-pagination small style="background-color: unset;margin-top:20px ;" @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page.sync="currentPage" :page-sizes="pageSizes" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper" :total=total>
        </el-pagination>
    </el-row>
    <div style="height: calc(100vh - 260px); width: 100%;">
      <el-scrollbar :native="false">
        <div class="默认视图" style="width: 98%">
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>
<script>
import Vue from "vue";
// import VCharts from "v-charts";
import {defaults} from "autoprefixer";


// Vue.use(VCharts);


export default {
  name: "ServiceList",
  data() {
    return {
      isAdmin: false,
      isSuperAdmin: false,
      serviceList: [],
      currentPage: 1, // 默认显示第几页
      pageSize: 10, // 每页显示条数
      total: 1, // 总条数
      pageSizes: [10, 20, 50, 100], // 个数选择器(可修改)
      // userviews: ['loading'],  //用户拥有的视图权限，列表形式
      pickerOptions: {
        shortcuts: [
          {
            text: "最近一周",
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit("pick", [start, end]);
            },
          },
          {
            text: "最近一个月",
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit("pick", [start, end]);
            },
          },
          {
            text: "最近三个月",
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit("pick", [start, end]);
            },
          },
        ],
      },
    };
  },


  created() {
    this.queryServiceList()
  },

  methods: {
    queryServiceList(){
      this.$axios.get('service/service_list',{params:{page_size : 10, page_no : 1}})
        .then((resp) => {
          var res = resp.data
          if (res.errno === 0) {
            this.total = res.data.total
            this.serviceList = res.data.list
            // // ...
          } else {
            this.$message.error('查询数据失败')
          }
        })
    },
    serviceAddHTTP(){

    },
    handleSizeChange(val) {
      // 改变每页显示的条数
      this.pageSize = val
      // 注意：在改变每页显示的条数时，要将页码显示到第一页
      this.currentPage = 1
    },
    // 显示第几页
    handleCurrentChange(val){
      // 改变默认的页数
      this.currentPage = val
    },
  },

};
</script>


<style>
.data-empty {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, .7);
  color: #888;
  font-size: 14px;
}
</style>
