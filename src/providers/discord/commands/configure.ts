import { CacheType, Client, CommandInteraction } from "discord.js";
import { CommandFunction } from ".";

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
      // TODO save the channel ID for this guild in a database
    }
  }
};

export default command;
