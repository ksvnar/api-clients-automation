/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.api

import com.algolia.client.configuration.*
import com.algolia.client.exception.*
import com.algolia.client.extensions.internal.*
import com.algolia.client.model.ingestion.*
import com.algolia.client.transport.*
import com.algolia.client.transport.internal.*
import kotlinx.serialization.json.*

public class IngestionClient(
  override val appId: String,
  override val apiKey: String,
  public val region: String,
  override val options: ClientOptions = ClientOptions(),
) : ApiClient {

  init {
    require(appId.isNotBlank()) { "`appId` is missing." }
    require(apiKey.isNotBlank()) { "`apiKey` is missing." }
  }

  override val requester: Requester = requesterOf(clientName = "Ingestion", appId = appId, apiKey = apiKey, options = options) {
    val allowedRegions = listOf("eu", "us")
    require(region in allowedRegions) { "`region` is required and must be one of the following: ${allowedRegions.joinToString()}" }
    val url = "data.$region.algolia.com"
    listOf(Host(url))
  }

  /**
   * Create a authentication.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param authenticationCreate
   * @param requestOptions additional request configuration.
   */
  public suspend fun createAuthentication(authenticationCreate: AuthenticationCreate, requestOptions: RequestOptions? = null): AuthenticationCreateResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "authentications"),
      body = authenticationCreate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Create a destination.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param destinationCreate
   * @param requestOptions additional request configuration.
   */
  public suspend fun createDestination(destinationCreate: DestinationCreate, requestOptions: RequestOptions? = null): DestinationCreateResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "destinations"),
      body = destinationCreate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Create a source.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceCreate
   * @param requestOptions additional request configuration.
   */
  public suspend fun createSource(sourceCreate: SourceCreate, requestOptions: RequestOptions? = null): SourceCreateResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "sources"),
      body = sourceCreate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Create a task.
   * @param taskCreate
   * @param requestOptions additional request configuration.
   */
  public suspend fun createTask(taskCreate: TaskCreate, requestOptions: RequestOptions? = null): TaskCreateResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "tasks"),
      body = taskCreate,
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
   * Soft delete the authentication of the given authenticationID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param authenticationID The authentication UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun deleteAuthentication(authenticationID: String, requestOptions: RequestOptions? = null): DeleteResponse {
    require(authenticationID.isNotBlank()) { "Parameter `authenticationID` is required when calling `deleteAuthentication`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = listOf("1", "authentications", "$authenticationID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Soft delete the destination of the given destinationID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param destinationID The destination UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun deleteDestination(destinationID: String, requestOptions: RequestOptions? = null): DeleteResponse {
    require(destinationID.isNotBlank()) { "Parameter `destinationID` is required when calling `deleteDestination`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = listOf("1", "destinations", "$destinationID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Soft delete the source of the given sourceID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceID The source UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun deleteSource(sourceID: String, requestOptions: RequestOptions? = null): DeleteResponse {
    require(sourceID.isNotBlank()) { "Parameter `sourceID` is required when calling `deleteSource`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = listOf("1", "sources", "$sourceID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Soft delete the task of the given taskID.
   * @param taskID The task UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun deleteTask(taskID: String, requestOptions: RequestOptions? = null): DeleteResponse {
    require(taskID.isNotBlank()) { "Parameter `taskID` is required when calling `deleteTask`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = listOf("1", "tasks", "$taskID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Disable the task of the given taskID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param taskID The task UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun disableTask(taskID: String, requestOptions: RequestOptions? = null): TaskUpdateResponse {
    require(taskID.isNotBlank()) { "Parameter `taskID` is required when calling `disableTask`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PUT,
      path = listOf("1", "tasks", "$taskID", "disable"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Enable the task of the given taskID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param taskID The task UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun enableTask(taskID: String, requestOptions: RequestOptions? = null): TaskUpdateResponse {
    require(taskID.isNotBlank()) { "Parameter `taskID` is required when calling `enableTask`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PUT,
      path = listOf("1", "tasks", "$taskID", "enable"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get the authentication of the given authenticationID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param authenticationID The authentication UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getAuthentication(authenticationID: String, requestOptions: RequestOptions? = null): Authentication {
    require(authenticationID.isNotBlank()) { "Parameter `authenticationID` is required when calling `getAuthentication`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "authentications", "$authenticationID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a list of authentications for the given query parameters, with pagination details.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param itemsPerPage The number of items per page to return.
   * @param page The page number to fetch, starting at 1.
   * @param type The type of the authentications to retrieve.
   * @param platform The platform of the authentications to retrieve.
   * @param sort The key by which the list should be sorted.
   * @param order The order of the returned list.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getAuthentications(itemsPerPage: Int? = null, page: Int? = null, type: List<AuthenticationType>? = null, platform: List<PlatformWithNone>? = null, sort: AuthenticationSortKeys? = null, order: OrderKeys? = null, requestOptions: RequestOptions? = null): ListAuthenticationsResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "authentications"),
      query = buildMap {
        itemsPerPage?.let { put("itemsPerPage", it) }
        page?.let { put("page", it) }
        type?.let { put("type", it.joinToString(",")) }
        platform?.let { put("platform", it.joinToString(",")) }
        sort?.let { put("sort", it) }
        order?.let { put("order", it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get the destination of the given destinationID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param destinationID The destination UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getDestination(destinationID: String, requestOptions: RequestOptions? = null): Destination {
    require(destinationID.isNotBlank()) { "Parameter `destinationID` is required when calling `getDestination`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "destinations", "$destinationID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a list of destinations for the given query parameters, with pagination details.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param itemsPerPage The number of items per page to return.
   * @param page The page number to fetch, starting at 1.
   * @param type The type of the destinations to retrive.
   * @param authenticationID The authenticationIDs of the destinations to retrive.
   * @param sort The key by which the list should be sorted.
   * @param order The order of the returned list.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getDestinations(itemsPerPage: Int? = null, page: Int? = null, type: List<DestinationType>? = null, authenticationID: List<String>? = null, sort: DestinationSortKeys? = null, order: OrderKeys? = null, requestOptions: RequestOptions? = null): ListDestinationsResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "destinations"),
      query = buildMap {
        itemsPerPage?.let { put("itemsPerPage", it) }
        page?.let { put("page", it) }
        type?.let { put("type", it.joinToString(",")) }
        authenticationID?.let { put("authenticationID", it.joinToString(",")) }
        sort?.let { put("sort", it) }
        order?.let { put("order", it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Retrieve a stream listing for a given Singer specification compatible docker type source ID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceID The source UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getDockerSourceStreams(sourceID: String, requestOptions: RequestOptions? = null): DockerSourceStreams {
    require(sourceID.isNotBlank()) { "Parameter `sourceID` is required when calling `getDockerSourceStreams`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "sources", "$sourceID", "discover"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a single event for a specific runID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param runID The run UUID.
   * @param eventID The event UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getEvent(runID: String, eventID: String, requestOptions: RequestOptions? = null): Event {
    require(runID.isNotBlank()) { "Parameter `runID` is required when calling `getEvent`." }
    require(eventID.isNotBlank()) { "Parameter `eventID` is required when calling `getEvent`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "runs", "$runID", "events", "$eventID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a list of events associated to the given runID, for the given query parameters.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param runID The run UUID.
   * @param itemsPerPage The number of items per page to return.
   * @param page The page number to fetch, starting at 1.
   * @param status Filter the status of the events.
   * @param type Filter the type of the events.
   * @param sort The key by which the list should be sorted.
   * @param order The order of the returned list.
   * @param startDate The start date (in RFC3339 format) of the events fetching window. Defaults to 'now'-3 hours if omitted.
   * @param endDate The end date (in RFC3339 format) of the events fetching window. Defaults to 'now' days if omitted.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getEvents(runID: String, itemsPerPage: Int? = null, page: Int? = null, status: List<EventStatus>? = null, type: List<EventType>? = null, sort: EventSortKeys? = null, order: OrderKeys? = null, startDate: String? = null, endDate: String? = null, requestOptions: RequestOptions? = null): ListEventsResponse {
    require(runID.isNotBlank()) { "Parameter `runID` is required when calling `getEvents`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "runs", "$runID", "events"),
      query = buildMap {
        itemsPerPage?.let { put("itemsPerPage", it) }
        page?.let { put("page", it) }
        status?.let { put("status", it.joinToString(",")) }
        type?.let { put("type", it.joinToString(",")) }
        sort?.let { put("sort", it) }
        order?.let { put("order", it) }
        startDate?.let { put("startDate", it) }
        endDate?.let { put("endDate", it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a single run for the given ID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param runID The run UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getRun(runID: String, requestOptions: RequestOptions? = null): Run {
    require(runID.isNotBlank()) { "Parameter `runID` is required when calling `getRun`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "runs", "$runID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a list of runs for the given query parameters, with pagination details.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param itemsPerPage The number of items per page to return.
   * @param page The page number to fetch, starting at 1.
   * @param status Filter the status of the runs.
   * @param taskID Filter by taskID.
   * @param sort The key by which the list should be sorted.
   * @param order The order of the returned list.
   * @param startDate The start date (in RFC3339 format) of the runs fetching window. Defaults to 'now'-7 days if omitted.
   * @param endDate The end date (in RFC3339 format) of the runs fetching window. Defaults to 'now' days if omitted.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getRuns(itemsPerPage: Int? = null, page: Int? = null, status: List<RunStatus>? = null, taskID: String? = null, sort: RunSortKeys? = null, order: OrderKeys? = null, startDate: String? = null, endDate: String? = null, requestOptions: RequestOptions? = null): RunListResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "runs"),
      query = buildMap {
        itemsPerPage?.let { put("itemsPerPage", it) }
        page?.let { put("page", it) }
        status?.let { put("status", it.joinToString(",")) }
        taskID?.let { put("taskID", it) }
        sort?.let { put("sort", it) }
        order?.let { put("order", it) }
        startDate?.let { put("startDate", it) }
        endDate?.let { put("endDate", it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get the source of the given sourceID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceID The source UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getSource(sourceID: String, requestOptions: RequestOptions? = null): Source {
    require(sourceID.isNotBlank()) { "Parameter `sourceID` is required when calling `getSource`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "sources", "$sourceID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a list of sources for the given query parameters, with pagination details.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param itemsPerPage The number of items per page to return.
   * @param page The page number to fetch, starting at 1.
   * @param type The type of the sources to retrieve.
   * @param authenticationID The authenticationIDs of the sources to retrieve. 'none' returns sources that doesn't have an authentication.
   * @param sort The key by which the list should be sorted.
   * @param order The order of the returned list.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getSources(itemsPerPage: Int? = null, page: Int? = null, type: List<SourceType>? = null, authenticationID: List<String>? = null, sort: SourceSortKeys? = null, order: OrderKeys? = null, requestOptions: RequestOptions? = null): ListSourcesResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "sources"),
      query = buildMap {
        itemsPerPage?.let { put("itemsPerPage", it) }
        page?.let { put("page", it) }
        type?.let { put("type", it.joinToString(",")) }
        authenticationID?.let { put("authenticationID", it.joinToString(",")) }
        sort?.let { put("sort", it) }
        order?.let { put("order", it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get the task of the given taskID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param taskID The task UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getTask(taskID: String, requestOptions: RequestOptions? = null): Task {
    require(taskID.isNotBlank()) { "Parameter `taskID` is required when calling `getTask`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "tasks", "$taskID"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a list of tasks for the given query parameters, with pagination details.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param itemsPerPage The number of items per page to return.
   * @param page The page number to fetch, starting at 1.
   * @param action The action of the tasks to retrieve.
   * @param enabled Whether the task is enabled or not.
   * @param sourceID The sourceIDs of the tasks to retrieve.
   * @param destinationID The destinationIDs of the tasks to retrieve.
   * @param triggerType The trigger type of the task.
   * @param sort The key by which the list should be sorted.
   * @param order The order of the returned list.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getTasks(itemsPerPage: Int? = null, page: Int? = null, action: List<ActionType>? = null, enabled: Boolean? = null, sourceID: List<String>? = null, destinationID: List<String>? = null, triggerType: List<TriggerType>? = null, sort: TaskSortKeys? = null, order: OrderKeys? = null, requestOptions: RequestOptions? = null): ListTasksResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "tasks"),
      query = buildMap {
        itemsPerPage?.let { put("itemsPerPage", it) }
        page?.let { put("page", it) }
        action?.let { put("action", it.joinToString(",")) }
        enabled?.let { put("enabled", it) }
        sourceID?.let { put("sourceID", it.joinToString(",")) }
        destinationID?.let { put("destinationID", it.joinToString(",")) }
        triggerType?.let { put("triggerType", it.joinToString(",")) }
        sort?.let { put("sort", it) }
        order?.let { put("order", it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Run the task of the given taskID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param taskID The task UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun runTask(taskID: String, requestOptions: RequestOptions? = null): RunResponse {
    require(taskID.isNotBlank()) { "Parameter `taskID` is required when calling `runTask`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "tasks", "$taskID", "run"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Search among authentications with a defined set of parameters.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param authenticationSearch
   * @param requestOptions additional request configuration.
   */
  public suspend fun searchAuthentications(authenticationSearch: AuthenticationSearch, requestOptions: RequestOptions? = null): List<Authentication> {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "authentications", "search"),
      body = authenticationSearch,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Search among destinations with a defined set of parameters.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param destinationSearch
   * @param requestOptions additional request configuration.
   */
  public suspend fun searchDestinations(destinationSearch: DestinationSearch, requestOptions: RequestOptions? = null): List<Destination> {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "destinations", "search"),
      body = destinationSearch,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Search among sources with a defined set of parameters.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceSearch
   * @param requestOptions additional request configuration.
   */
  public suspend fun searchSources(sourceSearch: SourceSearch, requestOptions: RequestOptions? = null): List<Source> {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "sources", "search"),
      body = sourceSearch,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Search among tasks with a defined set of parameters.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param taskSearch
   * @param requestOptions additional request configuration.
   */
  public suspend fun searchTasks(taskSearch: TaskSearch, requestOptions: RequestOptions? = null): List<Task> {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "tasks", "search"),
      body = taskSearch,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Trigger a stream listing request for a Singer specification compatible docker type source.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceID The source UUID.
   * @param requestOptions additional request configuration.
   */
  public suspend fun triggerDockerSourceDiscover(sourceID: String, requestOptions: RequestOptions? = null): DockerSourceDiscover {
    require(sourceID.isNotBlank()) { "Parameter `sourceID` is required when calling `triggerDockerSourceDiscover`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "sources", "$sourceID", "discover"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Update the authentication of the given authenticationID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param authenticationID The authentication UUID.
   * @param authenticationUpdate
   * @param requestOptions additional request configuration.
   */
  public suspend fun updateAuthentication(authenticationID: String, authenticationUpdate: AuthenticationUpdate, requestOptions: RequestOptions? = null): AuthenticationUpdateResponse {
    require(authenticationID.isNotBlank()) { "Parameter `authenticationID` is required when calling `updateAuthentication`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PATCH,
      path = listOf("1", "authentications", "$authenticationID"),
      body = authenticationUpdate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Update the destination of the given destinationID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param destinationID The destination UUID.
   * @param destinationUpdate
   * @param requestOptions additional request configuration.
   */
  public suspend fun updateDestination(destinationID: String, destinationUpdate: DestinationUpdate, requestOptions: RequestOptions? = null): DestinationUpdateResponse {
    require(destinationID.isNotBlank()) { "Parameter `destinationID` is required when calling `updateDestination`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PATCH,
      path = listOf("1", "destinations", "$destinationID"),
      body = destinationUpdate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Update the source of the given sourceID.
   *
   * Required API Key ACLs:
   *   - addObject
   *   - deleteIndex
   *   - editSettings
   * @param sourceID The source UUID.
   * @param sourceUpdate
   * @param requestOptions additional request configuration.
   */
  public suspend fun updateSource(sourceID: String, sourceUpdate: SourceUpdate, requestOptions: RequestOptions? = null): SourceUpdateResponse {
    require(sourceID.isNotBlank()) { "Parameter `sourceID` is required when calling `updateSource`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PATCH,
      path = listOf("1", "sources", "$sourceID"),
      body = sourceUpdate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Update the task of the given taskID.
   * @param taskID The task UUID.
   * @param taskUpdate
   * @param requestOptions additional request configuration.
   */
  public suspend fun updateTask(taskID: String, taskUpdate: TaskUpdate, requestOptions: RequestOptions? = null): TaskUpdateResponse {
    require(taskID.isNotBlank()) { "Parameter `taskID` is required when calling `updateTask`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PATCH,
      path = listOf("1", "tasks", "$taskID"),
      body = taskUpdate,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }
}
