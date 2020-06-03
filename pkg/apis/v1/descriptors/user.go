package descriptors

import (
	def "github.com/caicloud/nirvana/definition"
	"github.com/wangkailiang-caiyun/Nirvana-Testing/pkg/user"
)

func init() {
	register(
		[]def.Descriptor{
			{
				Path:        "/users",
				Definitions: []def.Definition{fetchUserListDef},
			},
			{
				Path: "/user",
				Definitions: []def.Definition{
					createUserDef,
					updateUserDef,
				},
			},
			{
				Path: "/user/{userID}",
				Definitions: []def.Definition{
					fetchUserDef,
					deleteUserDef,
				},
			},
		}...,
	)
}

var createUserDef = def.Definition{
	Method:      def.Create,
	Summary:     "create User information",
	Description: "create user information",
	Function:    user.CreateUser,
	Parameters: []def.Parameter{
		{
			Source:      def.Body,
			Name:        "user",
			Default:     nil,
			Description: "user information",
		},
	},
	Results: def.DataErrorResults("success flag"),
}

var deleteUserDef = def.Definition{
	Method:      def.Delete,
	Summary:     "delete a User",
	Description: "delete user from mongoDB",
	Function:    user.DeleteUser,
	Parameters: []def.Parameter{
		def.PathParameterFor("userID", "user Id"),
	},
	Results: def.DataErrorResults("success flag"),
}

var updateUserDef = def.Definition{
	Method:      def.Update,
	Summary:     "update User information",
	Description: "update user information",
	Function:    user.UpdateUser,
	Parameters: []def.Parameter{
		{
			Source:      def.Body,
			Name:        "user",
			Default:     nil,
			Description: "user information",
		},
	},
	Results: def.DataErrorResults("success flag"),
}

var fetchUserListDef = def.Definition{
	Method:      def.List,
	Summary:     "Get All User",
	Description: "fetch all user from mongoDB",
	Function:    user.FetchUserList,
	Parameters: []def.Parameter{
		{
			Source:      def.Query,
			Name:        "pageSize",
			Default:     10,
			Description: "page size",
		},
		{
			Source:      def.Query,
			Name:        "pageNo",
			Default:     1,
			Description: "page number",
		},
	},
	Results: def.DataErrorResults("a list of user"),
}

var fetchUserDef = def.Definition{
	Method:      def.Get,
	Summary:     "get one User information",
	Description: "get one User information",
	Function:    user.FetchUser,
	Parameters: []def.Parameter{
		def.PathParameterFor("userID", "user Id"),
	},
	Results: def.DataErrorResults("success flag"),
}
