# gin_realword

### 日志: logrus

### jwt: golang-jwt


### viper: 配置


### docker mysql

```markdown
docker run 
-p 3305:3306 // 端口映射
--name go_realworld 
-e MYSQL_ROOT_PASSWORD=123456 
-d mysql 
--character-set-server=utf8mb4 
--collation-server=utf8mb4_unicode_ci
```