package models

import (
	"github.com/graphql-go/graphql"
)

type Notifications struct {
	SenderID uint
	Name     string
	Time     string
	Image    string
	Type     string
	Status   string
	CommonID uint
}

var NotificationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Notifications",
	Fields: graphql.Fields{
		"CommonID": &graphql.Field{
			Type: graphql.Int,
		},
		"SenderID": &graphql.Field{
			Type: graphql.Int,
		},
		"Image": &graphql.Field{
			Type: graphql.String,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"Time": &graphql.Field{
			Type: graphql.String,
		},
		"Type": &graphql.Field{
			Type: graphql.String,
		},
		"Status": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var NotificationGraphqlObjectConfig = graphql.ObjectConfig{
	Name: "Notification query",
	Fields: graphql.Fields{
		"notifications": &graphql.Field{
			Type: graphql.NewList(NotificationType),
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"day": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
		},
	},
}

// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 	// Extract arguments from ResolveParams
// 	userID, _ := p.Args["userID"].(int)
// 	day, _ := p.Args["day"].(string)

// 	// Fetch notifications from the database based on userID and day
// 	filteredNotifications := getNotifications(userID, day)

// 	return filteredNotifications, nil
// },
