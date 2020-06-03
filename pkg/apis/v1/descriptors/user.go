package descriptors

import (
	"github.com/wangkailiang-caiyun/Nirvana-Testing/pkg/user"

	def "github.com/caicloud/nirvana/definition"
)

func init() {
	register(
		[]def.Descriptor{
			{
				Path: "/users",
				Definitions: []def.Definition{
					getUserList,
					createUser,
				},
			},
			{
				Path: "/user/{userID}",
				Definitions: []def.Definition{
					getUser,
					deleteUser,
					updateUser,
				},
			},
		}...,
	)
}

var createUser = def.Definition{
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

var deleteUser = def.Definition{
	Method:      def.Delete,
	Summary:     "delete a User",
	Description: "delete user",
	Function:    user.DeleteUser,
	Parameters: []def.Parameter{
		def.PathParameterFor("userID", "user Id"),
	},
	Results: def.DataErrorResults("success flag"),
}

var updateUser = def.Definition{
	Method:      def.Update,
	Summary:     "update User information",
	Description: "update user information",
	Function:    user.UpdateUser,
	Parameters: []def.Parameter{
		def.PathParameterFor("userID", "user Id"),
		{
			Source:      def.Body,
			Name:        "user",
			Default:     nil,
			Description: "user information",
		},
	},
	Results: def.DataErrorResults("success flag"),
}

var getUserList = def.Definition{
	Method:      def.List,
	Summary:     "Get All User",
	Description: "fetch all user from mongoDB",
	Function:    user.GetUserList,
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

var getUser = def.Definition{
	Method:      def.Get,
	Summary:     "get one User information",
	Description: "get one User information",
	Function:    user.FetchUser,
	Parameters: []def.Parameter{
		def.PathParameterFor("userID", "user Id"),
	},
	Results: def.DataErrorResults("success flag"),
}
