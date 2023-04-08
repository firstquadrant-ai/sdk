export * from './incomingWebhooks.service';
import { IncomingWebhooksService } from './incomingWebhooks.service';
export * from './teams.service';
import { TeamsService } from './teams.service';
export const APIS = [IncomingWebhooksService, TeamsService];
