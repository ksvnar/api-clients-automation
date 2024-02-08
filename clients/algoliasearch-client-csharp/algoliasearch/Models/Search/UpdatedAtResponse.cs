//
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
//
using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using Algolia.Search.Models;
using Algolia.Search.Models.Common;
using Algolia.Search.Serializer;

namespace Algolia.Search.Models.Search;

/// <summary>
/// Response, taskID, and update timestamp.
/// </summary>
[DataContract(Name = "updatedAtResponse")]
public partial class UpdatedAtResponse
{
  /// <summary>
  /// Initializes a new instance of the UpdatedAtResponse class.
  /// </summary>
  [JsonConstructor]
  public UpdatedAtResponse() { }
  /// <summary>
  /// Initializes a new instance of the UpdatedAtResponse class.
  /// </summary>
  /// <param name="taskID">Unique identifier of a task. A successful API response means that a task was added to a queue. It might not run immediately. You can check the task&#39;s progress with the &#x60;task&#x60; operation and this &#x60;taskID&#x60;.  (required).</param>
  /// <param name="updatedAt">Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. (required).</param>
  public UpdatedAtResponse(long taskID, string updatedAt)
  {
    TaskID = taskID;
    UpdatedAt = updatedAt ?? throw new ArgumentNullException(nameof(updatedAt));
  }

  /// <summary>
  /// Unique identifier of a task. A successful API response means that a task was added to a queue. It might not run immediately. You can check the task's progress with the `task` operation and this `taskID`. 
  /// </summary>
  /// <value>Unique identifier of a task. A successful API response means that a task was added to a queue. It might not run immediately. You can check the task's progress with the `task` operation and this `taskID`. </value>
  [DataMember(Name = "taskID")]
  public long TaskID { get; set; }

  /// <summary>
  /// Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
  /// </summary>
  /// <value>Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.</value>
  [DataMember(Name = "updatedAt")]
  public string UpdatedAt { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class UpdatedAtResponse {\n");
    sb.Append("  TaskID: ").Append(TaskID).Append("\n");
    sb.Append("  UpdatedAt: ").Append(UpdatedAt).Append("\n");
    sb.Append("}\n");
    return sb.ToString();
  }

  /// <summary>
  /// Returns the JSON string presentation of the object
  /// </summary>
  /// <returns>JSON string presentation of the object</returns>
  public virtual string ToJson()
  {
    return JsonConvert.SerializeObject(this, Formatting.Indented);
  }

}

