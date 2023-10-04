/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.insights

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * ObjectData
 *
 * @param price
 * @param quantity The quantity of the purchased or added-to-cart item. The total value of a purchase is the sum of `quantity` multiplied with the `price` for each purchased item.
 * @param discount
 */
@Serializable
public data class ObjectData(

  @SerialName(value = "price") val price: Price? = null,

  /** The quantity of the purchased or added-to-cart item. The total value of a purchase is the sum of `quantity` multiplied with the `price` for each purchased item. */
  @SerialName(value = "quantity") val quantity: Int? = null,

  @SerialName(value = "discount") val discount: Discount? = null,
)
