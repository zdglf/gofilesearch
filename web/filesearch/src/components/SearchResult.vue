<template>
  <el-container>
    <el-header>
      <el-row :gutter="20">
        <el-col :span="16">
          <el-input v-model="input" placeholder="请输入搜索内容" clearable @keyup.native.enter="clickFileSearch" ></el-input>
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
              <span>{{item.name}}</span>
              <el-button style="float: right; padding: 3px 0" type="text">复制地址</el-button>
            </div>
            <div v-for="hit in item.desc" :key="hit">
              {{hit}}
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
  name: 'SearchResult',
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
    requestFileSearch: function () {
      var self = this
      this.$http.post('/search/doc', {
        'keyword': this.input,
        'pageIndex': this.pageIndex
      }).then((response) => {
        if (response.data.code !== 0) {
          console.log(response)
          self.$notify({
            title: '检索失败',
            message: response.data.msg
          })
          return
        }
        self.total = response.data.total
        self.pageSize = response.data.count
        self.currentPage = self.pageIndex + 1
        self.search_result = response.data.data
        console.log('total: ' + self.total)
      }, (response) => {
        console.log(response)
        self.$notify({
          title: '检索失败',
          message: 'Server Status:' + response.status
        })
      }).catch((e) => {
        console.log(e)
        self.$notify({
          title: '检索失败',
          message: e
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
