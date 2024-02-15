// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

public struct MultipleBatchRequest: Codable, JSONEncodable, Hashable {
    public var action: Action
    /// Operation arguments (varies with specified `action`).
    public var body: AnyCodable
    /// Index to target for this operation.
    public var indexName: String

    public init(action: Action, body: AnyCodable, indexName: String) {
        self.action = action
        self.body = body
        self.indexName = indexName
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case action
        case body
        case indexName
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.action, forKey: .action)
        try container.encode(self.body, forKey: .body)
        try container.encode(self.indexName, forKey: .indexName)
    }
}
