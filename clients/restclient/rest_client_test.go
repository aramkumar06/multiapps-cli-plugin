package restclient_test

import (
	"net/http"
	"net/http/cookiejar"

	baseclient "github.com/cloudfoundry-incubator/multiapps-cli-plugin/clients/baseclient"
	restclient "github.com/cloudfoundry-incubator/multiapps-cli-plugin/clients/restclient"
	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/testutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RestClient", func() {

	Describe("PurgeConfiguration", func() {
		Context("when the backend returns not 204 No Content", func() {
			It("should return an error", func() {
				client := newRestClient(http.StatusInternalServerError, nil)
				err := client.PurgeConfiguration("org", "space")
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("when the backend returns 204 No Content", func() {
			It("should not return an error", func() {
				client := newRestClient(http.StatusNoContent, nil)
				err := client.PurgeConfiguration("org", "space")
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})

func newRestClient(statusCode int, v interface{}) restclient.RestClientOperations {
	tokenFactory := baseclient.NewCustomTokenFactory("test-token")
	cookieJar, _ := cookiejar.New(nil)
	roundTripper := testutil.NewCustomTransport(statusCode, v)
	return restclient.NewRestClient("http://localhost:1000", roundTripper, cookieJar, tokenFactory)
}
