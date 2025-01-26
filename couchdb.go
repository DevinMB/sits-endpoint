package main

import (
	"context"
	"fmt"
	"log"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
)

var couchClient *kivik.Client

func ConnectCouchDB(url, username, password string) error {
	connectionString := fmt.Sprintf("http://%s:%s@%s/",username,password,url)
	client, err := kivik.New("couch", connectionString)
	if err != nil {
		return err
	}
	couchClient = client
	log.Println("Connected to CouchDB!")
	return nil
}

func GetRowCount() (int64, error) {
    db := couchClient.DB("sits")
    if err := db.Err(); err != nil {
        return 0, err
    }
    stats, err := db.Stats(context.Background())
    if err != nil {
        return 0, err
    }
    return stats.DocCount, nil
}

func GetStackedHistoricalEvents() (map[string]map[string]float64, error){
	db := couchClient.DB("sits")
	if err := db.Err(); err != nil {
		return nil, err
	}
	opts := kivik.Param("group", true)   // "group_level": 3 
	rows := db.Query(context.Background(), "_design/sits","_view/by_day_device", opts)
     
	if err := rows.Err(); err != nil {
			panic(err)
	}
	defer rows.Close()

	result := make(map[string]map[string]float64)
    //[year, month, day, device_id] if group=true
	
	for rows.Next(){
		var keyArr []interface{}
    	if err := rows.ScanKey(&keyArr); err != nil {
        	return nil, err
    	}

		year := int(keyArr[0].(float64))
        month := int(keyArr[1].(float64))
        day := int(keyArr[2].(float64))
        deviceID := keyArr[3].(string)

		var durationSum float64
    	if err := rows.ScanValue(&durationSum); err != nil {
        	return nil, err
    	}

		dateKey := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

		if _, exists := result[dateKey]; !exists {
			result[dateKey] = make(map[string]float64)
		}

		result[dateKey][deviceID] = durationSum
	}

	if err := rows.Err(); err != nil {
        return nil, err
    }

    return result, nil
}



// func GetSensorSitCount(sensorName string) (int64, error) {

// }

