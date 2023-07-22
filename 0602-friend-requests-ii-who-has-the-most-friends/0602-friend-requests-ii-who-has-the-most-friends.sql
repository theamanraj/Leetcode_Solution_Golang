# Write your MySQL query statement below
WITH union_cte AS (
    SELECT requester_id AS id
    FROM RequestAccepted
    UNION ALL
    SELECT accepter_id
    FROM RequestAccepted
)

SELECT id, COUNT(id) AS num
FROM union_cte
GROUP BY id
ORDER BY num DESC LIMIT 1