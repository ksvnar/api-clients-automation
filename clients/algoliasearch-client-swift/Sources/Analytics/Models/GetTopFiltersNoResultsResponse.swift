// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

// MARK: - GetTopFiltersNoResultsResponse

public struct GetTopFiltersNoResultsResponse: Codable, JSONEncodable, Hashable {
    // MARK: Lifecycle

    public init(values: [GetTopFiltersNoResultsValues]) {
        self.values = values
    }

    // MARK: Public

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case values
    }

    /// Filters with no results.
    public var values: [GetTopFiltersNoResultsValues]

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.values, forKey: .values)
    }
}
