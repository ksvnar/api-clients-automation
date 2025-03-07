<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Ingestion;

/**
 * TaskUpdate Class Doc Comment.
 *
 * @category Class
 * @description API request body for updating a task.
 */
class TaskUpdate extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'destinationID' => 'string',
        'trigger' => '\Algolia\AlgoliaSearch\Model\Ingestion\TriggerUpdateInput',
        'input' => '\Algolia\AlgoliaSearch\Model\Ingestion\TaskInput',
        'enabled' => 'bool',
        'failureThreshold' => 'int',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'destinationID' => null,
        'trigger' => null,
        'input' => null,
        'enabled' => null,
        'failureThreshold' => null,
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'destinationID' => 'destinationID',
        'trigger' => 'trigger',
        'input' => 'input',
        'enabled' => 'enabled',
        'failureThreshold' => 'failureThreshold',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static $setters = [
        'destinationID' => 'setDestinationID',
        'trigger' => 'setTrigger',
        'input' => 'setInput',
        'enabled' => 'setEnabled',
        'failureThreshold' => 'setFailureThreshold',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static $getters = [
        'destinationID' => 'getDestinationID',
        'trigger' => 'getTrigger',
        'input' => 'getInput',
        'enabled' => 'getEnabled',
        'failureThreshold' => 'getFailureThreshold',
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
        if (isset($data['destinationID'])) {
            $this->container['destinationID'] = $data['destinationID'];
        }
        if (isset($data['trigger'])) {
            $this->container['trigger'] = $data['trigger'];
        }
        if (isset($data['input'])) {
            $this->container['input'] = $data['input'];
        }
        if (isset($data['enabled'])) {
            $this->container['enabled'] = $data['enabled'];
        }
        if (isset($data['failureThreshold'])) {
            $this->container['failureThreshold'] = $data['failureThreshold'];
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

        if (isset($this->container['failureThreshold']) && ($this->container['failureThreshold'] > 100)) {
            $invalidProperties[] = "invalid value for 'failureThreshold', must be smaller than or equal to 100.";
        }

        if (isset($this->container['failureThreshold']) && ($this->container['failureThreshold'] < 0)) {
            $invalidProperties[] = "invalid value for 'failureThreshold', must be bigger than or equal to 0.";
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
     * Gets destinationID.
     *
     * @return null|string
     */
    public function getDestinationID()
    {
        return $this->container['destinationID'] ?? null;
    }

    /**
     * Sets destinationID.
     *
     * @param null|string $destinationID universally unique identifier (UUID) of a destination resource
     *
     * @return self
     */
    public function setDestinationID($destinationID)
    {
        $this->container['destinationID'] = $destinationID;

        return $this;
    }

    /**
     * Gets trigger.
     *
     * @return null|\Algolia\AlgoliaSearch\Model\Ingestion\TriggerUpdateInput
     */
    public function getTrigger()
    {
        return $this->container['trigger'] ?? null;
    }

    /**
     * Sets trigger.
     *
     * @param null|\Algolia\AlgoliaSearch\Model\Ingestion\TriggerUpdateInput $trigger trigger
     *
     * @return self
     */
    public function setTrigger($trigger)
    {
        $this->container['trigger'] = $trigger;

        return $this;
    }

    /**
     * Gets input.
     *
     * @return null|\Algolia\AlgoliaSearch\Model\Ingestion\TaskInput
     */
    public function getInput()
    {
        return $this->container['input'] ?? null;
    }

    /**
     * Sets input.
     *
     * @param null|\Algolia\AlgoliaSearch\Model\Ingestion\TaskInput $input input
     *
     * @return self
     */
    public function setInput($input)
    {
        $this->container['input'] = $input;

        return $this;
    }

    /**
     * Gets enabled.
     *
     * @return null|bool
     */
    public function getEnabled()
    {
        return $this->container['enabled'] ?? null;
    }

    /**
     * Sets enabled.
     *
     * @param null|bool $enabled whether the task is enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets failureThreshold.
     *
     * @return null|int
     */
    public function getFailureThreshold()
    {
        return $this->container['failureThreshold'] ?? null;
    }

    /**
     * Sets failureThreshold.
     *
     * @param null|int $failureThreshold maximum accepted percentage of failures for a task run to finish successfully
     *
     * @return self
     */
    public function setFailureThreshold($failureThreshold)
    {
        if (!is_null($failureThreshold) && ($failureThreshold > 100)) {
            throw new \InvalidArgumentException('invalid value for $failureThreshold when calling TaskUpdate., must be smaller than or equal to 100.');
        }
        if (!is_null($failureThreshold) && ($failureThreshold < 0)) {
            throw new \InvalidArgumentException('invalid value for $failureThreshold when calling TaskUpdate., must be bigger than or equal to 0.');
        }

        $this->container['failureThreshold'] = $failureThreshold;

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
