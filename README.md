# openid-connect-proxy

`openid-connect-proxy` 是一个用于衔接 [OpenID Connect](https://openid.net/developers/how-connect-works/) 与 [MediaWiki OAuth Extension](https://www.mediawiki.org/wiki/Extension:OAuth) 交互流程的兼容层及代理服务器。它允许用户将 OpenID Connect 请求转发到指定的上游 [MediaWiki](https://www.mediawiki.org/wiki/MediaWiki) 服务器，并对响应进行必要的转换。兼容层亦会对 `"redirect_uri"` 进行替换以方便在代理环境中使用。

## 功能

1. 监听指定端口，接收并处理 OpenID Connect 请求。
2. 将请求转发到指定的上游服务器。
3. 根据配置，对请求和响应进行必要的转换。
4. 返回处理后的响应给客户端。

### 请求转换

- 修改请求的查询参数：
  - 如果查询参数中包含 `"redirect_uri"` ，则将其值转换为新的重定向 URI。
  - 如果查询参数中包含 `"scope"` ，则将其值中的`"profile"`替换为`"basic"`，将`"email"`替换为`"mwoauth-authonlyprivate"`，并丢弃其他值。
- 如果请求方法为 `POST` ，根据请求的 `Content-Type` 对请求体进行转换：
  - 如果 `Content-Type` 为 `"application/json"` ，则解析 JSON 请求体，修改 `"redirect_uri"` 的值为新的重定向 URI。
  - 如果 `Content-Type` 不是 `"application/json"` ，则解析表单请求体，修改 `"redirect_uri"` 的值为新的重定向 URI。

### 响应转换

- 如果响应体中包含 `"blocked"` 字段且其值为 `true` ，则修改为一个空的 JSON 对象。
- 否则，根据下列映射关系，将响应体中的旧键替换为新键。
  - `"username"` -> `"name"`
  - `"realname"` -> `"preferred_username"`
  - `"confirmed_email"` -> `"email_verified"`

## 用法

```bash
go run main.go --port=8071 --redirect=api.thwiki.cc --upstream=thwiki.cc
```

### 参数说明

- `--port`: 监听的端口号，默认为 `8000` 。
- `--redirect`: 重定向主机名，默认为 `""` 。
- `--upstream`: 上游服务器主机名，默认为 `""` 。

## 注意事项

1. 确保已安装 [Go 语言环境](https://go.dev/)。
2. 在运行程序之前，请确保已正确设置端口、重定向主机名和上游服务器主机名。
3. 上游服务器已安装并设置 [MediaWiki](https://www.mediawiki.org/wiki/Manual:Installing_MediaWiki) 及 [OAuth Extension](https://www.mediawiki.org/wiki/Extension:OAuth)
