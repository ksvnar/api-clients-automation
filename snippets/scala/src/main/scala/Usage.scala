// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package algoliasearch.methods.snippets

import scala.concurrent.duration.Duration

import algoliasearch.api.UsageClient
import algoliasearch.usage.*

import org.json4s.*
import org.json4s.native.JsonParser.*
import scala.concurrent.{Await, ExecutionContextExecutor}

class SnippetUsageClient {
  implicit val ec: ExecutionContextExecutor = scala.concurrent.ExecutionContext.global
  implicit val formats: Formats = org.json4s.DefaultFormats

  /** Snippet for the customDelete method.
    *
    * allow del method for a custom path with minimal parameters
    */
  def snippetForUsageClientCustomDelete(): Unit = {
    // >SEPARATOR customDelete
    // Initialize the client
    val client = UsageClient(appId = "YOUR_APP_ID", apiKey = "YOUR_API_KEY")

    // Call the API
    val res = client.customDelete[JObject](
      path = "test/minimal"
    )

    // Use the response
    val value = Await.result(res, Duration(100, "sec"))
    // SEPARATOR<
  }

  /** Snippet for the customGet method.
    *
    * allow get method for a custom path with minimal parameters
    */
  def snippetForUsageClientCustomGet(): Unit = {
    // >SEPARATOR customGet
    // Initialize the client
    val client = UsageClient(appId = "YOUR_APP_ID", apiKey = "YOUR_API_KEY")

    // Call the API
    val res = client.customGet[JObject](
      path = "test/minimal"
    )

    // Use the response
    val value = Await.result(res, Duration(100, "sec"))
    // SEPARATOR<
  }

  /** Snippet for the customPost method.
    *
    * allow post method for a custom path with minimal parameters
    */
  def snippetForUsageClientCustomPost(): Unit = {
    // >SEPARATOR customPost
    // Initialize the client
    val client = UsageClient(appId = "YOUR_APP_ID", apiKey = "YOUR_API_KEY")

    // Call the API
    val res = client.customPost[JObject](
      path = "test/minimal"
    )

    // Use the response
    val value = Await.result(res, Duration(100, "sec"))
    // SEPARATOR<
  }

  /** Snippet for the customPut method.
    *
    * allow put method for a custom path with minimal parameters
    */
  def snippetForUsageClientCustomPut(): Unit = {
    // >SEPARATOR customPut
    // Initialize the client
    val client = UsageClient(appId = "YOUR_APP_ID", apiKey = "YOUR_API_KEY")

    // Call the API
    val res = client.customPut[JObject](
      path = "test/minimal"
    )

    // Use the response
    val value = Await.result(res, Duration(100, "sec"))
    // SEPARATOR<
  }

  /** Snippet for the getIndexUsage method.
    *
    * getIndexUsage with minimal parameters
    */
  def snippetForUsageClientGetIndexUsage(): Unit = {
    // >SEPARATOR getIndexUsage
    // Initialize the client
    val client = UsageClient(appId = "YOUR_APP_ID", apiKey = "YOUR_API_KEY")

    // Call the API
    val res = client.getIndexUsage(
      statistic = Statistic.withName("queries_operations"),
      indexName = "myIndexName",
      startDate = "2024-04-03T12:46:43Z",
      endDate = "2024-04-05T12:46:43Z"
    )

    // Use the response
    val value = Await.result(res, Duration(100, "sec"))
    // SEPARATOR<
  }

  /** Snippet for the getUsage method.
    *
    * getUsage with minimal parameters
    */
  def snippetForUsageClientGetUsage(): Unit = {
    // >SEPARATOR getUsage
    // Initialize the client
    val client = UsageClient(appId = "YOUR_APP_ID", apiKey = "YOUR_API_KEY")

    // Call the API
    val res = client.getUsage(
      statistic = Statistic.withName("queries_operations"),
      startDate = "2024-04-03T12:46:43Z",
      endDate = "2024-04-05T12:46:43Z"
    )

    // Use the response
    val value = Await.result(res, Duration(100, "sec"))
    // SEPARATOR<
  }

}
