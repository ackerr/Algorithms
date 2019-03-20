# Write your MySQL query statement below

# select Name as Customers from Customers where Id not in (select distinct CustomerId from Orders)

select Name Customers from Customers c left join Orders o on c.id = o.CustomerId where o.CustomerId is null