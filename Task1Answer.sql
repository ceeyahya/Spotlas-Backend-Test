-- Task 1.1 - Changing the `website` to only include the domain, e.g.:
-- https://gelatorino.com/covent-garden-london => gelatorino.com
-- Assumption: a URL that has subdomains will be modified with said subdomain e.g:
-- https://m.facebook.com/twocentspizza/?locale2=en_GB => m.facebook.com
-- The proposed solution uses the `SUBSTRING()` function and Regular Expressions
-- `SUBSTRING(string, position, substring_length)` Function: return a part of a string.
UPDATE "MY_TABLE"
SET website = SUBSTRING(website FROM '(?:.*://)?(?:www\.)?([^/?]*)');

-- Task 1.2 - Count how many spots contain the same domain
-- Interpretation 1: return the number of spots that have non-unique domains
-- (SELECT COUNT(website) FROM "MY_TABLE") => Total Number of Domains
-- (SELECT COUNT(DISTINCT (website)) => Number Unique of Domains
-- Substract the unique domains from the total number of domains to get spots that have repeating domains
SELECT (SELECT COUNT(website) FROM "MY_TABLE") - (SELECT COUNT(DISTINCT (website)) FROM "MY_TABLE") AS non_unique_domains FROM "MY_TABLE" GROUP BY 1;

-- Interpretation 2: return the spots' domains and how many times they appear
SELECT string_agg(name, ', ') as names, website, COUNT(website)
FROM "MY_TABLE"
GROUP BY website
HAVING COUNT(website) > 1
ORDER BY COUNT(website);

-- Task 1.3 - Return spots that have a domain that is greater to 1
SELECT website, COUNT(website)
FROM "MY_TABLE"
GROUP BY website
HAVING COUNT(website) > 1
ORDER BY COUNT(website);


-- Task 1.4 - Write a PL/SQL Script for 1.1
BEGIN
UPDATE
    website
SET
    website = SUBSTRING(website FROM '(?:.*://)?(?:www\.)?([^/?]*)');
END;