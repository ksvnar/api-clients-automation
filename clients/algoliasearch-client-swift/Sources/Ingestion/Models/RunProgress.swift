// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Core
import Foundation
#if canImport(AnyCodable)
    import AnyCodable
#endif

// MARK: - RunProgress

public struct RunProgress: Codable, JSONEncodable, Hashable {
    // MARK: Lifecycle

    public init(expectedNbOfEvents: Int? = nil, receivedNbOfEvents: Int? = nil) {
        self.expectedNbOfEvents = expectedNbOfEvents
        self.receivedNbOfEvents = receivedNbOfEvents
    }

    // MARK: Public

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case expectedNbOfEvents
        case receivedNbOfEvents
    }

    public var expectedNbOfEvents: Int?
    public var receivedNbOfEvents: Int?

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(self.expectedNbOfEvents, forKey: .expectedNbOfEvents)
        try container.encodeIfPresent(self.receivedNbOfEvents, forKey: .receivedNbOfEvents)
    }
}
