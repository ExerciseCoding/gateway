<template>
  <div id="app">
    <el-container v-if="$route.meta.keepAlive" style=" border: 1px solid #eee; width: 100%;position: fixed;">
      <!-- 导航栏 -->
      <el-header style="text-align: right; font-size: 12px">
        <keep-alive>
          <headerNav></headerNav>
        </keep-alive>
      </el-header>
      <el-container>
        <!-- 导航栏 -->
        <keep-alive>
          <left></left>
        </keep-alive>
        <el-main>
            <!-- 组件渲染后会替代router-view，默认"/"路径下组件替代-->
            <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
    <router-view v-if="!$route.meta.keepAlive"></router-view>
  </div>
</template>
<script>
import left from "./components/Left";
import header from "./components/Headers";

export default {
  name: 'App',
  data() {
    return{
      note: {
            backgroundImage: "url(" + require("./assets/logo-apollo.png") + ")",
            backgroundRepeat: "no-repeat",
            backgroundSize: "25px auto",
            marginTop: "5px",
          },
    }
  },
  components: {
    headerNav: header,
    left: left,
  }
}

</script>
<style>
.app {
  background-image: url("./assets/logo-apollo.png");
}
.el-main {
  background-color: #f0f2f5;
  /*box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)*/

}


.el-scrollbar {
  height: 100%;
}

.el-scrollbar__wrap {
  width: calc(100% + 17px);
  /*将垂直滚动条挤到可视区域之外*/
  height: calc(100% + 17px);
  /*将水平滚动条挤到可视区域之外*/
}

</style>
