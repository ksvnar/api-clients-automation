# Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

require 'cgi'

module Algolia
  class MonitoringClient
    attr_accessor :api_client

    def initialize(config = nil)
      @api_client = Algolia::ApiClient.new(config)
    end

    def self.create(app_id, api_key)
      hosts = []

      hosts << Transport::StatefulHost.new("#{app_id}-dsn.algolia.net", accept: CallType::READ)
      hosts << Transport::StatefulHost.new("#{app_id}.algolia.net", accept: CallType::WRITE)

      hosts += 1.upto(3).map do |i|
        Transport::StatefulHost.new("#{app_id}-#{i}.algolianet.com", accept: CallType::READ | CallType::WRITE)
      end.shuffle

      config = Algolia::Configuration.new(app_id, api_key, hosts, 'Monitoring')
      create_with_config(config)
    end

    def self.create_with_config(config)
      new(config)
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def custom_delete_with_http_info(path, parameters = nil, request_options = {})
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        raise ArgumentError, "Missing the required parameter 'path' when calling MonitoringClient.custom_delete"
      end

      path = '/1{path}'.sub('{' + 'path' + '}', path.to_s)
      query_params = {}
      query_params = query_params.merge(parameters) unless parameters.nil?
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.custom_delete',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:DELETE, path, new_options)
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Object]
    def custom_delete(path, parameters = nil, request_options = {})
      response = custom_delete_with_http_info(path, parameters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Object')
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def custom_get_with_http_info(path, parameters = nil, request_options = {})
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        raise ArgumentError, "Missing the required parameter 'path' when calling MonitoringClient.custom_get"
      end

      path = '/1{path}'.sub('{' + 'path' + '}', path.to_s)
      query_params = {}
      query_params = query_params.merge(parameters) unless parameters.nil?
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.custom_get',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Object]
    def custom_get(path, parameters = nil, request_options = {})
      response = custom_get_with_http_info(path, parameters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Object')
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param body [Object] Parameters to send with the custom request.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def custom_post_with_http_info(path, parameters = nil, body = nil, request_options = {})
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        raise ArgumentError, "Missing the required parameter 'path' when calling MonitoringClient.custom_post"
      end

      path = '/1{path}'.sub('{' + 'path' + '}', path.to_s)
      query_params = {}
      query_params = query_params.merge(parameters) unless parameters.nil?
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body] || @api_client.object_to_http_body(body)

      new_options = request_options.merge(
        :operation => :'MonitoringClient.custom_post',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:POST, path, new_options)
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param body [Object] Parameters to send with the custom request.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Object]
    def custom_post(path, parameters = nil, body = nil, request_options = {})
      response = custom_post_with_http_info(path, parameters, body, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Object')
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param body [Object] Parameters to send with the custom request.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def custom_put_with_http_info(path, parameters = nil, body = nil, request_options = {})
      # verify the required parameter 'path' is set
      if @api_client.config.client_side_validation && path.nil?
        raise ArgumentError, "Missing the required parameter 'path' when calling MonitoringClient.custom_put"
      end

      path = '/1{path}'.sub('{' + 'path' + '}', path.to_s)
      query_params = {}
      query_params = query_params.merge(parameters) unless parameters.nil?
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body] || @api_client.object_to_http_body(body)

      new_options = request_options.merge(
        :operation => :'MonitoringClient.custom_put',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:PUT, path, new_options)
    end

    # Send requests to the Algolia REST API.
    # This method allow you to send requests to the Algolia REST API.
    # @param path [String] Path of the endpoint, anything after \&quot;/1\&quot; must be specified. (required)
    # @param parameters [Hash<String, Object>] Query parameters to apply to the current query.
    # @param body [Object] Parameters to send with the custom request.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Object]
    def custom_put(path, parameters = nil, body = nil, request_options = {})
      response = custom_put_with_http_info(path, parameters, body, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Object')
    end

    # List incidents for selected clusters.
    # List known incidents for selected clusters.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_cluster_incidents_with_http_info(clusters, request_options = {})
      # verify the required parameter 'clusters' is set
      if @api_client.config.client_side_validation && clusters.nil?
        raise ArgumentError, "Missing the required parameter 'clusters' when calling MonitoringClient.get_cluster_incidents"
      end

      path = '/1/incidents/{clusters}'.sub('{' + 'clusters' + '}', CGI.escape(clusters.to_s))
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_cluster_incidents',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # List incidents for selected clusters.
    # List known incidents for selected clusters.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [IncidentsResponse]
    def get_cluster_incidents(clusters, request_options = {})
      response = get_cluster_incidents_with_http_info(clusters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::IncidentsResponse')
    end

    # List statuses of selected clusters.
    # Report whether a cluster is operational.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_cluster_status_with_http_info(clusters, request_options = {})
      # verify the required parameter 'clusters' is set
      if @api_client.config.client_side_validation && clusters.nil?
        raise ArgumentError, "Missing the required parameter 'clusters' when calling MonitoringClient.get_cluster_status"
      end

      path = '/1/status/{clusters}'.sub('{' + 'clusters' + '}', CGI.escape(clusters.to_s))
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_cluster_status',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # List statuses of selected clusters.
    # Report whether a cluster is operational.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [StatusResponse]
    def get_cluster_status(clusters, request_options = {})
      response = get_cluster_status_with_http_info(clusters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::StatusResponse')
    end

    # List incidents.
    # List known incidents for all clusters.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_incidents_with_http_info(request_options = {})
      path = '/1/incidents'
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_incidents',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # List incidents.
    # List known incidents for all clusters.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [IncidentsResponse]
    def get_incidents(request_options = {})
      response = get_incidents_with_http_info(request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::IncidentsResponse')
    end

    # Get indexing times.
    # List the average times for indexing operations for selected clusters.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_indexing_time_with_http_info(clusters, request_options = {})
      # verify the required parameter 'clusters' is set
      if @api_client.config.client_side_validation && clusters.nil?
        raise ArgumentError, "Missing the required parameter 'clusters' when calling MonitoringClient.get_indexing_time"
      end

      path = '/1/indexing/{clusters}'.sub('{' + 'clusters' + '}', CGI.escape(clusters.to_s))
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_indexing_time',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # Get indexing times.
    # List the average times for indexing operations for selected clusters.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [IndexingTimeResponse]
    def get_indexing_time(clusters, request_options = {})
      response = get_indexing_time_with_http_info(clusters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::IndexingTimeResponse')
    end

    # List servers.
    # List the servers belonging to clusters.  The response depends on whether you authenticate your API request:  - With authentication, the response lists the servers assigned to your Algolia application&#39;s cluster.  - Without authentication, the response lists the servers for all Algolia clusters.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_inventory_with_http_info(request_options = {})
      path = '/1/inventory/servers'
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_inventory',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # List servers.
    # List the servers belonging to clusters.  The response depends on whether you authenticate your API request:  - With authentication, the response lists the servers assigned to your Algolia application's cluster.  - Without authentication, the response lists the servers for all Algolia clusters.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [InventoryResponse]
    def get_inventory(request_options = {})
      response = get_inventory_with_http_info(request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::InventoryResponse')
    end

    # Get search latency times.
    # List the average latency for search requests for selected clusters.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_latency_with_http_info(clusters, request_options = {})
      # verify the required parameter 'clusters' is set
      if @api_client.config.client_side_validation && clusters.nil?
        raise ArgumentError, "Missing the required parameter 'clusters' when calling MonitoringClient.get_latency"
      end

      path = '/1/latency/{clusters}'.sub('{' + 'clusters' + '}', CGI.escape(clusters.to_s))
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_latency',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # Get search latency times.
    # List the average latency for search requests for selected clusters.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [LatencyResponse]
    def get_latency(clusters, request_options = {})
      response = get_latency_with_http_info(clusters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::LatencyResponse')
    end

    # Get metrics for a given period.
    # Report the aggregate value of a metric for a selected period of time.
    # @param metric [Metric] Metric to report.  For more information about the individual metrics, see the response. To include all metrics, use &#x60;*&#x60; as the parameter.  (required)
    # @param period [Period] Period over which to aggregate the metrics:  - &#x60;minute&#x60;. Aggregate the last minute. 1 data point per 10 seconds. - &#x60;hour&#x60;. Aggregate the last hour. 1 data point per minute. - &#x60;day&#x60;. Aggregate the last day. 1 data point per 10 minutes. - &#x60;week&#x60;. Aggregate the last week. 1 data point per hour. - &#x60;month&#x60;. Aggregate the last month. 1 data point per day.  (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_metrics_with_http_info(metric, period, request_options = {})
      # verify the required parameter 'metric' is set
      if @api_client.config.client_side_validation && metric.nil?
        raise ArgumentError, "Missing the required parameter 'metric' when calling MonitoringClient.get_metrics"
      end
      # verify the required parameter 'period' is set
      if @api_client.config.client_side_validation && period.nil?
        raise ArgumentError, "Missing the required parameter 'period' when calling MonitoringClient.get_metrics"
      end

      path = '/1/infrastructure/{metric}/period/{period}'.sub('{' + 'metric' + '}', CGI.escape(metric.to_s)).sub('{' + 'period' + '}', CGI.escape(period.to_s))
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_metrics',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # Get metrics for a given period.
    # Report the aggregate value of a metric for a selected period of time.
    # @param metric [Metric] Metric to report.  For more information about the individual metrics, see the response. To include all metrics, use &#x60;*&#x60; as the parameter.  (required)
    # @param period [Period] Period over which to aggregate the metrics:  - &#x60;minute&#x60;. Aggregate the last minute. 1 data point per 10 seconds. - &#x60;hour&#x60;. Aggregate the last hour. 1 data point per minute. - &#x60;day&#x60;. Aggregate the last day. 1 data point per 10 minutes. - &#x60;week&#x60;. Aggregate the last week. 1 data point per hour. - &#x60;month&#x60;. Aggregate the last month. 1 data point per day.  (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [InfrastructureResponse]
    def get_metrics(metric, period, request_options = {})
      response = get_metrics_with_http_info(metric, period, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::InfrastructureResponse')
    end

    # Test the reachability of clusters.
    # Test whether clusters are reachable or not.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_reachability_with_http_info(clusters, request_options = {})
      # verify the required parameter 'clusters' is set
      if @api_client.config.client_side_validation && clusters.nil?
        raise ArgumentError, "Missing the required parameter 'clusters' when calling MonitoringClient.get_reachability"
      end

      path = '/1/reachability/{clusters}/probes'.sub('{' + 'clusters' + '}', CGI.escape(clusters.to_s))
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_reachability',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # Test the reachability of clusters.
    # Test whether clusters are reachable or not.
    # @param clusters [String] Subset of clusters, separated by comma. (required)
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Hash<String, Hash>]
    def get_reachability(clusters, request_options = {})
      response = get_reachability_with_http_info(clusters, request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::Hash<String, Hash>')
    end

    # List cluster statuses.
    # Report whether clusters are operational.  The response depends on whether you authenticate your API request.  - With authentication, the response includes the status of the cluster assigned to your Algolia application.  - Without authentication, the response lists the statuses of all public Algolia clusters.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [Http::Response] the response
    def get_status_with_http_info(request_options = {})
      path = '/1/status'
      query_params = {}
      query_params = query_params.merge(request_options[:query_params]) unless request_options[:query_params].nil?
      header_params = {}
      header_params = header_params.merge(request_options[:header_params]) unless request_options[:header_params].nil?

      post_body = request_options[:debug_body]

      new_options = request_options.merge(
        :operation => :'MonitoringClient.get_status',
        :header_params => header_params,
        :query_params => query_params,
        :body => post_body,
        :use_read_transporter => false
      )

      @api_client.call_api(:GET, path, new_options)
    end

    # List cluster statuses.
    # Report whether clusters are operational.  The response depends on whether you authenticate your API request.  - With authentication, the response includes the status of the cluster assigned to your Algolia application.  - Without authentication, the response lists the statuses of all public Algolia clusters.
    # @param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
    # @return [StatusResponse]
    def get_status(request_options = {})
      response = get_status_with_http_info(request_options)
      deserialize(response.body, request_options[:debug_return_type] || 'Monitoring::StatusResponse')
    end
  end
end
