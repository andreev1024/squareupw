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

#### Role
* Create Role
* Update Role
* List Roles
* Retrieve Role

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

    p := squareupw.RoleParams{
      Name:        "Test",
      Permissions: []string{squareupw.RegisterAccessSalesHistory},
      IsOwner:     false,
    }
    resp, err := api.CreateRole(p)
}
```

### Pagination

List endpoints might paginate the results they return. This means that instead of returning all results in a single response, these endpoints might return some of the results, along with a response header that links to the next set of results.

```
p := squareupw.ListRolesParams{Order: "desc", Limit: "5"}
resp, link, err := api.ListRoles(p)

if len(link) > 1 {
  additionalResp, link, e := api.ListRolesByLink(link)
}
```
