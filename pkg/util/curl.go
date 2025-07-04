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

// getFile performs the basic HTTP GET operation
// (private because it's our internal building block)
func getFile2(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("❌ Error: failed to get URL: %w", err)
	}
	defer resp.Body.Close()

	// check
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("❌ Error: unexpected status code: %d for URL: %s", resp.StatusCode, url)
	}

	// check
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorx.ByteError("Get Response Body from URL (%s), even status code is 200", url, err)
	}
	// success
	return body, nil
}

// gets a file and handles local-specific concerns
func GetFile(url string, path string) (string, error) {

	// get file in memory
	data, err := getFile2(url)
	/// error
	if err != nil {
		return "", err
	}

	// save file from memory to FS
	_, err = SaveToFile(data, path)
	/// error
	if err != nil {
		return "", err
	}

	// success
	return path, nil
}

// // gets a file and handles local-specific concerns
// func GetFileRemote(url string, path string, vm string) ([]byte, error) {
// 	cli := fmt.Sprintf("luc util url get %s %s --local", url, path)
// 	_, err := RunCLIRemote(cli, vm)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }
