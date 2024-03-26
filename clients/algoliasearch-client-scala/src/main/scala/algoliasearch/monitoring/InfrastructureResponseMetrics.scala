/** Algolia Monitoring API The Monitoring API lets you check the status of your Algolia infrastructure. # Base URLs The
  * base URL for making requests to the Monitoring API is: - `https://status.algolia.com` **All requests must use
  * HTTPS.** # Availability and authentication Access to the [Infrastructure](#tag/infrastructure) endpoints is
  * available as part of the [Premium or Elevate plans](https://www.algolia.com/pricing). To authenticate requests to
  * the Infrastructure endpoints, add these headers: <dl> <dt><code>x-algolia-application-id</code></dt> <dd>Your
  * Algolia application ID.</dd> <dt><code>x-algolia-api-key</code></dt> <dd>Your Monitoring API key.</dd> </dl> You can
  * find your application ID and API key in the [Algolia dashboard](https://dashboard.algolia.com/account). Other
  * endpoints don't require authentication. # Response status and errors The Monitoring API returns JSON responses.
  * Since JSON doesn't guarantee any specific ordering, don't rely on the order of attributes in the API response.
  * Successful responses return a `2xx` status. Client errors return a `4xx` status. Server errors are indicated by a
  * `5xx` status. Error responses have a `message` property with more information. # Version The current version of the
  * Monitoring API is version 1, as indicated by the `/1/` in each endpoint's URL.
  *
  * The version of the OpenAPI document: 1.0.0
  *
  * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
  * https://openapi-generator.tech Do not edit the class manually.
  */
package algoliasearch.monitoring

/** InfrastructureResponseMetrics
  *
  * @param cpuUsage
  *   CPU idleness in %.
  * @param ramIndexingUsage
  *   RAM used for indexing in MB.
  * @param ramSearchUsage
  *   RAM used for search in MB.
  * @param ssdUsage
  *   Solid-state disk (SSD) usage expressed as % of RAM. 0% means no SSD usage. A value of 50% indicates 32&nbsp;GB SSD
  *   usage for a machine with 64&nbsp;RAM.
  * @param avgBuildTime
  *   Average build time of the indices in seconds.
  */
case class InfrastructureResponseMetrics(
    cpuUsage: Option[Map[String, Seq[ProbesMetric]]] = scala.None,
    ramIndexingUsage: Option[Map[String, Seq[ProbesMetric]]] = scala.None,
    ramSearchUsage: Option[Map[String, Seq[ProbesMetric]]] = scala.None,
    ssdUsage: Option[Map[String, Seq[ProbesMetric]]] = scala.None,
    avgBuildTime: Option[Map[String, Seq[ProbesMetric]]] = scala.None
)
