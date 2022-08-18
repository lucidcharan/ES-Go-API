package models

type Testplan struct {
	Vertical string `json:"vertical"`
	Product string `json:"product"`
	Testplan_Name string `json:"testplan_name"`
	S3_Location string `json:"s3_location"`
	Version string `json:"version"`
	Testdata_Location string `json:"testdata_location"`
	Preferred_tool string `json:"preferred tool"`
	Creation_Date string `json:"Date"`
	TestPlan_Id string `json:"testplan_id"`
}