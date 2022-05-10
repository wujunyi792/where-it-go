# 大数据行程卡接口

## 声明
该程序仅仅用于学习用途！！！ 

该程序仅仅用于学习用途！！！ 

该程序仅仅用于学习用途！！！ 

请严禁用作非法用途！！！ 

法网恢恢，疏而不漏！！！ 

若不遵守！ 后果请自负！！！（~~~监狱饭真香~~~）

## 项目介绍
本服务相当于是个 API 服务器，可直接通过 API 查询到大数据通行码数据。目前的缺陷是，在官方返回的数据中 message 字段为 base64 图像编码，其内容为`您于14天内到达或途径....`，需要使用该字段数据需要配合OCR

详细接口参数即接口可见 internal/service/trace

官网地址`https://xc.caict.ac.cn`，前端是 Vue3-cli 项目，逆向不是很难。接口即 POST 参数均为随机字符串，不清楚有何用意。

## 服务接口
res.code==20000 为正常，其余均为失败（但是 http 响应码是 200）

### GET `/trace/cms/:phone` 发送验证码

### GET `/trace/:phone/:token` 获取行程卡数据

## 使用方式
```
go run cmd/main.go
```
此时会在 config 目录下生成一份配置文件，按照需要填写配置文件
```
{
  "MODE": "release",    // debug,release
  "ProgramName": "行程卡数据接口",
  "AUTHOR": "",
  "VERSION": "",
  "REDIS": {            // 是否启用redis,若不启用，则会使用内存缓存
    "Use": false,
    "Config": {
      "IP": "",
      "PORT": "",
      "PASSWORD": "",
      "DB": 0
    }
  },
  "OSS": {              // 是否启用阿里云oss，若不启用，行程信息图片base64编码将不会转换为图片url返回
    "Use": false,
    "Config": {
      "AccessKeySecret": "",
      "AccessKeyId": "",
      "EndPoint": "",
      "BucketName": "",
      "BaseURL": "",
      "Path": ""
    }
  },
  "OCR": {              // 是否启用行程信息图片base64，若不启用，则不会识别行程信息图片base64内容
    "Use": false,
    "Url": ""           // url为ocr接口地址，这边可以根据实际自定义 当前服务使用serveless参考 https://github.com/hduhelp/paddlehub_ppocr
  }
}
```

配置编辑完成后重新运行 `go run cmd/main.go` 启动 web 服务器