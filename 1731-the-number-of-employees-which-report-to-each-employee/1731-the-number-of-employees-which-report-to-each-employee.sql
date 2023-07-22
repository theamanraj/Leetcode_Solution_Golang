# Write your MySQL query statement below
SELECT e1.employee_id,e1.name,count(e2.name) reports_count, round(AVG(e2.age)) average_age
FROM Employees e1
LEFT JOIN Employees e2
ON e1.employee_id=e2.reports_to
where e2.age is not NULL
GROUP BY e1.employee_id,e1.name
ORDER BY e1.employee_id;