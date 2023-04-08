import typing_extensions

from openapi_client.paths import PathValues
from openapi_client.apis.paths.teams_me import TeamsMe
from openapi_client.apis.paths.incoming_webhooks_mark_conversations_as_won import IncomingWebhooksMarkConversationsAsWon

PathToApi = typing_extensions.TypedDict(
    'PathToApi',
    {
        PathValues.TEAMS_ME: TeamsMe,
        PathValues.INCOMINGWEBHOOKS_MARKCONVERSATIONSASWON: IncomingWebhooksMarkConversationsAsWon,
    }
)

path_to_api = PathToApi(
    {
        PathValues.TEAMS_ME: TeamsMe,
        PathValues.INCOMINGWEBHOOKS_MARKCONVERSATIONSASWON: IncomingWebhooksMarkConversationsAsWon,
    }
)
