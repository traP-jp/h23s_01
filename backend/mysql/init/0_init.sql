DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(36) not null primary key,
  `name` varchar(32) not null
)ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels` (
  `id` varchar(36) not null primary key,
  `name` varchar(128) not null
)ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci