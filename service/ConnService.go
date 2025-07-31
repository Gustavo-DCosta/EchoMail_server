package service

import "fmt"

func CreateAccSupabase(phoneNumber, email string) error {
	uuid := GenerateUUID()
	err := WriteRedis2ways(uuid, phoneNumber)
	if err != nil {
		fmt.Println("Error writing data to redis | ERROR CODE: ", err)
		return err
	}
	return nil
}

func LoginAccSupabase(phoneNumber, email string) error {
	return nil
}
