import type {
  Headers,
  Host,
  QueryParameters,
  Request,
  RequestOptions,
  Response,
  StackFrame,
} from '../types';

import { ApiError, DeserializationError, DetailedApiError } from './errors';

export function shuffle<TData>(array: TData[]): TData[] {
  const shuffledArray = array;

  for (let c = array.length - 1; c > 0; c--) {
    const b = Math.floor(Math.random() * (c + 1));
    const a = array[c];

    shuffledArray[c] = array[b];
    shuffledArray[b] = a;
  }

  return shuffledArray;
}

export function serializeUrl(
  host: Host,
  path: string,
  queryParameters: QueryParameters
): string {
  const queryParametersAsString = serializeQueryParameters(queryParameters);
  let url = `${host.protocol}://${host.url}${host.port ? `:${host.port}` : ''}/${
    path.charAt(0) === '/' ? path.substring(1) : path
  }`;

  if (queryParametersAsString.length) {
    url += `?${queryParametersAsString}`;
  }

  return url;
}

export function serializeQueryParameters(parameters: QueryParameters): string {
  const isObjectOrArray = (value: any): boolean =>
    Object.prototype.toString.call(value) === '[object Object]' ||
    Object.prototype.toString.call(value) === '[object Array]';

  return Object.keys(parameters)
    .map(
      (key) =>
        `${key}=${encodeURIComponent(
          isObjectOrArray(parameters[key])
            ? JSON.stringify(parameters[key])
            : parameters[key]
        ).replaceAll('+', '%20')}`
    )
    .join('&');
}

export function serializeData(
  request: Request,
  requestOptions: RequestOptions
): string | undefined {
  if (
    request.method === 'GET' ||
    (request.data === undefined && requestOptions.data === undefined)
  ) {
    return undefined;
  }

  const data = Array.isArray(request.data)
    ? request.data
    : { ...request.data, ...requestOptions.data };

  return JSON.stringify(data);
}

export function serializeHeaders(
  baseHeaders: Headers,
  requestHeaders: Headers,
  requestOptionsHeaders?: Headers
): Headers {
  const headers: Headers = {
    Accept: 'application/json',
    ...baseHeaders,
    ...requestHeaders,
    ...requestOptionsHeaders,
  };
  const serializedHeaders: Headers = {};

  Object.keys(headers).forEach((header) => {
    const value = headers[header];
    serializedHeaders[header.toLowerCase()] = value;
  });

  return serializedHeaders;
}

export function deserializeSuccess<TObject>(response: Response): TObject {
  try {
    return JSON.parse(response.content);
  } catch (e) {
    throw new DeserializationError((e as Error).message, response);
  }
}

export function deserializeFailure(
  { content, status }: Response,
  stackFrame: StackFrame[]
): Error {
  try {
    const parsed = JSON.parse(content);
    if ('error' in parsed) {
      return new DetailedApiError(
        parsed.message,
        status,
        parsed.error,
        stackFrame
      );
    }
    return new ApiError(parsed.message, status, stackFrame);
  } catch (e) {
    // ..
  }
  return new ApiError(content, status, stackFrame);
}
