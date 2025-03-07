/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * API request body for updating a destination.
 *
 * @param type
 * @param name Descriptive name for the resource.
 * @param input
 * @param authenticationID Universally unique identifier (UUID) of an authentication resource.
 */
@Serializable
public data class DestinationUpdate(

  @SerialName(value = "type") val type: DestinationType? = null,

  /** Descriptive name for the resource. */
  @SerialName(value = "name") val name: String? = null,

  @SerialName(value = "input") val input: DestinationInput? = null,

  /** Universally unique identifier (UUID) of an authentication resource. */
  @SerialName(value = "authenticationID") val authenticationID: String? = null,
)
