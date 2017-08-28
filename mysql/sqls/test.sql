
drop schema if exists `test`;
create schema `test`;
use test;
DROp table if exists `userinfo`;
CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`departname` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
DROp table if exists `userdetail`;
CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
