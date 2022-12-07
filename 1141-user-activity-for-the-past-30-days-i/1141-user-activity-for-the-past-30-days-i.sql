# Write your MySQL query statement below
SELECT a.activity_date "day", count(distinct user_id) "active_users" FROM Activity a
WHERE a.activity_date > '2019-07-27' - INTERVAL 30 DAY
AND a.activity_date <= '2019-07-27'
GROUP BY a.activity_date;