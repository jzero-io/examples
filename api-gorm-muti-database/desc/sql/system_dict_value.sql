use ntls;

CREATE TABLE if not exists `T_system_dict_value` (
  id BIGINT(11) NOT NULL AUTO_INCREMENT,
  key_uuid VARCHAR(36) NOT NULL COMMENT 'key uuid',
  uuid VARCHAR(36) UNIQUE NOT NULL COMMENT 'uuid',
  label VARCHAR(100) NOT NULL COMMENT '字典名称（展示）',
  value VARCHAR(100) NOT NULL COMMENT '字典值',
  language VARCHAR(100) NOT NULL COMMENT '字典名称所属语言',
  sort INT NOT NULL COMMENT '排序（升序）',
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
  update_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  remarks VARCHAR(255) DEFAULT '',
  default_flag int DEFAULT 0 COMMENT '默认数据，0：默认  1：非默认',
  PRIMARY KEY(id)
) ENGINE=InnoDB auto_increment=1 DEFAULT CHARSET=utf8mb4;