# Write your MySQL query statement below
select rq.title as results from (select ur.movie_id,ur.title,avg(mo.rating) as av from movierating as mo left join movies as ur on mo.movie_id = ur.movie_id where mo.created_at between '2020-02-01' and '2020-02-29' group by ur.movie_id order by av desc,ur.title limit 1) as rq 
union all
select ra.name as results from (select mi.user_id,mi.name , count(mo.rating) as coun from movierating as mo left join users as mi on mo.user_id = mi.user_id  group by mo.user_id order by coun desc,mi.name limit 1) as ra