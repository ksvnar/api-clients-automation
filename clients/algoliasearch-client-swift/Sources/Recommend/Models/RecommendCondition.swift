// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

/// Condition that triggers the rule. If not specified, the rule is triggered for all recommendations.
public struct RecommendCondition: Codable, JSONEncodable {
    /// Filter expression to only include items that match the filter criteria in the response.  You can use these
    /// filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`,
    /// `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and
    /// upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet
    /// attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>`
    /// (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and
    /// `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.
    /// **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not
    /// supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not
    /// supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet
    /// attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an
    /// array, the filter matches if it matches at least one element of the array.  For more information, see
    /// [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/).
    public var filters: String?
    /// An additional restriction that only triggers the rule, when the search has the same value as `ruleContexts`
    /// parameter. For example, if `context: mobile`, the rule is only triggered when the search request has a matching
    /// `ruleContexts: mobile`. A rule context must only contain alphanumeric characters.
    public var context: String?

    public init(filters: String? = nil, context: String? = nil) {
        self.filters = filters
        self.context = context
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case filters
        case context
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(self.filters, forKey: .filters)
        try container.encodeIfPresent(self.context, forKey: .context)
    }
}

extension RecommendCondition: Equatable {
    public static func ==(lhs: RecommendCondition, rhs: RecommendCondition) -> Bool {
        lhs.filters == rhs.filters &&
            lhs.context == rhs.context
    }
}

extension RecommendCondition: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.filters?.hashValue)
        hasher.combine(self.context?.hashValue)
    }
}
