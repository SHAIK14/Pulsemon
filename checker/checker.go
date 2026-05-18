package checker

import (
	"net/http"
	"time"

	"github.com/SHAIK14/pulsemon/config"
)

type Result struct {
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	IsUp         bool          `json:"is_up"`
	ResponseTime time.Duration `json:"responseTime"`
	StatusCode   int           `json:"statusCode"`
	Error        error         `json:"error"`
}

func Check(service config.Service) Result {
	start := time.Now()
	resp, err := http.Get(service.URL)
	endtime := time.Since(start)
	if err != nil {
		return Result{
			Name:  service.Name,
			URL:   service.URL,
			IsUp:  false,
			Error: err,
		}
	}
	defer resp.Body.Close()
	return Result{
		Name:         service.Name,
		URL:          service.URL,
		IsUp:         true,
		ResponseTime: endtime,
		StatusCode:   resp.StatusCode,
	}

}
