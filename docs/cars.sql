CREATE TABLE `gz_qc_car` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city_name` varchar(255) DEFAULT '' COMMENT '城市名称',
  `title` varchar(255) DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) unsigned DEFAULT NULL COMMENT '价格(万)',
  `oldprice` decimal(10,2) unsigned DEFAULT NULL COMMENT '原价(万)',
  `kilometer` decimal(10,2) unsigned DEFAULT NULL COMMENT '公里(万)',
  `year` int(10) unsigned DEFAULT '0' COMMENT '年份',
  `transfer_city` varchar(255) DEFAULT '' COMMENT '过户城市',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

