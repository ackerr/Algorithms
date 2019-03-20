# Write your MySQL query statement below
select Email from Person GROUP BY Email HAVING count(Email) > 1