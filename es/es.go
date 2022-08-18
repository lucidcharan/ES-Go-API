package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Olderest/ES-Go/models"
	"github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error){
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialised")

	return client, err
}

func ESInsert() {

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	//creating student object
	newTestPlan := models.Testplan{
		Vertical			:  "Marketplace",
		Product 			:  "Router",
		Testplan_Name 		:  "Loadtest for Router",
		S3_Location 		:  "aws s3 location",
		Version 			:  "v2",
		Testdata_Location	:  "S3 Location",
		Preferred_tool		:  "Jmeter",
		Creation_Date		:  "current date",
		TestPlan_Id		:  "ID",
	}

	dataJSON, err := json.Marshal(newTestPlan)
	js := string(dataJSON)
	ind, err := esclient.Index().
		Index("students").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(ind)

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")

}