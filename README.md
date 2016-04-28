Package **squareupw** implements a wrapper for [SquareUp API](https://docs.connect.squareup.com) for Go (Golang).

### API endpoints
Current version support following API endpoints:

#### Business & Locations
* Retrieve Business
* List Locations

#### Employee
* Create Employee
* Update Employee
* List Employees
* Retrieve Employee

### Example
```
    //  Notice: error handling skipped for simplicity

    token := "123456789"
    api := squareupw.NewAPI(token)

    p := squareupw.UpdateEmployeeParams{
      FirstName: "Andrey",
      LastName:  "Andreev",
      CommonOptionalEmployeeParams: &squareupw.CommonOptionalEmployeeParams{
        RoleIds: []string{"ZCZ0AzjPpVDUTMdMpdF3"}},
    }

    resp, err := api.UpdateEmployee("G7cAQQKTMk05R78mIQvZ", p)

    //  resp, _ := api.RetrieveEmployee("G7cAQQKTMk05R78mIQvZ")
}
```
