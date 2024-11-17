package utils

import (
	"fmt"

	"github.com/jichenssg/tikmall/app/auth/config"
)

func Enforce(userID int64, path, method string) (bool, error) {
	return config.Enforcer.Enforce(fmt.Sprint(userID), path, method)
}

func EnforceAnonymous(path, method string) (bool, error) {
	return config.Enforcer.Enforce("anonymous", path, method)
}

func AddUserRole(userID int64) error {
	return addRole(userID, "user")
}

func AddAdminRole(userID int64) error {
	return addRole(userID, "admin")
}

func RemoveUserRole(userID int64) error {
	return removeRole(userID, "user")
}

func RemoveAdminRole(userID int64) error {
	return removeRole(userID, "admin")
}

func RemoveAllRoles(userID int64) error {
	// Remove all roles for the user
	_, err := config.Enforcer.DeleteRolesForUser(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	// Save the policy back to the file
	err = config.Enforcer.SavePolicy()
	if err != nil {
		return err
	}

	return nil
}

func addRole(userID int64, role string) error {
	// Add the role for the user
	_, err := config.Enforcer.AddRoleForUser(fmt.Sprint(userID), role)
	if err != nil {
		return err
	}

	// Save the policy back to the file
	err = config.Enforcer.SavePolicy()
	if err != nil {
		return err
	}

	return nil
}

func removeRole(userID int64, role string) error {
	// Remove the role for the user
	_, err := config.Enforcer.DeleteRoleForUser(fmt.Sprint(userID), role)
	if err != nil {
		return err
	}

	// Save the policy back to the file
	err = config.Enforcer.SavePolicy()
	if err != nil {
		return err
	}

	return nil
}
