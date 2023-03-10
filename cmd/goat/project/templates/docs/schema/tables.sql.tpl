CREATE TABLE IF NOT EXISTS `books` (
  `id` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对象Id',
  `status` smallint Not NULL COMMENT '0:删除,1:创建,2:更新',
  `create_at` bigint NOT NULL COMMENT '创建时间(13位时间戳)',
  `create_by` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '创建人',
  `update_at` bigint DEFAULT 0 COMMENT '更新时间',
  `update_by` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '更新人',
  `delete_at` bigint DEFAULT 0 COMMENT '删除时间',
  `delete_by` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '删除人',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '书名',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE COMMENT '用于书名搜索',
  KEY `idx_author` (`author`) USING BTREE COMMENT '用于作者搜索'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;