--
-- @lc app=leetcode id=175 lang=mysql
--
-- [175] Combine Two Tables
--
# Write your MySQL query statement below
   SELECT p.FirstName FirstName,
          p.LastName LastName,
          a.City City,
          a.State State
     FROM Person p
LEFT JOIN Address a
       ON p.PersonId = a.PersonId;
