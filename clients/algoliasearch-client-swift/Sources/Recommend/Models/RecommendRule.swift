// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import Foundation
#if canImport(Core)
    import Core
#endif

/// Recommend rule.
public struct RecommendRule: Codable, JSONEncodable {
    public var metadata: RecommendRuleMetadata?
    /// Unique identifier of a rule object.
    public var objectID: String?
    public var condition: RecommendCondition?
    public var consequence: RecommendConsequence?
    /// Description of the rule's purpose. This can be helpful for display in the Algolia dashboard.
    public var description: String?
    /// Indicates whether to enable the rule. If it isn't enabled, it isn't applied at query time.
    public var enabled: Bool?

    public init(
        metadata: RecommendRuleMetadata? = nil,
        objectID: String? = nil,
        condition: RecommendCondition? = nil,
        consequence: RecommendConsequence? = nil,
        description: String? = nil,
        enabled: Bool? = nil
    ) {
        self.metadata = metadata
        self.objectID = objectID
        self.condition = condition
        self.consequence = consequence
        self.description = description
        self.enabled = enabled
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case metadata = "_metadata"
        case objectID
        case condition
        case consequence
        case description
        case enabled
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(self.metadata, forKey: .metadata)
        try container.encodeIfPresent(self.objectID, forKey: .objectID)
        try container.encodeIfPresent(self.condition, forKey: .condition)
        try container.encodeIfPresent(self.consequence, forKey: .consequence)
        try container.encodeIfPresent(self.description, forKey: .description)
        try container.encodeIfPresent(self.enabled, forKey: .enabled)
    }
}

extension RecommendRule: Equatable {
    public static func ==(lhs: RecommendRule, rhs: RecommendRule) -> Bool {
        lhs.metadata == rhs.metadata &&
            lhs.objectID == rhs.objectID &&
            lhs.condition == rhs.condition &&
            lhs.consequence == rhs.consequence &&
            lhs.description == rhs.description &&
            lhs.enabled == rhs.enabled
    }
}

extension RecommendRule: Hashable {
    public func hash(into hasher: inout Hasher) {
        hasher.combine(self.metadata?.hashValue)
        hasher.combine(self.objectID?.hashValue)
        hasher.combine(self.condition?.hashValue)
        hasher.combine(self.consequence?.hashValue)
        hasher.combine(self.description?.hashValue)
        hasher.combine(self.enabled?.hashValue)
    }
}
