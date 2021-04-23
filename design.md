# zero_blog


# 功能设计:
用户相关
评论及留言、反馈
分类管理 二级标签
文章管理
日志管理
赞赏

# 表设计：
## 用户表 user
id         用户id
nick_name  用户名
account    账号
password   密码
open_id    微信唯一id
level      身份级别

## 文章表 essay
id               文章id
user_id          用户id(预留字段)
title            标题
describe         描述
photo_uerl       图片
display_statu    显示状态    


## 评价留言反馈表 estimate
essay_id与estimate_id互斥

id
user_id         用户id
essay_id        文章id,0为系统
estimate_id     评论表id, 
type            类型 1-评论,2-留言，3-反馈
content         文本有上限好了

## 文章分类表 classify
id                  分类
classify_parent_id  父节点id
level               级别 1,2,3,4,5

## 文章内容表 essay_text
id
essay_id        文章id
content         文章内容

## 赞赏记录表
id
user_id         账号id,可以为空
amount          多少钱
way             途径,支付宝、微信

## 日志表
id
user_id
url
status
code
