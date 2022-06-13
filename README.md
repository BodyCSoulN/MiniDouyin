# MiniDouyin ![](https://img.shields.io/badge/go--version-1.18.x-green)

A simplifed implementation of Douyin.

# 相关资料

- [ `APP` 下载](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
- [ `API` 详情文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)

# 项目结构

```
├── config
├── controller    
├── middleware
├── model
├── public
├── router
├── service
├── storage
└── utils
```

# 开发规范

# 分工 （不代表最终，后续可能会更改）

API | 负责
:--- | :---:
测试 | 彭，谷
/douyin/user/register/<br/>/douyin/user/login/<br/>/douyin/user/ | 祝
/douyin/publish/action/<br/>/douyin/publish/list/ | 刘
/douyin/comment/action/<br/>/douyin/comment/list/ | 李
/douyin/relation/action/<br/>/douyin/relation/follow/list/<br/>/douyin/relation/follower/list/<br/>/douyin/feed/<br/>/douyin/favorite/action/<br/>/douyin/favorite/list/ | 张

# 相关技术

- [Gin](https://gin-gonic.com/)  
- [Mysql](https://www.mysql.com/)  
- [Redis](https://redis.io/)  ~~可能要用吧~~

# 如何运行

进入项目文件夹，直接编译运行
```shell
go build && ./MiniDouyin
```
