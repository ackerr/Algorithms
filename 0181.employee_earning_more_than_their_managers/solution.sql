# Write your MySQL query statement below
select Name as Employee from Employee as e where Salary > (select Salary from Employee where Id = e.ManagerId)