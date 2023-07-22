# Write your MySQL query statement below
SELECT s.machine_id, ROUND(AVG(e.timestamp-s.timestamp),3) as processing_time 
FROM Activity s JOIN Activity e USING (machine_id) WHERE s.activity_type='start' AND e.activity_type='end' GROUP BY s.machine_id;