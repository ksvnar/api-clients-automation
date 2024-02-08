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

namespace Algolia.Search.Models.Recommend;

/// <summary>
/// Additional search parameters.
/// </summary>
[DataContract(Name = "varParams")]
public partial class Params
{
  /// <summary>
  /// Initializes a new instance of the Params class.
  /// </summary>
  public Params()
  {
  }

  /// <summary>
  /// Gets or Sets Query
  /// </summary>
  [DataMember(Name = "query")]
  public ConsequenceQuery Query { get; set; }

  /// <summary>
  /// Gets or Sets AutomaticFacetFilters
  /// </summary>
  [DataMember(Name = "automaticFacetFilters")]
  public AutomaticFacetFilters AutomaticFacetFilters { get; set; }

  /// <summary>
  /// Gets or Sets AutomaticOptionalFacetFilters
  /// </summary>
  [DataMember(Name = "automaticOptionalFacetFilters")]
  public AutomaticFacetFilters AutomaticOptionalFacetFilters { get; set; }

  /// <summary>
  /// Gets or Sets RenderingContent
  /// </summary>
  [DataMember(Name = "renderingContent")]
  public RenderingContent RenderingContent { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class Params {\n");
    sb.Append("  Query: ").Append(Query).Append("\n");
    sb.Append("  AutomaticFacetFilters: ").Append(AutomaticFacetFilters).Append("\n");
    sb.Append("  AutomaticOptionalFacetFilters: ").Append(AutomaticOptionalFacetFilters).Append("\n");
    sb.Append("  RenderingContent: ").Append(RenderingContent).Append("\n");
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

