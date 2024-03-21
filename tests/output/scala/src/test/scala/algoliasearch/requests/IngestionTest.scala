package algoliasearch.requests

import algoliasearch.EchoInterceptor
import algoliasearch.api.IngestionClient
import algoliasearch.config.*
import algoliasearch.ingestion.*
import org.json4s.*
import org.json4s.native.JsonParser.*
import org.scalatest.funsuite.AnyFunSuite
import io.github.cdimascio.dotenv.Dotenv
import org.skyscreamer.jsonassert.JSONCompare.compareJSON
import org.skyscreamer.jsonassert.JSONCompareMode
import org.json4s.native.Serialization
import org.json4s.native.Serialization.write

import scala.concurrent.duration.Duration
import scala.concurrent.{Await, ExecutionContextExecutor}

class IngestionTest extends AnyFunSuite {
  implicit val ec: ExecutionContextExecutor = scala.concurrent.ExecutionContext.global
  implicit val formats: Formats = org.json4s.DefaultFormats

  def testClient(): (IngestionClient, EchoInterceptor) = {
    val echo = EchoInterceptor()
    (
      IngestionClient(
        appId = "appId",
        apiKey = "apiKey",
        region = "us",
        clientOptions = ClientOptions
          .builder()
          .withRequesterConfig(requester => requester.withInterceptor(echo))
          .build()
      ),
      echo
    )
  }

  def testE2EClient(): IngestionClient = {
    val region = "us"
    if (System.getenv("CI") == "true") {
      IngestionClient(
        appId = System.getenv("ALGOLIA_APPLICATION_ID"),
        apiKey = System.getenv("ALGOLIA_ADMIN_KEY"),
        region = region
      )
    } else {
      val dotenv = Dotenv.configure.directory("../../").load
      IngestionClient(
        appId = dotenv.get("ALGOLIA_APPLICATION_ID"),
        apiKey = dotenv.get("ALGOLIA_ADMIN_KEY"),
        region = region
      )
    }
  }

  test("createAuthenticationOAuth") {
    val (client, echo) = testClient()
    val future = client.createAuthentication(
      authenticationCreate = AuthenticationCreate(
        `type` = AuthenticationType.withName("oauth"),
        name = "authName",
        input = AuthOAuth(
          url = "http://test.oauth",
          client_id = "myID",
          client_secret = "mySecret"
        )
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"type":"oauth","name":"authName","input":{"url":"http://test.oauth","client_id":"myID","client_secret":"mySecret"}}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("createAuthenticationAlgolia") {
    val (client, echo) = testClient()
    val future = client.createAuthentication(
      authenticationCreate = AuthenticationCreate(
        `type` = AuthenticationType.withName("algolia"),
        name = "authName",
        input = AuthAlgolia(
          appID = "myappID",
          apiKey = "randomApiKey"
        )
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications")
    assert(res.method == "POST")
    val expectedBody =
      parse("""{"type":"algolia","name":"authName","input":{"appID":"myappID","apiKey":"randomApiKey"}}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("createDestination") {
    val (client, echo) = testClient()
    val future = client.createDestination(
      destinationCreate = DestinationCreate(
        `type` = DestinationType.withName("search"),
        name = "destinationName",
        input = DestinationIndexPrefix(
          indexPrefix = "prefix_"
        ),
        authenticationID = Some("6c02aeb1-775e-418e-870b-1faccd4b2c0f")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/destinations")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"type":"search","name":"destinationName","input":{"indexPrefix":"prefix_"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("createSource") {
    val (client, echo) = testClient()
    val future = client.createSource(
      sourceCreate = SourceCreate(
        `type` = SourceType.withName("commercetools"),
        name = "sourceName",
        input = SourceCommercetools(
          storeKeys = Some(Seq("myStore")),
          locales = Some(Seq("de")),
          url = "http://commercetools.com",
          projectKey = "keyID"
        ),
        authenticationID = Some("6c02aeb1-775e-418e-870b-1faccd4b2c0f")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"type":"commercetools","name":"sourceName","input":{"storeKeys":["myStore"],"locales":["de"],"url":"http://commercetools.com","projectKey":"keyID"},"authenticationID":"6c02aeb1-775e-418e-870b-1faccd4b2c0f"}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("createTaskOnDemand") {
    val (client, echo) = testClient()
    val future = client.createTask(
      taskCreate = TaskCreate(
        sourceID = "search",
        destinationID = "destinationName",
        trigger = OnDemandTriggerInput(
          `type` = OnDemandTriggerType.withName("onDemand")
        ),
        action = ActionType.withName("replace")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"onDemand"},"action":"replace"}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("createTaskSchedule") {
    val (client, echo) = testClient()
    val future = client.createTask(
      taskCreate = TaskCreate(
        sourceID = "search",
        destinationID = "destinationName",
        trigger = ScheduleTriggerInput(
          `type` = ScheduleTriggerType.withName("schedule"),
          cron = "* * * * *"
        ),
        action = ActionType.withName("replace")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"schedule","cron":"* * * * *"},"action":"replace"}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("createTaskSubscription") {
    val (client, echo) = testClient()
    val future = client.createTask(
      taskCreate = TaskCreate(
        sourceID = "search",
        destinationID = "destinationName",
        trigger = OnDemandTriggerInput(
          `type` = OnDemandTriggerType.withName("onDemand")
        ),
        action = ActionType.withName("replace")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"sourceID":"search","destinationID":"destinationName","trigger":{"type":"onDemand"},"action":"replace"}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("allow del method for a custom path with minimal parameters") {
    val (client, echo) = testClient()
    val future = client.customDelete[JObject](
      path = "test/minimal"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/minimal")
    assert(res.method == "DELETE")
    assert(res.body.isEmpty)
  }

  test("allow del method for a custom path with all parameters") {
    val (client, echo) = testClient()
    val future = client.customDelete[JObject](
      path = "test/all",
      parameters = Some(Map("query" -> "parameters"))
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/all")
    assert(res.method == "DELETE")
    assert(res.body.isEmpty)
    val expectedQuery = parse("""{"query":"parameters"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("allow get method for a custom path with minimal parameters") {
    val (client, echo) = testClient()
    val future = client.customGet[JObject](
      path = "test/minimal"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/minimal")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("allow get method for a custom path with all parameters") {
    val (client, echo) = testClient()
    val future = client.customGet[JObject](
      path = "test/all",
      parameters = Some(Map("query" -> "parameters with space"))
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/all")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
    val expectedQuery = parse("""{"query":"parameters%20with%20space"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions should be escaped too") {
    val (client, echo) = testClient()
    val future = client.customGet[JObject](
      path = "test/all",
      parameters = Some(Map("query" -> "to be overriden")),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("query", "parameters with space")
          .withQueryParameter("and an array", Seq("array", "with spaces"))
          .withHeader("x-header-1", "spaces are left alone")
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/all")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
    val expectedQuery = parse("""{"query":"parameters%20with%20space","and%20an%20array":"array%2Cwith%20spaces"}""")
      .asInstanceOf[JObject]
      .obj
      .toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
    val expectedHeaders = parse("""{"x-header-1":"spaces are left alone"}""").asInstanceOf[JObject].obj.toMap
    val actualHeaders = res.headers
    for ((k, v) <- expectedHeaders) {
      assert(actualHeaders.contains(k))
      assert(actualHeaders(k) == v.asInstanceOf[JString].s)
    }
  }

  test("allow post method for a custom path with minimal parameters") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/minimal"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/minimal")
    assert(res.method == "POST")
    val expectedBody = parse("""{}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("allow post method for a custom path with all parameters") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/all",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("body", JString("parameters")))))
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/all")
    assert(res.method == "POST")
    val expectedBody = parse("""{"body":"parameters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions can override default query parameters") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("query", "myQueryParameter")
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"myQueryParameter"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions merges query parameters with default ones") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("query2", "myQueryParameter")
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters","query2":"myQueryParameter"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions can override default headers") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withHeader("x-algolia-api-key", "myApiKey")
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
    val expectedHeaders = parse("""{"x-algolia-api-key":"myApiKey"}""").asInstanceOf[JObject].obj.toMap
    val actualHeaders = res.headers
    for ((k, v) <- expectedHeaders) {
      assert(actualHeaders.contains(k))
      assert(actualHeaders(k) == v.asInstanceOf[JString].s)
    }
  }

  test("requestOptions merges headers with default ones") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withHeader("x-algolia-api-key", "myApiKey")
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
    val expectedHeaders = parse("""{"x-algolia-api-key":"myApiKey"}""").asInstanceOf[JObject].obj.toMap
    val actualHeaders = res.headers
    for ((k, v) <- expectedHeaders) {
      assert(actualHeaders.contains(k))
      assert(actualHeaders(k) == v.asInstanceOf[JString].s)
    }
  }

  test("requestOptions queryParameters accepts booleans") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("isItWorking", true)
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters","isItWorking":"true"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions queryParameters accepts integers") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("myParam", 2)
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters","myParam":"2"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions queryParameters accepts list of string") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("myParam", Seq("b and c", "d"))
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters","myParam":"b%20and%20c%2Cd"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions queryParameters accepts list of booleans") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("myParam", Seq(true, true, false))
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery =
      parse("""{"query":"parameters","myParam":"true%2Ctrue%2Cfalse"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("requestOptions queryParameters accepts list of integers") {
    val (client, echo) = testClient()
    val future = client.customPost[JObject](
      path = "test/requestOptions",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("facet", JString("filters"))))),
      requestOptions = Some(
        RequestOptions
          .builder()
          .withQueryParameter("myParam", Seq(1, 2))
          .build()
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/requestOptions")
    assert(res.method == "POST")
    val expectedBody = parse("""{"facet":"filters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters","myParam":"1%2C2"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("allow put method for a custom path with minimal parameters") {
    val (client, echo) = testClient()
    val future = client.customPut[JObject](
      path = "test/minimal"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/minimal")
    assert(res.method == "PUT")
    val expectedBody = parse("""{}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("allow put method for a custom path with all parameters") {
    val (client, echo) = testClient()
    val future = client.customPut[JObject](
      path = "test/all",
      parameters = Some(Map("query" -> "parameters")),
      body = Some(JObject(List(JField("body", JString("parameters")))))
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/test/all")
    assert(res.method == "PUT")
    val expectedBody = parse("""{"body":"parameters"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val expectedQuery = parse("""{"query":"parameters"}""").asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
  }

  test("deleteAuthentication") {
    val (client, echo) = testClient()
    val future = client.deleteAuthentication(
      authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "DELETE")
    assert(res.body.isEmpty)
  }

  test("deleteDestination") {
    val (client, echo) = testClient()
    val future = client.deleteDestination(
      destinationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "DELETE")
    assert(res.body.isEmpty)
  }

  test("deleteSource") {
    val (client, echo) = testClient()
    val future = client.deleteSource(
      sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "DELETE")
    assert(res.body.isEmpty)
  }

  test("deleteTask") {
    val (client, echo) = testClient()
    val future = client.deleteTask(
      taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "DELETE")
    assert(res.body.isEmpty)
  }

  test("disableTask") {
    val (client, echo) = testClient()
    val future = client.disableTask(
      taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f/disable")
    assert(res.method == "PUT")
    assert(res.body.contains("{}"))
  }

  test("enable task e2e") {
    val (client, echo) = testClient()
    val future = client.enableTask(
      taskID = "76ab4c2a-ce17-496f-b7a6-506dc59ee498"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/76ab4c2a-ce17-496f-b7a6-506dc59ee498/enable")
    assert(res.method == "PUT")
    assert(res.body.contains("{}"))
    val e2eClient = testE2EClient()
    val e2eFuture = e2eClient.enableTask(
      taskID = "76ab4c2a-ce17-496f-b7a6-506dc59ee498"
    )

    val response = Await.result(e2eFuture, Duration.Inf)
    compareJSON("""{"taskID":"76ab4c2a-ce17-496f-b7a6-506dc59ee498"}""", write(response), JSONCompareMode.LENIENT)
  }

  test("getAuthentication") {
    val (client, echo) = testClient()
    val future = client.getAuthentication(
      authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getAuthentications") {
    val (client, echo) = testClient()
    val future = client.getAuthentications(
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getAuthentications with query params") {
    val (client, echo) = testClient()
    val future = client.getAuthentications(
      itemsPerPage = Some(10),
      page = Some(1),
      `type` = Some(Seq(AuthenticationType.withName("basic"), AuthenticationType.withName("algolia"))),
      platform = Some(Seq(PlatformNone.withName("none"))),
      sort = Some(AuthenticationSortKeys.withName("createdAt")),
      order = Some(OrderKeys.withName("desc"))
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
    val expectedQuery = parse(
      """{"itemsPerPage":"10","page":"1","type":"basic%2Calgolia","platform":"none","sort":"createdAt","order":"desc"}"""
    ).asInstanceOf[JObject].obj.toMap
    val actualQuery = res.queryParameters
    assert(actualQuery.size == expectedQuery.size)
    for ((k, v) <- actualQuery) {
      assert(expectedQuery.contains(k))
      assert(expectedQuery(k).values == v)
    }
    val e2eClient = testE2EClient()
    val e2eFuture = e2eClient.getAuthentications(
      itemsPerPage = Some(10),
      page = Some(1),
      `type` = Some(Seq(AuthenticationType.withName("basic"), AuthenticationType.withName("algolia"))),
      platform = Some(Seq(PlatformNone.withName("none"))),
      sort = Some(AuthenticationSortKeys.withName("createdAt")),
      order = Some(OrderKeys.withName("desc"))
    )

    val response = Await.result(e2eFuture, Duration.Inf)
    compareJSON(
      """{"pagination":{"page":1,"itemsPerPage":10},"authentications":[{"authenticationID":"b57a7ea5-8592-493b-b75b-6c66d77aee7f","type":"algolia","name":"Auto-generated Authentication for T8JK9S7I7X - 1704732447751","input":{},"createdAt":"2024-01-08T16:47:31Z","updatedAt":"2024-01-08T16:47:31Z"},{},{},{},{},{},{},{}]}""",
      write(response),
      JSONCompareMode.LENIENT
    )
  }

  test("getDestination") {
    val (client, echo) = testClient()
    val future = client.getDestination(
      destinationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getDestinations") {
    val (client, echo) = testClient()
    val future = client.getDestinations(
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/destinations")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getDockerSourceStreams") {
    val (client, echo) = testClient()
    val future = client.getDockerSourceStreams(
      sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f/discover")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getEvent") {
    val (client, echo) = testClient()
    val future = client.getEvent(
      runID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
      eventID = "6c02aeb1-775e-418e-870b-1faccd4b2c0c"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f/events/6c02aeb1-775e-418e-870b-1faccd4b2c0c")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getEvents") {
    val (client, echo) = testClient()
    val future = client.getEvents(
      runID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f/events")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getRun") {
    val (client, echo) = testClient()
    val future = client.getRun(
      runID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/runs/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getRuns") {
    val (client, echo) = testClient()
    val future = client.getRuns(
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/runs")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getSource") {
    val (client, echo) = testClient()
    val future = client.getSource(
      sourceID = "75eeb306-51d3-4e5e-a279-3c92bd8893ac"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources/75eeb306-51d3-4e5e-a279-3c92bd8893ac")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
    val e2eClient = testE2EClient()
    val e2eFuture = e2eClient.getSource(
      sourceID = "75eeb306-51d3-4e5e-a279-3c92bd8893ac"
    )

    val response = Await.result(e2eFuture, Duration.Inf)
    compareJSON(
      """{"sourceID":"75eeb306-51d3-4e5e-a279-3c92bd8893ac","name":"cts_e2e_browse","type":"json","input":{"url":"https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"}}""",
      write(response),
      JSONCompareMode.LENIENT
    )
  }

  test("getSources") {
    val (client, echo) = testClient()
    val future = client.getSources(
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getTask") {
    val (client, echo) = testClient()
    val future = client.getTask(
      taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("getTasks") {
    val (client, echo) = testClient()
    val future = client.getTasks(
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks")
    assert(res.method == "GET")
    assert(res.body.isEmpty)
  }

  test("runTask") {
    val (client, echo) = testClient()
    val future = client.runTask(
      taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f/run")
    assert(res.method == "POST")
    assert(res.body.contains("{}"))
  }

  test("searchAuthentications") {
    val (client, echo) = testClient()
    val future = client.searchAuthentications(
      authenticationSearch = AuthenticationSearch(
        authenticationIDs = Seq("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications/search")
    assert(res.method == "POST")
    val expectedBody =
      parse("""{"authenticationIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("searchDestinations") {
    val (client, echo) = testClient()
    val future = client.searchDestinations(
      destinationSearch = DestinationSearch(
        destinationIDs = Seq("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/destinations/search")
    assert(res.method == "POST")
    val expectedBody =
      parse("""{"destinationIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("searchSources") {
    val (client, echo) = testClient()
    val future = client.searchSources(
      sourceSearch = SourceSearch(
        sourceIDs = Seq("6c02aeb1-775e-418e-870b-1faccd4b2c0f", "947ac9c4-7e58-4c87-b1e7-14a68e99699a")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources/search")
    assert(res.method == "POST")
    val expectedBody =
      parse("""{"sourceIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a"]}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("searchTasks") {
    val (client, echo) = testClient()
    val future = client.searchTasks(
      taskSearch = TaskSearch(
        taskIDs = Seq(
          "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
          "76ab4c2a-ce17-496f-b7a6-506dc59ee498"
        )
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/search")
    assert(res.method == "POST")
    val expectedBody = parse(
      """{"taskIDs":["6c02aeb1-775e-418e-870b-1faccd4b2c0f","947ac9c4-7e58-4c87-b1e7-14a68e99699a","76ab4c2a-ce17-496f-b7a6-506dc59ee498"]}"""
    )
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
    val e2eClient = testE2EClient()
    val e2eFuture = e2eClient.searchTasks(
      taskSearch = TaskSearch(
        taskIDs = Seq(
          "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
          "947ac9c4-7e58-4c87-b1e7-14a68e99699a",
          "76ab4c2a-ce17-496f-b7a6-506dc59ee498"
        )
      )
    )

    val response = Await.result(e2eFuture, Duration.Inf)
    compareJSON(
      """[{"taskID":"76ab4c2a-ce17-496f-b7a6-506dc59ee498","sourceID":"75eeb306-51d3-4e5e-a279-3c92bd8893ac","destinationID":"506d79fa-e29d-4bcf-907c-6b6a41172153","trigger":{"type":"onDemand"},"enabled":true,"failureThreshold":0,"action":"replace","createdAt":"2024-01-08T16:47:41Z"}]""",
      write(response),
      JSONCompareMode.LENIENT
    )
  }

  test("triggerDockerSourceDiscover") {
    val (client, echo) = testClient()
    val future = client.triggerDockerSourceDiscover(
      sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f"
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f/discover")
    assert(res.method == "POST")
    assert(res.body.contains("{}"))
  }

  test("updateAuthentication") {
    val (client, echo) = testClient()
    val future = client.updateAuthentication(
      authenticationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
      authenticationUpdate = AuthenticationUpdate(
        name = Some("newName")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/authentications/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "PATCH")
    val expectedBody = parse("""{"name":"newName"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("updateDestination") {
    val (client, echo) = testClient()
    val future = client.updateDestination(
      destinationID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
      destinationUpdate = DestinationUpdate(
        name = Some("newName")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/destinations/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "PATCH")
    val expectedBody = parse("""{"name":"newName"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("updateSource") {
    val (client, echo) = testClient()
    val future = client.updateSource(
      sourceID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
      sourceUpdate = SourceUpdate(
        name = Some("newName")
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/sources/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "PATCH")
    val expectedBody = parse("""{"name":"newName"}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

  test("updateTask") {
    val (client, echo) = testClient()
    val future = client.updateTask(
      taskID = "6c02aeb1-775e-418e-870b-1faccd4b2c0f",
      taskUpdate = TaskUpdate(
        enabled = Some(false)
      )
    )

    Await.ready(future, Duration.Inf)
    val res = echo.lastResponse.get

    assert(res.path == "/1/tasks/6c02aeb1-775e-418e-870b-1faccd4b2c0f")
    assert(res.method == "PATCH")
    val expectedBody = parse("""{"enabled":false}""")
    val actualBody = parse(res.body.get)
    assert(actualBody == expectedBody)
  }

}
