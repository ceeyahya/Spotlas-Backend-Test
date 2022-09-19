# Spotlas Backend Engineer Test

## Task 1:

### Task 1.1

**Change the website field, so it only contains the domain**:

the proposed solution uses postgres' `SUBSTRING()` function and Regular Expression and takes into account multiple cases as well as makes some assumptions about the desired outcome

**Assumptions**:

- If a domain has a subdomain, the function will keep the subdomain e.g.:

`https://m.facebook.com/test/foo/bar` => `m.facebook.com`

**Cases**

- The domain contains either `http` or `https`
- The domain contains `www` or not

### Task 1.2

Count how many spots contain the same domain:

I was not sure how to interpret this task as in my opinion it seemed similar to 1.3 so I implemented both interpretations:

**Interpretation 1**: return the number of spots that have non-unique domains.

**Interpretation 2**: return the spots' domains and how many times they appear

### Task 1.3

Return spots which have a domain with a count greater than 1:

This task was implemented using the `GROUP BY` and `HAVING` functions

### Task 1.4

Make a PL/SQL function for point 1

## Task 2:

This task was implemented using as little libraries as possible, as such the only used library is `gorm` to manage the connection to the Database and the queries.

the folder's structure is as follows:

- `database` folder: contains the code responsible for connecting to the Database.
- `models` folder: contains the shape of a `Spot`.
- `main.go`: the starting point of the program.
