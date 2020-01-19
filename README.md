### 文件搜索


#### ES设计
 
> INDEX: file_search

> Document: doc

|es字段|字段类型|说明|
|:----:|:----:|:----:|
|id        |text   | 文件sha256值 [记录唯一值] |
|content   |text   | 文件内容 [加入中文搜索]    |
|url       |text   | 文件URI                  |
|click_count|int    | 搜索内容点击次数           | 
|create_at  |date   | 记录创建时间              |
|file_name  |text   | 文件名 [加入中文搜索]     |


#### 数据库设计

> DB: file_search

> table : file_spider

|字段|字段类型|说明|
|:----:|:----:|:----:|
|id       |varchar(32)  | uuid|
|type     |varchar(32)  | 文件类型,ftp,file,svn,nfs,等|
|folder   |text         | 爬虫目录|
|username |varchar(64)  | 用户名|
|password |varchar(64)  | 密码|
|create_at |date   | 记录创建时间|
|enable   |int    | 是否启用|
|regular  |text   | 文件名匹配正则表达式|
|timing   |int    | 启用后间隔天数执行  |
|last_running_time|date    | 上次执行时间|
|size_limit|int    | 文件大小限制 |
|process_size|int|同时处理的协程|

#### mysql 启动

> sudo docker run --name mysql -p 3306:3306 -v /Users/zhangmike/Documents/data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

#### es7.x 启动

> sudo docker run -p 9200:9200 -p 9300:9300 -v /Users/zhangmike/Documents/data/es:/usr/share/elasticsearch/data -e "discovery.type=single-node" -d elasticsearch-ik-pinyin:7.4.2

##### 初始化es7.x mapping

```
curl 'http://localhost:9200/file_search'  -H "Content-Type: application/json" -X PUT --data '{"settings":{"index.analysis.analyzer.default.type":"ik_max_word"},"mappings":{"properties":{"url":{"type":"text","analyzer":"ik_max_word"},"content":{"type":"text","analyzer":"ik_max_word"},"file_name":{"type":"text","analyzer":"ik_max_word"}}}}' 
```



#### 接口

##### 创建文件爬虫任务
    method:Post
    format:Json
    param:
    {

    }

##### 搜索接口

    method:Post 
    format:Json
    param: 
    {
      key: 'keyword',
      pageIndex: '0'
    }

    response:
    {
      code:'',
      msg:'',
      page: {
        index:0,
        total:10
      }
      data:[
        {
          'id':'',
          'name': '',
          'createAt': '',
          'desc': ‘’,
          'url': '',
        }
      ]
    }

### TODO
1. 文件GBK 转 UTF8
2. 多关键字查询
3. 定义实现接口
4. 实现前段查询
5. 实现后台控制管理
6. 实现后台管理登录
7. 实现后台用户管理
