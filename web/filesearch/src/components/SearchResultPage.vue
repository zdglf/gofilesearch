<template>
  <el-container>
    <el-header>
      <el-row :gutter="20">
        <el-col :span="16">
          <el-input v-model="input" placeholder="请输入搜索内容"  clearable @keyup.native.enter="clickFileSearch" >
            <i slot="prefix" class="el-input__icon el-icon-search"></i>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="clickFileSearch">搜索</el-button>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <div v-show="search_result.length===0">
        <el-image :src="not_found_image_url" fit></el-image>
      </div>
      <div v-show="search_result.length > 0">
        <div  v-for="item in search_result" :key="item.id">
          <el-card>
            <div slot="header" class="clearfix">
              <span style="float: left; padding: 3px 0" >{{item.name}}</span>
              <el-button style="float: right; padding: 3px 0" type="text" v-clipboard:copy="item.url" v-clipboard:success="onCopySuccess" v-clipboard:error="onCopyError">复制地址</el-button>
            </div>
            <div v-for="hit in item.desc" :key="hit">
              <div v-html="hit" style="float: left; padding: 3px 0" >
              {{hit}}
              </div>
            </div>
          </el-card>
        </div>
      </div>
    </el-main>
    <el-footer>
      <el-row :gutter="20">
        <el-col :span="20">
          <el-pagination
            @current-change="handleCurrengChange"
            @prev-click="handleCurrengChange"
            @next-click="handleCurrengChange"
            layout="prev, pager, next"
            :current-page="currentPage"
            :page-size="pageSize"
            :total="total">
          </el-pagination>
        </el-col>
      </el-row>
    </el-footer>
  </el-container>
</template>

<script>
export default {
  name: 'SearchResultPage',
  created: function () {
    var keyword = this.$route.params.keyword
    if (keyword !== undefined && keyword !== null) {
      this.input = keyword
      this.requestFileSearch()
    }
  },
  methods: {
    clickFileSearch: function () {
      if (this.input !== undefined && this.input !== null && this.input !== '') {
        this.requestFileSearch()
      }
    },
    handleCurrengChange: function (selectedPageNo) {
      this.pageIndex = selectedPageNo - 1
      this.requestFileSearch()
    },
    onCopySuccess: function (e) {
      this.$notify({
        title: '复制成功',
        message: e.text,
        duration: 1500
      })
    },
    onCopyError: function (e) {
      this.$notify({
        title: '复制失败',
        message: '浏览器复制内容失败',
        duration: 1500
      })
    },
    requestFileSearch: function () {
      var self = this
      this.$http.post('/search/doc', {
        'keyword': this.input,
        'pageIndex': this.pageIndex
      }).then((response) => {
        if (response.data.code !== 0) {
          console.log(response)
          self.$notify({
            title: '提示',
            message: '检索失败, ' + response.data.msg,
            duration: 1500
          })
          return
        }
        self.total = response.data.total
        self.pageSize = response.data.count
        self.currentPage = self.pageIndex + 1
        self.search_result = response.data.data
      }, (response) => {
        console.log(response)
        self.$notify({
          title: '提示',
          message: '检索失败，Server Status:' + response.status,
          duration: 1500
        })
      }).catch((e) => {
        console.log(e)
        self.$notify({
          title: '提示',
          message: '检索失败,' + e,
          duration: 1500
        })
      })
    }
  },
  data () {
    return {
      input: '',
      pageIndex: 0,
      pageSize: 0,
      total: 0,
      currentPage: 0,
      not_found_image_url: require('../assets/no_found.png'),
      msg: 'Welcom to Vue App Webpage!',
      search_result: [
      ]
    }
  }
}
</script>

<style>
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }

  .box-card {
    width: 480px;
  }
</style>
