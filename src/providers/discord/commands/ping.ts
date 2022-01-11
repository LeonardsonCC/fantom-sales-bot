import { CacheType, CommandInteraction } from "discord.js";

const command = async (interaction: CommandInteraction<CacheType>) => {
  return await interaction.reply("pong");
};

export default command;
