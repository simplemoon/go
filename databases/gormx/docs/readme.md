### 框架使用前的准备

### 文档地址

gorm 文档 [https://gorm.io/zh_CN/gen/sql_annotation.html]

####　安装的包

1. `gen`安装 `go install gorm.io/gen@latest` or `go get -u gorm.io/gen@latest`

2. 引用包安装

   ```shell
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/sqlite

   ```

#### 结论经验

1. 带有`Model(&data)` 或者 `&data` 这种类型的操作，都会将最新的值返回，体现在 `&data` 之中，不用重新查询一遍。

2. 权限控制

   - `->:opt`: `->` 代表读权限，`opt` 代表允许的类型
   - `<-:opt`: `<-` 代表写权限， `opt` 代表写的类型
   - `-:opt` `-` 代表忽略的操作，`opt` 代表忽略操作的类型

   ```golang
   type User struct {
   Name string `gorm:"<-:create"` // 允许读和创建
   Name string `gorm:"<-:update"` // 允许读和更新
   Name string `gorm:"<-"`        // 允许读和写（创建和更新）
   Name string `gorm:"<-:false"`  // 允许读，禁止写
   Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
   Name string `gorm:"->;<-:create"` // 允许读和写
   Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
   Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
   Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段
   Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段
   }
   ```

3. 字段说明

   - `embedded` 标签等效与将字段的等级提升以及，嵌入的结构带有此标签，自动嵌入到外层。
   - `autoCreateTime` 标签自动设置更新时间
   - `autoUpdateTime:milli` 标签使用毫秒时间戳，类型需要改成 int
   - `autoUpdateTime:nano` 标签使用纳秒时间戳，类型需要改成 int

4. 字段可选的 tag 说明

   | 标签名                 | 说明                                                                                                                                                                                                                                                                                                                                  |
   | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
   | column                 | 指定 db 列名                                                                                                                                                                                                                                                                                                                          |
   | type                   | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用,例如：not null、size, autoIncrement… 像 varbinary(8) 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT |
   | serializer             | 指定将数据序列化或反序列化到数据库中的序列化器, 例如: serializer:json/gob/unixtime                                                                                                                                                                                                                                                    |
   | size                   | 定义列数据类型的大小或长度，例如 size: 256                                                                                                                                                                                                                                                                                            |
   | primaryKey             | 将列定义为主键                                                                                                                                                                                                                                                                                                                        |
   | unique                 | 将列定义为唯一键                                                                                                                                                                                                                                                                                                                      |
   | default                | 定义列的默认值                                                                                                                                                                                                                                                                                                                        |
   | precision              | 指定列的精度                                                                                                                                                                                                                                                                                                                          |
   | scale                  | 指定列大小                                                                                                                                                                                                                                                                                                                            |
   | not null               | 指定列为 NOT NULL                                                                                                                                                                                                                                                                                                                     |
   | autoIncrement          | 指定列为自动增长                                                                                                                                                                                                                                                                                                                      |
   | autoIncrementIncrement | 自动步长，控制连续记录之间的间隔                                                                                                                                                                                                                                                                                                      |
   | embedded               | 嵌套字段                                                                                                                                                                                                                                                                                                                              |
   | embeddedPrefix         | 嵌入字段的列名前缀                                                                                                                                                                                                                                                                                                                    |
   | autoCreateTime         | 创建时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano                                                                                                                                                                                                        |
   | autoUpdateTime         | 创建/更新时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli                                                                                                                                                                                                  |
   | index                  | 根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情                                                                                                                                                                                                                                                            |
   | uniqueIndex            | 与 index 相同，但创建的是唯一索引                                                                                                                                                                                                                                                                                                     |
   | check                  | 创建检查约束，例如 check:age > 13，查看 约束 获取详情                                                                                                                                                                                                                                                                                 |
   | <-                     | 设置字段写入的权限， <-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限                                                                                                                                                                                                                                       |
   | ->                     | 设置字段读的权限，->:false 无读权限                                                                                                                                                                                                                                                                                                   |
   | -                      | 忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限                                                                                                                                                                                                                                                        |
   | comment                | 迁移时为字段添加注释                                                                                                                                                                                                                                                                                                                  |

5. tag 注意

   1. `autoIncrement` 会自动推断类型,并且设置称为主键

   2. 踩坑记录 `https://www.cnblogs.com/rickiyang/p/14517120.html`

   3. Omit之中传入的字段，忽略大小写。

   4. 
