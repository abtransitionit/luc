/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package errorx

import (
	"fmt"
)

// Generic error for any system-level errors rather than an application logic
//
// # Example Errors:
//   - Database connection failures
//   - Network issues during checks
//   - Permission denied errors
//
// # Parameters:
//   - objectAction         : including the object type (e.g., "user", "file", "folder")
//   - objectId 						: Name/ID/property of the resource
//   - err                  : Original system error
//
// # Returns
//   - false, wrappedError
//
// # Usage
//
//	if err != nil {
//			return errorx.BoolError("read user input", "from stdin", err)
//	}
//
// Note: Function intentionally only returns and does not include logging. Caller should handle logging if needed based on their context
func BoolError(objectAction string, objectId string, err error) (bool, error) {
	return false, fmt.Errorf("❌ Errror : failed to  %s : %s > %w",
		objectAction,
		objectId,
		err,
	)
}

// Generic error for any system-level errors rather than an application logic
//
// # Example Errors:
//   - Database connection failures
//   - Network issues during checks
//   - Permission denied errors
//
// # Parameters:
//   - objectAction         : including the object type (e.g., "user", "file", "folder")
//   - objectId 						: Name/ID/property of the resource
//   - err                  : Original system error
//
// # Returns
//   - "", wrappedError
//
// # Usage
//
//	if err != nil {
//			return errorx.StringError("read user input", "from stdin", err)
//	}
//
// Note: Function intentionally only returns and does not include logging. Caller should handle logging if needed based on their context
func StringError(objectAction string, objectId string, err error) (string, error) {
	return "", fmt.Errorf("❌ Errror : failed to  %s : %s > %w",
		objectAction,
		objectId,
		err,
	)
}

// Generic error for any system-level errors rather than an application logic
//
// # Example Errors:
//   - Database connection failures
//   - Network issues during checks
//   - Permission denied errors
//
// # Parameters:
//   - objectAction         : including the object type (e.g., "user", "file", "folder")
//   - objectId 						: Name/ID/property of the resource
//   - err                  : Original system error
//
// # Returns
//   - nil, wrappedError
//
// # Usage
//
//	if err != nil {
//			return errorx.ByteError("read user input", "from stdin", err)
//	}
//
// Note: Function intentionally only returns and does not include logging. Caller should handle logging if needed based on their context
func ByteError(objectAction string, objectId string, err error) ([]byte, error) {
	return nil, fmt.Errorf("❌ Errror : failed to  %s : %s > %w",
		objectAction,
		objectId,
		err,
	)
}
