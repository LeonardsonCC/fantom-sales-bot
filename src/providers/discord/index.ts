import { Client, Intents, MessageEmbed } from "discord.js";
import {
  default as commands,
  COMMANDS_OPTIONS,
  COMMANDS,
} from "./commands/index";
import * as Database from "../database";
import { Sale } from "../../types/sale";
import { Action } from "../twitter/message";
import { makeMessage } from "./message";

const client = new Client({
  intents: [Intents.FLAGS.GUILD_MESSAGES],
});

const init = () => {
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

const post = async (
  sale: Sale,
  beforeEvent: Sale,
  action: Action,
  imageUrl?: string,
  image?: string,
  imageType: "PNG" | "JPG" | "GIF" = "PNG"
) => {
  const db = Database.connect();
  const guildsCursor = db.collection("bot").find({
    guildId: { $exists: true },
    salesChannelId: { $exists: true },
  });

  const message = await makeMessage(sale, beforeEvent, action);

  guildsCursor.forEach((guild) => {
    if (guild) {
      const guildId = guild.guildId;
      const channelId = guild.salesChannelId;

      console.log("DATA", guildId, channelId);
      if (message)
        send(guildId, channelId, message, imageUrl, image, imageType);
    }
  });
};

const send = async (
  guildId: string,
  channelId: string,
  message: MessageEmbed,
  imageUrl?: string,
  image?: string,
  imageType: "PNG" | "JPG" | "GIF" = "PNG"
) => {
  if (imageUrl) {
    message.setImage(imageUrl);
  }

  const guild = client.guilds.cache.get(guildId);
  if (guild) {
    const channel = await guild.channels.fetch(channelId);
    if (channel && channel.isText()) {
      channel
        .send({
          embeds: [message],
        })
        .then(() => console.log("SUCCESS"))
        .catch((err) => console.log(err));
    }
  }
};

export default { init, post };
