# Write your MySQL query statement below
select employee_id from Employees as emp where emp.manager_id is not null and (select count(*) from Employees as emp2 where emp2.employee_id = emp.manager_id) = 0 and emp.salary < 30000 order by emp.employee_id;