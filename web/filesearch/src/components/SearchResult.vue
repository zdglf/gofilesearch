<template>
  <el-container>
    <el-header>
      <el-row :gutter="20">
        <el-col :span="16">
          <el-input v-model="input" placeholder="请输入搜索内容" clearable></el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary">搜索</el-button>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <div v-if="search_result.length===0">
        <el-image :src="not_found_image_url" fit></el-image>
      </div>
      <div v-else>
        <div  v-for="item in search_result" :key="item.id">
          <el-row :gutter="20">
            <el-col :span="10">
              <el-button type="text">{{item.title}}</el-button>
            </el-col>
            <el-col :span="10">
              <el-button type="text">{{item.hit}}</el-button>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="text">{{item.url}}</el-button>
            </el-col>
          </el-row>
          <el-divider></el-divider>
        </div>
      </div>
    </el-main>
    <el-footer>
      <el-row :gutter="20">
        <el-col :span="20">
          <el-pagination
            layout="prev, pager, next"
            :total="1000">
          </el-pagination>
        </el-col>
      </el-row>
    </el-footer>
  </el-container>
</template>

<script>
export default {
  name: 'SearchResult',
  data () {
    return {
      input: '',
      not_found_image_url: require('../assets/no_found.png'),
      msg: 'Welcom to Vue App Webpage!',
      search_result: [
        {
          id: '1',
          title: 'name',
          url: 'http://baidu.com',
          hit: 'text'
        },
        {
          id: '2',
          title: 'name',
          url: 'http://baidu.com',
          hit: 'text'
        },
        {
          id: '4',
          title: 'name',
          url: 'http://baidu.com',
          hit: 'text'
        }
      ]
    }
  }
}
</script>
