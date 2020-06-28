CREATE TABLE IF NOT EXISTS `tb_example` (
  `id_order` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(100) NOT NULL,
  `order_number` varchar(50) NOT NULL,
  `user_uuid` varchar(100) NOT NULL,
  `id_order_type` int(11) NOT NULL,
  `id_order_status` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id_order`),
  KEY `tb_orders_FK` (`id_order_type`),
  KEY `tb_orders_FK_1` (`id_order_status`),
  CONSTRAINT `tb_orders_FK` FOREIGN KEY (`id_order_type`) REFERENCES `tb_order_type` (`id_order_type`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `tb_orders_FK_1` FOREIGN KEY (`id_order_status`) REFERENCES `tb_order_status` (`id_order_status`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8