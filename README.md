# simple douyin demo
严格采用典型的接口层、业务层、持久化层方式，使用mysql数据库持久化，常见查询操作已设置索引，gorm处理CRUD操作，gin配置路由，视频和封面文件皆保存本地提供下载服务，所有接口请求都已实现，其中使用ffmpeg命令获取视频封面，略为简略的实现，可改进之处较多，尚无云服务器支持，本地测试良好。
