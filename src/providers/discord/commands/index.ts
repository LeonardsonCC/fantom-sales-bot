import {
  ApplicationCommandDataResolvable,
  Client,
  CommandInteraction,
} from "discord.js";
import Ping from "./ping";
import SetSalesChannel from "./set-sales-channel";

export type CommandFunction = (
  interaction: CommandInteraction,
  client: Client
) => void;
export enum COMMANDS {
  PING = "ping",
  SET_SALES_CHANNEL = "set-sales-channel",
}

const COMMANDS_OPTIONS: ApplicationCommandDataResolvable[] = [
  {
    name: COMMANDS.PING,
    description: "ping",
  },
  {
    name: COMMANDS.SET_SALES_CHANNEL,
    description: "Configure discord channel to the bot",
    options: [
      {
        name: "channel",
        description: "channel that bot will post",
        type: "CHANNEL",
        required: true,
      },
    ],
  },
];

export { COMMANDS_OPTIONS };
export default {
  ping: Ping,
  "set-sales-channel": SetSalesChannel,
} as {
  [key in COMMANDS]: CommandFunction;
};
