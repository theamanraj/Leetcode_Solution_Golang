# Write your MySQL query statement below
with cte as (
    select
        count(case when income < 20000 then 1 end) as count1,
        count(case when income >= 20000 and income <= 50000 then 1 end) as count2,
        count(case when income > 50000 then 1 end) as count3
    from accounts
)
select lj.*
from cte, lateral (values
    row('Low Salary', count1),
    row('Average Salary', count2),
    row('High Salary', count3)
) as lj(category, accounts_count)