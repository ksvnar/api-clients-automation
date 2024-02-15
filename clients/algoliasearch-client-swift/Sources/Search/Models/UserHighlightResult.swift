// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

// MARK: - UserHighlightResult

public struct UserHighlightResult: Codable, JSONEncodable, Hashable {
    // MARK: Lifecycle

    public init(userID: HighlightResult, clusterName: HighlightResult) {
        self.userID = userID
        self.clusterName = clusterName
    }

    // MARK: Public

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case userID
        case clusterName
    }

    public var userID: HighlightResult
    public var clusterName: HighlightResult

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.userID, forKey: .userID)
        try container.encode(self.clusterName, forKey: .clusterName)
    }
}
