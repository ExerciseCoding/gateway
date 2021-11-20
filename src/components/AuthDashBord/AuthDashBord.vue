<template>
  <div>
    <el-tabs type="border-card" style="height: 100%;">
      <el-tab-pane label="用户角色管理">
        <el-row>
          <el-col :span="24">
            <el-form :model="queryUserRoles" :inline="true" label-width="120px">
              <el-form-item  label="请输入用户名">
                <el-input v-model="queryUserRoles.username"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="queryByParams('userroles')">查询</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="24">
            <el-table :row-key="getRowKey" :data="userRoleList.slice((currentPage-1)*PageSize,currentPage*PageSize)" border max-height="550px" style="width: 100%;" @selection-change="handleSelectionChange">
              <el-table-column type="selection" :reserve-selection="true" width="55">
              </el-table-column>
              <!-- <el-table-column prop="id" label="序号" width="80">
                <template slot-scope="scope"> {{ scope.row.id }}</template>
              </el-table-column> -->
              <el-table-column prop="username" label="用户名">
                <template slot-scope="scope"> {{ scope.row.username }}</template>
              </el-table-column>
              <el-table-column prop="roles" label="关联角色">
                <template slot-scope="scope"> {{ scope.row.roles }}</template>
              </el-table-column>
              <el-table-column prop="updatetm" label="更新时间" sortable> 
                <template slot-scope="scope"> {{ scope.row.updatetm }}</template>
              </el-table-column>
            </el-table>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <div class="block" style="background-color: unset;text-align: right;">
              <el-pagination small style="background-color: unset;margin-top:20px ;" @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page.sync="currentPage" :page-sizes="pageSizes" :page-size="PageSize" layout="total, sizes, prev, pager, next, jumper" :total=userRoleListCount>
              </el-pagination>
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-col style="text-align: center; padding-top: 50px" :span="24">
            <el-button style="width: 180px;margin-right: 20px" type="primary" @click="ModifyRolesFormVisible = true">修改用户角色</el-button>
            <el-dialog title="修改用户角色" :visible.sync="ModifyRolesFormVisible" :append-to-body="true">
              <el-table :data="modifyUserRoleList" border style="width: 100%">
                <el-table-column prop="username" label="用户名">
                  <template slot-scope="scope"> {{ scope.row.username }}</template>
                </el-table-column>
                <el-table-column prop="roles" label="关联角色">
                  <template slot-scope="scope"> 
                    <el-select v-model="scope.row.roles" placeholder="请选择要关联的角色" multiple filterable remote clearable :remote-method="remoteRoles" :loading="loading" style="width: 100%">
                      <el-option v-for="item in rolelist" :key="item.index" :label="item.role" :value="item.role">
                      </el-option>
                  </el-select>
                  </template>
                </el-table-column>
              </el-table>
              <div slot="footer" class="dialog-footer">
                <el-button @click="ModifyRolesFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="updateInfo('userroles')">确 定</el-button>
              </div>
            </el-dialog>

            <el-button style="width: 180px;" type="primary" @click="AddUserFormVisible = true">新增用户</el-button>    
            <el-dialog title="新增用户" :visible.sync="AddUserFormVisible" :append-to-body="true">
              <el-form :model="addUserRole" :inline="true">
                <el-form-item label="新增用户" :label-width="formLabelWidth">
                  <el-input v-model="addUserRole.username" placeholder="新增用户">
                  </el-input>
                </el-form-item>
                <el-form-item label="请选择要关联的角色" :label-width="formLabelWidth">
                   <el-select v-model="addUserRole.roles" placeholder="请选择要关联的角色" multiple filterable remote clearable :remote-method="remoteRoles" :loading="loading" style="width: 100%">
                    <el-option v-for="item in rolelist" :key="item.index" :label="item.role" :value="item.role">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-form>
              <div slot="footer" class="dialog-footer">
                <el-button @click="AddUserFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="updateInfo('adduserroles')">确 定</el-button>
              </div>  
            </el-dialog>

          </el-col>
        </el-row>
      </el-tab-pane>

      <el-tab-pane label="视图权限管理">
        <el-row>
          <el-col :span="24">
            <el-form :model="queryViewRoles" :inline="true" label-width="120px">
              <el-form-item  label="请输入视图名">
                <el-input v-model="queryViewRoles.viewname"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="queryByParams('viewroles')">查询</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="24">
            <el-table :row-key="getRowKey" :data="viewRoleList.slice((currentPage_1-1)*PageSize_1,currentPage_1*PageSize_1)" border max-height="550px" @selection-change="handleSelectionChange_1">
              <el-table-column type="selection" :reserve-selection="true" width="55">
              </el-table-column>
              <!-- <el-table-column prop="id" label="序号" width="80">
                <template slot-scope="scope"> {{ scope.row.id }}</template>
              </el-table-column> -->
              <el-table-column prop="viewname" label="视图名">
                <template slot-scope="scope"> {{ scope.row.viewname }}</template>
              </el-table-column>
              <el-table-column prop="view" label="视图value">
                <template slot-scope="scope"> {{ scope.row.view }}</template>
              </el-table-column>
              <el-table-column prop="role" label="视图关联角色">
                <template slot-scope="scope"> {{ scope.row.roles }}</template>
              </el-table-column>
              <el-table-column prop="updatetm" label="更新时间" sortable> 
                <template slot-scope="scope"> {{ scope.row.updatetm }}</template>
              </el-table-column>
            </el-table>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <div class="block" style="background-color: unset;text-align: right;">
              <el-pagination small style="background-color: unset;margin-top:20px ;" @size-change="handleSizeChange_1" @current-change="handleCurrentChange_1" :current-page.sync="currentPage_1" :page-sizes="pageSizes_1" :page-size="PageSize_1" layout="total, sizes, prev, pager, next, jumper" :total=viewRoleListCount>
              </el-pagination>
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-col style="text-align: center; padding-top: 50px" :span="24">              
              <el-button style="width: 180px;" type="primary" @click="AddViewsRolesFormVisible = true">设置视图关联角色</el-button>    
              <el-dialog title="视图" :visible.sync="AddViewsRolesFormVisible" :append-to-body="true">
                <el-table :data="modifyViewRoleList" border style="width: 100%">
                  <el-table-column prop="view" label="视图名">
                    <template slot-scope="scope"> {{ scope.row.viewname }}</template>
                  </el-table-column>
                  <el-table-column prop="view" label="视图value">
                    <template slot-scope="scope"> {{ scope.row.view }}</template>
                  </el-table-column>
                  <el-table-column prop="roles" label="关联角色">
                    <template slot-scope="scope"> 
                      <el-select v-model="scope.row.roles" placeholder="请选择要关联的角色" multiple filterable remote clearable :remote-method="remoteRoles" :loading="loading" style="width: 100%">
                        <el-option v-for="item in rolelist" :key="item.index" :label="item.role" :value="item.role">
                        </el-option>
                    </el-select>
                    </template>
                  </el-table-column>
                </el-table>
                <div slot="footer" class="dialog-footer">
                  <el-button @click="AddViewsRolesFormVisible = false">取 消</el-button>
                  <el-button type="primary" @click="updateInfo('viewroles')">确 定</el-button>
                </div>
              </el-dialog>
          </el-col>
        </el-row>
      </el-tab-pane>

      <el-tab-pane label="操作日志">
        <el-row>
          <el-col :span="24">
            <el-form :model="queryOperLog" :inline="true" label-width="120px">
              <el-form-item label="操作日志类型">
                <el-select v-model="queryOperLog.type" placeholder="操作日志类型">
                  <el-option label="更新用户关联角色" value="删除用户关联角色"></el-option>
                  <el-option label="添加用户" value="添加用户"></el-option>
                  <el-option label="更新视图关联角色" value="删除视图关联角色"></el-option>
                  <el-option label="全部" value="all"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item  label="操作对象">
                <el-input v-model="queryOperLog.object"></el-input>
              </el-form-item>
              <el-form-item  label="操作者">
                <el-input v-model="queryOperLog.operator"></el-input>
              </el-form-item>
              <el-form-item  label="操作时间">
                <el-date-picker
                  v-model="queryOperLog.updatetm"
                  type="datetime"
                  placeholder="选择日期时间">
                </el-date-picker>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="queryByParams('operlogs')">查询</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="24">
            <el-table :data="operList.slice((currentPage_2-1)*PageSize_2,currentPage_2*PageSize_2)" border max-height="550px" style="width: 100%">
              <el-table-column type="selection" width="55">
              </el-table-column>
              <!-- <el-table-column prop="id" label="序号" width="80">
                <template slot-scope="scope"> {{ scope.row.id }}</template>
              </el-table-column> -->
              <el-table-column prop="type" label="操作日志类型">
                <template slot-scope="scope"> {{ scope.row.type }}</template>
              </el-table-column>
              <el-table-column prop="object" label="操作对象">
                <template slot-scope="scope"> {{ scope.row.object }}</template>
              </el-table-column>
              <el-table-column prop="value" label="操作内容"> 
                <template slot-scope="scope"> {{ scope.row.value }}</template>
              </el-table-column>
              <el-table-column prop="operator" label="操作者">
                <template slot-scope="scope"> {{ scope.row.operator }}</template>
              </el-table-column>
              <el-table-column prop="updatetm" label="操作时间" sortable> 
                <template slot-scope="scope"> {{ scope.row.updatetm }}</template>
              </el-table-column>
            </el-table>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <div class="block" style="background-color: unset;text-align: right;">
              <el-pagination small style="background-color: unset;margin-top:20px ;" @size-change="handleSizeChange_2" @current-change="handleCurrentChange_2" :current-page.sync="currentPage_2" :page-sizes="pageSizes_2" :page-size="PageSize_2" layout="total, sizes, prev, pager, next, jumper" :total=operListCount>
              </el-pagination>
            </div>
          </el-col>
        </el-row>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
export default {
  data() {
    return {
      rolelist: [],
      viewlist: [],
      userRoleList: [],
      userRoleListCount: 0,
      modifyUserRoleList: [],  //点击修改弹窗的数据
      addUserRole: {
        username: '',
        roles: []
      },
      viewRoleList: [],
      viewRoleListCount: 0,
      modifyViewRoleList: [],   //点击修改弹窗的数据
      addViewRole: {
        view: '',
        roles: []
      },
      operList: [],
      operListCount: 0,
      ModifyRolesFormVisible: false,
      AddUserFormVisible: false,
      AddViewsRolesFormVisible: false,
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      },
      formLabelWidth: '160px',

      queryUserRoles: {
        username: ''
      },
      queryViewRoles: {
        viewname: ''
      },
      queryOperLog: {
        type: '',
        object: '',
        operator: '',
        updatetm: ''
      },
      total: 1, //总条数
      currentPage: 1, //默认显示第几页
      PageSize: 10, //每页显示条数
      pageSizes: [10, 20, 50, 100, 1000], // 个数选择器（可修改）

      total_1: 1, //总条数
      currentPage_1: 1, //默认显示第几页
      PageSize_1: 10, //每页显示条数
      pageSizes_1: [10, 20, 50, 100, 1000], // 个数选择器（可修改）

      total_2: 1, //总条数
      currentPage_2: 1, //默认显示第几页
      PageSize_2: 10, //每页显示条数
      pageSizes_2: [10, 20, 50, 100, 1000], // 个数选择器（可修改）

      loading: false
    };
  },
  created: function() {
    this.queryRoles()
    this.queryViews()
  },
  mounted: function() {
    this.queryInfo('userroles')
    this.queryInfo('viewroles')
    this.queryInfo('operlogs')
  },
  methods: {
    queryInfo(flag) {
      this.$axios.post('api/auth/query', { flag: flag, action :"query"})
        .then((resp) => {
          let res = resp.data;
          if (res.error_code === 0) {
            for (let i = 0; i < res.data.length; i++) {
              // 修改utc时间显示
              res.data[i].updatetm = this.utc2local(res.data[i].updatetm)
            }
            if( flag === 'userroles'){
              this.userRoleList = res.data
              this.userRoleListCount = res.total
            } else if (flag === 'viewroles'){
              this.viewRoleList = res.data
              this.viewRoleListCount = res.total
            } else if (flag === 'operlogs'){
              this.operList = res.data
              this.operListCount = res.total
            }
          } else {
            this.$message.error('查询失败，请重试')
          }
        })
    },
    queryRoles(){
      this.$axios.get('api/auth/roles')
        .then((resp) => {
        if (resp.data.error_code === 0) {
          this.rolelist = resp.data.data
        } else {
          this.$message.error('查询失败 :' + resp.data.message)
          this.rolelist = []
        }
        })
    },
    queryViews(){
      this.$axios.get('api/auth/views')
        .then((resp) => {
        if (resp.data.error_code === 0) {
          this.viewlist = resp.data.data
        } else {
          this.$message.error('查询失败 :' + resp.data.message)
          this.viewlist = []
        }
        })
    },
    //远程搜索角色
    remoteRoles(query) {
      if (query !== '') {
        this.loading = true;
        setTimeout(() => {
          this.loading = false;
          this.rolelist = this.rolelist.filter(item => {
            return item.role.toLowerCase().indexOf(query.toLowerCase()) > -1;
          });
        }, 200);
      } else {
        this.rolelist = [];
      }
    },
    //远程搜索视图
    remoteViews(query) {
      if (query !== '') {
        this.loading = true;
        setTimeout(() => {
          this.loading = false;
          this.viewlist = this.viewlist.filter(item => {
            return item.view.toLowerCase().indexOf(query.toLowerCase()) > -1;
          });
        }, 200);
      } else {
        this.viewlist = [];
      }
    },
    handleSelectionChange(val) {
        this.modifyUserRoleList = JSON.parse(JSON.stringify(val))
        for (var i = 0; i < this.modifyUserRoleList.length; i++) {
            this.modifyUserRoleList[i].roles = this.modifyUserRoleList[i].roles.split(';').slice(1)
        }
    },
    handleSelectionChange_1(val) {
        console.log(val)
        this.modifyViewRoleList = JSON.parse(JSON.stringify(val))
        for (var i = 0; i < this.modifyViewRoleList.length; i++) {
            this.modifyViewRoleList[i].roles = this.modifyViewRoleList[i].roles.split(';').slice(1)
        }
    },
    updateInfo(flag) {
      let data = []
      if( flag === 'userroles'){
        data = this.modifyUserRoleList
        console.log('this.modifyUserRoleList')
        console.log(this.modifyUserRoleList)
      } else if (flag === 'adduserroles'){
        data = [this.addUserRole]
        // flag = 'userroles'  注释掉，因为可以直接传给后端，方便存储操作日志
        console.log('this.addUserRole')
        console.log(this.addUserRole)
      } else if (flag === 'viewroles'){
        data = this.modifyViewRoleList
        console.log('this.modifyViewRoleList')
        console.log(this.modifyViewRoleList)
      }
      this.$axios.post('api/auth/update', { flag: flag, data: data})
        .then((resp) => {
        if (resp.data.error_code === 0) {
          this.viewlist = resp.data.data
        } else {
          this.$message.error('查询失败 :' + resp.data.message)
          this.viewlist = []
        }
        })
      this.ModifyRolesFormVisible = false
      this.AddUserFormVisible = false
      this.AddViewsRolesFormVisible = false
      this.queryInfo('userroles')
      this.queryInfo('viewroles')
      this.queryInfo('operlogs')
    },
    queryByParams(type) {
      if( type === 'userroles'){
        let params = this.queryUserRoles
        console.log('this.queryUserRoles')
        console.log(this.queryUserRoles)
        if ((params.username === undefined) || (params.username === '')){
          this.queryInfo('userroles')
        } else {
          this.$axios.post('api/auth/query_by_param', { type: type, params: params})
            .then((resp) => {
              let res = resp.data;
              if (res.error_code === 0) {
                for (let i = 0; i < res.data.length; i++) {
                  // 修改utc时间显示
                  res.data[i].updatetm = this.utc2local(res.data[i].updatetm)
                }
                this.userRoleList = res.data
                this.userRoleListCount = res.total
              } else {
                this.$message.error('查询失败，请重试')
              }
            })
        }
      } else if (type === 'viewroles'){
        let params = this.queryViewRoles
        console.log('this.queryViewRoles')
        console.log(this.queryViewRoles)
        if ((params.viewname === undefined) || (params.viewname === '')){
          this.queryInfo('viewroles')
        } else {
          this.$axios.post('api/auth/query_by_param', { type: type, params: params})
            .then((resp) => {
              let res = resp.data;
              if (res.error_code === 0) {
                for (let i = 0; i < res.data.length; i++) {
                  // 修改utc时间显示
                  res.data[i].updatetm = this.utc2local(res.data[i].updatetm)
                }
                this.viewRoleList = res.data
                this.viewRoleListCount = res.total
              } else {
                this.$message.error('查询失败，请重试')
              }
            })
        }
      } else if (type === 'operlogs'){
        let params = this.queryOperLog
        console.log('this.queryOperLog')
        console.log(this.queryOperLog)
        if (((params.type === undefined) || (params.type === '') || (params.type === 'all')) && ((params.object === undefined) || (params.object === '')) && ((params.operator === undefined) || (params.operator === '')) && ((params.updatetm === undefined) || (params.updatetm === ''))){
          this.queryInfo('operlogs')
        } else {
          this.$axios.post('api/auth/query_by_param', { type: type, params: params})
            .then((resp) => {
              let res = resp.data;
              if (res.error_code === 0) {
                for (let i = 0; i < res.data.length; i++) {
                  // 修改utc时间显示
                  res.data[i].updatetm = this.utc2local(res.data[i].updatetm)
                }
                this.operList = res.data
                this.operListCount = res.total
              } else {
                this.$message.error('查询失败，请重试')
              }
            })
        }
      } 
    },

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

    handleSizeChange_1(val) {
      // 改变每页显示的条数
      this.PageSize_1 = val
      // 注意：在改变每页显示的条数时，要将页码显示到第一页
      this.currentPage_1 = 1
    },
    // 显示第几页
    handleCurrentChange_1(val) {
      // 改变默认的页数
      this.currentPage_1 = val
    },

    handleSizeChange_2(val) {
      // 改变每页显示的条数
      this.PageSize_2 = val
      // 注意：在改变每页显示的条数时，要将页码显示到第一页
      this.currentPage_2 = 1
    },
    // 显示第几页
    handleCurrentChange_2(val) {
      // 改变默认的页数
      this.currentPage_2 = val
    },
    getRowKey(row){
         return row.id;
    }, 


  }
}
</script>
<style>
</style>