# Write your MySQL query statement below
SELECT person_name
FROM (SELECT *, SUM(weight) OVER (ORDER BY turn) AS Total_weight
FROM  Queue) AS a
WHERE Total_weight<=1000
ORDER BY turn DESC
LIMIT 1