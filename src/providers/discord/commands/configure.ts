import { CacheType, Client, CommandInteraction } from "discord.js";
import { CommandFunction } from ".";
import * as Database from "../../database";

const command: CommandFunction = async (
  interaction: CommandInteraction<CacheType>,
  client: Client
) => {
  const channelId = interaction.options.getChannel("channel")?.id;
  if (channelId) {
    await interaction.reply(`channel ${channelId}!`);
    const channel = client.channels.cache.get(channelId);
    if (channel && channel.isText()) {
      channel.send("The sales will appear here!");
      const db = Database.connect();
      db.collection("bot").insertOne({
        guildId: interaction.guildId,
        salesChannelId: interaction.channelId,
      });
    }
  }
};

export default command;
