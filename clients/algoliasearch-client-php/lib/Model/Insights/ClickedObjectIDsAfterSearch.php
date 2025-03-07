<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Insights;

/**
 * ClickedObjectIDsAfterSearch Class Doc Comment.
 *
 * @category Class
 * @description Click event after an Algolia request.  Use this event to track when users click items in the search results. If you&#39;re building your category pages with Algolia, you&#39;ll also use this event.
 */
class ClickedObjectIDsAfterSearch extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'eventName' => 'string',
        'eventType' => '\Algolia\AlgoliaSearch\Model\Insights\ClickEvent',
        'index' => 'string',
        'objectIDs' => 'string[]',
        'positions' => 'int[]',
        'queryID' => 'string',
        'userToken' => 'string',
        'authenticatedUserToken' => 'string',
        'timestamp' => 'int',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'eventName' => null,
        'eventType' => null,
        'index' => null,
        'objectIDs' => null,
        'positions' => null,
        'queryID' => null,
        'userToken' => null,
        'authenticatedUserToken' => null,
        'timestamp' => 'int64',
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'eventName' => 'eventName',
        'eventType' => 'eventType',
        'index' => 'index',
        'objectIDs' => 'objectIDs',
        'positions' => 'positions',
        'queryID' => 'queryID',
        'userToken' => 'userToken',
        'authenticatedUserToken' => 'authenticatedUserToken',
        'timestamp' => 'timestamp',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'eventName' => 'setEventName',
        'eventType' => 'setEventType',
        'index' => 'setIndex',
        'objectIDs' => 'setObjectIDs',
        'positions' => 'setPositions',
        'queryID' => 'setQueryID',
        'userToken' => 'setUserToken',
        'authenticatedUserToken' => 'setAuthenticatedUserToken',
        'timestamp' => 'setTimestamp',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'eventName' => 'getEventName',
        'eventType' => 'getEventType',
        'index' => 'getIndex',
        'objectIDs' => 'getObjectIDs',
        'positions' => 'getPositions',
        'queryID' => 'getQueryID',
        'userToken' => 'getUserToken',
        'authenticatedUserToken' => 'getAuthenticatedUserToken',
        'timestamp' => 'getTimestamp',
    ];

    /**
     * Associative array for storing property values.
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor.
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['eventName'])) {
            $this->container['eventName'] = $data['eventName'];
        }
        if (isset($data['eventType'])) {
            $this->container['eventType'] = $data['eventType'];
        }
        if (isset($data['index'])) {
            $this->container['index'] = $data['index'];
        }
        if (isset($data['objectIDs'])) {
            $this->container['objectIDs'] = $data['objectIDs'];
        }
        if (isset($data['positions'])) {
            $this->container['positions'] = $data['positions'];
        }
        if (isset($data['queryID'])) {
            $this->container['queryID'] = $data['queryID'];
        }
        if (isset($data['userToken'])) {
            $this->container['userToken'] = $data['userToken'];
        }
        if (isset($data['authenticatedUserToken'])) {
            $this->container['authenticatedUserToken'] = $data['authenticatedUserToken'];
        }
        if (isset($data['timestamp'])) {
            $this->container['timestamp'] = $data['timestamp'];
        }
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if (!isset($this->container['eventName']) || null === $this->container['eventName']) {
            $invalidProperties[] = "'eventName' can't be null";
        }
        if (mb_strlen($this->container['eventName']) > 64) {
            $invalidProperties[] = "invalid value for 'eventName', the character length must be smaller than or equal to 64.";
        }

        if (mb_strlen($this->container['eventName']) < 1) {
            $invalidProperties[] = "invalid value for 'eventName', the character length must be bigger than or equal to 1.";
        }

        if (!preg_match('/[\\x20-\\x7E]{1,64}/', $this->container['eventName'])) {
            $invalidProperties[] = "invalid value for 'eventName', must be conform to the pattern /[\\x20-\\x7E]{1,64}/.";
        }

        if (!isset($this->container['eventType']) || null === $this->container['eventType']) {
            $invalidProperties[] = "'eventType' can't be null";
        }
        if (!isset($this->container['index']) || null === $this->container['index']) {
            $invalidProperties[] = "'index' can't be null";
        }
        if (!isset($this->container['objectIDs']) || null === $this->container['objectIDs']) {
            $invalidProperties[] = "'objectIDs' can't be null";
        }
        if (count($this->container['objectIDs']) > 20) {
            $invalidProperties[] = "invalid value for 'objectIDs', number of items must be less than or equal to 20.";
        }

        if (count($this->container['objectIDs']) < 1) {
            $invalidProperties[] = "invalid value for 'objectIDs', number of items must be greater than or equal to 1.";
        }

        if (!isset($this->container['positions']) || null === $this->container['positions']) {
            $invalidProperties[] = "'positions' can't be null";
        }
        if (count($this->container['positions']) > 20) {
            $invalidProperties[] = "invalid value for 'positions', number of items must be less than or equal to 20.";
        }

        if (count($this->container['positions']) < 1) {
            $invalidProperties[] = "invalid value for 'positions', number of items must be greater than or equal to 1.";
        }

        if (!isset($this->container['queryID']) || null === $this->container['queryID']) {
            $invalidProperties[] = "'queryID' can't be null";
        }
        if (mb_strlen($this->container['queryID']) > 32) {
            $invalidProperties[] = "invalid value for 'queryID', the character length must be smaller than or equal to 32.";
        }

        if (mb_strlen($this->container['queryID']) < 32) {
            $invalidProperties[] = "invalid value for 'queryID', the character length must be bigger than or equal to 32.";
        }

        if (!preg_match('/[0-9a-f]{32}/', $this->container['queryID'])) {
            $invalidProperties[] = "invalid value for 'queryID', must be conform to the pattern /[0-9a-f]{32}/.";
        }

        if (!isset($this->container['userToken']) || null === $this->container['userToken']) {
            $invalidProperties[] = "'userToken' can't be null";
        }
        if (mb_strlen($this->container['userToken']) > 129) {
            $invalidProperties[] = "invalid value for 'userToken', the character length must be smaller than or equal to 129.";
        }

        if (mb_strlen($this->container['userToken']) < 1) {
            $invalidProperties[] = "invalid value for 'userToken', the character length must be bigger than or equal to 1.";
        }

        if (!preg_match('/[a-zA-Z0-9_=\\/+-]{1,129}/', $this->container['userToken'])) {
            $invalidProperties[] = "invalid value for 'userToken', must be conform to the pattern /[a-zA-Z0-9_=\\/+-]{1,129}/.";
        }

        if (isset($this->container['authenticatedUserToken']) && (mb_strlen($this->container['authenticatedUserToken']) > 129)) {
            $invalidProperties[] = "invalid value for 'authenticatedUserToken', the character length must be smaller than or equal to 129.";
        }

        if (isset($this->container['authenticatedUserToken']) && (mb_strlen($this->container['authenticatedUserToken']) < 1)) {
            $invalidProperties[] = "invalid value for 'authenticatedUserToken', the character length must be bigger than or equal to 1.";
        }

        if (isset($this->container['authenticatedUserToken']) && !preg_match('/[a-zA-Z0-9_=\\/+-]{1,129}/', $this->container['authenticatedUserToken'])) {
            $invalidProperties[] = "invalid value for 'authenticatedUserToken', must be conform to the pattern /[a-zA-Z0-9_=\\/+-]{1,129}/.";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed.
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return 0 === count($this->listInvalidProperties());
    }

    /**
     * Gets eventName.
     *
     * @return string
     */
    public function getEventName()
    {
        return $this->container['eventName'] ?? null;
    }

    /**
     * Sets eventName.
     *
     * @param string $eventName Event name, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
     *
     * @return self
     */
    public function setEventName($eventName)
    {
        if (mb_strlen($eventName) > 64) {
            throw new \InvalidArgumentException('invalid length for $eventName when calling ClickedObjectIDsAfterSearch., must be smaller than or equal to 64.');
        }
        if (mb_strlen($eventName) < 1) {
            throw new \InvalidArgumentException('invalid length for $eventName when calling ClickedObjectIDsAfterSearch., must be bigger than or equal to 1.');
        }
        if (!preg_match('/[\\x20-\\x7E]{1,64}/', $eventName)) {
            throw new \InvalidArgumentException("invalid value for {$eventName} when calling ClickedObjectIDsAfterSearch., must conform to the pattern /[\\x20-\\x7E]{1,64}/.");
        }

        $this->container['eventName'] = $eventName;

        return $this;
    }

    /**
     * Gets eventType.
     *
     * @return \Algolia\AlgoliaSearch\Model\Insights\ClickEvent
     */
    public function getEventType()
    {
        return $this->container['eventType'] ?? null;
    }

    /**
     * Sets eventType.
     *
     * @param \Algolia\AlgoliaSearch\Model\Insights\ClickEvent $eventType eventType
     *
     * @return self
     */
    public function setEventType($eventType)
    {
        $this->container['eventType'] = $eventType;

        return $this;
    }

    /**
     * Gets index.
     *
     * @return string
     */
    public function getIndex()
    {
        return $this->container['index'] ?? null;
    }

    /**
     * Sets index.
     *
     * @param string $index index name (case-sensitive) to which the event's items belong
     *
     * @return self
     */
    public function setIndex($index)
    {
        $this->container['index'] = $index;

        return $this;
    }

    /**
     * Gets objectIDs.
     *
     * @return string[]
     */
    public function getObjectIDs()
    {
        return $this->container['objectIDs'] ?? null;
    }

    /**
     * Sets objectIDs.
     *
     * @param string[] $objectIDs object IDs of the records that are part of the event
     *
     * @return self
     */
    public function setObjectIDs($objectIDs)
    {
        if (count($objectIDs) > 20) {
            throw new \InvalidArgumentException('invalid value for $objectIDs when calling ClickedObjectIDsAfterSearch., number of items must be less than or equal to 20.');
        }
        if (count($objectIDs) < 1) {
            throw new \InvalidArgumentException('invalid length for $objectIDs when calling ClickedObjectIDsAfterSearch., number of items must be greater than or equal to 1.');
        }
        $this->container['objectIDs'] = $objectIDs;

        return $this;
    }

    /**
     * Gets positions.
     *
     * @return int[]
     */
    public function getPositions()
    {
        return $this->container['positions'] ?? null;
    }

    /**
     * Sets positions.
     *
     * @param int[] $positions Position of the clicked item the search results.  You must provide 1 `position` for each `objectID`.
     *
     * @return self
     */
    public function setPositions($positions)
    {
        if (count($positions) > 20) {
            throw new \InvalidArgumentException('invalid value for $positions when calling ClickedObjectIDsAfterSearch., number of items must be less than or equal to 20.');
        }
        if (count($positions) < 1) {
            throw new \InvalidArgumentException('invalid length for $positions when calling ClickedObjectIDsAfterSearch., number of items must be greater than or equal to 1.');
        }
        $this->container['positions'] = $positions;

        return $this;
    }

    /**
     * Gets queryID.
     *
     * @return string
     */
    public function getQueryID()
    {
        return $this->container['queryID'] ?? null;
    }

    /**
     * Sets queryID.
     *
     * @param string $queryID Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
     *
     * @return self
     */
    public function setQueryID($queryID)
    {
        if (mb_strlen($queryID) > 32) {
            throw new \InvalidArgumentException('invalid length for $queryID when calling ClickedObjectIDsAfterSearch., must be smaller than or equal to 32.');
        }
        if (mb_strlen($queryID) < 32) {
            throw new \InvalidArgumentException('invalid length for $queryID when calling ClickedObjectIDsAfterSearch., must be bigger than or equal to 32.');
        }
        if (!preg_match('/[0-9a-f]{32}/', $queryID)) {
            throw new \InvalidArgumentException("invalid value for {$queryID} when calling ClickedObjectIDsAfterSearch., must conform to the pattern /[0-9a-f]{32}/.");
        }

        $this->container['queryID'] = $queryID;

        return $this;
    }

    /**
     * Gets userToken.
     *
     * @return string
     */
    public function getUserToken()
    {
        return $this->container['userToken'] ?? null;
    }

    /**
     * Sets userToken.
     *
     * @param string $userToken Anonymous or pseudonymous user identifier.  Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
     *
     * @return self
     */
    public function setUserToken($userToken)
    {
        if (mb_strlen($userToken) > 129) {
            throw new \InvalidArgumentException('invalid length for $userToken when calling ClickedObjectIDsAfterSearch., must be smaller than or equal to 129.');
        }
        if (mb_strlen($userToken) < 1) {
            throw new \InvalidArgumentException('invalid length for $userToken when calling ClickedObjectIDsAfterSearch., must be bigger than or equal to 1.');
        }
        if (!preg_match('/[a-zA-Z0-9_=\\/+-]{1,129}/', $userToken)) {
            throw new \InvalidArgumentException("invalid value for {$userToken} when calling ClickedObjectIDsAfterSearch., must conform to the pattern /[a-zA-Z0-9_=\\/+-]{1,129}/.");
        }

        $this->container['userToken'] = $userToken;

        return $this;
    }

    /**
     * Gets authenticatedUserToken.
     *
     * @return null|string
     */
    public function getAuthenticatedUserToken()
    {
        return $this->container['authenticatedUserToken'] ?? null;
    }

    /**
     * Sets authenticatedUserToken.
     *
     * @param null|string $authenticatedUserToken Identifier for authenticated users.  When the user signs in, you can get an identifier from your system and send it as `authenticatedUserToken`. This lets you keep using the `userToken` from before the user signed in, while providing a reliable way to identify users across sessions. Don't use personally identifiable information in user tokens. For more information, see [User token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
     *
     * @return self
     */
    public function setAuthenticatedUserToken($authenticatedUserToken)
    {
        if (!is_null($authenticatedUserToken) && (mb_strlen($authenticatedUserToken) > 129)) {
            throw new \InvalidArgumentException('invalid length for $authenticatedUserToken when calling ClickedObjectIDsAfterSearch., must be smaller than or equal to 129.');
        }
        if (!is_null($authenticatedUserToken) && (mb_strlen($authenticatedUserToken) < 1)) {
            throw new \InvalidArgumentException('invalid length for $authenticatedUserToken when calling ClickedObjectIDsAfterSearch., must be bigger than or equal to 1.');
        }
        if (!is_null($authenticatedUserToken) && (!preg_match('/[a-zA-Z0-9_=\\/+-]{1,129}/', $authenticatedUserToken))) {
            throw new \InvalidArgumentException("invalid value for {$authenticatedUserToken} when calling ClickedObjectIDsAfterSearch., must conform to the pattern /[a-zA-Z0-9_=\\/+-]{1,129}/.");
        }

        $this->container['authenticatedUserToken'] = $authenticatedUserToken;

        return $this;
    }

    /**
     * Gets timestamp.
     *
     * @return null|int
     */
    public function getTimestamp()
    {
        return $this->container['timestamp'] ?? null;
    }

    /**
     * Sets timestamp.
     *
     * @param null|int $timestamp Timestamp of the event, measured in milliseconds since the Unix epoch. By default, the Insights API uses the time it receives an event as its timestamp.
     *
     * @return self
     */
    public function setTimestamp($timestamp)
    {
        $this->container['timestamp'] = $timestamp;

        return $this;
    }

    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return null|mixed
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param null|int $offset Offset
     * @param mixed    $value  Value to be set
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}
