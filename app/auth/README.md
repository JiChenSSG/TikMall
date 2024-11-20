# 鉴权文档

## 结构

鉴权模块集成了casbin，实现了动态的权限管理，支持多种权限模型。

访问接口流程

1. 用户访问接口
2. 网关拦截请求，将token发送给鉴权模块
3. 鉴权模块解析token，获取用户信息
4. 返回是否有权限

同时，使用长短token的模式实现token的续费与用户登出

登录流程

1. 用户访问登录接口
2. 确认用户信息后，生成access_token和refresh_token，并返回给用户，同时将refresh_token存入redis
3. 用户使用access_token访问接口
4. access_token过期，用户将refresh_token发送给鉴权模块
5. 鉴权模块验证refresh_token，返回新的access_token
6. 用户使用新的access_token访问接口
7. 用户登出，删除refresh_token

## 使用

权限文件位置: `app/auth/config/xxx/policy.csv`

```csv
p, user, /product/get, GET
p, user, /product/search, GET
```

如图配置权限即可，url支持正则表达式