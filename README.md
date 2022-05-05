# 大数据行程卡接口

## 项目介绍
本服务相当于是个 API 服务器，可直接通过 API 查询到大数据通行码数据。目前的缺陷是，在官方返回的数据中 message 字段为 base64 图像编码，其内容为`您于14天内到达或途径....`，需要使用该字段数据需要配合OCR

详细接口参数即接口可见 internal/service/trace

官网地址`https://xc.caict.ac.cn`，前端是 Vue3-cli 项目，逆向不是很难。接口即 POST 参数均为随机字符串，不清楚有何用意。

## 服务接口
res.code==20000 为正常，其余均为失败（但是 http 响应码是 200）

### GET `/trace/cms/:phone` 发送验证码

### GET `/trace/:phone/:token` 获取行程卡数据