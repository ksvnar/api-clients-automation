// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package requests

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"

	"github.com/joho/godotenv"

	"gotests/tests"

	suggestions "github.com/algolia/algoliasearch-client-go/v4/algolia/query-suggestions"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/transport"
)

func createSuggestionsClient(t *testing.T) (*suggestions.APIClient, *tests.EchoRequester) {
	t.Helper()

	echo := &tests.EchoRequester{}
	cfg := suggestions.Configuration{
		Configuration: transport.Configuration{
			AppID:     "appID",
			ApiKey:    "apiKey",
			Requester: echo,
		},
		Region: suggestions.US,
	}
	client, err := suggestions.NewClientWithConfig(cfg)
	require.NoError(t, err)

	return client, echo
}

func createE2ESuggestionsClient(t *testing.T) *suggestions.APIClient {
	t.Helper()

	appID := os.Getenv("ALGOLIA_APPLICATION_ID")
	if appID == "" && os.Getenv("CI") != "true" {
		err := godotenv.Load("../../../../.env")
		require.NoError(t, err)
		appID = os.Getenv("ALGOLIA_APPLICATION_ID")
	}
	apiKey := os.Getenv("ALGOLIA_ADMIN_KEY")
	client, err := suggestions.NewClient(appID, apiKey, suggestions.US)
	require.NoError(t, err)

	return client
}

func TestSuggestions_CreateConfig(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("createConfig0", func(t *testing.T) {
		_, err := client.CreateConfig(client.NewApiCreateConfigRequest(

			suggestions.NewEmptyQuerySuggestionsConfigurationWithIndex().SetIndexName("theIndexName").SetSourceIndices(
				[]suggestions.SourceIndex{*suggestions.NewEmptySourceIndex().SetIndexName("testIndex").SetFacets(
					[]suggestions.Facet{*suggestions.NewEmptyFacet().SetAttribute("test")}).SetGenerate(
					[][]string{
						[]string{"facetA", "facetB"},
						[]string{"facetC"}})}).SetLanguages(suggestions.ArrayOfStringAsLanguages(
				[]string{"french"})).SetExclude(
				[]string{"test"}),
		))
		require.NoError(t, err)

		require.Equal(t, "/1/configs", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"indexName":"theIndexName","sourceIndices":[{"indexName":"testIndex","facets":[{"attribute":"test"}],"generate":[["facetA","facetB"],["facetC"]]}],"languages":["french"],"exclude":["test"]}`)
	})
}

func TestSuggestions_CustomDelete(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("allow del method for a custom path with minimal parameters", func(t *testing.T) {
		_, err := client.CustomDelete(client.NewApiCustomDeleteRequest(
			"test/minimal",
		))
		require.NoError(t, err)

		require.Equal(t, "/test/minimal", echo.Path)
		require.Equal(t, "DELETE", echo.Method)

		require.Nil(t, echo.Body)
	})
	t.Run("allow del method for a custom path with all parameters", func(t *testing.T) {
		_, err := client.CustomDelete(client.NewApiCustomDeleteRequest(
			"test/all",
		).WithParameters(map[string]any{"query": "parameters"}))
		require.NoError(t, err)

		require.Equal(t, "/test/all", echo.Path)
		require.Equal(t, "DELETE", echo.Method)

		require.Nil(t, echo.Body)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
}

func TestSuggestions_CustomGet(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("allow get method for a custom path with minimal parameters", func(t *testing.T) {
		_, err := client.CustomGet(client.NewApiCustomGetRequest(
			"test/minimal",
		))
		require.NoError(t, err)

		require.Equal(t, "/test/minimal", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
	})
	t.Run("allow get method for a custom path with all parameters", func(t *testing.T) {
		_, err := client.CustomGet(client.NewApiCustomGetRequest(
			"test/all",
		).WithParameters(map[string]any{"query": "parameters with space"}))
		require.NoError(t, err)

		require.Equal(t, "/test/all", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters%20with%20space"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions should be escaped too", func(t *testing.T) {
		_, err := client.CustomGet(client.NewApiCustomGetRequest(
			"test/all",
		).WithParameters(map[string]any{"query": "to be overriden"}),
			suggestions.QueryParamOption("query", "parameters with space"), suggestions.QueryParamOption("and an array",
				[]string{"array", "with spaces"}), suggestions.HeaderParamOption("x-header-1", "spaces are left alone"),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/all", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
		headers := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"x-header-1":"spaces are left alone"}`), &headers))
		for k, v := range headers {
			require.Equal(t, v, echo.Header.Get(k))
		}
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters%20with%20space","and%20an%20array":"array%2Cwith%20spaces"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
}

func TestSuggestions_CustomPost(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("allow post method for a custom path with minimal parameters", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/minimal",
		))
		require.NoError(t, err)

		require.Equal(t, "/test/minimal", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{}`)
	})
	t.Run("allow post method for a custom path with all parameters", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/all",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"body": "parameters"}))
		require.NoError(t, err)

		require.Equal(t, "/test/all", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"body":"parameters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions can override default query parameters", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("query", "myQueryParameter"),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"myQueryParameter"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions merges query parameters with default ones", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("query2", "myQueryParameter"),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters","query2":"myQueryParameter"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions can override default headers", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.HeaderParamOption("x-algolia-api-key", "myApiKey"),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		headers := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"x-algolia-api-key":"myApiKey"}`), &headers))
		for k, v := range headers {
			require.Equal(t, v, echo.Header.Get(k))
		}
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions merges headers with default ones", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.HeaderParamOption("x-algolia-api-key", "myApiKey"),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		headers := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"x-algolia-api-key":"myApiKey"}`), &headers))
		for k, v := range headers {
			require.Equal(t, v, echo.Header.Get(k))
		}
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions queryParameters accepts booleans", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("isItWorking", true),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters","isItWorking":"true"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions queryParameters accepts integers", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("myParam", 2),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters","myParam":"2"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions queryParameters accepts list of string", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("myParam",
				[]string{"b and c", "d"}),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters","myParam":"b%20and%20c%2Cd"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions queryParameters accepts list of booleans", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("myParam",
				[]bool{true, true, false}),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters","myParam":"true%2Ctrue%2Cfalse"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
	t.Run("requestOptions queryParameters accepts list of integers", func(t *testing.T) {
		_, err := client.CustomPost(client.NewApiCustomPostRequest(
			"test/requestOptions",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"facet": "filters"}),
			suggestions.QueryParamOption("myParam",
				[]int32{1, 2}),
		)
		require.NoError(t, err)

		require.Equal(t, "/test/requestOptions", echo.Path)
		require.Equal(t, "POST", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"facet":"filters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters","myParam":"1%2C2"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
}

func TestSuggestions_CustomPut(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("allow put method for a custom path with minimal parameters", func(t *testing.T) {
		_, err := client.CustomPut(client.NewApiCustomPutRequest(
			"test/minimal",
		))
		require.NoError(t, err)

		require.Equal(t, "/test/minimal", echo.Path)
		require.Equal(t, "PUT", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{}`)
	})
	t.Run("allow put method for a custom path with all parameters", func(t *testing.T) {
		_, err := client.CustomPut(client.NewApiCustomPutRequest(
			"test/all",
		).WithParameters(map[string]any{"query": "parameters"}).WithBody(map[string]any{"body": "parameters"}))
		require.NoError(t, err)

		require.Equal(t, "/test/all", echo.Path)
		require.Equal(t, "PUT", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"body":"parameters"}`)
		queryParams := map[string]string{}
		require.NoError(t, json.Unmarshal([]byte(`{"query":"parameters"}`), &queryParams))
		require.Len(t, queryParams, len(echo.Query))
		for k, v := range queryParams {
			require.Equal(t, v, echo.Query.Get(k))
		}
	})
}

func TestSuggestions_DeleteConfig(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("deleteConfig0", func(t *testing.T) {
		_, err := client.DeleteConfig(client.NewApiDeleteConfigRequest(
			"theIndexName",
		))
		require.NoError(t, err)

		require.Equal(t, "/1/configs/theIndexName", echo.Path)
		require.Equal(t, "DELETE", echo.Method)

		require.Nil(t, echo.Body)
	})
}

func TestSuggestions_GetAllConfigs(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("getAllConfigs0", func(t *testing.T) {
		_, err := client.GetAllConfigs()
		require.NoError(t, err)

		require.Equal(t, "/1/configs", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
	})
}

func TestSuggestions_GetConfig(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("Retrieve QS config e2e", func(t *testing.T) {
		_, err := client.GetConfig(client.NewApiGetConfigRequest(
			"cts_e2e_browse_query_suggestions",
		))
		require.NoError(t, err)

		require.Equal(t, "/1/configs/cts_e2e_browse_query_suggestions", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
		clientE2E := createE2ESuggestionsClient(t)
		res, err := clientE2E.GetConfig(client.NewApiGetConfigRequest(
			"cts_e2e_browse_query_suggestions",
		))
		require.NoError(t, err)
		_ = res

		rawBody, err := json.Marshal(res)
		require.NoError(t, err)

		var rawBodyMap any
		err = json.Unmarshal(rawBody, &rawBodyMap)
		require.NoError(t, err)

		expectedBodyRaw := `{"appID":"T8JK9S7I7X","allowSpecialCharacters":true,"enablePersonalization":false,"exclude":["^cocaines$"],"indexName":"cts_e2e_browse_query_suggestions","languages":[],"sourceIndices":[{"facets":[{"amount":1,"attribute":"title"}],"generate":[["year"]],"indexName":"cts_e2e_browse","minHits":5,"minLetters":4,"replicas":false}]}`
		var expectedBody any
		err = json.Unmarshal([]byte(expectedBodyRaw), &expectedBody)
		require.NoError(t, err)

		unionBody := tests.Union(expectedBody, rawBodyMap)
		unionBodyRaw, err := json.Marshal(unionBody)
		require.NoError(t, err)

		jaE2E := jsonassert.New(t)
		jaE2E.Assertf(expectedBodyRaw, strings.ReplaceAll(string(unionBodyRaw), "%", "%%"))
	})
}

func TestSuggestions_GetConfigStatus(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("getConfigStatus0", func(t *testing.T) {
		_, err := client.GetConfigStatus(client.NewApiGetConfigStatusRequest(
			"theIndexName",
		))
		require.NoError(t, err)

		require.Equal(t, "/1/configs/theIndexName/status", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
	})
}

func TestSuggestions_GetLogFile(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("getLogFile0", func(t *testing.T) {
		_, err := client.GetLogFile(client.NewApiGetLogFileRequest(
			"theIndexName",
		))
		require.NoError(t, err)

		require.Equal(t, "/1/logs/theIndexName", echo.Path)
		require.Equal(t, "GET", echo.Method)

		require.Nil(t, echo.Body)
	})
}

func TestSuggestions_UpdateConfig(t *testing.T) {
	client, echo := createSuggestionsClient(t)
	_ = echo

	t.Run("updateConfig0", func(t *testing.T) {
		_, err := client.UpdateConfig(client.NewApiUpdateConfigRequest(
			"theIndexName",
			suggestions.NewEmptyQuerySuggestionsConfiguration().SetSourceIndices(
				[]suggestions.SourceIndex{*suggestions.NewEmptySourceIndex().SetIndexName("testIndex").SetFacets(
					[]suggestions.Facet{*suggestions.NewEmptyFacet().SetAttribute("test")}).SetGenerate(
					[][]string{
						[]string{"facetA", "facetB"},
						[]string{"facetC"}})}).SetLanguages(suggestions.ArrayOfStringAsLanguages(
				[]string{"french"})).SetExclude(
				[]string{"test"}),
		))
		require.NoError(t, err)

		require.Equal(t, "/1/configs/theIndexName", echo.Path)
		require.Equal(t, "PUT", echo.Method)

		ja := jsonassert.New(t)
		ja.Assertf(*echo.Body, `{"sourceIndices":[{"indexName":"testIndex","facets":[{"attribute":"test"}],"generate":[["facetA","facetB"],["facetC"]]}],"languages":["french"],"exclude":["test"]}`)
	})
}
