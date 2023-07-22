# Write your MySQL query statement below
select e.name, b.bonus
from Bonus b 
right Join employee e
on e.empId = b.empId
where bonus is null or bonus <1000