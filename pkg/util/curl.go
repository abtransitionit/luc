/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/abtransitionit/luc/pkg/errorx"
)

// # Purpose
//
//   - Downloads a file form a public URL into memory and returns it as a byte slice.
//   - performs an HTTP GET request.
//
// Parameters:
//   - url: The HTTPS URL of the file to download. Must be a valid URL.
//
// Returns:
//   - []byte : the downloaded file on suuccee (nil on failure)
//   - error  : Failure details (nil on success)
//
// Possible returns:
//
//   - ([]byte, nil): On success
//   - (nil, error) : On failure
//     -- Network errors
//     -- Non-200 HTTP status codes
//     -- Response body reading errors
//
// Example usage:
//
//	fileInMemory, err := util.GetPublicFile(logx.L, cliUrl)
//	if err != nil {
//			return
//	}
//
// Example usage:
//
//	content, err := CurlPublicFile("https://example.com/file.txt")
//	if err != nil {
//	    log.Fatalf("Download failed: %v", err)
//	}
//
// Notes:
//   - The function follows HTTP redirects automatically (via http.DefaultClient)
//   - There is a default 10-second timeout (via http.DefaultClient)
//   - The response body is automatically closed after reading
//   - The caller is responsible for handling the returned data.
func GetPublicFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return errorx.ByteError("Get URL", url, err)
	}
	// close the response body at the end
	defer resp.Body.Close()

	// manage status code
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Get correct HTTP status code (%s) for url %s", resp.Status, url)
		return errorx.ByteError(msg, "", errors.New(""))
	}
	// here: status code is 200

	// Read the response body - Get the file content
	body, err := io.ReadAll(resp.Body)
	// handle system FAILURE
	if err != nil {
		return errorx.ByteError("Get Response Body from URL (%s), even status code is 200", url, err)
	}
	// handle applogic SUCCESS - here file content exists
	// log.Infof("✅ data donwloaded into memory")
	return body, nil
}

// out, err := os.Create(srcPath)
// if err != nil {
// 	fmt.Println("File creation error:", err)
// 	return
// }
// defer out.Close()

// _, err = io.Copy(out, resp.Body)
// if err != nil {
// 	fmt.Println("File save error:", err)
// 	return
// }

// 1. Download file
// err = os.WriteFile(srcPath, data, 0644)
// if err != nil {
// 	fmt.Println("write error:", err)
// 	return
// }

// // 2. Make file executable
// err = os.Chmod(srcPath, 0755)
// if err != nil {
// 	fmt.Println("chmod error:", err)
// 	return
// }

// // 3. Move file to final destination (requires root privileges)
// err = os.Rename(srcPath, dstPath)
// if err != nil {
// 	fmt.Println("move error:", err)
// 	return
// }
