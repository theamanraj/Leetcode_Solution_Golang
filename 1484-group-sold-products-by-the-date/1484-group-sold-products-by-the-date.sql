# Write your MySQL query statement below
SELECT sell_date, count(DISTINCT product) AS num_sold,
GROUP_CONCAT(DISTINCT product) AS products
FROM Activities
GROUP BY sell_date

