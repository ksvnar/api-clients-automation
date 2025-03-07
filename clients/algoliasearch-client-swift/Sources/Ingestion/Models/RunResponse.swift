// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

/// API response for running a task.
public struct RunResponse: Codable, JSONEncodable {
    /// Universally unique identifier (UUID) of a task run.
    public var runID: String
    /// Date of creation in RFC3339 format.
    public var createdAt: String

    public init(runID: String, createdAt: String) {
        self.runID = runID
        self.createdAt = createdAt
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case runID
        case createdAt
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(self.runID, forKey: .runID)
        try container.encode(self.createdAt, forKey: .createdAt)
    }
}

extension RunResponse: Equatable {
    public static func ==(lhs: RunResponse, rhs: RunResponse) -> Bool {
        lhs.runID == rhs.runID &&
            lhs.createdAt == rhs.createdAt
    }
}

extension RunResponse: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.runID.hashValue)
        hasher.combine(self.createdAt.hashValue)
    }
}
