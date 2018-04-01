package handler

import (
	"fmt"
	"log"
)

//Response do a backend process and respond ajax request from caller
func (m *Module) Response(input string) error {
	if err := m.validateResponse(input); err != nil {
		return err
	}

	return nil
}

//validateResponse validate response, hold multiple / spam response with redis
func (m *Module) validateResponse(input string) error {
	//let's use "anti-spam:test" key as redis key
	key := fmt.Sprintf("anti-spam:%s", input)

	//check key, is it exist?
	value, err := m.Redis.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	//value not nil, means the redis is already exists
	//therefore return error
	if value != nil {
		return fmt.Errorf("Value Exists, Please Wait")
	}

	//if there is no redis key
	//set new key with 10 second expiration time
	err = m.Redis.Setex(key, 10, "1")
	if err != nil {
		return err
	}

	return nil
}
