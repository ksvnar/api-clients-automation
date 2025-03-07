// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.ingestion;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** Time window by which to filter the observability data. */
public class Window {

  @JsonProperty("startDate")
  private String startDate;

  @JsonProperty("endDate")
  private String endDate;

  public Window setStartDate(String startDate) {
    this.startDate = startDate;
    return this;
  }

  /** Date in RFC3339 format representing the oldest data in the time window. */
  @javax.annotation.Nonnull
  public String getStartDate() {
    return startDate;
  }

  public Window setEndDate(String endDate) {
    this.endDate = endDate;
    return this;
  }

  /** Date in RFC3339 format representing the newest data in the time window. */
  @javax.annotation.Nonnull
  public String getEndDate() {
    return endDate;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Window window = (Window) o;
    return Objects.equals(this.startDate, window.startDate) && Objects.equals(this.endDate, window.endDate);
  }

  @Override
  public int hashCode() {
    return Objects.hash(startDate, endDate);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Window {\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
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
