import {
  ApplicationCommandDataResolvable,
  Client,
  CommandInteraction,
} from "discord.js";
import Ping from "./ping";
import SetSalesChannel from "./set-sales-channel";
import SetMintChannel from "./set-mint-channel";

export type CommandFunction = (
  interaction: CommandInteraction,
  client: Client
) => void;
export enum COMMANDS {
  PING = "ping",
  SET_SALES_CHANNEL = "set-sales-channel",
  SET_MINT_CHANNEL = "set-mint-channel",
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
  {
    name: COMMANDS.SET_MINT_CHANNEL,
    description: "Track a contract mint",
    options: [
      {
        name: "channel",
        description: "channel that bot will post",
        type: "CHANNEL",
        required: true,
      },
      {
        name: "collection",
        description: "collection address to track",
        type: "STRING",
        required: true,
      },
    ],
  },
];

export { COMMANDS_OPTIONS };
export default {
  ping: Ping,
  "set-sales-channel": SetSalesChannel,
  "set-mint-channel": SetMintChannel,
} as {
  [key in COMMANDS]: CommandFunction;
};
