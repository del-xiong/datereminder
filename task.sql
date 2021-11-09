-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2021-11-09 19:00:00
-- 服务器版本： 8.0.20
-- PHP 版本： 7.3.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `datereminder`
--

-- --------------------------------------------------------

--
-- 表的结构 `task`
--

CREATE TABLE `task` (
  `task_id` int NOT NULL,
  `task_name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务名',
  `task_desc` varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `reminder_date` varchar(100) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '提醒日期',
  `is_lunar` int NOT NULL DEFAULT '0' COMMENT '是否农历日\r\n默认公历',
  `is_loop` int NOT NULL DEFAULT '0' COMMENT '是否循环提醒',
  `pre_day` int NOT NULL DEFAULT '15' COMMENT '提前多少天开始提醒',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `task`
--

INSERT INTO `task` (`task_id`, `task_name`, `task_desc`, `reminder_date`, `is_lunar`, `is_loop`, `pre_day`, `create_time`) VALUES
(1, '妈生日', '老妈生日，她一般喜欢花或者漂亮衣服', '08-12', 0, 0, 15, '2021-11-09 10:59:16'),
(2, '老婆生日', '忘了会死，送的礼物不合意会死，看着办', '07-22', 1, 0, 15, '2021-11-09 10:59:52');

--
-- 转储表的索引
--

--
-- 表的索引 `task`
--
ALTER TABLE `task`
  ADD PRIMARY KEY (`task_id`),
  ADD KEY `is_loop` (`is_loop`),
  ADD KEY `pre_day` (`pre_day`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `task`
--
ALTER TABLE `task`
  MODIFY `task_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
