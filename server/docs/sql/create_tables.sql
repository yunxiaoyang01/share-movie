-- 影视表
CREATE TABLE movie (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
  title VARCHAR(100) NOT NULL COMMENT '影视标题',
  cover VARCHAR(255) COMMENT '封面图地址',
  play_count INT DEFAULT 0 COMMENT '播放次数',
  share_time DATETIME COMMENT '分享时间',
  expire_time DATETIME COMMENT '过期时间',
  FULLTEXT INDEX idx_title (title),
  INDEX idx_expire_time (expire_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 播放源表
CREATE TABLE play_source (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  movie_id BIGINT NOT NULL COMMENT '影视ID',
  source_type VARCHAR(20) NOT NULL COMMENT '网盘类型',
  episode VARCHAR(50) COMMENT '集数/版本',
  resource_url VARCHAR(500) NOT NULL COMMENT '资源链接',
  extract_code VARCHAR(20) COMMENT '提取码',
  file_size BIGINT COMMENT '文件大小(字节)',
  file_format VARCHAR(10) COMMENT '文件格式',
  valid TINYINT DEFAULT 1 COMMENT '是否有效',
  FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
  INDEX idx_source_type (source_type),
  INDEX idx_valid (valid),
  FULLTEXT INDEX idx_resource_url (resource_url),
  INDEX idx_movie_source_valid (movie_id, source_type, valid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 分类表
CREATE TABLE category (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(20) NOT NULL UNIQUE COMMENT '分类名称'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 影视-分类关联表
CREATE TABLE movie_category (
  movie_id BIGINT NOT NULL,
  category_id BIGINT NOT NULL,
  PRIMARY KEY (movie_id, category_id),
  FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;