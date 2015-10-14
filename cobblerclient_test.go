package cobblerclient

import (
	"testing"

	"github.com/ContainerSolutions/go-utils"
)

var config = ClientConfig{
	Url:      "http://example.org:1234",
	Username: "john",
	Password: "doe",
}

func TestLogin(t *testing.T) {
	expectedReq, err := utils.Fixture("login-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("login-res.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response
	c := NewClient(hc, config)
	ok, err := c.Login()
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}

	expected := "sa/1EWr40BWU+Pq3VEOOpD4cQtxkeMuFUw=="
	if c.token != expected {
		t.Errorf(`"%s" expected; got "%s"`, expected, c.token)
	}
}

func TestLoginWithError(t *testing.T) {
	expected := "error 1: <class 'cobbler.cexceptions.CX'>:'login failed (cobbler)'"
	response, err := utils.Fixture("login-res-err.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Response = response
	hc.ShouldVerify = false

	c := NewClient(hc, config)
	ok, err := c.Login()
	if ok {
		t.Errorf("false expected; got true")
	}

	if err.Error() != expected {
		t.Errorf("%s expected; got %s", expected, err)
	}
}

func TestSync(t *testing.T) {
	expectedReq, err := utils.Fixture("sync-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("sync-res.xml")
	utils.FailOnError(t, err)

	expected := true
	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "securetoken99"
	result, err := c.Sync()
	utils.FailOnError(t, err)

	if result != expected {
		t.Errorf("%s expected; got %s", expected, result)
	}
}
