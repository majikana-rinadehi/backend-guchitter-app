CREATE TABLE `avatar` (
  `avatar_id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `avatar_name` varchar(45) NOT NULL,
  PRIMARY KEY (`avatar_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `complaints` (
  `complaint_id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `complaint_text` varchar(255) NOT NULL,
  `avatar_id` smallint unsigned NOT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`complaint_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;