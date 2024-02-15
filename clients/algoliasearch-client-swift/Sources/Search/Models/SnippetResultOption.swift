// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

// MARK: - SnippetResultOption

/// Snippeted attributes show parts of the matched attributes. Only returned when attributesToSnippet is non-empty.
public struct SnippetResultOption: Codable, JSONEncodable, Hashable {
    // MARK: Lifecycle

    public init(value: String, matchLevel: MatchLevel) {
        self.value = value
        self.matchLevel = matchLevel
    }

    // MARK: Public

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case value
        case matchLevel
    }

    /// Markup text with `facetQuery` matches highlighted.
    public var value: String
    public var matchLevel: MatchLevel

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.value, forKey: .value)
        try container.encode(self.matchLevel, forKey: .matchLevel)
    }
}
