package squareupw

import (
	"encoding/json"
	"log"
)

//EmployeeRole.Permission
const (
	RegisterAccessSalesHistory        = "REGISTER_ACCESS_SALES_HISTORY"
	RegisterApplyRestrictedDiscounts  = "REGISTER_APPLY_RESTRICTED_DISCOUNTS"
	RegisterChangeSettings            = "REGISTER_CHANGE_SETTINGS"
	RegisterEditItem                  = "REGISTER_EDIT_ITEM"
	RegisterIssueRefunds              = "REGISTER_ISSUE_REFUNDS"
	RegisterOpenCashDrawerOutsideSale = "REGISTER_OPEN_CASH_DRAWER_OUTSIDE_SALE"
	RegisterViewSummaryReports        = "REGISTER_VIEW_SUMMARY_REPORTS"
)

//Merchant represents Merchant model.
type Merchant struct {
	ID                  string                  `json:"id"`
	Name                string                  `json:"name"`
	Email               string                  `json:"email"`
	CountryCode         string                  `json:"country_code"`
	LanguageCode        string                  `json:"language_code"`
	AccountType         string                  `json:"account_type"`
	AccountCapabilities []string                `json:"account_capabilities"`
	CurrencyCode        string                  `json:"currency_code"`
	BusinessName        string                  `json:"business_name"`
	BusinessAddress     GlobalAddress           `json:"business_address"`
	BusinessPhone       PhoneNumber             `json:"business_phone"`
	BusinessType        string                  `json:"business_type"`
	ShippingAddress     GlobalAddress           `json:"shipping_address"`
	LocationDetails     MerchantLocationDetails `json:"location_details"`
	MarketURL           string                  `json:"market_url"`
}

//GlobalAddress represents GlobalAddress model.
type GlobalAddress struct {
	AddressLine1                 string      `json:"address_line_1"`
	AddressLine2                 string      `json:"address_line_2"`
	AddressLine3                 string      `json:"address_line_3"`
	AddressLine4                 string      `json:"address_line_4"`
	AddressLine5                 string      `json:"address_line_5"`
	Locality                     string      `json:"locality"`
	Sublocality                  string      `json:"sublocality"`
	Sublocality1                 string      `json:"sublocality_1"`
	Sublocality2                 string      `json:"sublocality_2"`
	Sublocality3                 string      `json:"sublocality_3"`
	Sublocality4                 string      `json:"sublocality_4"`
	Sublocality5                 string      `json:"sublocality_5"`
	AdministrativeDistrictLevel1 string      `json:"administrative_district_level_1"`
	AdministrativeDistrictLevel2 string      `json:"administrative_district_level_2"`
	AdministrativeDistrictLevel3 string      `json:"administrative_district_level_3"`
	PostalCode                   string      `json:"postal_code"`
	CountryCode                  string      `json:"country_code"`
	AddressCoordinates           Coordinates `json:"address_coordinates"`
}

//Employee represents Employee model.
type Employee struct {
	ID                    string   `json:"id"`
	FirstName             string   `json:"first_name"`
	LastName              string   `json:"last_name"`
	RoleIds               []string `json:"role_ids"`
	AuthorizedLocationIds []string `json:"authorized_location_ids"`
	Email                 string   `json:"email"`
	Status                string   `json:"status"`
	ExternalID            string   `json:"external_id"`
	CreatedAt             string   `json:"created_at"`
	UpdatedAt             string   `json:"updated_at"`
}

//EmployeeRole represents Employee role model.
type EmployeeRole struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	IsOwner     bool     `json:"is_owner"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

//PhoneNumber represents PhoneNumber model.
type PhoneNumber struct {
	CallingCode string `json:"calling_code"`
	Number      string `json:"number"`
}

//MerchantLocationDetails represents MerchantLocationDetails model.
type MerchantLocationDetails struct {
	Nickname string `json:"nickname"`
}

//Coordinates represents Coordinates model.
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

//CreateEmployeeParams represents params for CreateEmployee method.
type CreateEmployeeParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,omitempty"`
	*CommonOptionalEmployeeParams
}

//UpdateEmployeeParams represents params for UpdateEmployee method.
type UpdateEmployeeParams struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	*CommonOptionalEmployeeParams
}

//CommonOptionalEmployeeParams represents common params for other *EmployeeParams.
type CommonOptionalEmployeeParams struct {
	ExternalID string   `json:"external_id,omitempty"`
	RoleIds    []string `json:"role_ids,omitempty"`
}

//ListEmployeesParams represents params for ListEmployees method.
type ListEmployeesParams struct {
	Order          string `param:"order"`
	BeginUpdatedAt string `param:"begin_updated_at"`
	EndUpdatedAt   string `param:"end_updated_at"`
	BeginCreatedAt string `param:"begin_created_at"`
	EndCreatedAt   string `param:"end_created_at"`
	Status         string `param:"status"`
	ExternalID     string `param:"external_id"`
	Limit          string `param:"limit"`
}

//RoleParams represents params for *Role methods.
type RoleParams struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	IsOwner     bool     `json:"is_owner"`
}

//ListRolesParams represents params for ListRoles methods.
type ListRolesParams struct {
	Order string `param:"order"`
	Limit string `param:"limit"`
}

func apiURL() string {
	return BaseURL + "/v1"
}

//RetrieveBusiness provides a business's account information.
func (a API) RetrieveBusiness() (resp Merchant, err error) {
	url := apiURL() + "/me"
	method := MethodGet

	_, body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//ListLocations provides details for a business's locations, including their IDs.
func (a API) ListLocations() (resp []Merchant, err error) {
	url := apiURL() + "/me/locations"
	method := MethodGet

	_, body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//CreateEmployee creates an employee for a business.
func (a API) CreateEmployee(p CreateEmployeeParams) (resp Employee, err error) {
	url := apiURL() + "/me/employees"
	method := MethodPost
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	_, body, err := a.Send(method, url, data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//UpdateEmployee modifies the details of an employee.
func (a API) UpdateEmployee(id string, p UpdateEmployeeParams) (resp Employee, err error) {
	url := apiURL() + "/me/employees/" + id
	method := MethodPut
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	_, body, err := a.Send(method, url, data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//ListEmployees provides summary information for all of a business's employees.
//This endpoint might paginate its results. You should check `link`.
//If it not empty - you can perform additional request via ListEmployeesByLink method.
func (a API) ListEmployees(p ListEmployeesParams) (resp []Employee, link string, err error) {
	endpointURL := apiURL() + "/me/employees"
	body, link, err := a.list(&p, MethodGet, endpointURL, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//ListEmployeesByLink is ListEmployees alias which get explicit url.
func (a API) ListEmployeesByLink(url string) (resp []Employee, link string, err error) {
	p := ListEmployeesParams{}
	body, link, err := a.list(&p, MethodGet, "", url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//RetrieveEmployee provides the details for a single employee.
func (a API) RetrieveEmployee(id string) (resp Employee, err error) {
	url := apiURL() + "/me/employees/" + id
	method := MethodGet

	_, body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//CreateRole creates an employee role you can then assign to employees.
func (a API) CreateRole(p RoleParams) (resp EmployeeRole, err error) {
	url := apiURL() + "/me/roles"
	method := MethodPost
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	log.Printf("%s", data)

	_, body, err := a.Send(method, url, data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//ListRoles provides summary information for all of a business's employee roles.
//This endpoint might paginate its results. You should check `link`.
//If it not empty - you can perform additional request via ListRolesByLink method.
func (a API) ListRoles(p ListRolesParams) (resp []EmployeeRole, link string, err error) {
	endpointURL := apiURL() + "/me/roles"
	body, link, err := a.list(&p, MethodGet, endpointURL, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//ListRolesByLink is ListRoles alias which get explicit url.
func (a API) ListRolesByLink(url string) (resp []EmployeeRole, link string, err error) {
	p := ListRolesParams{}
	body, link, err := a.list(&p, MethodGet, "", url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//RetrieveRole provides the details for a single employee role.
func (a API) RetrieveRole(id string) (resp EmployeeRole, err error) {
	url := apiURL() + "/me/roles/" + id
	method := MethodGet

	_, body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//UpdateRole modifies the details of an employee role.
func (a API) UpdateRole(id string, p RoleParams) (resp EmployeeRole, err error) {
	url := apiURL() + "/me/roles/" + id
	method := MethodPut
	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	_, body, err := a.Send(method, url, data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

//list method is DRY method for all other List* methods.
func (a API) list(p interface{}, method, endpointURL, url string) (body []byte, link string, err error) {
	if len(url) < 1 {
		queryString, e := GetQueryStringByStruct(p, "param", true)
		if e != nil {
			err = e
			return
		}

		if len(queryString) > 1 {
			queryString = "?" + queryString
		}

		url = endpointURL + queryString
	}

	log.Println(url)

	resp, body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	if linkHeader, ok := resp.Header["Link"]; ok {
		link, err = ExtractURLFromLinkHeader(linkHeader)
	}
	return
}
