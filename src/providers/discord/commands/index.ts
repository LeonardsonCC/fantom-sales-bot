import {
  ApplicationCommandDataResolvable,
  Client,
  CommandInteraction,
} from "discord.js";
import Ping from "./ping";
import Configure from "./configure";

export type CommandFunction = (
  interaction: CommandInteraction,
  client: Client
) => void;
export enum COMMANDS {
  PING = "ping",
  CONFIGURE = "configure",
}

const COMMANDS_OPTIONS: ApplicationCommandDataResolvable[] = [
  {
    name: COMMANDS.PING,
    description: "ping",
  },
  {
    name: COMMANDS.CONFIGURE,
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
  configure: Configure,
} as {
  [key in COMMANDS]: CommandFunction;
};
