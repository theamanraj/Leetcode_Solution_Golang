# Write your MySQL query statement below
select e.employee_id, e.salary as bonus from Employees e
where e.employee_id % 2 <> 0 and e.name not like 'm%'

union 

select e.employee_id, 0 as bonus from Employees e
where e.employee_id % 2 = 0 or e.name like 'm%'

order by employee_id
