// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

// MARK: - PersonalizationStrategyParams

public struct PersonalizationStrategyParams: Codable, JSONEncodable, Hashable {
    // MARK: Lifecycle

    public init(eventScoring: [EventScoring], facetScoring: [FacetScoring], personalizationImpact: Int) {
        self.eventScoring = eventScoring
        self.facetScoring = facetScoring
        self.personalizationImpact = personalizationImpact
    }

    // MARK: Public

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case eventScoring
        case facetScoring
        case personalizationImpact
    }

    /// Scores associated with the events.
    public var eventScoring: [EventScoring]
    /// Scores associated with the facets.
    public var facetScoring: [FacetScoring]
    /// The impact that personalization has on search results: a number between 0 (personalization disabled) and 100
    /// (personalization fully enabled).
    public var personalizationImpact: Int

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.eventScoring, forKey: .eventScoring)
        try container.encode(self.facetScoring, forKey: .facetScoring)
        try container.encode(self.personalizationImpact, forKey: .personalizationImpact)
    }
}
