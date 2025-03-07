package com.algolia.codegen.cts.tests;

import com.algolia.codegen.AlgoliaSwiftGenerator;
import com.algolia.codegen.exceptions.*;
import com.algolia.codegen.utils.*;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;
import io.swagger.util.Json;
import java.util.*;
import java.util.Map.Entry;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import org.apache.commons.lang3.StringUtils;
import org.openapitools.codegen.*;

@SuppressWarnings("unchecked")
public class ParametersWithDataType {

  private final Map<String, CodegenModel> models;
  private final String language;
  private final String client;

  public ParametersWithDataType(Map<String, CodegenModel> models, String language, String client) {
    this.models = models;
    this.language = language;
    this.client = client;
  }

  public void enhanceParameters(Map<String, Object> parameters, Map<String, Object> bundle)
    throws CTSException, JsonMappingException, JsonProcessingException {
    this.enhanceParameters(parameters, bundle, null);
  }

  /**
   * @param parameters The object to traverse and annotate with type
   * @param bundle The output object
   * @param operation (optional) The model in which to look for spec
   */
  public void enhanceParameters(Map<String, Object> parameters, Map<String, Object> bundle, CodegenOperation operation)
    throws CTSException, JsonMappingException, JsonProcessingException {
    if (parameters != null && parameters.size() == 0) {
      bundle.put("parameters", "{}");
      return;
    }
    IJsonSchemaValidationProperties spec = null;
    String paramName = null;
    // special case if there is only bodyParam which is not an array
    if (operation != null && operation.allParams.size() == 1 && operation.bodyParams.size() == 1 && !operation.bodyParam.isArray) {
      spec = operation.bodyParam;
      paramName = operation.bodyParam.paramName;
    }

    List<Map<String, Object>> parametersWithDataType = new ArrayList<>();
    Map<String, Object> parametersWithDataTypeMap = new HashMap<>();

    if (paramName == null) {
      if (parameters != null) {
        for (Entry<String, Object> param : parameters.entrySet()) {
          CodegenParameter specParam = null;
          if (operation != null) {
            for (CodegenParameter sp : operation.allParams) {
              if (sp.paramName.equals(param.getKey())) {
                specParam = sp;
                break;
              }
            }
            if (specParam == null) {
              throw new CTSException("Parameter " + param.getKey() + " not found in the root parameter");
            }
          }
          Map<String, Object> paramWithType = traverseParams(param.getKey(), param.getValue(), specParam, "", 0, false);
          parametersWithDataType.add(paramWithType);
          parametersWithDataTypeMap.put((String) paramWithType.get("key"), paramWithType);
        }
      }
    } else {
      Map<String, Object> paramWithType = traverseParams(paramName, parameters, spec, "", 0, false);
      parametersWithDataType.add(paramWithType);
      parametersWithDataTypeMap.put((String) paramWithType.get("key"), paramWithType);
    }

    // Give the stringified version to mustache, for js
    bundle.put("parameters", Json.mapper().writeValueAsString(parameters));
    bundle.put("parametersWithDataType", parametersWithDataType);
    // Also provide a map version for those who know which keys to look for
    bundle.put("parametersWithDataTypeMap", parametersWithDataTypeMap);
  }

  private Map<String, Object> traverseParams(
    String paramName,
    Object param,
    IJsonSchemaValidationProperties spec,
    String parent,
    int suffix,
    boolean isParentFreeFormObject
  ) throws CTSException {
    if (spec == null) {
      return traverseParamsWithoutSpec(paramName, param, parent, suffix);
    }
    String baseType = getTypeName(spec);
    if (baseType == null) {
      throw new CTSException("Cannot determine type of " + paramName + " (value: " + param + ")");
    }

    boolean isCodegenModel = spec instanceof CodegenModel;

    Map<String, Object> testOutput = createDefaultOutput();

    if (spec instanceof CodegenParameter parameter) {
      testOutput.put("required", parameter.required);
    } else if (spec instanceof CodegenProperty property) {
      testOutput.put("required", property.required);
    }

    if (!isCodegenModel) {
      // don't overwrite it if it's already a model sometimes it's in lowercase for some reason
      String lowerBaseType = baseType.substring(0, 1).toLowerCase() + baseType.substring(1);
      if (models.containsKey(baseType)) {
        // get the real model if possible
        spec = models.get(baseType);
      } else if (models.containsKey(lowerBaseType)) {
        spec = models.get(lowerBaseType);
      }
      isCodegenModel = spec instanceof CodegenModel;
    }

    String finalParamName = getFinalParamName(paramName);

    testOutput.put("key", finalParamName);
    testOutput.put("isKeyAllUpperCase", StringUtils.isAllUpperCase(finalParamName));
    testOutput.put("parentSuffix", suffix - 1);
    testOutput.put("useAnonymousKey", !finalParamName.matches("(.*)_[0-9]$") && suffix != 0);
    testOutput.put("suffix", suffix);
    testOutput.put("parent", parent);
    testOutput.put("isRoot", "".equals(parent));
    testOutput.put("objectName", getObjectNameForLanguage(baseType));
    testOutput.put("isParentFreeFormObject", isParentFreeFormObject);

    if (param == null) {
      handleNull(spec, testOutput);
    } else if (spec.getIsArray()) {
      handleArray(paramName, param, testOutput, spec, suffix);
    } else if (spec.getIsEnum()) {
      handleEnum(param, testOutput);
    } else if (spec.getIsModel() || isCodegenModel) {
      // recursive object
      handleModel(paramName, param, testOutput, spec, baseType, parent, suffix, isParentFreeFormObject);
    } else if (baseType.equals("Object")) {
      // not var, no item, pure free form
      handleObject(paramName, param, testOutput, true, suffix);
    } else if (spec.getIsMap()) {
      // free key but only one type
      handleMap(paramName, param, testOutput, spec, suffix);
    } else {
      handlePrimitive(param, testOutput, spec);
    }

    return testOutput;
  }

  /** Same method but with inference only */
  private Map<String, Object> traverseParamsWithoutSpec(String paramName, Object param, String parent, int suffix) throws CTSException {
    String finalParamName = getFinalParamName(paramName);

    Map<String, Object> testOutput = createDefaultOutput();
    testOutput.put("key", finalParamName);
    testOutput.put("parentSuffix", suffix - 1);
    testOutput.put("useAnonymousKey", !finalParamName.matches("(.*)_[0-9]$") && suffix != 0);
    testOutput.put("suffix", suffix);
    testOutput.put("parent", parent);
    testOutput.put("isRoot", "".equals(parent));
    // try to infer the type
    try {
      String dataType = inferDataType(param, null, testOutput);
      testOutput.put("objectName", getObjectNameForLanguage(dataType));
    } catch (CTSException e) {
      // ignore the error
    }

    if (param == null) {
      handleNull(null, testOutput);
    } else if (param instanceof List) {
      handleArray(paramName, param, testOutput, null, suffix);
    } else if (param instanceof Map) {
      handleObject(paramName, param, testOutput, false, suffix);
    } else {
      handlePrimitive(param, testOutput, null);
    }
    return testOutput;
  }

  private String getFinalParamName(String paramName) {
    switch (language) {
      case "java":
        return paramName.startsWith("_") ? paramName.substring(1) : paramName;
      case "go":
        return paramName.equals("type") ? "type_" : paramName;
      case "swift":
        return AlgoliaSwiftGenerator.removeReservedModelNamePrefix(paramName, client);
    }

    return paramName;
  }

  private Map<String, Object> createDefaultOutput() {
    Map<String, Object> testOutput = new HashMap<>();

    // we need to set all types to false otherwise mustache will read the one from the parent
    // context and run into a infinite loop
    testOutput.put("isObject", false);
    testOutput.put("isNullObject", false);
    testOutput.put("isArray", false);
    testOutput.put("isNull", false);
    testOutput.put("isFreeFormObject", false);
    testOutput.put("isParentFreeFormObject", false);
    testOutput.put("isAnyType", false);
    testOutput.put("isString", false);
    testOutput.put("isInteger", false);
    testOutput.put("isLong", false);
    testOutput.put("isDouble", false);
    testOutput.put("isNumber", false);
    testOutput.put("isBoolean", false);
    testOutput.put("isEnum", false);
    testOutput.put("isSimpleObject", false);
    testOutput.put("oneOfModel", false);
    testOutput.put("isAdditionalProperty", false);

    return testOutput;
  }

  private void handleNull(IJsonSchemaValidationProperties spec, Map<String, Object> testOutput) {
    testOutput.put("isNull", true);
    if (spec.getIsModel() || spec instanceof CodegenModel) {
      testOutput.put("isNullObject", true);
    }
  }

  private void handleArray(
    String paramName,
    Object param,
    Map<String, Object> testOutput,
    IJsonSchemaValidationProperties spec,
    int suffix
  ) throws CTSException {
    List<Object> items = (List<Object>) param;

    List<Object> values = new ArrayList<>();
    for (int i = 0; i < items.size(); i++) {
      values.add(traverseParams(paramName + "_" + i, items.get(i), spec == null ? null : spec.getItems(), paramName, suffix + 1, false));
    }

    testOutput.put("isArray", true);
    testOutput.put("value", values);
  }

  private void handleEnum(Object param, Map<String, Object> testOutput) {
    testOutput.put("isEnum", true);
    testOutput.put("value", param);

    if (language.equals("swift")) {
      testOutput.put("valueEscaped", ((String) param).contains("-") ? Helpers.camelize((String) param) : param);
    } else {
      testOutput.put("valueEscaped", param + "_");
    }
  }

  private void handleModel(
    String paramName,
    Object param,
    Map<String, Object> testOutput,
    IJsonSchemaValidationProperties spec,
    String baseType,
    String parent,
    int suffix,
    boolean isParentFreeFormObject
  ) throws CTSException {
    if (!spec.getHasVars()) {
      // In this case we might have a complex `allOf`, we will first check if it exists
      CodegenComposedSchemas composedSchemas = spec.getComposedSchemas();

      if (composedSchemas != null) {
        List<CodegenProperty> allOf = composedSchemas.getAllOf();

        if (allOf != null && !allOf.isEmpty()) {
          traverseParams(paramName, param, allOf.get(0), parent, suffix, false);

          return;
        }
      }
      // We only throw if there is no `composedSchemas`, because `oneOf` can also be handled below
      else {
        throw new CTSException("Spec has no vars.");
      }
    }

    if (spec.getItems() != null) {
      throw new CTSException("Spec has items.");
    }

    if (spec instanceof CodegenModel && ((CodegenModel) spec).oneOf.size() > 0) {
      // find a discriminator to handle oneOf
      CodegenModel model = (CodegenModel) spec;
      IJsonSchemaValidationProperties match = findMatchingOneOf(param, model);

      paramName = getTransformedParamName(paramName);

      testOutput.putAll(traverseParams(paramName, param, match, parent, suffix, isParentFreeFormObject));

      Map<String, Object> oneOfModel = new HashMap<>();
      IJsonSchemaValidationProperties current = match;
      String typeName = getTypeName(current);
      boolean isList = false;
      if (language.equals("go") || language.equals("swift")) {
        typeName = Helpers.capitalize(getObjectNameForLanguage(typeName));
        while (current.getItems() != null) {
          current = current.getItems();
          typeName += "Of" + Helpers.capitalize(getObjectNameForLanguage(getTypeName(current)));
          isList = true;
        }
      } else {
        while (current.getItems() != null) {
          current = current.getItems();
          typeName += "Of" + getTypeName(current);
          isList = true;
        }
      }

      boolean useExplicitName;
      CodegenComposedSchemas composedSchemas = model.getComposedSchemas();
      if (composedSchemas != null && composedSchemas.getOneOf() != null && !composedSchemas.getOneOf().isEmpty()) {
        useExplicitName = Helpers.shouldUseExplicitOneOfName(composedSchemas.getOneOf().stream().map(this::getTypeName).toList());
      } else {
        useExplicitName = Helpers.shouldUseExplicitOneOfName(model.oneOf);
      }

      typeName = getTransformedParamName(typeName);
      baseType = getTransformedParamName(baseType);

      oneOfModel.put("parentClassName", Helpers.capitalize(baseType));
      oneOfModel.put("type", typeName);
      oneOfModel.put("x-one-of-explicit-name", useExplicitName);
      oneOfModel.put("hasWrapper", isList || isString(current) || current.getIsNumber() || current.getIsBoolean());
      testOutput.put("oneOfModel", oneOfModel);
      return;
    }

    Map<String, Object> vars = (Map<String, Object>) param;
    List<Map<String, Object>> values = new ArrayList<>();
    List<Map<String, Object>> additionalPropertyValues = new ArrayList<>();
    for (Entry<String, Object> entry : vars.entrySet()) {
      IJsonSchemaValidationProperties varSpec = null;
      for (CodegenProperty vs : spec.getVars()) {
        if (vs.baseName.equals(entry.getKey())) {
          varSpec = vs;
          break;
        }
      }
      if (varSpec == null) {
        if (spec.getAdditionalPropertiesIsAnyType()) {
          // we hit an additionalProperties, infer it's type
          CodegenParameter additionalPropertiesSpec = new CodegenParameter();
          additionalPropertiesSpec.dataType = inferDataType(entry.getValue(), additionalPropertiesSpec, null);
          Map<String, Object> value = traverseParams(
            entry.getKey(),
            entry.getValue(),
            additionalPropertiesSpec,
            paramName,
            suffix + 1,
            false
          );
          value.put("isAdditionalProperty", true);
          values.add(value);
        } else {
          throw new CTSException(
            "Parameter '" +
            entry.getKey() +
            "' not found in '" +
            paramName +
            "'. Available properties are: " +
            spec.getVars().stream().map(v -> v.baseName).collect(Collectors.joining(", ")) +
            (spec.getAdditionalPropertiesIsAnyType() ? " (and any additional properties)" : "") +
            ". Or you might have a type conflict in the spec for '" +
            baseType +
            "'"
          );
        }
      } else {
        Map<String, Object> transformedParam = traverseParams(entry.getKey(), entry.getValue(), varSpec, paramName, suffix + 1, false);
        values.add(transformedParam);
      }
    }

    if (language.equals("swift")) {
      // Store ordered params from the spec
      var orderedParams = spec.getVars().stream().map(v -> v.baseName).toList();

      // Create a map to store the indices of each string in orderedParams
      Map<String, Integer> indexMap = IntStream
        .range(0, orderedParams.size())
        .boxed()
        .collect(Collectors.toMap(orderedParams::get, i -> i));

      values.sort(Comparator.comparing(value -> indexMap.getOrDefault((String) value.get("key"), Integer.MAX_VALUE)));
    }

    var hasAdditionalProperties = values
      .stream()
      .anyMatch(el -> el.containsKey("isAdditionalProperty") && (boolean) el.get("isAdditionalProperty"));
    testOutput.put("hasAdditionalProperties", hasAdditionalProperties);

    testOutput.put("isObject", true);
    testOutput.put("value", values);
    testOutput.put("additionalProperties", additionalPropertyValues);
  }

  private String getTransformedParamName(String paramName) {
    if (language.equals("swift")) {
      paramName = AlgoliaSwiftGenerator.prefixReservedModelName(paramName, client);
    }
    return paramName;
  }

  private void handleObject(String paramName, Object param, Map<String, Object> testOutput, boolean isSimpleObject, int suffix)
    throws CTSException {
    Map<String, Object> vars = (Map<String, Object>) param;

    List<Object> values = new ArrayList<>();
    for (Entry<String, Object> entry : vars.entrySet()) {
      CodegenParameter objSpec = new CodegenParameter();
      objSpec.dataType = inferDataType(entry.getValue(), objSpec, null);
      values.add(traverseParams(entry.getKey(), entry.getValue(), objSpec, paramName, suffix + 1, true));
    }

    testOutput.put("isSimpleObject", isSimpleObject);
    testOutput.put("isFreeFormObject", true);
    testOutput.put("value", values);
  }

  private void handleMap(String paramName, Object param, Map<String, Object> testOutput, IJsonSchemaValidationProperties spec, int suffix)
    throws CTSException {
    if (spec.getHasVars()) {
      throw new CTSException("Spec has vars.");
    }

    Map<String, Object> vars = (Map<String, Object>) param;

    List<Object> values = new ArrayList<>();

    CodegenProperty items = spec.getItems();

    for (Entry<String, Object> entry : vars.entrySet()) {
      IJsonSchemaValidationProperties itemType = items;

      // The generator consider a free form object type as an `object`, which
      // is wrong in our case, so we infer it to explore the right path in the traverseParams
      // function, but we keep the any type for the CTS.
      if (items == null || (items.openApiType.equals("object") && items.isFreeFormObject)) {
        CodegenParameter maybeMatch = new CodegenParameter();
        String paramType = inferDataType(entry.getValue(), maybeMatch, null);

        maybeMatch.dataType = paramType;
        maybeMatch.isAnyType = true;
        itemType = maybeMatch;
      }

      values.add(traverseParams(entry.getKey(), entry.getValue(), itemType, paramName, suffix + 1, true));
    }

    testOutput.put("isFreeFormObject", true);
    testOutput.put("value", values);
  }

  private void handlePrimitive(Object param, Map<String, Object> testOutput, IJsonSchemaValidationProperties spec) throws CTSException {
    if (spec != null && spec.getIsPrimitiveType()) {
      transferPrimitiveData(spec, testOutput);
    } else {
      inferDataType(param, null, testOutput);
      if (spec != null && spec.getIsAnyType()) {
        testOutput.put("isAnyType", true);
      }
    }
    testOutput.put("value", param);
  }

  private String getTypeName(IJsonSchemaValidationProperties param) {
    if (param instanceof CodegenModel parameter) {
      return parameter.classname;
    }

    return param.getDataType();
  }

  private boolean isString(IJsonSchemaValidationProperties param) {
    if (param instanceof CodegenParameter parameter) {
      return parameter.isString;
    }
    if (param instanceof CodegenProperty parameter) {
      return parameter.isString;
    }
    return false;
  }

  private String getObjectNameForLanguage(String objectName) {
    switch (language) {
      case "csharp":
        switch (objectName) {
          case "Map":
            return "Dictionary";
          case "Integer":
            return "int";
          case "String":
            return "string";
        }
        break;
      case "go":
        switch (objectName) {
          case "String":
            return "string";
          case "Double":
            return "float64";
          case "Integer":
            return "int32";
          case "Long":
            return "int64";
          case "Boolean":
            return "bool";
          case "List":
            return "Array";
          case "Object":
            return "map[string]any";
        }
        break;
      case "swift":
        switch (objectName) {
          case "List":
            return "array";
          default:
            return Helpers.capitalize(getTransformedParamName(objectName));
        }
    }

    return Helpers.capitalize(objectName);
  }

  private String inferDataType(Object param, CodegenParameter spec, Map<String, Object> output) throws CTSException {
    switch (param.getClass().getSimpleName()) {
      case "String":
        if (spec != null) spec.setIsString(true);
        if (output != null) output.put("isString", true);
        return "String";
      case "Integer":
        if (spec != null) spec.setIsNumber(true);
        if (output != null) {
          output.put("isInteger", true);
          output.put("isNumber", true);
        }
        return "Integer";
      case "Long":
        if (spec != null) spec.setIsNumber(true);
        if (output != null) {
          output.put("isLong", true);
          output.put("isNumber", true);
        }
        return "Long";
      case "Double":
        if (spec != null) spec.setIsNumber(true);
        if (output != null) {
          output.put("isDouble", true);
          output.put("isNumber", true);
        }
        return "Double";
      case "Boolean":
        if (spec != null) spec.setIsBoolean(true);
        if (output != null) output.put("isBoolean", true);
        return "Boolean";
      case "ArrayList":
        if (spec != null) {
          spec.setIsArray(true);
          // This is just to find the correct path in `handlePrimitive`, but it's not always the
          // real type
          CodegenProperty baseItems = new CodegenProperty();
          baseItems.dataType = "String";
          spec.setItems(baseItems);
        }
        if (output != null) output.put("isArray", true);
        return "List";
      case "LinkedHashMap":
        if (spec != null) spec.baseType = "Object";
        if (output != null) output.put("isFreeFormObject", true);
        return "Object";
      default:
        throw new CTSException("Unknown type: " + param.getClass().getSimpleName());
    }
  }

  private void transferPrimitiveData(IJsonSchemaValidationProperties spec, Map<String, Object> output) throws CTSException {
    switch (getTypeName(spec)) {
      case "String":
        output.put("isString", true);
        break;
      case "UUID":
        output.put("isString", true);
        break;
      case "Integer":
        output.put("isInteger", true);
        output.put("isNumber", true);
        break;
      case "Long":
        output.put("isLong", true);
        output.put("isNumber", true);
        break;
      case "Double":
        output.put("isDouble", true);
        output.put("isNumber", true);
        break;
      case "Boolean":
        output.put("isBoolean", true);
        break;
      default:
        throw new CTSException("Unknown primitive: " + getTypeName(spec));
    }
  }

  private IJsonSchemaValidationProperties findMatchingOneOf(Object param, CodegenModel model) throws CTSException {
    if (param instanceof Map) {
      // for object, check which has the most of property in common
      int maxCount = 0;
      CodegenModel bestOneOf = model.interfaceModels.get(0);
      oneOfLoop:for (CodegenModel oneOf : model.interfaceModels) {
        if (oneOf.vars.size() == 0) {
          continue;
        }

        Map<String, Object> map = (Map<String, Object>) param;

        // if a required property is not in param, it's not a match
        for (CodegenProperty prop : oneOf.requiredVars) {
          if (!map.containsKey(prop.baseName)) {
            continue oneOfLoop;
          }
        }

        int commonCount = 0;
        for (String prop : map.keySet()) {
          for (CodegenProperty propOneOf : oneOf.vars) {
            if (prop.equals(propOneOf.baseName) && couldMatchEnum(map.get(prop), propOneOf)) {
              commonCount++;
            }
          }
        }
        if (commonCount > maxCount) {
          maxCount = commonCount;
          bestOneOf = oneOf;
        }
      }
      return bestOneOf;
    }
    if (param instanceof List) {
      CodegenComposedSchemas composedSchemas = model.getComposedSchemas();

      if (composedSchemas != null) {
        List<CodegenProperty> oneOf = composedSchemas.getOneOf();

        // Somehow this is not yet enough
        if (oneOf != null && !oneOf.isEmpty()) {
          System.out.println("Choosing the first oneOf by default: " + oneOf.get(0).baseName + " (this won't stay correct forever)");
          return oneOf.get(0);
        }
      }

      return null;
    }

    // find the correct enum
    if (param instanceof String) {
      for (CodegenProperty prop : model.getComposedSchemas().getOneOf()) {
        if (prop.getIsEnumOrRef() && couldMatchEnum(param, prop)) {
          return prop;
        }
      }
    }

    CodegenParameter maybeMatch = new CodegenParameter();
    String paramType = inferDataType(param, maybeMatch, null);
    maybeMatch.dataType = paramType;

    boolean hasFloat = false;
    for (String oneOfName : model.oneOf) {
      if (oneOfName.equals(paramType)) {
        return maybeMatch;
      }
      if (oneOfName.equals("Float") || oneOfName.equals("Double")) {
        hasFloat = true;
      }
    }

    // If there is a number, try to use it as other number type, in the order
    // Integer, Long, Float, Double
    if (hasFloat && (paramType.equals("Integer") || paramType.equals("Long") || paramType.equals("Double"))) {
      return maybeMatch;
    }

    for (CodegenModel oneOf : model.interfaceModels) {
      // Somehow the dataType can be in lower case?
      if (oneOf.dataType.toLowerCase().equals(paramType.toLowerCase())) {
        return oneOf;
      }
    }
    return null;
  }

  // If the model is an enum and contains a valid list of allowed values,
  // it will check that 'value' is in the list.
  // Otherwise return true by default to avoid false negative.
  private boolean couldMatchEnum(Object value, CodegenProperty model) {
    if (!model.getIsEnumOrRef() || model.allowableValues == null || !model.allowableValues.containsKey("values")) return true;
    Object values = model.allowableValues.get("values");
    if (!(values instanceof List)) return true;

    return ((List) values).contains(value);
  }
}
