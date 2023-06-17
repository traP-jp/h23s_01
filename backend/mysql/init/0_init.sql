DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` varchar(36) not null,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP
)ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

INSERT INTO `users` (`user_id`) VALUES ("7265b13d-9e06-42f6-98e3-41ea742f8fb2");

DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels` (
  `id` varchar(36) not null primary key,
  `name` varchar(128) not null
)ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci