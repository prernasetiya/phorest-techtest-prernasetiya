## Problem Statement

According to me, the problem at hand was that I had a set of CSV files that needed to be operated based on a few commands essentially through rest APIs. The files needed to be put in a database to access it better. Addtitionally, the files had to be stored in a format that operations like import, delete or update can be run on it. Extra attention to provide an easier access to the client to get information about let’s say, top 50 client names, purchases and loyalty card points total to get more insights from the data.

  

## Why did I use Go?

The validation was a major focus point in the problem statement. Go makes it easier to validate request/response bodies and database schema structures.

  

## What did I do ?

I made the decision of creating a local database because there was no access to any external database. I started off with making a schema of the CSV files in Go struct and defined the relationships among the tables as supplied in the problem statement, eg- **id** column of the _clients_ table is associated to **client_id** of _appointments_ table.

I made API router and different APIs for different parts of the problem statement.


I created ```/import``` API for the first part of the. Problem statement which was to consume and parse csv files and import data into the database.

  
I created ```/top-clients``` API to fetch top X clients with maximum loyalty points since a specific date. The generic structure for this API is ``` /top-clients? limit=10&since=2016/01/01```

  

I also tried to implement APIs to cover the “Nice to Have” features of the problem statement.

  

## Testing

I tried to cover majority of the constraints through tests like:

- Checking the fact that the files are getting imported correctly

- Testing formatting of the dates

- Total number of customers being negative

- Check whether records are deleted perfectly or not

  

## Missing features or Scope of Improvements

- Nice to have features are working for only one table, its impact on the other tables is not implemented yet

- For now, If we delete an entity in the clients table, it doesn’t delete the associated appointments in the associated tables