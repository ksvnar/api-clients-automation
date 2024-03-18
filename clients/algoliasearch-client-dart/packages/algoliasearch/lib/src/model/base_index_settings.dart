// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/supported_language.dart';

import 'package:json_annotation/json_annotation.dart';

part 'base_index_settings.g.dart';

@JsonSerializable()
final class BaseIndexSettings {
  /// Returns a new [BaseIndexSettings] instance.
  const BaseIndexSettings({
    this.attributesForFaceting,
    this.replicas,
    this.paginationLimitedTo,
    this.unretrievableAttributes,
    this.disableTypoToleranceOnWords,
    this.attributesToTransliterate,
    this.camelCaseAttributes,
    this.decompoundedAttributes,
    this.indexLanguages,
    this.disablePrefixOnAttributes,
    this.allowCompressionOfIntegerArray,
    this.numericAttributesForFiltering,
    this.separatorsToIndex,
    this.searchableAttributes,
    this.userData,
    this.customNormalization,
    this.attributeForDistinct,
  });

  /// Attributes used for [faceting](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/).  Facets are ways to categorize search results based on attributes. Facets can be used to let user filter search results. By default, no attribute is used for faceting.  **Modifiers**  <dl> <dt><code>filterOnly(\"ATTRIBUTE\")</code></dt> <dd>Allows using this attribute as a filter, but doesn't evalue the facet values.</dd> <dt><code>searchable(\"ATTRIBUTE\")</code></dt> <dd>Allows searching for facet values.</dd> <dt><code>afterDistinct(\"ATTRIBUTE\")</code></dt> <dd>  Evaluates the facet count _after_ deduplication with `distinct`. This ensures accurate facet counts. You can apply this modifier to searchable facets: `afterDistinct(searchable(ATTRIBUTE))`.  </dd> </dl>  Without modifiers, the attribute is used as a regular facet.
  @JsonKey(name: r'attributesForFaceting')
  final List<String>? attributesForFaceting;

  /// Creates [replica indices](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/replicas/).  Replicas are copies of a primary index with the same records but different settings, synonyms, or rules. If you want to offer a different ranking or sorting of your search results, you'll use replica indices. All index operations on a primary index are automatically forwarded to its replicas. To add a replica index, you must provide the complete set of replicas to this parameter. If you omit a replica from this list, the replica turns into a regular, standalone index that will no longer by synced with the primary index.  **Modifier**  <dl> <dt><code>virtual(\"REPLICA\")</code></dt> <dd>  Create a virtual replica, Virtual replicas don't increase the number of records and are optimized for [Relevant sorting](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/relevant-sort/).  </dd> </dl>  Without modifier, a standard replica is created, which duplicates your record count and is used for strict, or [exhaustive sorting](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/exhaustive-sort/).
  @JsonKey(name: r'replicas')
  final List<String>? replicas;

  /// Maximum number of search results that can be obtained through pagination.  Higher pagination limits might slow down your search. For pagination limits above 1,000, the sorting of results beyond the 1,000th hit can't be guaranteed.
  // maximum: 20000
  @JsonKey(name: r'paginationLimitedTo')
  final int? paginationLimitedTo;

  /// Attributes that can't be retrieved at query time.  This can be useful if you want to use an attribute for ranking or to [restrict access](https://www.algolia.com/doc/guides/security/api-keys/how-to/user-restricted-access-to-data/), but don't want to include it in the search results.
  @JsonKey(name: r'unretrievableAttributes')
  final List<String>? unretrievableAttributes;

  /// Words for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/). This also turns off [word splitting and concatenation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/splitting-and-concatenation/) for the specified words.
  @JsonKey(name: r'disableTypoToleranceOnWords')
  final List<String>? disableTypoToleranceOnWords;

  /// Attributes, for which you want to support [Japanese transliteration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#japanese-transliteration-and-type-ahead).  Transliteration supports searching in any of the Japanese writing systems. To support transliteration, you must set the indexing language to Japanese.
  @JsonKey(name: r'attributesToTransliterate')
  final List<String>? attributesToTransliterate;

  /// Attributes for which to split [camel case](https://wikipedia.org/wiki/Camel_case) words.
  @JsonKey(name: r'camelCaseAttributes')
  final List<String>? camelCaseAttributes;

  /// Searchable attributes to which Algolia should apply [word segmentation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/how-to/customize-segmentation/) (decompounding).  Compound words are formed by combining two or more individual words, and are particularly prevalent in Germanic languages—for example, \"firefighter\". With decompounding, the individual components are indexed separately.  You can specify different lists for different languages. Decompounding is supported for these languages: Dutch (`nl`), German (`de`), Finnish (`fi`), Danish (`da`), Swedish (`sv`), and Norwegian (`no`).
  @JsonKey(name: r'decompoundedAttributes')
  final Object? decompoundedAttributes;

  /// Languages for language-specific processing steps, such as word detection and dictionary settings.  **You should always specify an indexing language.** If you don't specify an indexing language, the search engine uses all [supported languages](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/supported-languages/), or the languages you specified with the `ignorePlurals` or `removeStopWords` parameters. This can lead to unexpected search results. For more information, see [Language-specific configuration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/).
  @JsonKey(name: r'indexLanguages')
  final List<SupportedLanguage>? indexLanguages;

  /// Searchable attributes for which you want to turn off [prefix matching](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#adjusting-prefix-search).
  @JsonKey(name: r'disablePrefixOnAttributes')
  final List<String>? disablePrefixOnAttributes;

  /// Whether arrays with exclusively non-negative integers should be compressed for better performance. If true, the compressed arrays may be reordered.
  @JsonKey(name: r'allowCompressionOfIntegerArray')
  final bool? allowCompressionOfIntegerArray;

  /// Numeric attributes that can be used as [numerical filters](https://www.algolia.com/doc/guides/managing-results/rules/detecting-intent/how-to/applying-a-custom-filter-for-a-specific-query/#numerical-filters).  By default, all numeric attributes are available as numerical filters. For faster indexing, reduce the number of numeric attributes.  If you want to turn off filtering for all numeric attributes, specifiy an attribute that doesn't exist in your index, such as `NO_NUMERIC_FILTERING`.  **Modifier**  <dl> <dt><code>equalOnly(\"ATTRIBUTE\")</code></dt> <dd>  Support only filtering based on equality comparisons `=` and `!=`.  </dd> </dl>  Without modifier, all numeric comparisons are supported.
  @JsonKey(name: r'numericAttributesForFiltering')
  final List<String>? numericAttributesForFiltering;

  /// Controls which separators are indexed.  Separators are all non-letter characters except spaces and currency characters, such as $€£¥. By default, separator characters aren't indexed. With `separatorsToIndex`, Algolia treats separator characters as separate words. For example, a search for `C#` would report two matches.
  @JsonKey(name: r'separatorsToIndex')
  final String? separatorsToIndex;

  /// Attributes used for searching.  By default, all attributes are searchable and the [Attribute](https://www.algolia.com/doc/guides/managing-results/relevance-overview/in-depth/ranking-criteria/#attribute) ranking criterion is turned off. With a non-empty list, Algolia only returns results with matches in the selected attributes. In addition, the Attribute ranking criterion is turned on: matches in attributes that are higher in the list of `searchableAttributes` rank first. To make matches in two attributes rank equally, include them in a comma-separated string, such as `\"title,alternate_title\"`. Attributes with the same priority are always unordered.  For more information, see [Searchable attributes](https://www.algolia.com/doc/guides/sending-and-managing-data/prepare-your-data/how-to/setting-searchable-attributes/).  **Modifier**  <dl> <dt><code>unordered(\"ATTRIBUTE\")</code></dt> <dd> Ignore the position of a match within the attribute. </dd> </dl>  Without modifier, matches at the beginning of an attribute rank higer than matches at the end.
  @JsonKey(name: r'searchableAttributes')
  final List<String>? searchableAttributes;

  /// An object with custom data.  You can store up to 32&nbsp;kB as custom data.
  @JsonKey(name: r'userData')
  final Object? userData;

  /// Characters and their normalized replacements. This overrides Algolia's default [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).
  @JsonKey(name: r'customNormalization')
  final Map<String, Map<String, String>>? customNormalization;

  /// Attribute that should be used to establish groups of results.  All records with the same value for this attribute are considered a group. You can combine `attributeForDistinct` with the `distinct` search parameter to control how many items per group are included in the search results.  If you want to use the same attribute also for faceting, use the `afterDistinct` modifier of the `attributesForFaceting` setting. This applies faceting _after_ deduplication, which will result in accurate facet counts.
  @JsonKey(name: r'attributeForDistinct')
  final String? attributeForDistinct;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is BaseIndexSettings &&
          other.attributesForFaceting == attributesForFaceting &&
          other.replicas == replicas &&
          other.paginationLimitedTo == paginationLimitedTo &&
          other.unretrievableAttributes == unretrievableAttributes &&
          other.disableTypoToleranceOnWords == disableTypoToleranceOnWords &&
          other.attributesToTransliterate == attributesToTransliterate &&
          other.camelCaseAttributes == camelCaseAttributes &&
          other.decompoundedAttributes == decompoundedAttributes &&
          other.indexLanguages == indexLanguages &&
          other.disablePrefixOnAttributes == disablePrefixOnAttributes &&
          other.allowCompressionOfIntegerArray ==
              allowCompressionOfIntegerArray &&
          other.numericAttributesForFiltering ==
              numericAttributesForFiltering &&
          other.separatorsToIndex == separatorsToIndex &&
          other.searchableAttributes == searchableAttributes &&
          other.userData == userData &&
          other.customNormalization == customNormalization &&
          other.attributeForDistinct == attributeForDistinct;

  @override
  int get hashCode =>
      attributesForFaceting.hashCode +
      replicas.hashCode +
      paginationLimitedTo.hashCode +
      unretrievableAttributes.hashCode +
      disableTypoToleranceOnWords.hashCode +
      attributesToTransliterate.hashCode +
      camelCaseAttributes.hashCode +
      decompoundedAttributes.hashCode +
      indexLanguages.hashCode +
      disablePrefixOnAttributes.hashCode +
      allowCompressionOfIntegerArray.hashCode +
      numericAttributesForFiltering.hashCode +
      separatorsToIndex.hashCode +
      searchableAttributes.hashCode +
      (userData == null ? 0 : userData.hashCode) +
      customNormalization.hashCode +
      attributeForDistinct.hashCode;

  factory BaseIndexSettings.fromJson(Map<String, dynamic> json) =>
      _$BaseIndexSettingsFromJson(json);

  Map<String, dynamic> toJson() => _$BaseIndexSettingsToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
