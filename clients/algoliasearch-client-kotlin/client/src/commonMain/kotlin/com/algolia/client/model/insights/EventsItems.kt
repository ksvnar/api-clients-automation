/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.insights

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * EventsItems
 */
@Serializable(EventsItemsSerializer::class)
public sealed interface EventsItems {

  public companion object {

    /**
     * Use this event to track when users add items to their shopping cart unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track add-to-cart events related to Algolia requests, use the \"Added to cart object IDs after search\" event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param eventSubtype
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param objectData Extra information about the records involved in the event—for example, to add price and quantities of purchased products.  If provided, must be the same length as `objectIDs`.
     * @param currency If you include pricing information in the `objectData` parameter, you must also specify the currency as ISO-4217 currency code, such as USD or EUR.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun AddedToCartObjectIDs(
      eventName: String,
      eventType: ConversionEvent,
      eventSubtype: AddToCartEvent,
      index: String,
      objectIDs: List<String>,
      userToken: String,
      objectData: List<ObjectData>? = null,
      currency: String? = null,
      timestamp: Long? = null,
    ): AddedToCartObjectIDs = com.algolia.client.model.insights.AddedToCartObjectIDs(
      eventName = eventName,
      eventType = eventType,
      eventSubtype = eventSubtype,
      index = index,
      objectIDs = objectIDs,
      userToken = userToken,
      objectData = objectData,
      currency = currency,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users add items to their shopping cart after a previous Algolia request. If you're building your category pages with Algolia, you'll also use this event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param eventSubtype
     * @param index Name of the Algolia index.
     * @param queryID Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param objectData Extra information about the records involved in the event—for example, to add price and quantities of purchased products.  If provided, must be the same length as `objectIDs`.
     * @param currency If you include pricing information in the `objectData` parameter, you must also specify the currency as ISO-4217 currency code, such as USD or EUR.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun AddedToCartObjectIDsAfterSearch(
      eventName: String,
      eventType: ConversionEvent,
      eventSubtype: AddToCartEvent,
      index: String,
      queryID: String,
      objectIDs: List<String>,
      userToken: String,
      objectData: List<ObjectDataAfterSearch>? = null,
      currency: String? = null,
      timestamp: Long? = null,
    ): AddedToCartObjectIDsAfterSearch = com.algolia.client.model.insights.AddedToCartObjectIDsAfterSearch(
      eventName = eventName,
      eventType = eventType,
      eventSubtype = eventSubtype,
      index = index,
      queryID = queryID,
      objectIDs = objectIDs,
      userToken = userToken,
      objectData = objectData,
      currency = currency,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users click facet filters in your user interface.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param filters Facet filters.  Each facet filter string must be URL-encoded, such as, `discount:10%25`.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ClickedFilters(
      eventName: String,
      eventType: ClickEvent,
      index: String,
      filters: List<String>,
      userToken: String,
      timestamp: Long? = null,
    ): ClickedFilters = com.algolia.client.model.insights.ClickedFilters(
      eventName = eventName,
      eventType = eventType,
      index = index,
      filters = filters,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users click items unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track click events related to Algolia requests, use the \"Clicked object IDs after search\" event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ClickedObjectIDs(
      eventName: String,
      eventType: ClickEvent,
      index: String,
      objectIDs: List<String>,
      userToken: String,
      timestamp: Long? = null,
    ): ClickedObjectIDs = com.algolia.client.model.insights.ClickedObjectIDs(
      eventName = eventName,
      eventType = eventType,
      index = index,
      objectIDs = objectIDs,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * Click event after an Algolia request.  Use this event to track when users click items in the search results. If you're building your category pages with Algolia, you'll also use this event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param positions Position of the clicked objects in the search results.  The first search result has a position of 1 (not 0). You must provide 1 `position` for each `objectID`.
     * @param queryID Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ClickedObjectIDsAfterSearch(
      eventName: String,
      eventType: ClickEvent,
      index: String,
      objectIDs: List<String>,
      positions: List<Int>,
      queryID: String,
      userToken: String,
      timestamp: Long? = null,
    ): ClickedObjectIDsAfterSearch = com.algolia.client.model.insights.ClickedObjectIDsAfterSearch(
      eventName = eventName,
      eventType = eventType,
      index = index,
      objectIDs = objectIDs,
      positions = positions,
      queryID = queryID,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * ConvertedFilters
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param filters Facet filters.  Each facet filter string must be URL-encoded, such as, `discount:10%25`.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ConvertedFilters(
      eventName: String,
      eventType: ConversionEvent,
      index: String,
      filters: List<String>,
      userToken: String,
      timestamp: Long? = null,
    ): ConvertedFilters = com.algolia.client.model.insights.ConvertedFilters(
      eventName = eventName,
      eventType = eventType,
      index = index,
      filters = filters,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users convert on items unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track conversion events related to Algolia requests, use the \"Converted object IDs after search\" event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ConvertedObjectIDs(
      eventName: String,
      eventType: ConversionEvent,
      index: String,
      objectIDs: List<String>,
      userToken: String,
      timestamp: Long? = null,
    ): ConvertedObjectIDs = com.algolia.client.model.insights.ConvertedObjectIDs(
      eventName = eventName,
      eventType = eventType,
      index = index,
      objectIDs = objectIDs,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users convert after a previous Algolia request. For example, a user clicks on an item in the search results to view the product detail page. Then, the user adds the item to their shopping cart.  If you're building your category pages with Algolia, you'll also use this event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param queryID Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ConvertedObjectIDsAfterSearch(
      eventName: String,
      eventType: ConversionEvent,
      index: String,
      objectIDs: List<String>,
      queryID: String,
      userToken: String,
      timestamp: Long? = null,
    ): ConvertedObjectIDsAfterSearch = com.algolia.client.model.insights.ConvertedObjectIDsAfterSearch(
      eventName = eventName,
      eventType = eventType,
      index = index,
      objectIDs = objectIDs,
      queryID = queryID,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users make a purchase unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track purchase events related to Algolia requests, use the \"Purchased object IDs after search\" event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param eventSubtype
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param objectData Extra information about the records involved in the event—for example, to add price and quantities of purchased products.  If provided, must be the same length as `objectIDs`.
     * @param currency If you include pricing information in the `objectData` parameter, you must also specify the currency as ISO-4217 currency code, such as USD or EUR.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun PurchasedObjectIDs(
      eventName: String,
      eventType: ConversionEvent,
      eventSubtype: PurchaseEvent,
      index: String,
      objectIDs: List<String>,
      userToken: String,
      objectData: List<ObjectData>? = null,
      currency: String? = null,
      timestamp: Long? = null,
    ): PurchasedObjectIDs = com.algolia.client.model.insights.PurchasedObjectIDs(
      eventName = eventName,
      eventType = eventType,
      eventSubtype = eventSubtype,
      index = index,
      objectIDs = objectIDs,
      userToken = userToken,
      objectData = objectData,
      currency = currency,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users make a purchase after a previous Algolia request. If you're building your category pages with Algolia, you'll also use this event.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param eventSubtype
     * @param index Name of the Algolia index.
     * @param queryID Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param objectData Extra information about the records involved in the event—for example, to add price and quantities of purchased products.  If provided, must be the same length as `objectIDs`.
     * @param currency If you include pricing information in the `objectData` parameter, you must also specify the currency as ISO-4217 currency code, such as USD or EUR.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun PurchasedObjectIDsAfterSearch(
      eventName: String,
      eventType: ConversionEvent,
      eventSubtype: PurchaseEvent,
      index: String,
      queryID: String,
      objectIDs: List<String>,
      userToken: String,
      objectData: List<ObjectDataAfterSearch>? = null,
      currency: String? = null,
      timestamp: Long? = null,
    ): PurchasedObjectIDsAfterSearch = com.algolia.client.model.insights.PurchasedObjectIDsAfterSearch(
      eventName = eventName,
      eventType = eventType,
      eventSubtype = eventSubtype,
      index = index,
      queryID = queryID,
      objectIDs = objectIDs,
      userToken = userToken,
      objectData = objectData,
      currency = currency,
      timestamp = timestamp,
    )

    /**
     * Use this method to capture active filters. For example, when browsing a category page, users see content filtered on that specific category.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param filters Facet filters.  Each facet filter string must be URL-encoded, such as, `discount:10%25`.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ViewedFilters(
      eventName: String,
      eventType: ViewEvent,
      index: String,
      filters: List<String>,
      userToken: String,
      timestamp: Long? = null,
    ): ViewedFilters = com.algolia.client.model.insights.ViewedFilters(
      eventName = eventName,
      eventType = eventType,
      index = index,
      filters = filters,
      userToken = userToken,
      timestamp = timestamp,
    )

    /**
     * Use this event to track when users viewed items in the search results.
     *
     * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     * @param eventType
     * @param index Name of the Algolia index.
     * @param objectIDs List of object identifiers for items of an Algolia index.
     * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
     * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
     */
    public fun ViewedObjectIDs(
      eventName: String,
      eventType: ViewEvent,
      index: String,
      objectIDs: List<String>,
      userToken: String,
      timestamp: Long? = null,
    ): ViewedObjectIDs = com.algolia.client.model.insights.ViewedObjectIDs(
      eventName = eventName,
      eventType = eventType,
      index = index,
      objectIDs = objectIDs,
      userToken = userToken,
      timestamp = timestamp,
    )
  }
}

internal class EventsItemsSerializer : KSerializer<EventsItems> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("EventsItems")

  override fun serialize(encoder: Encoder, value: EventsItems) {
    when (value) {
      is AddedToCartObjectIDs -> AddedToCartObjectIDs.serializer().serialize(encoder, value)
      is AddedToCartObjectIDsAfterSearch -> AddedToCartObjectIDsAfterSearch.serializer().serialize(encoder, value)
      is ClickedFilters -> ClickedFilters.serializer().serialize(encoder, value)
      is ClickedObjectIDs -> ClickedObjectIDs.serializer().serialize(encoder, value)
      is ClickedObjectIDsAfterSearch -> ClickedObjectIDsAfterSearch.serializer().serialize(encoder, value)
      is ConvertedFilters -> ConvertedFilters.serializer().serialize(encoder, value)
      is ConvertedObjectIDs -> ConvertedObjectIDs.serializer().serialize(encoder, value)
      is ConvertedObjectIDsAfterSearch -> ConvertedObjectIDsAfterSearch.serializer().serialize(encoder, value)
      is PurchasedObjectIDs -> PurchasedObjectIDs.serializer().serialize(encoder, value)
      is PurchasedObjectIDsAfterSearch -> PurchasedObjectIDsAfterSearch.serializer().serialize(encoder, value)
      is ViewedFilters -> ViewedFilters.serializer().serialize(encoder, value)
      is ViewedObjectIDs -> ViewedObjectIDs.serializer().serialize(encoder, value)
    }
  }

  override fun deserialize(decoder: Decoder): EventsItems {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize AddedToCartObjectIDs
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AddedToCartObjectIDs>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AddedToCartObjectIDs (error: ${e.message})")
      }
    }

    // deserialize AddedToCartObjectIDsAfterSearch
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AddedToCartObjectIDsAfterSearch>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AddedToCartObjectIDsAfterSearch (error: ${e.message})")
      }
    }

    // deserialize ClickedFilters
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ClickedFilters>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ClickedFilters (error: ${e.message})")
      }
    }

    // deserialize ClickedObjectIDs
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ClickedObjectIDs>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ClickedObjectIDs (error: ${e.message})")
      }
    }

    // deserialize ClickedObjectIDsAfterSearch
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ClickedObjectIDsAfterSearch>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ClickedObjectIDsAfterSearch (error: ${e.message})")
      }
    }

    // deserialize ConvertedFilters
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ConvertedFilters>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ConvertedFilters (error: ${e.message})")
      }
    }

    // deserialize ConvertedObjectIDs
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ConvertedObjectIDs>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ConvertedObjectIDs (error: ${e.message})")
      }
    }

    // deserialize ConvertedObjectIDsAfterSearch
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ConvertedObjectIDsAfterSearch>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ConvertedObjectIDsAfterSearch (error: ${e.message})")
      }
    }

    // deserialize PurchasedObjectIDs
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<PurchasedObjectIDs>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize PurchasedObjectIDs (error: ${e.message})")
      }
    }

    // deserialize PurchasedObjectIDsAfterSearch
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<PurchasedObjectIDsAfterSearch>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize PurchasedObjectIDsAfterSearch (error: ${e.message})")
      }
    }

    // deserialize ViewedFilters
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ViewedFilters>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ViewedFilters (error: ${e.message})")
      }
    }

    // deserialize ViewedObjectIDs
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<ViewedObjectIDs>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize ViewedObjectIDs (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
