package com.algolia.codegen.cts.tests;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.TreeNode;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import java.io.IOException;
import java.util.Map;

public class Request {

  public String testName;

  public Map<String, Object> parameters;
  public RequestOptions requestOptions;
  public RequestProp request;
  public ResponseProp response;

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Request {\n");
    sb.append("    testName: ").append(testName).append("\n");
    sb.append("    parameters: ").append(parameters).append("\n");
    sb.append("    requestOptions: ").append(requestOptions).append("\n");
    sb.append("    request: ").append(request).append("\n");
    sb.append("    response: ").append(response).append("\n");
    sb.append("}");
    return sb.toString();
  }
}

class RequestOptions {

  public Map<String, Object> queryParameters;
  public Map<String, String> headers;

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RequestOptions {\n");
    sb.append("    queryParameters: ").append(queryParameters).append("\n");
    sb.append("    headers: ").append(headers).append("\n");
    sb.append("}");
    return sb.toString();
  }
}

class RequestProp {

  public String path;
  public String method;

  @JsonDeserialize(using = RawDeserializer.class)
  public String body;

  @JsonDeserialize(using = RawDeserializer.class)
  public String queryParameters;

  @JsonDeserialize(using = RawDeserializer.class)
  public String headers;

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RequestProp {\n");
    sb.append("    path: ").append(path).append("\n");
    sb.append("    method: ").append(method).append("\n");
    sb.append("    body: ").append(body).append("\n");
    sb.append("    queryParameters: ").append(queryParameters).append("\n");
    sb.append("    headers: ").append(headers).append("\n");
    sb.append("}");
    return sb.toString();
  }
}

class ResponseProp {

  public int statusCode;

  @JsonDeserialize(using = RawDeserializer.class)
  public String body;

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ResponseProp {\n");
    sb.append("    statusCode: ").append(statusCode).append("\n");
    sb.append("    body: ").append(body).append("\n");
    sb.append("}");
    return sb.toString();
  }
}

// Output json to raw string with quotes
class RawDeserializer extends JsonDeserializer<String> {

  @Override
  public String deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
    TreeNode tree = jp.getCodec().readTree(jp);
    return tree.toString();
  }
}
