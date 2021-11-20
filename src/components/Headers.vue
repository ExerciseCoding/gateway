<template>
  <el-header>
    <div>
      <img style="float: left;padding-top: 10px" src="../assets/logo-apollo.png" height="35" width="45" class="logo">
    </div>
    <div><span style="float: left;font-size: 18px;padding-left: 10px">Apollo网关系统</span></div>
    <span><a href="https://github.com/ExerciseCoding/gateway">帮助文档</a>&nbsp;&nbsp;&nbsp;</span>
    <el-dropdown>
      <i class="el-icon-setting" style="margin-right: 25px"></i>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item @click.native="logout">登出</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
<!--    <span>{{ userinfo.username }}</span>-->
    <span>winston</span>
  </el-header>
</template>

<script>
import { Bus } from '../utils/bus'

export default {
  name: "headers",
  data() {
    return {
      userinfo: {
        username: '',
      },
    }
  },
  mounted: function() {
    //bus
    Bus.$on('change', () => {
      this.userItem = localStorage.getItem("userinfo")
      if (this.userItem != null) {
        this.$set(this.userinfo, 'username', JSON.parse(this.userItem).username)
      }
    })
  },
  methods: {
    logout() {
      this.$store.commit('changeFirst', false)
      console.log("logout")
      this.$axios.get('api/user/logout')
        .then((resp) => {
          console.log("resp: " + resp)
          var res = resp.data;
          console.log("res")
          console.log(res)
          if (res.code === 200) {
            localStorage.removeItem('userinfo');
            this.$router.push({ path: "/login" })
          }
        })
    }
  }
}
</script>

<style>
.el-header {
  background-color: white;
  /*box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);*/
  border-radius: 2px;
  color: #333333;
  line-height: 60px;

}

.logo {
  vertical-align: top
}
</style>
