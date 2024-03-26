/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.api

import com.algolia.client.configuration.*
import com.algolia.client.exception.*
import com.algolia.client.extensions.internal.*
import com.algolia.client.model.monitoring.*
import com.algolia.client.transport.*
import com.algolia.client.transport.internal.*
import kotlinx.serialization.json.*

public class MonitoringClient(
  override val appId: String,
  override val apiKey: String,
  override val options: ClientOptions = ClientOptions(),
) : ApiClient {

  init {
    require(appId.isNotBlank()) { "`appId` is missing." }
    require(apiKey.isNotBlank()) { "`apiKey` is missing." }
  }

  override val requester: Requester = requesterOf(clientName = "Monitoring", appId = appId, apiKey = apiKey, options = options) {
    listOf(Host("status.algolia.com"))
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   * @param path Path of the endpoint, anything after \"/1\" must be specified.
   * @param parameters Query parameters to apply to the current query.
   * @param requestOptions additional request configuration.
   */
  public suspend fun customDelete(path: String, parameters: Map<kotlin.String, Any>? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `customDelete`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = "/{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   * @param path Path of the endpoint, anything after \"/1\" must be specified.
   * @param parameters Query parameters to apply to the current query.
   * @param requestOptions additional request configuration.
   */
  public suspend fun customGet(path: String, parameters: Map<kotlin.String, Any>? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `customGet`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = "/{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   * @param path Path of the endpoint, anything after \"/1\" must be specified.
   * @param parameters Query parameters to apply to the current query.
   * @param body Parameters to send with the custom request.
   * @param requestOptions additional request configuration.
   */
  public suspend fun customPost(path: String, parameters: Map<kotlin.String, Any>? = null, body: JsonObject? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `customPost`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = "/{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
      body = body,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   * @param path Path of the endpoint, anything after \"/1\" must be specified.
   * @param parameters Query parameters to apply to the current query.
   * @param body Parameters to send with the custom request.
   * @param requestOptions additional request configuration.
   */
  public suspend fun customPut(path: String, parameters: Map<kotlin.String, Any>? = null, body: JsonObject? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `customPut`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PUT,
      path = "/{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
      body = body,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves known incidents for the selected clusters.
   * @param clusters Subset of clusters, separated by comma.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getClusterIncidents(clusters: String, requestOptions: RequestOptions? = null): IncidentsResponse {
    require(clusters.isNotBlank()) { "Parameter `clusters` is required when calling `getClusterIncidents`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "incidents", "$clusters"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves the status of selected clusters.
   * @param clusters Subset of clusters, separated by comma.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getClusterStatus(clusters: String, requestOptions: RequestOptions? = null): StatusResponse {
    require(clusters.isNotBlank()) { "Parameter `clusters` is required when calling `getClusterStatus`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "status", "$clusters"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves known incidents for all clusters.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getIncidents(requestOptions: RequestOptions? = null): IncidentsResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "incidents"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves average times for indexing operations for selected clusters.
   * @param clusters Subset of clusters, separated by comma.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getIndexingTime(clusters: String, requestOptions: RequestOptions? = null): IndexingTimeResponse {
    require(clusters.isNotBlank()) { "Parameter `clusters` is required when calling `getIndexingTime`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "indexing", "$clusters"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves the average latency for search requests for selected clusters.
   * @param clusters Subset of clusters, separated by comma.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getLatency(clusters: String, requestOptions: RequestOptions? = null): LatencyResponse {
    require(clusters.isNotBlank()) { "Parameter `clusters` is required when calling `getLatency`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "latency", "$clusters"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves metrics related to your Algolia infrastructure, aggregated over a selected time window.  Access to this API is available as part of the [Premium or Elevate plans](https://www.algolia.com/pricing). You must authenticate requests with the `x-algolia-application-id` and `x-algolia-api-key` headers (using the Monitoring API key).
   * @param metric Metric to report.  For more information about the individual metrics, see the description of the API response. To include all metrics, use `*`.
   * @param period Period over which to aggregate the metrics:  - `minute`. Aggregate the last minute. 1 data point per 10 seconds. - `hour`. Aggregate the last hour. 1 data point per minute. - `day`. Aggregate the last day. 1 data point per 10 minutes. - `week`. Aggregate the last week. 1 data point per hour. - `month`. Aggregate the last month. 1 data point per day.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getMetrics(metric: Metric, period: Period, requestOptions: RequestOptions? = null): InfrastructureResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "infrastructure", "$metric", "period", "$period"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Test whether clusters are reachable or not.
   * @param clusters Subset of clusters, separated by comma.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getReachability(clusters: String, requestOptions: RequestOptions? = null): Map<kotlin.String, Map<kotlin.String, Boolean>> {
    require(clusters.isNotBlank()) { "Parameter `clusters` is required when calling `getReachability`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "reachability", "$clusters", "probes"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves the servers that belong to clusters.  The response depends on whether you authenticate your API request:  - With authentication, the response lists the servers assigned to your Algolia application's cluster.  - Without authentication, the response lists the servers for all Algolia clusters.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getServers(requestOptions: RequestOptions? = null): InventoryResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "inventory", "servers"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieves the status of all Algolia clusters and instances.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getStatus(requestOptions: RequestOptions? = null): StatusResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "status"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }
}
