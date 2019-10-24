--
-- @lc app=leetcode id=196 lang=mysql
--
-- [196] Delete Duplicate Emails
--

-- @lc code=start
# Write your MySQL query statement below

DELETE FROM Person 
      WHERE Id NOT IN (SELECT temp.Id 
                         FROM (SELECT MIN(Id) AS Id,
                                      Email
                                 FROM Person
                             GROUP BY Email) temp)

-- @lc code=end
