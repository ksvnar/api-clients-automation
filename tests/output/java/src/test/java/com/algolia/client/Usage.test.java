package com.algolia.client;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

import com.algolia.EchoInterceptor;
import com.algolia.EchoResponse;
import com.algolia.api.UsageClient;
import com.algolia.config.*;
import com.algolia.model.usage.*;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.json.JsonMapper;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class UsageClientClientTests {

  private EchoInterceptor echo = new EchoInterceptor();
  private ObjectMapper json;

  @BeforeAll
  void init() {
    this.json = JsonMapper.builder().configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false).build();
  }

  UsageClient createClient() {
    return new UsageClient("appId", "apiKey", withEchoRequester());
  }

  private ClientOptions withEchoRequester() {
    return ClientOptions.builder().setRequesterConfig(requester -> requester.addInterceptor(echo)).build();
  }

  private ClientOptions withCustomHosts(List<Host> hosts, boolean gzipEncoding) {
    return ClientOptions.builder().setHosts(hosts).setCompressionType(gzipEncoding ? CompressionType.GZIP : CompressionType.NONE).build();
  }

  @Test
  @DisplayName("calls api with correct read host")
  void apiTest0() {
    UsageClient client = new UsageClient("test-app-id", "test-api-key", withEchoRequester());
    client.customGet("test");
    EchoResponse result = echo.getLastResponse();

    assertEquals("test-app-id-dsn.algolia.net", result.host);
  }

  @Test
  @DisplayName("calls api with correct write host")
  void apiTest1() {
    UsageClient client = new UsageClient("test-app-id", "test-api-key", withEchoRequester());
    client.customPost("test");
    EchoResponse result = echo.getLastResponse();

    assertEquals("test-app-id.algolia.net", result.host);
  }

  @Test
  @DisplayName("calls api with correct user agent")
  void commonApiTest0() {
    UsageClient client = createClient();

    client.customPost("1/test");
    EchoResponse result = echo.getLastResponse();
    {
      String regexp =
        "^Algolia for Java \\(\\d+\\.\\d+\\.\\d+(-?.*)?\\)(; [a-zA-Z. ]+" +
        " (\\(\\d+((\\.\\d+)?\\.\\d+)?(-?.*)?\\))?)*(; Usage" +
        " (\\(\\d+\\.\\d+\\.\\d+(-?.*)?\\)))(; [a-zA-Z. ]+" +
        " (\\(\\d+((\\.\\d+)?\\.\\d+)?(-?.*)?\\))?)*$";
      assertTrue(
        result.headers.get("user-agent").matches(regexp),
        "Expected " + result.headers.get("user-agent") + " to match the following regex: " + regexp
      );
    }
  }

  @Test
  @DisplayName("calls api with default read timeouts")
  void commonApiTest1() {
    UsageClient client = createClient();

    client.customGet("1/test");
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(5000, result.responseTimeout);
  }

  @Test
  @DisplayName("calls api with default write timeouts")
  void commonApiTest2() {
    UsageClient client = createClient();

    client.customPost("1/test");
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(30000, result.responseTimeout);
  }

  @Test
  @DisplayName("client throws with invalid parameters")
  void parametersTest0() {
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          UsageClient client = new UsageClient("", "", withEchoRequester());
        }
      );
      assertEquals("`appId` is missing.", exception.getMessage());
    }
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          UsageClient client = new UsageClient("", "my-api-key", withEchoRequester());
        }
      );
      assertEquals("`appId` is missing.", exception.getMessage());
    }
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          UsageClient client = new UsageClient("my-app-id", "", withEchoRequester());
        }
      );
      assertEquals("`apiKey` is missing.", exception.getMessage());
    }
  }
}
