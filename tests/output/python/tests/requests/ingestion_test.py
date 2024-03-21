from json import loads
from os import environ

from algoliasearch.http.transporter import EchoTransporter
from algoliasearch.ingestion.client import IngestionClient
from algoliasearch.ingestion.config import IngestionConfig
from dotenv import load_dotenv

from ..helpers import Helpers

load_dotenv("../../.env")


class TestIngestionClient:
    _config = IngestionConfig("test_app_id", "test_api_key", "us")
    _client = IngestionClient.create_with_config(
        config=_config, transporter=EchoTransporter(_config)
    )

    _helpers = Helpers()
    _e2e_app_id = environ.get("ALGOLIA_APPLICATION_ID")
    if _e2e_app_id is None:
        raise Exception(
            "please provide an `ALGOLIA_APPLICATION_ID` env var for e2e tests"
        )

    _e2e_api_key = environ.get("ALGOLIA_ADMIN_KEY")
    if _e2e_api_key is None:
        raise Exception("please provide an `ALGOLIA_ADMIN_KEY` env var for e2e tests")

    async def test_create_authentication_0(self):
        """
        createAuthenticationOAuth
        """
        _req = await self._client.create_authentication_with_http_info(
            authentication_create={
                "type": "oauth",
                "name": "authName",
                "input": {
                    "url": "http://test.oauth",
                    "client_id": "myID",
                    "client_secret": "mySecret",
                },
            },
        )

        assert _req.path == "/1/authentications"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"type":"oauth","name":"authName","input":{"url":"http://test.oauth","client_id":"myID","client_secret":"mySecret"}}"""
        )

    async def test_create_authentication_1(self):
        """
        createAuthenticationAlgolia
        """
        _req = await self._client.create_authentication_with_http_info(
            authentication_create={
                "type": "algolia",
                "name": "authName",
                "input": {
                    "appID": "myappID",
                    "apiKey": "randomApiKey",
                },
            },
        )

        assert _req.path == "/1/authentications"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"type":"algolia","name":"authName","input":{"appID":"myappID","apiKey":"randomApiKey"}}"""
        )

    async def test_create_destination_0(self):
        """
        createDestination
        """
        _req = await self._client.create_destination_with_http_info(
            destination_create={
                "type": "search",
                "name": "destinationName",
                "input": {
                    "indexPrefix": "prefix_",
                },
                "authenticationID": "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            },
        )

        assert _req.path == "/1/destinations"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"type":"search","name":"destinationName","input":{"indexPrefix":"prefix_"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}"""
        )

    async def test_create_source_0(self):
        """
        createSource
        """
        _req = await self._client.create_source_with_http_info(
            source_create={
                "type": "commercetools",
                "name": "sourceName",
                "input": {
                    "storeKeys": [
                        "myStore",
                    ],
                    "locales": [
                        "de",
                    ],
                    "url": "http://commercetools.com",
                    "projectKey": "keyID",
                },
                "authenticationID": "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            },
        )

        assert _req.path == "/1/sources"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"type":"commercetools","name":"sourceName","input":{"storeKeys":["myStore"],"locales":["de"],"url":"http://commercetools.com","projectKey":"keyID"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}"""
        )

    async def test_create_task_0(self):
        """
        createTaskOnDemand
        """
        _req = await self._client.create_task_with_http_info(
            task_create={
                "sourceID": "search",
                "destinationID": "destinationName",
                "trigger": {
                    "type": "onDemand",
                },
                "action": "replace",
            },
        )

        assert _req.path == "/1/tasks"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"onDemand"},"action":"replace"}"""
        )

    async def test_create_task_1(self):
        """
        createTaskSchedule
        """
        _req = await self._client.create_task_with_http_info(
            task_create={
                "sourceID": "search",
                "destinationID": "destinationName",
                "trigger": {
                    "type": "schedule",
                    "cron": "* * * * *",
                },
                "action": "replace",
            },
        )

        assert _req.path == "/1/tasks"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"schedule","cron":"* * * * *"},"action":"replace"}"""
        )

    async def test_create_task_2(self):
        """
        createTaskSubscription
        """
        _req = await self._client.create_task_with_http_info(
            task_create={
                "sourceID": "search",
                "destinationID": "destinationName",
                "trigger": {
                    "type": "onDemand",
                },
                "action": "replace",
            },
        )

        assert _req.path == "/1/tasks"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"onDemand"},"action":"replace"}"""
        )

    async def test_custom_delete_0(self):
        """
        allow del method for a custom path with minimal parameters
        """
        _req = await self._client.custom_delete_with_http_info(
            path="test/minimal",
        )

        assert _req.path == "/test/minimal"
        assert _req.verb == "DELETE"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_custom_delete_1(self):
        """
        allow del method for a custom path with all parameters
        """
        _req = await self._client.custom_delete_with_http_info(
            path="test/all",
            parameters={
                "query": "parameters",
            },
        )

        assert _req.path == "/test/all"
        assert _req.verb == "DELETE"
        assert _req.query_parameters.items() == {"query": "parameters"}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_custom_get_0(self):
        """
        allow get method for a custom path with minimal parameters
        """
        _req = await self._client.custom_get_with_http_info(
            path="test/minimal",
        )

        assert _req.path == "/test/minimal"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_custom_get_1(self):
        """
        allow get method for a custom path with all parameters
        """
        _req = await self._client.custom_get_with_http_info(
            path="test/all",
            parameters={
                "query": "parameters with space",
            },
        )

        assert _req.path == "/test/all"
        assert _req.verb == "GET"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters%20with%20space"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_custom_get_2(self):
        """
        requestOptions should be escaped too
        """
        _req = await self._client.custom_get_with_http_info(
            path="test/all",
            parameters={
                "query": "to be overriden",
            },
            request_options={
                "headers": loads("""{"x-header-1":"spaces are left alone"}"""),
                "query_parameters": loads(
                    """{"query":"parameters with space","and an array":["array","with spaces"]}"""
                ),
            },
        )

        assert _req.path == "/test/all"
        assert _req.verb == "GET"
        assert (
            _req.query_parameters.items()
            == {
                "query": "parameters%20with%20space",
                "and%20an%20array": "array%2Cwith%20spaces",
            }.items()
        )
        assert _req.headers.items() >= {"x-header-1": "spaces are left alone"}.items()
        assert _req.data is None

    async def test_custom_post_0(self):
        """
        allow post method for a custom path with minimal parameters
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/minimal",
        )

        assert _req.path == "/test/minimal"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{}""")

    async def test_custom_post_1(self):
        """
        allow post method for a custom path with all parameters
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/all",
            parameters={
                "query": "parameters",
            },
            body={
                "body": "parameters",
            },
        )

        assert _req.path == "/test/all"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {"query": "parameters"}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"body":"parameters"}""")

    async def test_custom_post_2(self):
        """
        requestOptions can override default query parameters
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"query":"myQueryParameter"}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {"query": "myQueryParameter"}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_3(self):
        """
        requestOptions merges query parameters with default ones
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"query2":"myQueryParameter"}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters", "query2": "myQueryParameter"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_4(self):
        """
        requestOptions can override default headers
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "headers": loads("""{"x-algolia-api-key":"myApiKey"}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {"query": "parameters"}.items()
        assert _req.headers.items() >= {"x-algolia-api-key": "myApiKey"}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_5(self):
        """
        requestOptions merges headers with default ones
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "headers": loads("""{"x-algolia-api-key":"myApiKey"}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {"query": "parameters"}.items()
        assert _req.headers.items() >= {"x-algolia-api-key": "myApiKey"}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_6(self):
        """
        requestOptions queryParameters accepts booleans
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"isItWorking":true}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters", "isItWorking": "true"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_7(self):
        """
        requestOptions queryParameters accepts integers
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"myParam":2}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters", "myParam": "2"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_8(self):
        """
        requestOptions queryParameters accepts list of string
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"myParam":["b and c","d"]}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters", "myParam": "b%20and%20c%2Cd"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_9(self):
        """
        requestOptions queryParameters accepts list of booleans
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"myParam":[true,true,false]}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters", "myParam": "true%2Ctrue%2Cfalse"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_post_10(self):
        """
        requestOptions queryParameters accepts list of integers
        """
        _req = await self._client.custom_post_with_http_info(
            path="test/requestOptions",
            parameters={
                "query": "parameters",
            },
            body={
                "facet": "filters",
            },
            request_options={
                "query_parameters": loads("""{"myParam":[1,2]}"""),
            },
        )

        assert _req.path == "/test/requestOptions"
        assert _req.verb == "POST"
        assert (
            _req.query_parameters.items()
            == {"query": "parameters", "myParam": "1%2C2"}.items()
        )
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"facet":"filters"}""")

    async def test_custom_put_0(self):
        """
        allow put method for a custom path with minimal parameters
        """
        _req = await self._client.custom_put_with_http_info(
            path="test/minimal",
        )

        assert _req.path == "/test/minimal"
        assert _req.verb == "PUT"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{}""")

    async def test_custom_put_1(self):
        """
        allow put method for a custom path with all parameters
        """
        _req = await self._client.custom_put_with_http_info(
            path="test/all",
            parameters={
                "query": "parameters",
            },
            body={
                "body": "parameters",
            },
        )

        assert _req.path == "/test/all"
        assert _req.verb == "PUT"
        assert _req.query_parameters.items() == {"query": "parameters"}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"body":"parameters"}""")

    async def test_delete_authentication_0(self):
        """
        deleteAuthentication
        """
        _req = await self._client.delete_authentication_with_http_info(
            authentication_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "DELETE"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_delete_destination_0(self):
        """
        deleteDestination
        """
        _req = await self._client.delete_destination_with_http_info(
            destination_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "DELETE"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_delete_source_0(self):
        """
        deleteSource
        """
        _req = await self._client.delete_source_with_http_info(
            source_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "DELETE"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_delete_task_0(self):
        """
        deleteTask
        """
        _req = await self._client.delete_task_with_http_info(
            task_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "DELETE"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_disable_task_0(self):
        """
        disableTask
        """
        _req = await self._client.disable_task_with_http_info(
            task_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f/disable"
        assert _req.verb == "PUT"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()

    async def test_enable_task_0(self):
        """
        enable task e2e
        """
        _req = await self._client.enable_task_with_http_info(
            task_id="76ab4c2a-ce17-496f-b7a6-506dc59ee498",
        )

        assert _req.path == "/1/tasks/76ab4c2a-ce17-496f-b7a6-506dc59ee498/enable"
        assert _req.verb == "PUT"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()

        raw_resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).enable_task_with_http_info(
            task_id="76ab4c2a-ce17-496f-b7a6-506dc59ee498",
        )
        assert raw_resp.status_code == 200

        resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).enable_task(
            task_id="76ab4c2a-ce17-496f-b7a6-506dc59ee498",
        )
        _expected_body = loads("""{"taskID":"76ab4c2a-ce17-496f-b7a6-506dc59ee498"}""")
        assert self._helpers.union(_expected_body, resp) == _expected_body

    async def test_get_authentication_0(self):
        """
        getAuthentication
        """
        _req = await self._client.get_authentication_with_http_info(
            authentication_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_authentications_0(self):
        """
        getAuthentications
        """
        _req = await self._client.get_authentications_with_http_info()

        assert _req.path == "/1/authentications"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_authentications_1(self):
        """
        getAuthentications with query params
        """
        _req = await self._client.get_authentications_with_http_info(
            items_per_page=10,
            page=1,
            type=[
                "basic",
                "algolia",
            ],
            platform=[
                "none",
            ],
            sort="createdAt",
            order="desc",
        )

        assert _req.path == "/1/authentications"
        assert _req.verb == "GET"
        assert (
            _req.query_parameters.items()
            == {
                "itemsPerPage": "10",
                "page": "1",
                "type": "basic%2Calgolia",
                "platform": "none",
                "sort": "createdAt",
                "order": "desc",
            }.items()
        )
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

        raw_resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).get_authentications_with_http_info(
            items_per_page=10,
            page=1,
            type=[
                "basic",
                "algolia",
            ],
            platform=[
                "none",
            ],
            sort="createdAt",
            order="desc",
        )
        assert raw_resp.status_code == 200

        resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).get_authentications(
            items_per_page=10,
            page=1,
            type=[
                "basic",
                "algolia",
            ],
            platform=[
                "none",
            ],
            sort="createdAt",
            order="desc",
        )
        _expected_body = loads(
            """{"pagination":{"page":1,"itemsPerPage":10},"authentications":[{"authenticationID":"b57a7ea5-8592-493b-b75b-6c66d77aee7f","type":"algolia","name":"Auto-generated Authentication for T8JK9S7I7X - 1704732447751","input":{},"createdAt":"2024-01-08T16:47:31Z","updatedAt":"2024-01-08T16:47:31Z"},{},{},{},{},{},{},{}]}"""
        )
        assert self._helpers.union(_expected_body, resp) == _expected_body

    async def test_get_destination_0(self):
        """
        getDestination
        """
        _req = await self._client.get_destination_with_http_info(
            destination_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_destinations_0(self):
        """
        getDestinations
        """
        _req = await self._client.get_destinations_with_http_info()

        assert _req.path == "/1/destinations"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_docker_source_streams_0(self):
        """
        getDockerSourceStreams
        """
        _req = await self._client.get_docker_source_streams_with_http_info(
            source_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f/discover"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_event_0(self):
        """
        getEvent
        """
        _req = await self._client.get_event_with_http_info(
            run_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            event_id="6c02aeb1-775e-418e-870b-1faccd4b2c0c",
        )

        assert (
            _req.path
            == "/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f/events/6c02aeb1-775e-418e-870b-1faccd4b2c0c"
        )
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_events_0(self):
        """
        getEvents
        """
        _req = await self._client.get_events_with_http_info(
            run_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f/events"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_run_0(self):
        """
        getRun
        """
        _req = await self._client.get_run_with_http_info(
            run_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_runs_0(self):
        """
        getRuns
        """
        _req = await self._client.get_runs_with_http_info()

        assert _req.path == "/1/runs"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_source_0(self):
        """
        getSource
        """
        _req = await self._client.get_source_with_http_info(
            source_id="75eeb306-51d3-4e5e-a279-3c92bd8893ac",
        )

        assert _req.path == "/1/sources/75eeb306-51d3-4e5e-a279-3c92bd8893ac"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

        raw_resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).get_source_with_http_info(
            source_id="75eeb306-51d3-4e5e-a279-3c92bd8893ac",
        )
        assert raw_resp.status_code == 200

        resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).get_source(
            source_id="75eeb306-51d3-4e5e-a279-3c92bd8893ac",
        )
        _expected_body = loads(
            """{"sourceID":"75eeb306-51d3-4e5e-a279-3c92bd8893ac","name":"cts_e2e_browse","type":"json","input":{"url":"https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"}}"""
        )
        assert self._helpers.union(_expected_body, resp) == _expected_body

    async def test_get_sources_0(self):
        """
        getSources
        """
        _req = await self._client.get_sources_with_http_info()

        assert _req.path == "/1/sources"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_task_0(self):
        """
        getTask
        """
        _req = await self._client.get_task_with_http_info(
            task_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_get_tasks_0(self):
        """
        getTasks
        """
        _req = await self._client.get_tasks_with_http_info()

        assert _req.path == "/1/tasks"
        assert _req.verb == "GET"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert _req.data is None

    async def test_run_task_0(self):
        """
        runTask
        """
        _req = await self._client.run_task_with_http_info(
            task_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f/run"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()

    async def test_search_authentications_0(self):
        """
        searchAuthentications
        """
        _req = await self._client.search_authentications_with_http_info(
            authentication_search={
                "authenticationIDs": [
                    "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
                    "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
                ],
            },
        )

        assert _req.path == "/1/authentications/search"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"authenticationIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}"""
        )

    async def test_search_destinations_0(self):
        """
        searchDestinations
        """
        _req = await self._client.search_destinations_with_http_info(
            destination_search={
                "destinationIDs": [
                    "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
                    "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
                ],
            },
        )

        assert _req.path == "/1/destinations/search"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"destinationIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}"""
        )

    async def test_search_sources_0(self):
        """
        searchSources
        """
        _req = await self._client.search_sources_with_http_info(
            source_search={
                "sourceIDs": [
                    "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
                    "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
                ],
            },
        )

        assert _req.path == "/1/sources/search"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"sourceIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}"""
        )

    async def test_search_tasks_0(self):
        """
        searchTasks
        """
        _req = await self._client.search_tasks_with_http_info(
            task_search={
                "taskIDs": [
                    "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
                    "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
                    "76ab4c2a-ce17-496f-b7a6-506dc59ee498",
                ],
            },
        )

        assert _req.path == "/1/tasks/search"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads(
            """{"taskIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a","76ab4c2a-ce17-496f-b7a6-506dc59ee498"]}"""
        )

        raw_resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).search_tasks_with_http_info(
            task_search={
                "taskIDs": [
                    "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
                    "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
                    "76ab4c2a-ce17-496f-b7a6-506dc59ee498",
                ],
            },
        )
        assert raw_resp.status_code == 200

        resp = await IngestionClient(
            self._e2e_app_id, self._e2e_api_key, "us"
        ).search_tasks(
            task_search={
                "taskIDs": [
                    "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
                    "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
                    "76ab4c2a-ce17-496f-b7a6-506dc59ee498",
                ],
            },
        )
        _expected_body = loads(
            """[{"taskID":"76ab4c2a-ce17-496f-b7a6-506dc59ee498","sourceID":"75eeb306-51d3-4e5e-a279-3c92bd8893ac","destinationID":"506d79fa-e29d-4bcf-907c-6b6a41172153","trigger":{"type":"onDemand"},"enabled":true,"failureThreshold":0,"action":"replace","createdAt":"2024-01-08T16:47:41Z"}]"""
        )
        assert self._helpers.union(_expected_body, resp) == _expected_body

    async def test_trigger_docker_source_discover_0(self):
        """
        triggerDockerSourceDiscover
        """
        _req = await self._client.trigger_docker_source_discover_with_http_info(
            source_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
        )

        assert _req.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f/discover"
        assert _req.verb == "POST"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()

    async def test_update_authentication_0(self):
        """
        updateAuthentication
        """
        _req = await self._client.update_authentication_with_http_info(
            authentication_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            authentication_update={
                "name": "newName",
            },
        )

        assert _req.path == "/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "PATCH"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"name":"newName"}""")

    async def test_update_destination_0(self):
        """
        updateDestination
        """
        _req = await self._client.update_destination_with_http_info(
            destination_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            destination_update={
                "name": "newName",
            },
        )

        assert _req.path == "/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "PATCH"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"name":"newName"}""")

    async def test_update_source_0(self):
        """
        updateSource
        """
        _req = await self._client.update_source_with_http_info(
            source_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            source_update={
                "name": "newName",
            },
        )

        assert _req.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "PATCH"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"name":"newName"}""")

    async def test_update_task_0(self):
        """
        updateTask
        """
        _req = await self._client.update_task_with_http_info(
            task_id="6c02aeb1-775e-418e-870b-1faccd4b2c0f",
            task_update={
                "enabled": False,
            },
        )

        assert _req.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f"
        assert _req.verb == "PATCH"
        assert _req.query_parameters.items() == {}.items()
        assert _req.headers.items() >= {}.items()
        assert loads(_req.data) == loads("""{"enabled":false}""")
