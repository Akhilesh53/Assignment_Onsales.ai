To handle the task of evaluating multiple mathematical expressions using a web API that has rate limitations, following steps can be considered:

1) Group the Input expressions to Batches: 
   - Since the API supports a maximum of 50 requests per second per client, what we can do is to divide the expressions into smaller batches 
     of 50 or fewer expressions. This ensures we don't exceed the rate limit imposed by the API.

2) To control the number of requests sent to the API per second, we can introduce a delay between consecutive batches. 
   For example, introduce a delay of 1 second between each batch of requests to ensure that the rate limit is not exceeded.

=======================================================
============== HOW TO RUN THE PROGRAM =================

*TO RUN THE MAIN FILE*

1) Go to the task1 directory
2) Run the below command

**go run main.go**

*TO RUN THE TEST FILE*

1) Go to the task1 directory
2) Run the below command

**go test**