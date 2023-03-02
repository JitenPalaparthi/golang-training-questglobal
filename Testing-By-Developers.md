- Unit Test : is used to test the unit of functinality.
- Expected result and actual result.Both must be same otherwise it is a failure
- Requirements for unit test is input and output data of that unit.

        func HelloWorld(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintln(w, "Hello World!")
        }

- to perform unit test there is lot of dependency. Http is one of the dependencies.

- addEmployee(emp *Employee)(bool,error) {
    if emp!=nil{
        return false, errors.New("empty employee object")
    }
    result,err:= database.insert(emp) // assume that database.insert inserts record in  database
    return result,err
}
- Database is the dependency in above method. So you dont need to write code to store in database.

- How do you write unit test for the above func?

- Integration Testing : Integrate various components and test your component.Usually this is part of automation.

- System testing: Test whole system and make sure that as a system it is working. This is done by automation.

- End to End testing: Test a flow. Example shopping cart --> Search for items - add an item to the cart- make payment - get it delivered (Email|SMS). This is done by automation.