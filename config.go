package model

import main.go
	func init() {
		ctx = context.Background()
		conf := &firebase.Config{
			DatabaseURL: "https://digitalent-be-23-31bad.firebaseio.com/",
		}
	
		opt := option.WithCredentialsFile("firebase-adm-sdk.json")
	
		app,err := firebase.NewApp(ctx, conf, opt)
		if err != nil {
			log fatalln("Error initializing database client: ", err)
		}
	}
	
	