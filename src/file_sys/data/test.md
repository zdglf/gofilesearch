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