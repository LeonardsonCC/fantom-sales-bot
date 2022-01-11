import { Client, Intents } from "discord.js";
import { default as commands, COMMANDS_OPTIONS } from "./commands/index";

const init = () => {
  const client = new Client({
    intents: [Intents.FLAGS.GUILD_MESSAGES],
  });

  client.on("ready", async () => {
    if (client.application) {
      COMMANDS_OPTIONS.forEach((command) => {
        client.application!.commands.create(command);
      });
    }

    console.log(`Logged in as ${client?.user?.tag}!`);
  });

  client.on("interactionCreate", async (interaction) => {
    if (!interaction.isCommand()) return;

    // @ts-ignore
    commands[interaction.commandName](interaction, client);
  });

  client.login(process.env.DISCORD_TOKEN);
};

export default { init };
