<template>
  <el-aside  style="width:220px; height: 100%; background-color : #f0f2f5; padding-top: 3px;">
    <el-menu :router="true" :default-active="this.$route.path" :unique-opened="true" style="width: 220px;height: 100%;">
      <el-submenu key="2" index="2">
        <template slot="title" index="/"><i class="el-icon-tickets"></i>总览</template>
        <el-menu-item key="2-1" index='/'><i class="el-icon-data-analysis"></i> 数据分析大盘</el-menu-item>
      </el-submenu>
      <el-submenu key="0" index="0">
          <template slot="title"><i class="el-icon-tickets"></i> 服务管理</template>
          <el-menu-item key="0-1" index='/ServiceList'><i class="el-icon-data-analysis"></i> 服务列表</el-menu-item>

          <el-submenu v-for="(group,i) in ServiceNavGroup" :key="i" :index="group.group">
            <template slot="title"><i :class="group.icon"></i> {{ group.service }} </template>
            <el-menu-item v-for="(item,i) in ServiceNavItem" :key="i" :index="'/' + group.group + item.name"> {{ item.navTitle }}</el-menu-item>
          </el-submenu>

      </el-submenu>
      <el-submenu key="1" index="1">
          <template slot="title"><i class="el-icon-tickets"></i> 租户列表</template>
          <el-menu-item key="1-1" index='/OverviewEvent'><i class="el-icon-data-analysis"></i> 总览</el-menu-item>
      </el-submenu>


      <el-menu-item key="2" index=""> <i class="el-icon-paperclip"></i> 更多功能（建设中）
      </el-menu-item>
      <el-menu-item key="3" v-if="isSuperAdmin" index="subscribe"> <i class="el-icon-bell"></i> 视图订阅（建设中）
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>
<script>

export default {
  name: "left",
  data() {
    return {
      ServiceNavItem: [
        { name: 'Services', navTitle: 'Feeds' },
	    ],
      ServiceNavGroup: [
        {group: 'HTTP', icon: 'el-icon-user', service: 'HTTP服务'},
      ],


      peopleAdmin: false,
      peopleAdminList: [],
      isSuperAdmin: false,
      user: '',
      ticket_value: 0,
    }
  },
  created: function() {
    // this.$axios.get('api/auth/permission')
    // .then((resp) => {
    //   var res = resp.data
    //   if (res.error_code === 0) {
    //     console.log(res)
    //     this.isSuperAdmin = res.data.superadmin
    //   } else {
    //     this.$message.error('查询用户权限数据失败')
    //   }
    // })
  },

  methods: {

  },

  mounted() {

  },

  beforeDestroy() {

  }

}

</script>
<style>
.el-container {
  height: 100%
}

.el-aside {
  color: #333;
  border-radius: 2px
}

.el-badge__content {
  position: absolute;
  top: 28px !important;
  right: -14px !important;

}

</style>
