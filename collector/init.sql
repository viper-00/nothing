-- server table
CREATE TABLE `server` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `timezone` varchar(30)
);

-- system_metrics table
CREATE TABLE `system_metrics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `server_id` int,
  `log_time` bigint(20),
  `log_type` varchar(255),
  `log_name` VARCHAR(255),
  `log_text` text
);

-- custom_metrics table
CREATE TABLE `custom_metrics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `server_id` int,
  `log_time` bigint(20),
  `log_type` varchar(255),
  `log_name` VARCHAR(255),
  `log_text` text
);

-- alert table
CREATE TABLE `alert` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `server_id` int,
  `type` int,
  `expected` varchar(50),
  `actual` varchar(50),
  `time` bigint(20),
  `start_log_id` int,
  `end_log_id` int
);

-- server_ping_time table
CREATE TABLE `server_ping_time` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `server_id` int,
  `time` bigint(20)
);

ALTER TABLE `system_metrics` ADD FOREIGN KEY (`server_id`) REFERENCES `server` (`id`);

ALTER TABLE `custom_metrics` ADD FOREIGN KEY (`server_id`) REFERENCES `server` (`id`);

ALTER TABLE `alert` ADD FOREIGN KEY (`server_id`) REFERENCES `server` (`id`);

ALTER TABLE `alert` ADD FOREIGN KEY (`start_log_id`) REFERENCES `system_metrics` (`id`);

ALTER TABLE `alert` ADD FOREIGN KEY (`end_log_id`) REFERENCES `system_metrics` (`id`);

ALTER TABLE `server_ping_time` ADD FOREIGN KEY (`server_id`) REFERENCES `server` (`id`);

CREATE INDEX `log_time` ON `system_metrics`(`log_time`);
CREATE INDEX `log_name` ON `system_metrics`(`log_name`);

CREATE INDEX `log_time` ON `custom_metrics`(`log_time`);
CREATE INDEX `log_name` ON `custom_metrics`(`log_name`);