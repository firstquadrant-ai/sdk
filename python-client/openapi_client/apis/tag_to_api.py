import typing_extensions

from openapi_client.apis.tags import TagValues
from openapi_client.apis.tags.teams_api import TeamsApi
from openapi_client.apis.tags.incoming_webhooks_api import IncomingWebhooksApi

TagToApi = typing_extensions.TypedDict(
    'TagToApi',
    {
        TagValues.TEAMS: TeamsApi,
        TagValues.INCOMINGWEBHOOKS: IncomingWebhooksApi,
    }
)

tag_to_api = TagToApi(
    {
        TagValues.TEAMS: TeamsApi,
        TagValues.INCOMINGWEBHOOKS: IncomingWebhooksApi,
    }
)
