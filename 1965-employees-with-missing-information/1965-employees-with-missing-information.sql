# Write your MySQL query statement below
select newtable.employee_id from (select * from employees 
union
select * from salaries) as newtable where newtable.employee_id not in (
select e.employee_id from employees e join salaries s on e.employee_id = s.employee_id) order by  newtable.employee_id