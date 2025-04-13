package config

import "os"

func GetMongoURI() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb+srv://aryanprashant123:yJUrDdDEOyxR1AnO@tasks.q0whgvp.mongodb.net/"
	}
	return uri
}
