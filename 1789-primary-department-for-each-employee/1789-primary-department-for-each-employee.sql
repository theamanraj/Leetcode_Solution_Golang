# Write your MySQL query statement below
with cte as 
(select employee_id , count(*) as cnt
from employee
group by employee_id
)

select x.employee_id , x.department_id 
from
(select e.employee_id , 
case
    when cte.cnt = 1  then department_id
    when cte.cnt > 1 and primary_flag = 'Y' then department_id
  end as department_id
  from employee e
join cte 
on e.employee_id = cte.employee_id
) x

where x.department_id is not null