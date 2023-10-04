// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.insights;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/**
 * Use this event to track when users make a purchase unrelated to a previous Algolia request. For
 * example, if you don't use Algolia to build your category pages, use this event. To track purchase
 * events related to Algolia requests, use the \"Purchased object IDs after search\" event.
 */
@JsonDeserialize(as = PurchasedObjectIDs.class)
public class PurchasedObjectIDs implements EventsItems {

  @JsonProperty("eventName")
  private String eventName;

  @JsonProperty("eventType")
  private ConversionEvent eventType;

  @JsonProperty("eventSubtype")
  private PurchaseEvent eventSubtype;

  @JsonProperty("index")
  private String index;

  @JsonProperty("objectIDs")
  private List<String> objectIDs = new ArrayList<>();

  @JsonProperty("objectData")
  private List<ObjectData> objectData;

  @JsonProperty("currency")
  private String currency;

  @JsonProperty("userToken")
  private String userToken;

  @JsonProperty("timestamp")
  private Long timestamp;

  public PurchasedObjectIDs setEventName(String eventName) {
    this.eventName = eventName;
    return this;
  }

  /**
   * Can contain up to 64 ASCII characters. Consider naming events consistently—for example, by
   * adopting Segment's
   * [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework)
   * framework.
   */
  @javax.annotation.Nonnull
  public String getEventName() {
    return eventName;
  }

  public PurchasedObjectIDs setEventType(ConversionEvent eventType) {
    this.eventType = eventType;
    return this;
  }

  /** Get eventType */
  @javax.annotation.Nonnull
  public ConversionEvent getEventType() {
    return eventType;
  }

  public PurchasedObjectIDs setEventSubtype(PurchaseEvent eventSubtype) {
    this.eventSubtype = eventSubtype;
    return this;
  }

  /** Get eventSubtype */
  @javax.annotation.Nonnull
  public PurchaseEvent getEventSubtype() {
    return eventSubtype;
  }

  public PurchasedObjectIDs setIndex(String index) {
    this.index = index;
    return this;
  }

  /** Name of the Algolia index. */
  @javax.annotation.Nonnull
  public String getIndex() {
    return index;
  }

  public PurchasedObjectIDs setObjectIDs(List<String> objectIDs) {
    this.objectIDs = objectIDs;
    return this;
  }

  public PurchasedObjectIDs addObjectIDs(String objectIDsItem) {
    this.objectIDs.add(objectIDsItem);
    return this;
  }

  /** List of object identifiers for items of an Algolia index. */
  @javax.annotation.Nonnull
  public List<String> getObjectIDs() {
    return objectIDs;
  }

  public PurchasedObjectIDs setObjectData(List<ObjectData> objectData) {
    this.objectData = objectData;
    return this;
  }

  public PurchasedObjectIDs addObjectData(ObjectData objectDataItem) {
    if (this.objectData == null) {
      this.objectData = new ArrayList<>();
    }
    this.objectData.add(objectDataItem);
    return this;
  }

  /**
   * Extra information about the records involved in the event—for example, to add price and
   * quantities of purchased products. If provided, must be the same length as `objectIDs`.
   */
  @javax.annotation.Nullable
  public List<ObjectData> getObjectData() {
    return objectData;
  }

  public PurchasedObjectIDs setCurrency(String currency) {
    this.currency = currency;
    return this;
  }

  /**
   * If you include pricing information in the `objectData` parameter, you must also specify the
   * currency as ISO-4217 currency code, such as USD or EUR.
   */
  @javax.annotation.Nullable
  public String getCurrency() {
    return currency;
  }

  public PurchasedObjectIDs setUserToken(String userToken) {
    this.userToken = userToken;
    return this;
  }

  /**
   * Anonymous or pseudonymous user identifier. > **Note**: Never include personally identifiable
   * information in user tokens.
   */
  @javax.annotation.Nonnull
  public String getUserToken() {
    return userToken;
  }

  public PurchasedObjectIDs setTimestamp(Long timestamp) {
    this.timestamp = timestamp;
    return this;
  }

  /**
   * Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time).
   * By default, the Insights API uses the time it receives an event as its timestamp.
   */
  @javax.annotation.Nullable
  public Long getTimestamp() {
    return timestamp;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PurchasedObjectIDs purchasedObjectIDs = (PurchasedObjectIDs) o;
    return (
      Objects.equals(this.eventName, purchasedObjectIDs.eventName) &&
      Objects.equals(this.eventType, purchasedObjectIDs.eventType) &&
      Objects.equals(this.eventSubtype, purchasedObjectIDs.eventSubtype) &&
      Objects.equals(this.index, purchasedObjectIDs.index) &&
      Objects.equals(this.objectIDs, purchasedObjectIDs.objectIDs) &&
      Objects.equals(this.objectData, purchasedObjectIDs.objectData) &&
      Objects.equals(this.currency, purchasedObjectIDs.currency) &&
      Objects.equals(this.userToken, purchasedObjectIDs.userToken) &&
      Objects.equals(this.timestamp, purchasedObjectIDs.timestamp)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(eventName, eventType, eventSubtype, index, objectIDs, objectData, currency, userToken, timestamp);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PurchasedObjectIDs {\n");
    sb.append("    eventName: ").append(toIndentedString(eventName)).append("\n");
    sb.append("    eventType: ").append(toIndentedString(eventType)).append("\n");
    sb.append("    eventSubtype: ").append(toIndentedString(eventSubtype)).append("\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    objectIDs: ").append(toIndentedString(objectIDs)).append("\n");
    sb.append("    objectData: ").append(toIndentedString(objectData)).append("\n");
    sb.append("    currency: ").append(toIndentedString(currency)).append("\n");
    sb.append("    userToken: ").append(toIndentedString(userToken)).append("\n");
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
