WITH cte as (
SELECT user_id,action, COUNT(*) as count,SUM(CASE WHEN action = "confirmed" THEN 1 ELSE 0 END) AS con_cnt
from Confirmations
Group by user_id
)

SELECT s.user_id, CASE WHEN ROUND(con_cnt/count,2) IS NOT NULL
THEN ROUND(con_cnt/count,2) 
ELSE 0 END As confirmation_rate
FROM Signups as s
LEFT JOIN cte on
s.user_id = cte.user_id
