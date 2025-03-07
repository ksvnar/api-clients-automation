// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

public struct SourceUpdateResponse: Codable, JSONEncodable {
    /// Universally uniqud identifier (UUID) of a source.
    public var sourceID: String
    /// Descriptive name of the source.
    public var name: String
    /// Date of last update in RFC3339 format.
    public var updatedAt: String

    public init(sourceID: String, name: String, updatedAt: String) {
        self.sourceID = sourceID
        self.name = name
        self.updatedAt = updatedAt
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case sourceID
        case name
        case updatedAt
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.sourceID, forKey: .sourceID)
        try container.encode(self.name, forKey: .name)
        try container.encode(self.updatedAt, forKey: .updatedAt)
    }
}

extension SourceUpdateResponse: Equatable {
    public static func ==(lhs: SourceUpdateResponse, rhs: SourceUpdateResponse) -> Bool {
        lhs.sourceID == rhs.sourceID &&
            lhs.name == rhs.name &&
            lhs.updatedAt == rhs.updatedAt
    }
}

extension SourceUpdateResponse: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.sourceID.hashValue)
        hasher.combine(self.name.hashValue)
        hasher.combine(self.updatedAt.hashValue)
    }
}
