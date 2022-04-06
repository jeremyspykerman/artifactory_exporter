package artifactory

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/log/level"
)

const quotaProperty = "repository.path.quota"

type Properties struct {
	Properties map[string][]string `json:"properties"`
	Uri        string              `json:"uri"`
}

// Fetch the repo quota in bytes or 0 on failure
func (c *Client) FetchRepoQuota(repository string) uint64 {
	bytes, err := c.FetchHTTP(getPropertiesPath(repository))

	if err == nil {
		quota, err := extractQuota(bytes)
		if err == nil {
			return quota
		} else {
			// Because we fetch the quota properties separately for every repo, it seems like
			// a mistake to set Artifactory up value to 0 for any one of them failing
			level.Warn(c.logger).Log("msg", "attempt to extract quota bytes failed", "err", err)
		}
	}

	return 0
}

func extractQuota(bytes []byte) (uint64, error) {
	parsed := Properties{}
	err := json.Unmarshal(bytes, &parsed)

	if err == nil {
		value := parsed.Properties[quotaProperty]
		if len(value) > 0 {
			result, err := strconv.ParseUint(value[0], 10, 0)
			return result, err
		}
	}

	return 0, err
}

func getPropertiesPath(repository string) string {
	return fmt.Sprintf("storage/%s/?properties", repository)
}
