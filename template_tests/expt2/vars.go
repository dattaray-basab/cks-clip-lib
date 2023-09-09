package expt2

import (

)
	var DataVal []Data = []Data{
		{
			Name: "Server",
			Desc: "Details about a server",
			Fields: []Field{
				{Name: "Id", TypeName: "int"},
				{Name: "Url", TypeName: "string"},
				{Name: "Active", TypeName: "boolean"},
			},
		},
		{
			Name: "Channel",
			Desc: "A channel to chat in the server",
			Fields: []Field{
				{Name: "Id", TypeName: "int"},
				{Name: "Name", TypeName: "string"},
				{Name: "Private", TypeName: "boolean"},
			},
		},
		{
			Name: "Role",
			Desc: "Used to group permissions for users",
			Fields: []Field{
				{Name: "Id", TypeName: "int"},
				{Name: "Name", TypeName: "string"},
				{Name: "Color", TypeName: "string"},
			},
		},
	}