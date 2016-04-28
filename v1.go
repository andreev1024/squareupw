package squareupw

import "encoding/json"

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
//If it not empty - you can perform additional request.
func (a API) ListEmployees(p ListEmployeesParams) (resp []Employee, link string, err error) {
	queryString, err := GetQueryStringByStruct(&p, "param", true)
	if err != nil {
		return
	}

	if len(queryString) > 1 {
		queryString = "?" + queryString
	}

	url := apiURL() + "/me/employees" + queryString
	method := MethodGet

	httpResp, body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	if linkHeader, ok := httpResp.Header["Link"]; ok {
		link = linkHeader[0]
	}

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
