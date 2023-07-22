# Write your MySQL query statement below
with cte as (
  select d.visited_on
  , sum(amount) over(order by visited_on range interval 6 day preceding) as amount
  , avg(amount) over(order by visited_on range interval 6 day preceding) as average_amount
  from (select visited_on, sum(amount) as amount from Customer group by visited_on) as d
)

select
  visited_on
  , round(amount, 2) as amount
  , round(average_amount, 2) as average_amount
  from cte
  , (select min(visited_on) as min_visit from Customer) as c2
where datediff(visited_on, min_visit) >= 6