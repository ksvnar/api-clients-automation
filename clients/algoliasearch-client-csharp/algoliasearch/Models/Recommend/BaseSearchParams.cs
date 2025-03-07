//
// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
//
using System;
using System.Text;
using System.Linq;
using System.Text.Json.Serialization;
using System.Collections.Generic;
using Algolia.Search.Serializer;
using System.Text.Json;

namespace Algolia.Search.Models.Recommend;

/// <summary>
/// BaseSearchParams
/// </summary>
public partial class BaseSearchParams
{
  /// <summary>
  /// Initializes a new instance of the BaseSearchParams class.
  /// </summary>
  public BaseSearchParams()
  {
  }

  /// <summary>
  /// Search query.
  /// </summary>
  /// <value>Search query.</value>
  [JsonPropertyName("query")]
  public string Query { get; set; }

  /// <summary>
  /// Keywords to be used instead of the search query to conduct a more broader search.  Using the `similarQuery` parameter changes other settings:  - `queryType` is set to `prefixNone`. - `removeStopWords` is set to true. - `words` is set as the first ranking criterion. - All remaining words are treated as `optionalWords`.  Since the `similarQuery` is supposed to do a broad search, they usually return many results. Combine it with `filters` to narrow down the list of results. 
  /// </summary>
  /// <value>Keywords to be used instead of the search query to conduct a more broader search.  Using the `similarQuery` parameter changes other settings:  - `queryType` is set to `prefixNone`. - `removeStopWords` is set to true. - `words` is set as the first ranking criterion. - All remaining words are treated as `optionalWords`.  Since the `similarQuery` is supposed to do a broad search, they usually return many results. Combine it with `filters` to narrow down the list of results. </value>
  [JsonPropertyName("similarQuery")]
  public string SimilarQuery { get; set; }

  /// <summary>
  /// Filter expression to only include items that match the filter criteria in the response.  You can use these filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`, `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>` (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.   **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an array, the filter matches if it matches at least one element of the array.  For more information, see [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/). 
  /// </summary>
  /// <value>Filter expression to only include items that match the filter criteria in the response.  You can use these filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`, `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>` (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.   **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an array, the filter matches if it matches at least one element of the array.  For more information, see [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/). </value>
  [JsonPropertyName("filters")]
  public string Filters { get; set; }

  /// <summary>
  /// Gets or Sets FacetFilters
  /// </summary>
  [JsonPropertyName("facetFilters")]
  public FacetFilters FacetFilters { get; set; }

  /// <summary>
  /// Gets or Sets OptionalFilters
  /// </summary>
  [JsonPropertyName("optionalFilters")]
  public OptionalFilters OptionalFilters { get; set; }

  /// <summary>
  /// Gets or Sets NumericFilters
  /// </summary>
  [JsonPropertyName("numericFilters")]
  public NumericFilters NumericFilters { get; set; }

  /// <summary>
  /// Gets or Sets TagFilters
  /// </summary>
  [JsonPropertyName("tagFilters")]
  public TagFilters TagFilters { get; set; }

  /// <summary>
  /// Whether to sum all filter scores.  If true, all filter scores are summed. Otherwise, the maximum filter score is kept. For more information, see [filter scores](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/in-depth/filter-scoring/#accumulating-scores-with-sumorfiltersscores). 
  /// </summary>
  /// <value>Whether to sum all filter scores.  If true, all filter scores are summed. Otherwise, the maximum filter score is kept. For more information, see [filter scores](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/in-depth/filter-scoring/#accumulating-scores-with-sumorfiltersscores). </value>
  [JsonPropertyName("sumOrFiltersScores")]
  public bool? SumOrFiltersScores { get; set; }

  /// <summary>
  /// Restricts a search to a subset of your searchable attributes. Attribute names are case-sensitive. 
  /// </summary>
  /// <value>Restricts a search to a subset of your searchable attributes. Attribute names are case-sensitive. </value>
  [JsonPropertyName("restrictSearchableAttributes")]
  public List<string> RestrictSearchableAttributes { get; set; }

  /// <summary>
  /// Facets for which to retrieve facet values that match the search criteria and the number of matching facet values.  To retrieve all facets, use the wildcard character `*`. For more information, see [facets](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/#contextual-facet-values-and-counts). 
  /// </summary>
  /// <value>Facets for which to retrieve facet values that match the search criteria and the number of matching facet values.  To retrieve all facets, use the wildcard character `*`. For more information, see [facets](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/#contextual-facet-values-and-counts). </value>
  [JsonPropertyName("facets")]
  public List<string> Facets { get; set; }

  /// <summary>
  /// Whether faceting should be applied after deduplication with `distinct`.  This leads to accurate facet counts when using faceting in combination with `distinct`. It's usually better to use `afterDistinct` modifiers in the `attributesForFaceting` setting, as `facetingAfterDistinct` only computes correct facet counts if all records have the same facet values for the `attributeForDistinct`. 
  /// </summary>
  /// <value>Whether faceting should be applied after deduplication with `distinct`.  This leads to accurate facet counts when using faceting in combination with `distinct`. It's usually better to use `afterDistinct` modifiers in the `attributesForFaceting` setting, as `facetingAfterDistinct` only computes correct facet counts if all records have the same facet values for the `attributeForDistinct`. </value>
  [JsonPropertyName("facetingAfterDistinct")]
  public bool? FacetingAfterDistinct { get; set; }

  /// <summary>
  /// Page of search results to retrieve.
  /// </summary>
  /// <value>Page of search results to retrieve.</value>
  [JsonPropertyName("page")]
  public int? Page { get; set; }

  /// <summary>
  /// Position of the first hit to retrieve.
  /// </summary>
  /// <value>Position of the first hit to retrieve.</value>
  [JsonPropertyName("offset")]
  public int? Offset { get; set; }

  /// <summary>
  /// Number of hits to retrieve (used in combination with `offset`).
  /// </summary>
  /// <value>Number of hits to retrieve (used in combination with `offset`).</value>
  [JsonPropertyName("length")]
  public int? Length { get; set; }

  /// <summary>
  /// Coordinates for the center of a circle, expressed as a comma-separated string of latitude and longitude.  Only records included within circle around this central location are included in the results. The radius of the circle is determined by the `aroundRadius` and `minimumAroundRadius` settings. This parameter is ignored if you also specify `insidePolygon` or `insideBoundingBox`. 
  /// </summary>
  /// <value>Coordinates for the center of a circle, expressed as a comma-separated string of latitude and longitude.  Only records included within circle around this central location are included in the results. The radius of the circle is determined by the `aroundRadius` and `minimumAroundRadius` settings. This parameter is ignored if you also specify `insidePolygon` or `insideBoundingBox`. </value>
  [JsonPropertyName("aroundLatLng")]
  public string AroundLatLng { get; set; }

  /// <summary>
  /// Whether to obtain the coordinates from the request's IP address.
  /// </summary>
  /// <value>Whether to obtain the coordinates from the request's IP address.</value>
  [JsonPropertyName("aroundLatLngViaIP")]
  public bool? AroundLatLngViaIP { get; set; }

  /// <summary>
  /// Gets or Sets AroundRadius
  /// </summary>
  [JsonPropertyName("aroundRadius")]
  public AroundRadius AroundRadius { get; set; }

  /// <summary>
  /// Gets or Sets AroundPrecision
  /// </summary>
  [JsonPropertyName("aroundPrecision")]
  public AroundPrecision AroundPrecision { get; set; }

  /// <summary>
  /// Minimum radius (in meters) for a search around a location when `aroundRadius` isn't set.
  /// </summary>
  /// <value>Minimum radius (in meters) for a search around a location when `aroundRadius` isn't set.</value>
  [JsonPropertyName("minimumAroundRadius")]
  public int? MinimumAroundRadius { get; set; }

  /// <summary>
  /// Coordinates for a rectangular area in which to search.  Each bounding box is defined by the two opposite points of its diagonal, and expressed as latitude and longitude pair: `[p1 lat, p1 long, p2 lat, p2 long]`. Provide multiple bounding boxes as nested arrays. For more information, see [rectangular area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). 
  /// </summary>
  /// <value>Coordinates for a rectangular area in which to search.  Each bounding box is defined by the two opposite points of its diagonal, and expressed as latitude and longitude pair: `[p1 lat, p1 long, p2 lat, p2 long]`. Provide multiple bounding boxes as nested arrays. For more information, see [rectangular area](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). </value>
  [JsonPropertyName("insideBoundingBox")]
  public List<List<double>> InsideBoundingBox { get; set; }

  /// <summary>
  /// Coordinates of a polygon in which to search.  Polygons are defined by 3 to 10,000 points. Each point is represented by its latitude and longitude. Provide multiple polygons as nested arrays. For more information, see [filtering inside polygons](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). This parameter is ignored if you also specify `insideBoundingBox`. 
  /// </summary>
  /// <value>Coordinates of a polygon in which to search.  Polygons are defined by 3 to 10,000 points. Each point is represented by its latitude and longitude. Provide multiple polygons as nested arrays. For more information, see [filtering inside polygons](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). This parameter is ignored if you also specify `insideBoundingBox`. </value>
  [JsonPropertyName("insidePolygon")]
  public List<List<double>> InsidePolygon { get; set; }

  /// <summary>
  /// ISO language codes that adjust settings that are useful for processing natural language queries (as opposed to keyword searches):  - Sets `removeStopWords` and `ignorePlurals` to the list of provided languages. - Sets `removeWordsIfNoResults` to `allOptional`. - Adds a `natural_language` attribute to `ruleContexts` and `analyticsTags`. 
  /// </summary>
  /// <value>ISO language codes that adjust settings that are useful for processing natural language queries (as opposed to keyword searches):  - Sets `removeStopWords` and `ignorePlurals` to the list of provided languages. - Sets `removeWordsIfNoResults` to `allOptional`. - Adds a `natural_language` attribute to `ruleContexts` and `analyticsTags`. </value>
  [JsonPropertyName("naturalLanguages")]
  public List<SupportedLanguage> NaturalLanguages { get; set; }

  /// <summary>
  /// Assigns a rule context to the search query.  [Rule contexts](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/how-to/customize-search-results-by-platform/#whats-a-context) are strings that you can use to trigger matching rules. 
  /// </summary>
  /// <value>Assigns a rule context to the search query.  [Rule contexts](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/how-to/customize-search-results-by-platform/#whats-a-context) are strings that you can use to trigger matching rules. </value>
  [JsonPropertyName("ruleContexts")]
  public List<string> RuleContexts { get; set; }

  /// <summary>
  /// Impact that Personalization should have on this search.  The higher this value is, the more Personalization determines the ranking compared to other factors. For more information, see [Understanding Personalization impact](https://www.algolia.com/doc/guides/personalization/personalizing-results/in-depth/configuring-personalization/#understanding-personalization-impact). 
  /// </summary>
  /// <value>Impact that Personalization should have on this search.  The higher this value is, the more Personalization determines the ranking compared to other factors. For more information, see [Understanding Personalization impact](https://www.algolia.com/doc/guides/personalization/personalizing-results/in-depth/configuring-personalization/#understanding-personalization-impact). </value>
  [JsonPropertyName("personalizationImpact")]
  public int? PersonalizationImpact { get; set; }

  /// <summary>
  /// Unique pseudonymous or anonymous user identifier.  This helps with analytics and click and conversion events. For more information, see [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/). 
  /// </summary>
  /// <value>Unique pseudonymous or anonymous user identifier.  This helps with analytics and click and conversion events. For more information, see [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/). </value>
  [JsonPropertyName("userToken")]
  public string UserToken { get; set; }

  /// <summary>
  /// Whether the search response should include detailed ranking information.
  /// </summary>
  /// <value>Whether the search response should include detailed ranking information.</value>
  [JsonPropertyName("getRankingInfo")]
  public bool? GetRankingInfo { get; set; }

  /// <summary>
  /// Whether to take into account an index's synonyms for this search.
  /// </summary>
  /// <value>Whether to take into account an index's synonyms for this search.</value>
  [JsonPropertyName("synonyms")]
  public bool? Synonyms { get; set; }

  /// <summary>
  /// Whether to include a `queryID` attribute in the response.  The query ID is a unique identifier for a search query and is required for tracking [click and conversion events](https://www.algolia.com/guides/sending-events/getting-started/). 
  /// </summary>
  /// <value>Whether to include a `queryID` attribute in the response.  The query ID is a unique identifier for a search query and is required for tracking [click and conversion events](https://www.algolia.com/guides/sending-events/getting-started/). </value>
  [JsonPropertyName("clickAnalytics")]
  public bool? ClickAnalytics { get; set; }

  /// <summary>
  /// Whether this search will be included in Analytics.
  /// </summary>
  /// <value>Whether this search will be included in Analytics.</value>
  [JsonPropertyName("analytics")]
  public bool? Analytics { get; set; }

  /// <summary>
  /// Tags to apply to the query for [segmenting analytics data](https://www.algolia.com/doc/guides/search-analytics/guides/segments/).
  /// </summary>
  /// <value>Tags to apply to the query for [segmenting analytics data](https://www.algolia.com/doc/guides/search-analytics/guides/segments/).</value>
  [JsonPropertyName("analyticsTags")]
  public List<string> AnalyticsTags { get; set; }

  /// <summary>
  /// Whether to include this search when calculating processing-time percentiles.
  /// </summary>
  /// <value>Whether to include this search when calculating processing-time percentiles.</value>
  [JsonPropertyName("percentileComputation")]
  public bool? PercentileComputation { get; set; }

  /// <summary>
  /// Whether to enable A/B testing for this search.
  /// </summary>
  /// <value>Whether to enable A/B testing for this search.</value>
  [JsonPropertyName("enableABTest")]
  public bool? EnableABTest { get; set; }

  /// <summary>
  /// Returns the string presentation of the object
  /// </summary>
  /// <returns>String presentation of the object</returns>
  public override string ToString()
  {
    StringBuilder sb = new StringBuilder();
    sb.Append("class BaseSearchParams {\n");
    sb.Append("  Query: ").Append(Query).Append("\n");
    sb.Append("  SimilarQuery: ").Append(SimilarQuery).Append("\n");
    sb.Append("  Filters: ").Append(Filters).Append("\n");
    sb.Append("  FacetFilters: ").Append(FacetFilters).Append("\n");
    sb.Append("  OptionalFilters: ").Append(OptionalFilters).Append("\n");
    sb.Append("  NumericFilters: ").Append(NumericFilters).Append("\n");
    sb.Append("  TagFilters: ").Append(TagFilters).Append("\n");
    sb.Append("  SumOrFiltersScores: ").Append(SumOrFiltersScores).Append("\n");
    sb.Append("  RestrictSearchableAttributes: ").Append(RestrictSearchableAttributes).Append("\n");
    sb.Append("  Facets: ").Append(Facets).Append("\n");
    sb.Append("  FacetingAfterDistinct: ").Append(FacetingAfterDistinct).Append("\n");
    sb.Append("  Page: ").Append(Page).Append("\n");
    sb.Append("  Offset: ").Append(Offset).Append("\n");
    sb.Append("  Length: ").Append(Length).Append("\n");
    sb.Append("  AroundLatLng: ").Append(AroundLatLng).Append("\n");
    sb.Append("  AroundLatLngViaIP: ").Append(AroundLatLngViaIP).Append("\n");
    sb.Append("  AroundRadius: ").Append(AroundRadius).Append("\n");
    sb.Append("  AroundPrecision: ").Append(AroundPrecision).Append("\n");
    sb.Append("  MinimumAroundRadius: ").Append(MinimumAroundRadius).Append("\n");
    sb.Append("  InsideBoundingBox: ").Append(InsideBoundingBox).Append("\n");
    sb.Append("  InsidePolygon: ").Append(InsidePolygon).Append("\n");
    sb.Append("  NaturalLanguages: ").Append(NaturalLanguages).Append("\n");
    sb.Append("  RuleContexts: ").Append(RuleContexts).Append("\n");
    sb.Append("  PersonalizationImpact: ").Append(PersonalizationImpact).Append("\n");
    sb.Append("  UserToken: ").Append(UserToken).Append("\n");
    sb.Append("  GetRankingInfo: ").Append(GetRankingInfo).Append("\n");
    sb.Append("  Synonyms: ").Append(Synonyms).Append("\n");
    sb.Append("  ClickAnalytics: ").Append(ClickAnalytics).Append("\n");
    sb.Append("  Analytics: ").Append(Analytics).Append("\n");
    sb.Append("  AnalyticsTags: ").Append(AnalyticsTags).Append("\n");
    sb.Append("  PercentileComputation: ").Append(PercentileComputation).Append("\n");
    sb.Append("  EnableABTest: ").Append(EnableABTest).Append("\n");
    sb.Append("}\n");
    return sb.ToString();
  }

  /// <summary>
  /// Returns the JSON string presentation of the object
  /// </summary>
  /// <returns>JSON string presentation of the object</returns>
  public virtual string ToJson()
  {
    return JsonSerializer.Serialize(this, JsonConfig.Options);
  }

  /// <summary>
  /// Returns true if objects are equal
  /// </summary>
  /// <param name="obj">Object to be compared</param>
  /// <returns>Boolean</returns>
  public override bool Equals(object obj)
  {
    if (obj is not BaseSearchParams input)
    {
      return false;
    }

    return
        (Query == input.Query || (Query != null && Query.Equals(input.Query))) &&
        (SimilarQuery == input.SimilarQuery || (SimilarQuery != null && SimilarQuery.Equals(input.SimilarQuery))) &&
        (Filters == input.Filters || (Filters != null && Filters.Equals(input.Filters))) &&
        (FacetFilters == input.FacetFilters || (FacetFilters != null && FacetFilters.Equals(input.FacetFilters))) &&
        (OptionalFilters == input.OptionalFilters || (OptionalFilters != null && OptionalFilters.Equals(input.OptionalFilters))) &&
        (NumericFilters == input.NumericFilters || (NumericFilters != null && NumericFilters.Equals(input.NumericFilters))) &&
        (TagFilters == input.TagFilters || (TagFilters != null && TagFilters.Equals(input.TagFilters))) &&
        (SumOrFiltersScores == input.SumOrFiltersScores || SumOrFiltersScores.Equals(input.SumOrFiltersScores)) &&
        (RestrictSearchableAttributes == input.RestrictSearchableAttributes || RestrictSearchableAttributes != null && input.RestrictSearchableAttributes != null && RestrictSearchableAttributes.SequenceEqual(input.RestrictSearchableAttributes)) &&
        (Facets == input.Facets || Facets != null && input.Facets != null && Facets.SequenceEqual(input.Facets)) &&
        (FacetingAfterDistinct == input.FacetingAfterDistinct || FacetingAfterDistinct.Equals(input.FacetingAfterDistinct)) &&
        (Page == input.Page || Page.Equals(input.Page)) &&
        (Offset == input.Offset || Offset.Equals(input.Offset)) &&
        (Length == input.Length || Length.Equals(input.Length)) &&
        (AroundLatLng == input.AroundLatLng || (AroundLatLng != null && AroundLatLng.Equals(input.AroundLatLng))) &&
        (AroundLatLngViaIP == input.AroundLatLngViaIP || AroundLatLngViaIP.Equals(input.AroundLatLngViaIP)) &&
        (AroundRadius == input.AroundRadius || (AroundRadius != null && AroundRadius.Equals(input.AroundRadius))) &&
        (AroundPrecision == input.AroundPrecision || (AroundPrecision != null && AroundPrecision.Equals(input.AroundPrecision))) &&
        (MinimumAroundRadius == input.MinimumAroundRadius || MinimumAroundRadius.Equals(input.MinimumAroundRadius)) &&
        (InsideBoundingBox == input.InsideBoundingBox || InsideBoundingBox != null && input.InsideBoundingBox != null && InsideBoundingBox.SequenceEqual(input.InsideBoundingBox)) &&
        (InsidePolygon == input.InsidePolygon || InsidePolygon != null && input.InsidePolygon != null && InsidePolygon.SequenceEqual(input.InsidePolygon)) &&
        (NaturalLanguages == input.NaturalLanguages || NaturalLanguages != null && input.NaturalLanguages != null && NaturalLanguages.SequenceEqual(input.NaturalLanguages)) &&
        (RuleContexts == input.RuleContexts || RuleContexts != null && input.RuleContexts != null && RuleContexts.SequenceEqual(input.RuleContexts)) &&
        (PersonalizationImpact == input.PersonalizationImpact || PersonalizationImpact.Equals(input.PersonalizationImpact)) &&
        (UserToken == input.UserToken || (UserToken != null && UserToken.Equals(input.UserToken))) &&
        (GetRankingInfo == input.GetRankingInfo || GetRankingInfo.Equals(input.GetRankingInfo)) &&
        (Synonyms == input.Synonyms || Synonyms.Equals(input.Synonyms)) &&
        (ClickAnalytics == input.ClickAnalytics || ClickAnalytics.Equals(input.ClickAnalytics)) &&
        (Analytics == input.Analytics || Analytics.Equals(input.Analytics)) &&
        (AnalyticsTags == input.AnalyticsTags || AnalyticsTags != null && input.AnalyticsTags != null && AnalyticsTags.SequenceEqual(input.AnalyticsTags)) &&
        (PercentileComputation == input.PercentileComputation || PercentileComputation.Equals(input.PercentileComputation)) &&
        (EnableABTest == input.EnableABTest || EnableABTest.Equals(input.EnableABTest));
  }

  /// <summary>
  /// Gets the hash code
  /// </summary>
  /// <returns>Hash code</returns>
  public override int GetHashCode()
  {
    unchecked // Overflow is fine, just wrap
    {
      int hashCode = 41;
      if (Query != null)
      {
        hashCode = (hashCode * 59) + Query.GetHashCode();
      }
      if (SimilarQuery != null)
      {
        hashCode = (hashCode * 59) + SimilarQuery.GetHashCode();
      }
      if (Filters != null)
      {
        hashCode = (hashCode * 59) + Filters.GetHashCode();
      }
      if (FacetFilters != null)
      {
        hashCode = (hashCode * 59) + FacetFilters.GetHashCode();
      }
      if (OptionalFilters != null)
      {
        hashCode = (hashCode * 59) + OptionalFilters.GetHashCode();
      }
      if (NumericFilters != null)
      {
        hashCode = (hashCode * 59) + NumericFilters.GetHashCode();
      }
      if (TagFilters != null)
      {
        hashCode = (hashCode * 59) + TagFilters.GetHashCode();
      }
      hashCode = (hashCode * 59) + SumOrFiltersScores.GetHashCode();
      if (RestrictSearchableAttributes != null)
      {
        hashCode = (hashCode * 59) + RestrictSearchableAttributes.GetHashCode();
      }
      if (Facets != null)
      {
        hashCode = (hashCode * 59) + Facets.GetHashCode();
      }
      hashCode = (hashCode * 59) + FacetingAfterDistinct.GetHashCode();
      hashCode = (hashCode * 59) + Page.GetHashCode();
      hashCode = (hashCode * 59) + Offset.GetHashCode();
      hashCode = (hashCode * 59) + Length.GetHashCode();
      if (AroundLatLng != null)
      {
        hashCode = (hashCode * 59) + AroundLatLng.GetHashCode();
      }
      hashCode = (hashCode * 59) + AroundLatLngViaIP.GetHashCode();
      if (AroundRadius != null)
      {
        hashCode = (hashCode * 59) + AroundRadius.GetHashCode();
      }
      if (AroundPrecision != null)
      {
        hashCode = (hashCode * 59) + AroundPrecision.GetHashCode();
      }
      hashCode = (hashCode * 59) + MinimumAroundRadius.GetHashCode();
      if (InsideBoundingBox != null)
      {
        hashCode = (hashCode * 59) + InsideBoundingBox.GetHashCode();
      }
      if (InsidePolygon != null)
      {
        hashCode = (hashCode * 59) + InsidePolygon.GetHashCode();
      }
      if (NaturalLanguages != null)
      {
        hashCode = (hashCode * 59) + NaturalLanguages.GetHashCode();
      }
      if (RuleContexts != null)
      {
        hashCode = (hashCode * 59) + RuleContexts.GetHashCode();
      }
      hashCode = (hashCode * 59) + PersonalizationImpact.GetHashCode();
      if (UserToken != null)
      {
        hashCode = (hashCode * 59) + UserToken.GetHashCode();
      }
      hashCode = (hashCode * 59) + GetRankingInfo.GetHashCode();
      hashCode = (hashCode * 59) + Synonyms.GetHashCode();
      hashCode = (hashCode * 59) + ClickAnalytics.GetHashCode();
      hashCode = (hashCode * 59) + Analytics.GetHashCode();
      if (AnalyticsTags != null)
      {
        hashCode = (hashCode * 59) + AnalyticsTags.GetHashCode();
      }
      hashCode = (hashCode * 59) + PercentileComputation.GetHashCode();
      hashCode = (hashCode * 59) + EnableABTest.GetHashCode();
      return hashCode;
    }
  }

}

