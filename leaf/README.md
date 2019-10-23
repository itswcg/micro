### leaf
基于[leaf](https://tech.meituan.com/2017/04/21/mt-leaf.html)的分布式ID生成系统.

### Mysql
```
create database leaf

CREATE TABLE `segment` (
  `biz_tag` varchar(128) NOT NULL,
  `max_id` bigint(20) NOT NULL DEFAULT 1,
  `step` int(11) NOT NULL,
  `desc` varchar(256),
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`biz_tag`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

insert into segment(biz_tag, max_id, step, `desc`) values('account', 10000, 10, 'account.service')
```