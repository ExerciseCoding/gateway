<template>
  <!-- <el-scrollbar wrapClass="scrollbar-wrap" :style="{height: scrollHeight}" ref="scrollbarContainer"> -->
    <div class="home">
      <el-breadcrumb separator-class="el-icon-arrow-right" style="padding:5px;margin-bottom: 15px">
        <el-breadcrumb-item :to="{ path: '/' }"><i class="el-icon-s-home"></i>&nbsp;首页</el-breadcrumb-item>
      </el-breadcrumb>
      <div style="padding-bottom: 20px;padding-top: 10px">
        <div>
          <span style="font-size: 16px">&nbsp;<i class="el-icon-position"></i>
            内容建设中
          </span>
        </div>
      </div>

      <div class="block" style="background-color: unset;text-align: center;">
        <el-pagination style="background-color: unset;margin-top:20px ;" small @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage" :page-sizes="pageSizes" :page-size="PageSize" layout="total, sizes, prev, pager, next, jumper" :total="totalCount">
        </el-pagination>
      </div>
    </div>
  <!-- </el-scrollbar> -->
</template>
<script>

export default {
  name: 'Index',
  data() {
    return {
      scrollHeight: '0px',
      input: '',
      ticketList: [],
      // 默认显示第几页
      currentPage: 1,
      // 总条数，根据接口获取数据长度(注意：这里不能为空)
      totalCount: 1,
      // 个数选择器（可修改）
      pageSizes: [10, 20, 50, 100],
      // 默认每页显示的条数（可修改）
      PageSize: 10,
      NOCApprove: false,
      AMApprove: false,
      ManagerApprove: false,
      root: false,
      pic: '', //pic组别flag
      public: false, //pic可编辑public
    }
  },
  mounted: function() {
    this.scrollHeight = window.innerHeight * 0.9 + 'px';
    let userinfo = JSON.parse(localStorage.getItem('userinfo'))
    if (userinfo != null) {
      this.user = userinfo.username
    } else {
      this.user = 'None'
    }
    this.queryTicket()
    this.queryPic()
  },
  methods: {
    queryTicket() {

    },
    Edit(ticket_id) {
      this.$router.push({ path: "/updateTicket", query: { ticket_id: ticket_id } })
    },

    View(ticket_id) {
      this.$router.push({ path: "/viewTicket", query: { ticket_id: ticket_id } })
    },

    Handle(ticket_id) {
      this.$router.push({ path: "/toProblem", query: { ticket_id: ticket_id } })
    },

    // 分页
    // 每页显示的条数
    handleSizeChange(val) {
      // 改变每页显示的条数
      this.PageSize = val
      // 注意：在改变每页显示的条数时，要将页码显示到第一页
      this.currentPage = 1
    },
    // 显示第几页
    handleCurrentChange(val) {
      // 改变默认的页数
      this.currentPage = val
    },

    //查询组别
    queryPic() {
      // this.$axios.post('/api/user/message', { "action": "query", "flag": "pic" })
      //   .then((resp) => {
      //     if (resp.data.error_code === 0) {
      //       let piclist = resp.data.data
      //       for (let i = 0; i < piclist.length; i++) {
      //         if (this.user === piclist[i]['username']) {
      //           this.pic = piclist[i]['group']
      //           this.public = true
      //         }
      //       }
      //     } else {
      //       this.$message.error('查询pic列表失败 :' + resp.data.message)
      //     }
      //   })
    },

  },
  created: function() {
    // let userinfo = JSON.parse(localStorage.getItem('userinfo'))
    // console.log("userinfo: " + userinfo)
    // if (userinfo != null) {
    // console.log("userinfo is not null")
    // this.user = userinfo.username
    // } else {
    // console.log("userinfo is null")
    // this.user = 'None'
    // }
    // console.log("当前登录用户: " + this.user)
    // this.queryTicket()
  }
}

</script>
<!-- Add "scoped" attribute to li
mit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.el-scrollbar {
  height: 90%;

}



</style>
