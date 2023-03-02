1- Physical Layer

2- Data Link layer (Mac address)

3- Network layer  (IP address)

4- Transmission layer (TCP UDP)

5- Session layer ( TCP + HTTP)

6- Presentation layer (HTTP)

7- Application layer  (HTTP)


1- GET      http://localhost:50060/employee : It fetches all employees
1.1-GET     http://localhost:50060/employee?id=101
2- POST     http://localhost:50060/employee : It creates a new employee
3- PUT      http://localhost:50060/employee : it updates existing employee
4- DELETE   http://localhost:50060/employee : it deletes an existing employee
5- PATCH    http://localhost:50060/employee : it does partial update of employee

1- GET      http://localhost:50060/employee/all : It fetches all employees
1.1-GET     http://localhost:50060/employee?id=101
2- POST     http://localhost:50060/employee/add : It creates a new employee
3- PUT      http://localhost:50060/employee/edit : it updates existing employee
4- DELETE   http://localhost:50060/employee/delete : it deletes an existing employee
5- PATCH    http://localhost:50060/employee/partial : it does partial update of employee


/ 
/employee
/health
/ping
/employee/salaries