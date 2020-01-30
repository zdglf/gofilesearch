### 文件搜索



#### 简介

平时接触的文档较多，找起来也不是方便，有时想不起来。决定结合ES,Gin，Vue，搞个文档搜索。方便自己查找文档。
目前支持的文件格式为，docx,pdf,txt,md. 
文件存储方式，file(本地)，ftp。


#### ES设计
 
> INDEX: file_search

> Document: _doc

|es字段|字段类型|说明|
|:----:|:----:|:----:|
|id        |text   | 文件sha256值 [记录唯一值] |
|content   |text   | 文件内容 [加入中文搜索]    |
|url       |text   | 文件URI[加入中文搜索]    |
|click_count|int    | 搜索内容点击次数         | 
|create_at  |date   | 记录创建时间            |
|file_name  |text   | 文件名                 |


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
|enable   |int    | 是否启用 (未实现)|
|regular  |text   | 文件名匹配正则表达式|
|timing   |int    | 启用后间隔天数执行(未实现)  |
|last_running_time|date    | 上次执行时间|
|size_limit|int    | 文件大小限制 |
|process_size|int|同时处理的协程|

#### mysql 启动

> sudo docker run -p 3306:3306 -v /Users/zhangmike/Documents/data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest


#### es7.x 启动

```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -d docker.elastic.co/elasticsearch/elasticsearch:7.4.2

docker exec -it [容器Id] /bin/sh
elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.4.2/elasticsearch-analysis-ik-7.4.2.zip 
elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-pinyin/releases/download/v7.4.2/elasticsearch-analysis-pinyin-7.4.2.zip  
docker commit [容器Id] elasticsearch-ik-pinyin:7.4.2
```

> sudo docker run -p 9200:9200 -p 9300:9300 -v /Users/zhangmike/Documents/data/es:/usr/share/elasticsearch/data -e "discovery.type=single-node" -d elasticsearch-ik-pinyin:7.4.2

##### 初始化es7.x mapping

```
curl 'http://localhost:9200/file_search'  -H "Content-Type: application/json" -X PUT --data '{"settings":{"index.analysis.analyzer.default.type":"ik_max_word"},"mappings":{"properties":{"url":{"type":"text","analyzer":"ik_max_word"},"content":{"type":"text","analyzer":"ik_max_word"},"file_name":{"type":"text","analyzer":"ik_max_word"}}}}' 
```

#### vue 

##### 安装

```
cd web/filesearch
//还需要elementui
npm install
npm run build
```

#### 运行

```
MYSQL_USER=root MYSQL_PASSWORD=123456 MYSQL_HOST=localhost MYSQL_PORT=3306 MYSQL_DBNAME=file_search go run main.go

//浏览器打开
//添加任务
http://localhost:8090/build/#/admin
//搜索
http://localhost:8090/build/#/
```

**文件目录修改，文档sha256值不变，不会重新提交ES,建议重新清除ES Index**

#### 接口

##### 创建文件爬虫任务
    url:/admin/task/create
    method:Post
    format:Json
    param:
    {
      type:'file',//文件类型目前只支持file,ftp
      folder:'',//文件目录
      userName:'',
      password:'',
      enable:1|0,//是否启动定时(未实现)
      regular:''//文件匹配正则表达式
      timing:1//定时启动（未实现）
      sizeLimit:1024//文件大小限制，字节
      processSize:10//协程处理大小

    }
    response:
    {
      code:'',
      msg:'',
    }

##### 执行文件爬虫任务
    url: /admin/task/exec
    method:Post
    format:Json
    param:
    {
      id: "id"
    }
    response:
    {
      code:'',
      msg:'',
    }
##### 查看文件爬虫任务列表
    url: /admin/task/list
    method:Post
    format:Json
    param:
    {
      pageIndex: 0
    }
    response:
    {
      code:'',
      msg:'',
      index:0,
      total:10,
      count:10,
      data:[{
        id:'id'
        type:'file',//文件类型目前只支持file
        folder:'',//文件目录
        userName:'',
        password:'',
        enable:1|0,//是否启动定时(未实现)
        regular:''//文件匹配正则表达式
        createAt:'',
        lastRunningTime: '',
        timing:''//定时启动（未实现）
        sizeLimit:1024//文件大小限制，字节
        processSize:10//协程处理大小

      }]
    }

##### 删除文件爬虫任务

    url: /admin/task/delete
    method:Post
    format:Json
    param:
    {
      id: "id"
    }
    response:
    {
      code:'',
      msg:'',
    }

##### 搜索接口
    url:/search/doc
    method:Post 
    format:Json
    param: 
    {
      keyword: 'keyword',
      pageIndex: '0'
    }

    response:
    {
      code:'',
      msg:'',
      index:0,
      total:10,
      count:10,
      
      data:[
        {
          'id':'',
          'name': '',
          'createAt': '',
          'desc': [‘’,''],
          'url': '',
        }
      ]
    }

### TODO (未实现)
1. 后台控制管理
2. 后台管理登录
3. 后台用户管理
4. SVN,WINODWS共享文件
5. 过滤HTML标签，关键字高亮
6. 定时任务