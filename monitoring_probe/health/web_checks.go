package health

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "time"
)

// CheckHealth performs both liveness and readiness checks.
func CheckHealth(port int) {
    CheckLiveness(port, "liveness")
    CheckReadiness(port, "readiness")
}

// CheckLiveness performs a liveness check on the given port.
func CheckLiveness(port int, endpoint string) {
    current_time := time.Now()
    formatted_current_time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        current_time.Year(),
        current_time.Month(),
        current_time.Day(),
        current_time.Hour(),
        current_time.Minute(),
        current_time.Second())
    fmt.Printf("Port: %d  %s  %s\n", port, strings.ToUpper(endpoint), formatted_current_time)

    requestURL := fmt.Sprintf("http://127.0.0.1:%d/health/%s", port, endpoint)
    req, err := http.NewRequest(http.MethodGet, requestURL, nil)
    if err != nil {
        fmt.Printf("error creating http request: %s\n", err)
        os.Exit(1)
    }
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }

    fmt.Printf("client: got response!\n")
    fmt.Printf("client: status code: %d\n", res.StatusCode)
    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Printf("client: could not read response body: %s\n", err)
        os.Exit(1)
    }

    fmt.Printf("client: response body: %s\n", resBody)
    if res.StatusCode != http.StatusOK {
        fmt.Printf("client: liveness check failed\n")
        os.Exit(1)
    }
}

// CheckReadiness performs a readiness check on the given port.
func CheckReadiness(port int, endpoint string) {
    current_time := time.Now()
    formatted_current_time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
        current_time.Year(),
        current_time.Month(),
        current_time.Day(),
        current_time.Hour(),
        current_time.Minute(),
        current_time.Second())

    fmt.Printf("Port: %d  %s  %s\n", port, strings.ToUpper(endpoint), formatted_current_time)
    requestURL := fmt.Sprintf("http://127.0.0.1:%d/health/%s", port, endpoint)

    req, err := http.NewRequest(http.MethodGet, requestURL, nil)
    if err != nil {
        fmt.Printf("error creating http request: %s\n", err)
        os.Exit(1)
    }
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }

    fmt.Printf("client: got response!\n")
    fmt.Printf("client: status code: %d\n", res.StatusCode)
    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Printf("client: could not read response body: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("client: response body: %s\n", resBody)
    if res.StatusCode != http.StatusOK {
        fmt.Printf("client: Readiness check failed\n")
        os.Exit(1)
    }
}