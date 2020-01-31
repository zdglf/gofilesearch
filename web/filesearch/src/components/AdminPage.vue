<template>
  <el-tabs type="border-card">
    <el-tab-pane label="任务管理">
      <el-button type="text" @click="taskDialogVisable = true">创建任务</el-button>
      <el-dialog title="创建任务" :visible.sync="taskDialogVisable">
        <el-form :model="taskCreateForm">
          <el-form-item label="类型" :label-width="taskCreateFormWidth">
            <el-radio-group v-model="taskCreateForm.type">
              <el-radio label="file" name="type"></el-radio>
              <el-radio label="ftp" name="type"></el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="目录" :label-width="taskCreateFormWidth">
            <el-input v-model="taskCreateForm.folder" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item label="用户名" :label-width="taskCreateFormWidth">
            <el-input v-model="taskCreateForm.userName" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item label="密码" :label-width="taskCreateFormWidth">
            <el-input v-model="taskCreateForm.password" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item label="正则表达" :label-width="taskCreateFormWidth">
            <el-input v-model="taskCreateForm.regular" autocomplete="off" ></el-input>
          </el-form-item>
          <el-form-item label="大小限制" :label-width="taskCreateFormWidth">
            <el-input-number v-model="taskCreateForm.sizeLimit" autocomplete="off" ></el-input-number>
          </el-form-item>
          <el-form-item label="协程数限制" :label-width="taskCreateFormWidth">
            <el-input-number v-model="taskCreateForm.processSize" autocomplete="off" :max="50" :min="1"></el-input-number>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="taskDialogVisable = false">取 消</el-button>
          <el-button type="primary" @click="requestTaskCreate">确 定</el-button>
        </div>
      </el-dialog>
      <el-table
        :data="taskDataList"
        style="width: 100%"
        max-height="500">
        <el-table-column
          prop="createAt"
          label="创建日期"
          width="150">
        </el-table-column>
        <el-table-column
          prop="type"
          label="类型"
          width="60">
        </el-table-column>
        <el-table-column
          prop="folder"
          label="目录"
          width="250">
        </el-table-column>
        <el-table-column
          prop="regular"
          label="匹配正则"
          width="120">
        </el-table-column>
        <el-table-column
          prop="sizeLimit"
          label="限制大小"
          width="100">
        </el-table-column>
        <el-table-column
          prop="processSize"
          label="协程数"
          width="80">
        </el-table-column>
        <el-table-column
          prop="lastRunningTime"
          label="最后执行时间"
          width="150">
        </el-table-column>
        <el-table-column
          fixed="right"
          label="操作"
          width="120">
          <template slot-scope="scope">
            <el-button
              @click.native="deleteRow(scope.$index, taskDataList)"
              type="text"
              size="small">
              移除
            </el-button>
            <el-button
              @click.native="execTask(scope.$index, taskDataList)"
              type="text"
              size="small">
              执行
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        @current-change="handleCurrengChange"
        @prev-click="handleCurrengChange"
        @next-click="handleCurrengChange"
        layout="prev, pager, next"
        :current-page="taskCurrentPage"
        :page-size="taskPageSize"
        :total="taskTotal">
      </el-pagination>
    </el-tab-pane>
    <el-tab-pane label="用户管理">
      未实现
    </el-tab-pane>
    <el-tab-pane label="日志管理">
      未实现
    </el-tab-pane>
  </el-tabs>
</template>
<script>
export default {
  name: 'AdminPage',
  data () {
    return {
      taskDialogVisable: false,
      taskDataList: [],
      taskPageIndex: 0,
      taskCurrentPage: 0,
      taskTotal: 0,
      taskPageSize: 0,
      taskCreateFormWidth: '120px',
      taskCreateForm: {
        type: 'file',
        folder: '',
        userName: '',
        password: '',
        enable: 0,
        regular: '\\.(docx|pdf|txt|md)$',
        timing: 0,
        sizeLimit: 20 * 1024 * 1024,
        processSize: 20
      }
    }
  },
  created: function () {
    this.requestTaskList()
  },
  methods: {
    deleteRow: function (index, data) {
      console.log(data)
      this.requestTaskDelete(data[index].id)
    },
    execTask: function (index, data) {
      console.log(data)
      this.requestTaskExec(data[index].id)
    },
    handleCurrengChange: function (selectedPageNo) {
      this.taskPageIndex = selectedPageNo - 1
      this.requestTaskList()
    },
    requestTaskCreate: function () {
      var self = this
      this.$http.post('/admin/task/create', this.taskCreateForm)
        .then((response) => {
          if (response.data.code !== 0) {
            console.log(response)
            self.$notify({
              title: '提示',
              message: '创建任务失败,' + response.data.msg,
              duration: 1500
            })
            return
          }
          this.taskDialogVisable = false
          self.$notify({
            title: '提示',
            message: '创建任务成功',
            duration: 1500
          })
        }, (response) => {
          console.log(response)
          self.$notify({
            title: '提示',
            message: '创建任务失败, Server Status:' + response.status,
            duration: 1500
          })
        }).catch((e) => {
          console.log(e)
          self.$notify({
            title: '提示',
            message: '创建任务失败,' + e,
            duration: 1500
          })
        })
    },
    requestTaskList: function () {
      var self = this
      this.$http.post('/admin/task/list', {
        'pageIndex': this.taskPageIndex
      }).then((response) => {
        if (response.data.code !== 0) {
          console.log(response)
          self.$notify({
            title: '提示',
            message: '获取任务列表失败,' + response.data.msg,
            duration: 1500
          })
          return
        }
        self.taskDataList = response.data.data
        self.taskTotal = response.data.total
        self.taskPageSize = response.data.count
        self.taskCurrentPage = self.taskPageIndex + 1
      }, (response) => {
        console.log(response)
        self.$notify({
          title: '提示',
          message: '获取任务列表失败, Server Status:' + response.status,
          duration: 1500
        })
      }).catch((e) => {
        console.log(e)
        self.$notify({
          title: '提示',
          message: '获取任务列表失败,' + e,
          duration: 1500
        })
      })
    },
    requestTaskExec: function (id) {
      var self = this
      this.$http.post('/admin/task/exec', {
        'id': id
      }).then((response) => {
        if (response.data.code !== 0) {
          console.log(response)
          self.$notify({
            title: '提示',
            message: '执行任务失败,' + response.data.msg,
            duration: 1500
          })
          return
        }
        self.$notify({
          title: '提示',
          message: '执行任务成功',
          duration: 1500
        })
      }, (response) => {
        console.log(response)
        self.$notify({
          title: '提示',
          message: '执行任务失败，Server Status:' + response.status,
          duration: 1500
        })
      }).catch((e) => {
        console.log(e)
        self.$notify({
          title: '提示',
          message: '执行任务失败，' + e,
          duration: 1500
        })
      })
    },
    requestTaskDelete: function (id) {
      var self = this
      this.$http.post('/admin/task/delete', {
        'id': id
      }).then((response) => {
        if (response.data.code !== 0) {
          console.log(response)
          self.$notify({
            title: '提示',
            message: '删除任务失败,' + response.data.msg,
            duration: 1500
          })
          return
        }
        self.$notify({
          title: '提示',
          message: '删除任务成功',
          duration: 1500
        })
        self.taskPageIndex = 0
        self.requestTaskList()
      }, (response) => {
        console.log(response)
        self.$notify({
          title: '提示',
          message: '删除任务失败，Server Status:' + response.status,
          duration: 1500
        })
      }).catch((e) => {
        console.log(e)
        self.$notify({
          title: '提示',
          message: '删除任务失败，' + e,
          duration: 1500
        })
      })
    }
  }
}
</script>
