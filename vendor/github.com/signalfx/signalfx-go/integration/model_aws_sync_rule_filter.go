/*
 * Integrations API
 *
 * APIs for creating, retrieving, updating, and deleting SignalFx integrations to the systems you use.<br> An integration provides SignalFx with information from the external system that you're connecting to. You'll need to retrieve this information from the external system before you use the API. Each external system is different, so to see a summary of its requirements and procedures, view its request body description. # Authentication To create, update, delete, or validate an integration, you need to authenticate your request using a session token associated with a SignalFx administrator. To **retrieve** an integration, your session token doesn't need to be associated with an administrator. You can also retrieve integrations using an org token.<br> In the web UI, session tokens are known as <strong>user access</strong> tokens, and org tokens are known as <strong>access tokens</strong>. <br> To learn more about authentication tokens, see the topic [Authentication Tokens](https://developers.signalfx.com/administration/access_tokens_overview.html) in the Developers Guide. # Supported service types SignalFx offers integrations for the following:<br>   * Data collection from other monitoring systems such as AWS CloudWatch   * Authentication using your existing Single Sign-On (**SSO**) system   * Sending alerts using your preferred messaging, chat, or incident management service <br> To use one of these integrations, you first register it with SignalFx. After that, you configure the integration to communicate between the system you're using and SignalFx. ## Data collection SignalFx integrations APIs support data collection for the following services:<br>   * Amazon Web Services (**AWS**)   * Google Cloud Platform (**GCP**)   * Microsoft Azure   * NewRelic  ## Authentication using SSO SignalFx integration APIs support SAML-based SSO integrations for the following services:<br>   * Microsoft Active Directory Federation Services (**ADFS**)   * Bitium   * Okta   * OneLogin   * PingOne  ## Alerts using message, chat, or incident management services SignalFx integration APIs support alert notifications using the following services: <br>   * BigPanda   * Office 365   * Opsgenie   * PagerDuty   * ServiceNow   * Slack   * VictorOps   * Webhook   * xMatters<br>  **NOTE:** You can't create Office 365 integrations using the API, and your ability to update them in a **PUT** request is limited, but you can retrieve their data or delete them. To create an Office 365 integration, use the the web UI. <br> # Viewing request body documentation The *request* body format for the following operations depends on the type of integration you use:<br>   * POST `/integration`   * PUT `/integration/{id}`<br>  The *response* body format for the following operations also depends on the type of integration you use:<br>   * GET `/integration`   * GET `/integration/{id}`  <br>  To see the request or response body format for an integration: <br>   1. Find the endpoint and method.   2. For a request body, find the section *REQUEST BODY SCHEMA*. For a     response body, find the section *RESPONSE SCHEMA*.   3. Scroll down to the `type` property.   4. At the end of the description for `type`, find the dropdown box that      contains the integration type. By default, it's set to *AWSCloudWatch*.   5. To see a complete list of integrations, click the down arrow. A list      with a vertical scroll bar appears.   6. Select the integration type from the list. The request body properties      for this integration type now appear.
 *
 * API version: 3.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package integration

// Filter that SignalFx applies to data coming in from an AWS namespace. This gives you more fine-grained control over the incoming data. If you don't specify a filter, SignalFx brings in all the data from the namespace.
type AwsSyncRuleFilter struct {
	Action AwsSyncRuleFilterAction `json:"action,omitempty"`
	// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule.<br> The expression uses the syntax defined for the SignalFlow `filter()`  function; it can be any valid SignalFlow filter expression.<br> **NOTE:** Some filter options aren't visible in the web UI for AWS filters. The SignalFlow expression for a filter that SignalFx can't display in the filters UI appears in the web UI, but you can only edit it with API calls.<br> You can refer to AWS tags by prefacing the tag name with the string `aws_tag_`. The value doesn't need a preface.<br> To refer to AWS metrics, preface the metric name with the string `sf_metric`. The value doesn't need a preface. See the topic [Integrating With AWS](https://developers.signalfx.com/integrating/aws_integration_overview.html) for more information.
	Source string `json:"source,omitempty"`
}
