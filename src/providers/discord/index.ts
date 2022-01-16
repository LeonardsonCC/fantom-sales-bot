import { Client, Intents } from "discord.js";
import {
  default as commands,
  COMMANDS_OPTIONS,
  COMMANDS,
} from "./commands/index";

const init = () => {
  const client = new Client({
    intents: [Intents.FLAGS.GUILD_MESSAGES],
  });

  client.on("ready", async () => {
    if (client.application) {
      COMMANDS_OPTIONS.forEach((command) => {
        client.application!.commands.create(command);
        console.log(`Command ${command.name} registered`);
      });
    }

    console.log(`Logged in as ${client?.user?.tag}!`);
  });

  client.on("interactionCreate", async (interaction) => {
    if (!interaction.isCommand()) return;

    commands[interaction.commandName as COMMANDS](interaction, client);
  });

  client.login(process.env.DISCORD_TOKEN);
};

export default { init };
