package errors

import "log"

func RequiredPropertyNotFound(propertyName string) {
	log.Fatal("Missing required property: " + "\"" + propertyName + "\"")
}
