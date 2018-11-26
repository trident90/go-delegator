package metaservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"bitbucket.org/coinplugin/proxy/log"
)

type reCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
	Hostname    string    `json:"hostname"`     // the hostname of the site where the reCAPTCHA was solved
	ErrorCodes  []string  `json:"error-codes"`  //optional
}

func checkRecaptcha(response string) error {

	postStr := url.Values{"secret": {reCaptchaSecret}, "response": {response}}

	responsePost, err := http.PostForm(reCaptchaURL, postStr)

	if err != nil {
		log.Error("Post Error", err)
		return err
	}

	defer responsePost.Body.Close()
	body, err := ioutil.ReadAll(responsePost.Body)

	if err != nil {
		log.Error("Read data Error", err)
		return err
	}
	var result reCaptchaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error(" Read JSON Error", err)
		return err
	}
	if !result.Success {
		log.Debug(result.ErrorCodes)
		return fmt.Errorf("Fail to Verify reCaptcha")
	}
	return nil
}

func intToByte32(_val *big.Int) [32]byte {
	result := [32]byte{}
	if _val != nil {
		_bytes := _val.Bytes()
		copy(result[32-len(_bytes):], _bytes)
	}
	return result
}
func concatBytes(args ...[]byte) []byte {
	var ret []byte
	for _, a := range args {
		ret = append(ret, a...)
	}
	return ret
}
