package roleschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/roles"
)

// Schema returns a key/value map of the various fields that make up a OneLogin User.
func Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"apps": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"users": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"admins": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
	}
}

// Inflate takes a key/value map of interfaces and uses the fields to construct a Role
func Inflate(s map[string]interface{}) roles.Role {
	out := roles.Role{}
	if id, notNil := s["id"].(int); notNil {
		out.ID = oltypes.Int32(int32(id))
	}
	if name, notNil := s["name"].(string); notNil {
		out.Name = oltypes.String(name)
	}
	if s["apps"] != nil {
		out.Apps = make([]int32, len(s["apps"].([]interface{})))
		for i, appID := range s["apps"].([]interface{}) {
			out.Apps[i] = int32(appID.(int))
		}
	}
	if s["users"] != nil {
		out.Users = make([]int32, len(s["users"].([]interface{})))
		for i, userID := range s["users"].([]interface{}) {
			out.Users[i] = int32(userID.(int))
		}
	}
	if s["admins"] != nil {
		out.Admins = make([]int32, len(s["admins"].([]interface{})))
		for i, adminID := range s["admins"].([]interface{}) {
			out.Admins[i] = int32(adminID.(int))
		}
	}
	return out
}
