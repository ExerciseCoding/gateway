<template>
  <div class="page-container">
    <el-header style="height: 200px"></el-header>
    <el-container>
      <el-main>
        <div>
          <div class="logo-container">
            <img src="../assets/logo-apollo.png" height="45" width="60" class="logo">
            <div class="line-left">
              <span class="title">Operation Data Analysis Platform</span>
            </div>
          </div>
          <div class="button-container">
            <p>Welcome! Please sign in with Google account</p>
            <el-button class="login-button" type="primary" @click="login">SIGN IN</el-button>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>
<script>
// import { getUserinfo } from "../../axios/api";
// import { BASE_URL } from "../../axios/request"
import axios from "axios";
axios.defaults.baseURL = process.env.API_ROOT
console.log(process.env.API_ROOT)

export default {
  created() {
    //this.checkLogin()
  },
  methods: {
    async login() {
      window.location.href = axios.defaults.baseURL + "/api/user/login/";

    }
  }
};

</script>
<style>
.page-container {
  height: calc(100vh - 20px);
  background: #f1f2f5;
}

.logo-container {
  font-size: 25px;
  font-weight: 500;
  margin: 0 auto;
  /*align-content: center;*/
  text-align: center;
}

.button-container {
  width: 550px;
  height: 120px;
  background: #fff;
  border-radius: 0px;
  margin: 0 auto;
  padding: 25px 0;
  line-height: 50px;
  font-size: 16px;
  text-align: center;

}

.logo {
  vertical-align: top
}

.title {
  line-height: 2;
}

.line-left {
  display: inline;
  position: relative;
  left: 25px;
  top: 0;
  width: 50px;
  padding: 0 30px;
}

.line-left:after {
  content: '';
  position: absolute;
  left: 0;
  top: 20%;
  bottom: auto;
  right: auto;
  height: 60%;
  width: 3px;
  background-color: #dde0e7;
}

.login-button {
  border-color: #e05739;
  background-color: #e05739;
  width: 150px;
}

</style>
