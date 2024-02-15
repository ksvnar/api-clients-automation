// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

public struct ListSourcesResponse: Codable, JSONEncodable, Hashable {
    public var sources: [Source]
    public var pagination: Pagination

    public init(sources: [Source], pagination: Pagination) {
        self.sources = sources
        self.pagination = pagination
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case sources
        case pagination
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.sources, forKey: .sources)
        try container.encode(self.pagination, forKey: .pagination)
    }
}
