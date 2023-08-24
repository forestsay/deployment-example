# Deployment Example

Ryze 业务部署模板

## 环境变量注入

环境变量的注入方式是：在 `manifests` 文件夹里的所有 `deployment` 类型配置中，给 Containers 的 Env 数组添加环境变量。

仅对使用 Ryze 的容器生效

```text
// 静态注入
DREAM_APP_NAME:              应用名
DREAM_UNIT_NAME:             应用名
DREAM_APP_VERSION:           CommitHash 前八位
DREAM_APP_ROOT:              /data/apps/应用名
DREAM_ENV:                   环境名
DREAM_REGION_NAME:           部署地域
DREAM_SERVICE_DISCOVERY_URI: consul://consul-server.consul.svc.cluster.local:8500  consul EndPoint
DREAM_IMAGE_TAG:             镜像 Tag

// 引用注入
NODE_NAME:       部署节点名
POD_NAME:        POD 名
POD_IP:          POD IP
DREAM_NODE_NAME: 部署节点名
DREAM_POD_NAME:  POD 名
DREAM_POD_IP:    POD IP
```
