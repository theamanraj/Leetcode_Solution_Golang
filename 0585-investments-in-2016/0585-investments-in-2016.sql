# Write your MySQL query statement below
With  same_tiv_2015 AS
(Select i2.pid, i2.tiv_2016
From Insurance as i1 join Insurance as i2
On i1.tiv_2015 = i2.tiv_2015
Where i1.pid <> i2.pid 
Group by i2.pid),

same_lat_lon AS
(Select t1.pid 
From Insurance as t1 join Insurance as t2
On (t1.lat,t1.lon) = (t2.lat,t2.lon)
Where t1.pid <> t2.pid)

Select  Round(sum(tiv_2016),2) as tiv_2016
From same_tiv_2015 
Where pid not in (Select pid From same_lat_lon);