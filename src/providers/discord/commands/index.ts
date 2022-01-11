import { ApplicationCommandDataResolvable } from "discord.js";
import Ping from "./ping";

const COMMANDS_OPTIONS: ApplicationCommandDataResolvable[] = [
  {
    name: "ping",
    description: "ping",
  },
  {
    name: "configure",
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
};
