// utils/errors.go
package utils

import "log"

// LogError logs the error to the standard output
func LogError(err error) {
    if err != nil {
        log.Println("ERROR:", err)
    }
}
