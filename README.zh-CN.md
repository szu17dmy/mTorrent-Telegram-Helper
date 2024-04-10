# mTorrent Telegram 助手

项目状态：🚧 仍在开发中 Work-in-Progress (WIP)

M-Team 迁移至 mTorrent 后，给出了 OpenAPI 令牌和文档，故开此坑以整活。

## 已实现的功能

+ [x] 每周活动置顶自动推送

## Todolist

+ [ ] 推送通知
+ [ ] 私有部署时允许推送图片等信息

......

## 部署

公开部署见 [t.me/MTeam_Helper](https://t.me/MTeam_Helper)。

目前私有部署需要自行使用`build/package/Dockerfile`构建镜像，然后`kubectl apply -f deploy/deployment.yaml`。
