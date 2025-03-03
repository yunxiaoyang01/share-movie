# 影视资源共享平台数据库设计文档

## 核心表结构设计

### 1. 用户表 (sys_user)
| 字段名 | 类型 | 描述 | 关联关系 |
|--------|------|------|---------|
| id | bigint | 主键 | 关联评论/收藏表 |
| username | varchar(50) | 登录账号 | 唯一索引 |
| nickname | varchar(50) | 显示名称 | |
| role_id | bigint | 角色ID | 外键关联角色表 |

### 2. 影视主表 (movie)
| 字段名 | 类型 | 描述 | 关联关系 |
|--------|------|------|---------|
| id | bigint | 主键 | 关联分类/标签/播放源 |
| title | varchar(100) | 影视标题 | 全文索引 |
| cover | varchar(255) | 封面图地址 | |
| play_count | int | 播放次数 | |
| share_time | datetime | 分享时间 | |
| expire_time | datetime | 过期时间 | 索引 |

### 5. 播放源表 (play_source)
| 字段名 | 类型 | 描述 | 关联关系 |
|--------|------|------|---------|
| id | bigint | 主键 | |
| movie_id | bigint | 影视ID | 外键关联影视表 |
| source_type | varchar(20) | 网盘类型（百度云/阿里云）| 索引 |
| episode | varchar(50) | 集数/版本 | |
| resource_url | varchar(500) | 资源链接 | 全文索引 |
| extract_code | varchar(20) | 提取码 | |
| file_size | bigint | 文件大小(字节) | |
| file_format | varchar(10) | 文件格式 | |
| valid | tinyint | 是否有效(1有效 0失效) | 索引 |

### 3. 分类表 (category)
| 字段名 | 类型 | 描述 | 关联关系 |
|--------|------|------|---------|
| id | bigint | 主键 | 多对多关联影视表 |
| name | varchar(20) | 分类名称 | 唯一索引 |

### 4. 影视-分类关联表 (movie_category)
| 字段名 | 类型 | 描述 |
|--------|------|------|
| movie_id | bigint | 影视ID |
| category_id | bigint | 分类ID |

## 表关系说明
1. **用户体系**
   - 用户表 ↔ 角色表：多对一关系
   - 用户表 ↔ 评论表：一对多关系
   - 用户表 ↔ 收藏表：一对多关系

2. **内容体系**
   - 影视表 ↔ 播放源表：一对多关系
   - 播放源表 ↔ 影视表：多对一关系

## 索引优化建议
4. 播放源表建立联合索引(movie_id, source_type, valid)
5. 播放源表source_type字段单独建立普通索引
6. 播放源表resource_url字段建立全文索引