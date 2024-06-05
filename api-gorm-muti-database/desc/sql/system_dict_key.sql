use ntls;

CREATE TABLE IF NOT EXISTS `T_system_dict_key` (
  id BIGINT(11) NOT NULL AUTO_INCREMENT,
  uuid VARCHAR(36) UNIQUE NOT NULL COMMENT 'uuid',
  category_code VARCHAR(100) UNIQUE NOT NULL COMMENT '字典类型',
  category_desc VARCHAR(255) NOT NULL COMMENT '字典描述',
  sort INT NOT NULL COMMENT '排序（升序）',
  PRIMARY KEY(id)
) ENGINE=InnoDB auto_increment=1 DEFAULT CHARSET=utf8mb4;