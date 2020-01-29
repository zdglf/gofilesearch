<template>
  <el-container>
    <el-header></el-header>
    <el-main>
      <el-image :src="icon_url">
      </el-image>
      <el-row :gutter="20">
        <el-col :span="18">
          <el-input v-model="input"  @keyup.native.enter="clickFileSearchEvent" placeholder="请输入搜索内容"></el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="clickFileSearchEvent">搜索</el-button>
        </el-col>
      </el-row>
    </el-main>
    <el-footer></el-footer>
  </el-container>
</template>
<script>
export default {
  name: 'HomeSearchPage',
  data () {
    return {
      input: '',
      icon_url: require('../assets/logo.png')
    }
  },
  methods: {
    clickFileSearchEvent: function () {
      console.log(this.input)
      this.$router.push({
        name: 'SearchResultPage',
        params: {
          keyword: this.input
        }
      })
    }
  }
}
</script>
